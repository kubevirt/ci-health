package htmlreport

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/joshdk/go-junit"
	"github.com/kubevirt/ci-health/pkg/prow"
	"github.com/kubevirt/ci-health/pkg/sigretests"
	"github.com/kubevirt/ci-health/pkg/types"
	log "github.com/sirupsen/logrus"
)

var (
	//go:embed sig-failure-report.gohtml
	sigFailureReportTemplate string
)

type HTMLReportResults struct {
	Data struct {
		SIGRetests types.RunningAverageDataItem `json:"SIGRetests"`
	} `json:"Data"`
}

type JobWithFailures struct {
	types.FailedJob
	FailureRate float64
	Total       int
	SigFailures []SigFailure
}

type BranchStats struct {
	Branch      string
	Failures    int
	Successes   int
	Total       int
	FailureRate float64
	Jobs        []JobWithFailures
}

type TestFailureInstance struct {
	JobName    string
	FailureURL string
	Started    time.Time
	Error      error
}

type TestFailure struct {
	TestName  string
	Count     int
	Instances []TestFailureInstance
}

type OverallStats struct {
	TotalRuns   int
	Failures    int
	Successes   int
	FailureRate float64
}

type DateFailure struct {
	Date     string
	Count    int
	Failures []SigFailure
}

type ReportData struct {
	Sig          string
	GeneratedAt  time.Time
	Overall      OverallStats
	BranchStats  []BranchStats
	JobFailures  []JobWithFailures
	TestFailures []TestFailure
	DateFailures []DateFailure
}

type SigFailure struct {
	Sig        string
	JobName    string
	FailureURL string
	Started    time.Time
	Testcase   []junit.Test
}

func getJobTargetBranch(jobName string) string {
	parts := strings.Split(jobName, "-")
	lastPart := parts[len(parts)-1]
	if _, err := strconv.ParseFloat(lastPart, 32); err == nil {
		return fmt.Sprintf("release-%s", lastPart)
	}
	return "main"
}

func sortBranchKeys(branches []string) {
	sort.Slice(branches, func(i, j int) bool {
		bi, bj := branches[i], branches[j]
		if bi == "main" {
			return true
		}
		if bj == "main" {
			return false
		}
		vi := strings.TrimPrefix(bi, "release-")
		vj := strings.TrimPrefix(bj, "release-")
		fi, _ := strconv.ParseFloat(vi, 64)
		fj, _ := strconv.ParseFloat(vj, 64)
		return fi > fj
	})
}

var jobRegexAliases = map[string]string{
	"compute":    "sig-compute|sig-compute-serial|sig-compute-migrations|vgpu|sev",
	"operator":   "sig-operator",
	"monitoring": "sig-monitoring",
	"network":    "sig-network|sriov",
	"storage":    "sig-storage",
}

func fetchResults(resultsPath string) (*HTMLReportResults, error) {
	body, err := os.ReadFile(resultsPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read results.json file: %w", err)
	}

	var results HTMLReportResults
	if err := json.Unmarshal(body, &results); err != nil {
		return nil, fmt.Errorf("failed to unmarshal results.json: %w", err)
	}

	return &results, nil
}

func fetchJunit(url string) ([]junit.Suite, error) {
	resp, err := sigretests.HttpGetWithRetry(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch %s: %s", url, err)
	}
	defer resp.Body.Close()

	// Ignore missing junit files as it suggests an issue with the job
	if resp.StatusCode == http.StatusNotFound {
		return nil, nil
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch %s: status code %d", url, resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read %s body: %w", url, err)
	}

	testsuite, err := junit.Ingest(body)
	if err == nil {
		return testsuite, nil
	}

	return nil, fmt.Errorf("failed to unmarshal junit.functest.xml as <testsuites> or <testsuite>")
}

func constructReportFilePath(opt *types.Options) string {
	return fmt.Sprintf("%s/sig-%s-failure-report.html", opt.Path, opt.Sig)
}

func constructGCSBaseURL(failureURL string) string {
	return strings.Replace(failureURL, "https://prow.ci.kubevirt.io//view/gs/", "https://storage.googleapis.com/", 1)
}

