package cifailures

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"regexp"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

const maxDays = 28

var flakefinderBaseURL = "https://storage.googleapis.com/kubevirt-prow/reports/flakefinder/kubevirt/kubevirt"

type FlakefinderReport struct {
	StartOfReport string                            `json:"startOfReport"`
	EndOfReport   string                            `json:"endOfReport"`
	Headers       []string                          `json:"headers"`
	Tests         []string                          `json:"tests"`
	Data          map[string]map[string]*TestDetails `json:"data"`
}

type TestDetails struct {
	Succeeded int    `json:"succeeded"`
	Failed    int    `json:"failed"`
	Skipped   int    `json:"skipped"`
	Severity  string `json:"severity"`
}

type TestRateResult struct {
	ProwJobURL        string          `yaml:"prow_job_url"`
	JobName           string          `yaml:"job_name"`
	ReportDays        int             `yaml:"report_days"`
	ReportPeriodStart string          `yaml:"report_period_start"`
	ReportPeriodEnd   string          `yaml:"report_period_end"`
	FailedTests       []TestRateEntry `yaml:"failed_tests"`
}

type TestRateEntry struct {
	TestName       string  `yaml:"test_name"`
	MatchedName    string  `yaml:"matched_name,omitempty"`
	TotalSucceeded int     `yaml:"total_succeeded"`
	TotalFailed    int     `yaml:"total_failed"`
	TotalSkipped   int     `yaml:"total_skipped"`
	SuccessRate    float64 `yaml:"success_rate"`
	Severity       string  `yaml:"severity"`
}

func AnalyzeTestRate(prowJobURL string, days int) (*TestRateResult, error) {
	prowJobURL = normalizeJobURL(prowJobURL)

	if days < 1 {
		days = 7
	}
	if days > maxDays {
		days = maxDays
	}

	jobBuildErrors, err := AnalyzeBuild(prowJobURL)
	if err != nil {
		return nil, fmt.Errorf("failed to analyze build: %w", err)
	}

	failedTests := extractFailedTestNames(jobBuildErrors)
	if len(failedTests) == 0 {
		return nil, fmt.Errorf("no failed tests found in build log")
	}

	log.Infof("found %d failed test(s) in build log", len(failedTests))

	reports, err := fetchFlakefinderReports(days)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch flakefinder reports: %w", err)
	}

	merged := mergeReports(reports)

	log.Infof("merged %d flakefinder report(s) covering %s to %s (%d tests)", len(reports), merged.StartOfReport, merged.EndOfReport, len(merged.Tests))

	var entries []TestRateEntry
	for _, testName := range failedTests {
		entry := computeTestRate(testName, merged)
		entries = append(entries, entry)
	}

	return &TestRateResult{
		ProwJobURL:        prowJobURL,
		JobName:           jobBuildErrors.JobName,
		ReportDays:        days,
		ReportPeriodStart: merged.StartOfReport,
		ReportPeriodEnd:   merged.EndOfReport,
		FailedTests:       entries,
	}, nil
}

var ginkgoFailLineRegex = regexp.MustCompile(`\[FAIL\]\s+(.+)$`)

func extractFailedTestNames(jobBuildErrors *JobBuildErrors) []string {
	seen := map[string]bool{}
	var result []string

	for _, buildErr := range jobBuildErrors.BuildErrors {
		for _, snippet := range buildErr.BuildLogErrorSnippets {
			text := stripTimestamp(snippet.ErrorText)
			matches := ginkgoFailLineRegex.FindStringSubmatch(text)
			if matches == nil {
				continue
			}
			name := strings.TrimSpace(matches[1])
			norm := normalizeTestName(name)
			if name == "" || seen[norm] {
				continue
			}
			seen[norm] = true
			result = append(result, name)
		}
	}

	return result
}

var timestampPrefixRegex = regexp.MustCompile(`^\d{2}:\d{2}:\d{2}:\s*`)

func stripTimestamp(s string) string {
	return timestampPrefixRegex.ReplaceAllString(strings.TrimSpace(s), "")
}

func fetchFlakefinderReport(date time.Time) (*FlakefinderReport, error) {
	url := fmt.Sprintf("%s/flakefinder-%s-168h.json", flakefinderBaseURL, date.Format(time.DateOnly))
	log.Infof("fetching flakefinder report from %s", url)

	raw, err := retrieveFileContentFromGCS(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch flakefinder report: %w", err)
	}

	var report FlakefinderReport
	if err := json.Unmarshal(raw, &report); err != nil {
		return nil, fmt.Errorf("failed to decode flakefinder JSON: %w", err)
	}

	return &report, nil
}

