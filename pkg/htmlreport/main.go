package htmlreport

import (
	_ "embed"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/kubevirt/ci-health/pkg/types"
	log "github.com/sirupsen/logrus"
	"html/template"
	"io"
	"net"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)

var (
	jobRegex    string
	SigFailures []SigFailure

	//go:embed sig-failure-report.gohtml
	sigFailureReportTemplate string
)

type Results struct {
	Data struct {
		SIGRetests struct {
			FailedJobLeaderBoard []Job `json:"FailedJobLeaderBoard"`
		} `json:"SIGRetests"`
	} `json:"Data"`
}

type Job struct {
	JobName      string   `json:"JobName"`
	FailureCount int      `json:"FailureCount"`
	SuccessCount int      `json:"SuccessCount"`
	FailureURLs  []string `json:"FailureURLs"`
}

type Failure struct {
	XMLName xml.Name `xml:"failure"`
	Message string   `xml:"message,attr"`
	Type    string   `xml:"type,attr"`
	Value   string   `xml:",chardata"`
}

type Testcase struct {
	XMLName   xml.Name `xml:"testcase"`
	Classname string   `xml:"classname,attr"`
	Name      string   `xml:"name,attr"`
	Time      string   `xml:"time,attr"`
	Failure   *Failure `xml:"failure,omitempty"`
	URL       string   `xml:"url,omitempty"`
}

type Testsuite struct {
	XMLName  xml.Name   `xml:"testsuite"`
	Failures string     `xml:"failures,attr"`
	Name     string     `xml:"name,attr"`
	Tests    string     `xml:"tests,attr"`
	Time     string     `xml:"time,attr"`
	Testcase []Testcase `xml:"testcase"`
}

type SigFailure struct {
	Sig        string
	JobName    string
	FailureURL string
	Testcase   []Testcase
}

var jobRegexAliases = map[string]string{
	"compute": "sig-compute$|sig-compute-serial$|sig-compute-migrations$|sig-operator$|vgpu$|sev$",
	"network": "sig-network$|sriov$",
	"storage": "sig-storage$",
}

func fetchResults(resultsPath string) (*Results, error) {
	body, err := os.ReadFile(resultsPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read results.json file: %w", err)
	}

	var results Results
	if err := json.Unmarshal(body, &results); err != nil {
		return nil, fmt.Errorf("failed to unmarshal results.json: %w", err)
	}

	return &results, nil
}

func fetchJunit(url string) (*Testsuite, error) {
	client := &http.Client{
		Timeout: 60 * time.Second, // Increased timeout to 60 seconds
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return nil
		},
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout:   15 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			MaxIdleConns:        100,
			IdleConnTimeout:     90 * time.Second,
			TLSHandshakeTimeout: 15 * time.Second,
		},
	}

	resp, err := client.Get(url)
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

	var testsuite Testsuite
	if err := xml.Unmarshal(body, &testsuite); err == nil {
		return &testsuite, nil
	}

	return nil, fmt.Errorf("failed to unmarshal junit.functest.xml as <testsuites> or <testsuite>")
}

func constructReportFilePath(opt *types.Options) string {
	return fmt.Sprintf("%s/sig-%s-failure-report.html", opt.Path, opt.Sig)
}

func constructJunitURL(failureURL string) string {
	junitURL := strings.Replace(failureURL, "prow.ci.kubevirt.io//view/gs", "gcsweb.ci.kubevirt.io/gcs", 1)
	if !strings.HasSuffix(junitURL, "/") {
		junitURL += "/"
	}
	junitURL += "artifacts/junit.functest.xml"
	return junitURL
}

func Generate(opt *types.Options) {

	if opt.ResultsPath == "" {
		log.Errorf("the path to results.json is required")
	}

	if _, ok := jobRegexAliases[opt.Sig]; ok {
		jobRegex = jobRegexAliases[opt.Sig]
	}

	results, err := fetchResults(opt.ResultsPath)
	if err != nil {
		log.Errorf("failed to parse results.json")
	}

	jobRegex, err := regexp.Compile(jobRegex)
	if err != nil {
		log.Errorf("invalid job regex provided : %s", err)
	}

	for _, job := range results.Data.SIGRetests.FailedJobLeaderBoard {
		if !jobRegex.MatchString(job.JobName) {
			continue
		}
		for _, failureURL := range job.FailureURLs {
			var sigFail SigFailure
			junitURL := constructJunitURL(failureURL)
			testSuite, err := fetchJunit(junitURL)
			if err != nil {
				log.Errorf("failed to fetch junit results: %s", err)
			}
			if testSuite == nil {
				// SIG CI failure
				continue
			}
			sigFail.Sig = opt.Sig
			sigFail.JobName = job.JobName
			sigFail.FailureURL = failureURL

			for _, testcase := range testSuite.Testcase {
				if testcase.Failure == nil {
					continue
				}
				sigFail.Testcase = append(sigFail.Testcase, testcase)
			}
			SigFailures = append(SigFailures, sigFail)
		}

	}

	reportTemplate, err := template.New("sigFailures").Parse(sigFailureReportTemplate)
	if err != nil {
		log.Errorf("could not read template: %s", err)
	}
	outputFile, err := os.Create(constructReportFilePath(opt))
	if err != nil {
		log.Errorf("could not create report file: %s", err)
	}
	err = reportTemplate.Execute(outputFile, SigFailures)
	if err != nil {
		log.Errorf("could not execute template: %s", err)
	}
}