func constructJunitURL(failureURL string) string {
	junitURL := strings.Replace(failureURL, "prow.ci.kubevirt.io//view/gs", "gcsweb.ci.kubevirt.io/gcs", 1)
	if !strings.HasSuffix(junitURL, "/") {
		junitURL += "/"
	}
	junitURL += "artifacts/junit.functest.xml"
	return junitURL
}

func fetchStartedTime(failureURL string) time.Time {
	gcsBaseURL := constructGCSBaseURL(failureURL)
	startedURL := gcsBaseURL + "/started.json"

	resp, err := sigretests.HttpGetWithRetry(startedURL)
	if err != nil {
		log.Warnf("failed to fetch started.json: %s", err)
		return time.Time{}
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Warnf("failed to fetch started.json from %s: status code %d", startedURL, resp.StatusCode)
		return time.Time{}
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Warnf("failed to read started.json body: %s", err)
		return time.Time{}
	}

	var started prow.Started
	if err := json.Unmarshal(body, &started); err != nil {
		log.Warnf("failed to unmarshal started.json: %s", err)
		return time.Time{}
	}

	return started.Time()
}

func Generate(opt *types.Options) error {

	if opt.ResultsPath == "" {
		return fmt.Errorf("the path to results.json is required")
	}

	jobRegex, ok := jobRegexAliases[opt.Sig]
	if !ok {
		return fmt.Errorf("unknown SIG: %s", opt.Sig)
	}

	results, err := fetchResults(opt.ResultsPath)
	if err != nil {
		return fmt.Errorf("failed to parse results.json: %w", err)
	}

	compiledRegex, err := regexp.Compile(jobRegex)
	if err != nil {
		return fmt.Errorf("invalid job regex provided: %w", err)
	}

	// Compute per-branch stats and collect individual SIG failures grouped by branch and job
	branchStatsMap := make(map[string]*BranchStats)
	for _, job := range results.Data.SIGRetests.FailedJobLeaderBoard {
		if !compiledRegex.MatchString(job.JobName) {
			continue
		}
		branch := getJobTargetBranch(job.JobName)
		bs, ok := branchStatsMap[branch]
		if !ok {
			bs = &BranchStats{Branch: branch}
			branchStatsMap[branch] = bs
		}
		bs.Successes += job.SuccessCount

		jwf := JobWithFailures{FailedJob: job}
		jwf.FailureCount = 0

		for _, failureURL := range job.FailureURLs {
			var sigFail SigFailure
			junitURL := constructJunitURL(failureURL)
			testSuites, err := fetchJunit(junitURL)
			if err != nil {
				log.Warnf("failed to fetch junit results: %s", err)
				continue
			}
			if testSuites == nil {
				// SIG CI failure - no junit results
				continue
			}
			sigFail.Sig = opt.Sig
			sigFail.JobName = job.JobName
			sigFail.FailureURL = failureURL
			sigFail.Started = fetchStartedTime(failureURL)

			for _, suite := range testSuites {
				for _, test := range suite.Tests {
					if test.Status == junit.StatusFailed {
						sigFail.Testcase = append(sigFail.Testcase, test)
					}
				}
			}
			jwf.SigFailures = append(jwf.SigFailures, sigFail)
			jwf.FailureCount++
		}
		bs.Failures += jwf.FailureCount
		sort.Slice(jwf.SigFailures, func(i, j int) bool {
			return jwf.SigFailures[i].Started.After(jwf.SigFailures[j].Started)
		})
		bs.Jobs = append(bs.Jobs, jwf)
	}

	var branchKeys []string
	for k := range branchStatsMap {
		branchKeys = append(branchKeys, k)
	}
	sortBranchKeys(branchKeys)

	var branchStatsList []BranchStats
	for _, k := range branchKeys {
		bs := branchStatsMap[k]
		bs.Total = bs.Failures + bs.Successes
		if bs.Total > 0 {
			bs.FailureRate = float64(bs.Failures) / float64(bs.Total) * 100
		}
		for i := range bs.Jobs {
			j := &bs.Jobs[i]
			j.Total = j.FailureCount + j.SuccessCount
			if j.Total > 0 {
				j.FailureRate = float64(j.FailureCount) / float64(j.Total) * 100
			}
		}
		sort.Slice(bs.Jobs, func(i, j int) bool {
			return bs.Jobs[i].FailureCount > bs.Jobs[j].FailureCount
		})
		branchStatsList = append(branchStatsList, *bs)
	}

	var allJobs []JobWithFailures
	for _, bs := range branchStatsList {
		allJobs = append(allJobs, bs.Jobs...)
	}
	sort.Slice(allJobs, func(i, j int) bool {
		return allJobs[i].FailureCount > allJobs[j].FailureCount
	})

	// Aggregate test failures across all jobs and branches
	testFailureMap := make(map[string]*TestFailure)
	for _, jwf := range allJobs {
		for _, sf := range jwf.SigFailures {
			for _, tc := range sf.Testcase {
				tf, ok := testFailureMap[tc.Name]
				if !ok {
					tf = &TestFailure{TestName: tc.Name}
					testFailureMap[tc.Name] = tf
				}
				tf.Count++
				tf.Instances = append(tf.Instances, TestFailureInstance{
					JobName:    sf.JobName,
					FailureURL: sf.FailureURL,
					Started:    sf.Started,
					Error:      tc.Error,
				})
			}
		}
	}

	var testFailures []TestFailure
	for _, tf := range testFailureMap {
		testFailures = append(testFailures, *tf)
	}
	sort.Slice(testFailures, func(i, j int) bool {
		return testFailures[i].Count > testFailures[j].Count
	})

	// Use the badge-level SIG retest/total stats for the overall summary
	sigRetests := results.Data.SIGRetests
	type sigStats struct{ retest, total float64 }
	sigStatsMap := map[string]sigStats{
		"compute":    {sigRetests.SIGComputeRetest, sigRetests.SIGComputeTotal},
		"storage":    {sigRetests.SIGStorageRetest, sigRetests.SIGStorageTotal},
		"network":    {sigRetests.SIGNetworkRetest, sigRetests.SIGNetworkTotal},
		"operator":   {sigRetests.SIGOperatorRetest, sigRetests.SIGOperatorTotal},
		"monitoring": {sigRetests.SIGMonitoringRetest, sigRetests.SIGMonitoringTotal},
	}
	ss := sigStatsMap[opt.Sig]
	var overall OverallStats
	overall.Failures = int(ss.retest)
	overall.TotalRuns = int(ss.total)
	overall.Successes = overall.TotalRuns - overall.Failures
	if overall.TotalRuns > 0 {
		overall.FailureRate = ss.retest / ss.total * 100
	}

	// Aggregate failures by date
	dateFailureMap := make(map[string]*DateFailure)
	for _, jwf := range allJobs {
		for _, sf := range jwf.SigFailures {
			if sf.Started.IsZero() {
				continue
			}
			dateKey := sf.Started.Format("2006-01-02")
			df, ok := dateFailureMap[dateKey]
			if !ok {
				df = &DateFailure{Date: dateKey}
				dateFailureMap[dateKey] = df
			}
			df.Count++
			df.Failures = append(df.Failures, sf)
		}
	}

	var dateFailures []DateFailure
	for _, df := range dateFailureMap {
		sort.Slice(df.Failures, func(i, j int) bool {
			return df.Failures[i].Started.After(df.Failures[j].Started)
		})
		dateFailures = append(dateFailures, *df)
	}
	sort.Slice(dateFailures, func(i, j int) bool {
		return dateFailures[i].Date > dateFailures[j].Date
	})

	reportData := ReportData{
		Sig:          strings.ToUpper(opt.Sig[:1]) + opt.Sig[1:],
		GeneratedAt:  time.Now().UTC(),
		Overall:      overall,
		BranchStats:  branchStatsList,
		JobFailures:  allJobs,
		TestFailures: testFailures,
		DateFailures: dateFailures,
	}

	funcMap := template.FuncMap{
		"buildID": func(failureURL string) string {
			parts := strings.Split(failureURL, "/")
			return parts[len(parts)-1]
		},
	}
	reportTemplate, err := template.New("sigFailures").Funcs(funcMap).Parse(sigFailureReportTemplate)
	if err != nil {
		return fmt.Errorf("could not read template: %w", err)
	}

	outputFile, err := os.Create(constructReportFilePath(opt))
	if err != nil {
		return fmt.Errorf("could not create report file: %w", err)
	}
	defer outputFile.Close()

	err = reportTemplate.Execute(outputFile, reportData)
	if err != nil {
		return fmt.Errorf("could not execute template: %w", err)
	}

	return nil
}