func fetchFlakefinderReports(days int) ([]*FlakefinderReport, error) {
	numReports := int(math.Ceil(float64(days) / 7.0))
	var reports []*FlakefinderReport

	for r := range numReports {
		baseOffset := r * 7
		var report *FlakefinderReport
		var err error
		for try := 1; try <= 3; try++ {
			date := time.Now().UTC().AddDate(0, 0, -(baseOffset + try))
			report, err = fetchFlakefinderReport(date)
			if err == nil {
				break
			}
		}
		if err != nil {
			return nil, fmt.Errorf("failed to fetch flakefinder report for period %d: %w", r+1, err)
		}
		reports = append(reports, report)
	}

	return reports, nil
}

func mergeReports(reports []*FlakefinderReport) *FlakefinderReport {
	if len(reports) == 1 {
		return reports[0]
	}

	merged := &FlakefinderReport{
		StartOfReport: reports[0].StartOfReport,
		EndOfReport:   reports[0].EndOfReport,
		Data:          make(map[string]map[string]*TestDetails),
	}

	testSet := map[string]bool{}
	headerSet := map[string]bool{}

	for _, r := range reports {
		if r.StartOfReport < merged.StartOfReport {
			merged.StartOfReport = r.StartOfReport
		}
		if r.EndOfReport > merged.EndOfReport {
			merged.EndOfReport = r.EndOfReport
		}

		for _, h := range r.Headers {
			if !headerSet[h] {
				headerSet[h] = true
				merged.Headers = append(merged.Headers, h)
			}
		}

		for _, t := range r.Tests {
			if !testSet[t] {
				testSet[t] = true
				merged.Tests = append(merged.Tests, t)
			}
		}

		for testName, jobData := range r.Data {
			if merged.Data[testName] == nil {
				merged.Data[testName] = make(map[string]*TestDetails)
			}
			for jobName, details := range jobData {
				existing, ok := merged.Data[testName][jobName]
				if !ok {
					merged.Data[testName][jobName] = &TestDetails{
						Succeeded: details.Succeeded,
						Failed:    details.Failed,
						Skipped:   details.Skipped,
						Severity:  details.Severity,
					}
				} else {
					existing.Succeeded += details.Succeeded
					existing.Failed += details.Failed
					existing.Skipped += details.Skipped
				}
			}
		}
	}

	return merged
}

var (
	itMarkerRegex       = regexp.MustCompile(`\s*\[It\]\s*`)
	trailingLabelsRegex = regexp.MustCompile(`\s*\[[^\]]*,\s*[^\]]*\]\s*$`)
)

func normalizeTestName(raw string) string {
	name := strings.TrimSpace(raw)
	name = itMarkerRegex.ReplaceAllString(name, " ")
	name = trailingLabelsRegex.ReplaceAllString(name, "")
	return strings.TrimSpace(name)
}

func findMatchingTest(name string, report *FlakefinderReport) string {
	normalized := normalizeTestName(name)

	for _, candidate := range report.Tests {
		if candidate == name {
			return candidate
		}
	}

	for _, candidate := range report.Tests {
		if candidate == normalized {
			return candidate
		}
	}

	for _, candidate := range report.Tests {
		if strings.Contains(candidate, normalized) || strings.Contains(normalized, candidate) {
			return candidate
		}
	}

	return ""
}

func computeTestRate(testName string, report *FlakefinderReport) TestRateEntry {
	entry := TestRateEntry{
		TestName: testName,
		Severity: "unknown",
	}

	matched := findMatchingTest(testName, report)
	if matched == "" {
		log.Warnf("no flakefinder match for test: %s", testName)
		return entry
	}

	entry.MatchedName = matched

	jobData, ok := report.Data[matched]
	if !ok {
		return entry
	}

	for _, details := range jobData {
		entry.TotalSucceeded += details.Succeeded
		entry.TotalFailed += details.Failed
		entry.TotalSkipped += details.Skipped
	}

	total := entry.TotalSucceeded + entry.TotalFailed
	if total > 0 {
		entry.SuccessRate = float64(entry.TotalSucceeded) / float64(total) * 100.0
	}

	switch {
	case entry.SuccessRate >= 95:
		entry.Severity = "likely-pr-related"
	case entry.SuccessRate >= 80:
		entry.Severity = "inconclusive"
	default:
		entry.Severity = "likely-flaky"
	}

	return entry
}

func WriteTestRateResultYAML(outputPath string, result *TestRateResult) error {
	outputFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file %s: %v", outputPath, err)
	}
	defer outputFile.Close()

	encoder := yaml.NewEncoder(outputFile)
	if err := encoder.Encode(result); err != nil {
		return fmt.Errorf("failed to encode YAML: %v", err)
	}
	return nil
}
