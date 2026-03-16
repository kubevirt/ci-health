package htmlreport

import (
	_ "embed"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"regexp"
	"sort"
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
		SIGRetests struct {
			FailedJobLeaderBoard []types.FailedJob `json:"FailedJobLeaderBoard"`
		} `json:"SIGRetests"`
	} `json:"Data"`
}

type Failure struct {
	XMLName xml.Name `xml:"failure"`
	Message string   `xml:"message,attr"`
	Type    string   `xml:"type,attr"`
	Value   string   `xml:",chardata"`
}

type ReportData struct {
	GeneratedAt time.Time
	Failures    []SigFailure
}

type SigFailure struct {
	Sig        string
	JobName    string
	FailureURL string
	Started    time.Time
	Testcase   []junit.Test
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

	var sigFailures []SigFailure

	for _, job := range results.Data.SIGRetests.FailedJobLeaderBoard {
		if !compiledRegex.MatchString(job.JobName) {
			continue
		}
		for _, failureURL := range job.FailureURLs {
			var sigFail SigFailure
			junitURL := constructJunitURL(failureURL)
			testSuites, err := fetchJunit(junitURL)
			if err != nil {
				log.Warnf("failed to fetch junit results: %s", err)
				continue
			}
			if testSuites == nil {
				// SIG CI failure
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
			sigFailures = append(sigFailures, sigFail)
		}

	}

	sort.Slice(sigFailures, func(i, j int) bool {
		return sigFailures[i].Started.After(sigFailures[j].Started)
	})

	reportData := ReportData{
		GeneratedAt: time.Now().UTC(),
		Failures:    sigFailures,
	}

	reportTemplate, err := template.New("sigFailures").Parse(sigFailureReportTemplate)
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
