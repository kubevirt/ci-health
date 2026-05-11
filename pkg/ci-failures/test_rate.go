package cifailures

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

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
	ProwJobURL      string          `yaml:"prow_job_url"`
	JobName         string          `yaml:"job_name"`
	ReportPeriodEnd string          `yaml:"report_period_end"`
	FailedTests     []TestRateEntry `yaml:"failed_tests"`
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

func AnalyzeTestRate(prowJobURL string) (*TestRateResult, error) {
	prowJobURL = normalizeJobURL(prowJobURL)

	jobBuildErrors, err := AnalyzeBuild(prowJobURL)
	if err != nil {
		return nil, fmt.Errorf("failed to analyze build: %w", err)
	}

	failedTests := extractFailedTestNames(jobBuildErrors)
	if len(failedTests) == 0 {
		return nil, fmt.Errorf("no failed tests found in build log")
	}

	log.Infof("found %d failed test(s) in build log", len(failedTests))

	yesterday := time.Now().UTC().AddDate(0, 0, -1)
	report, err := fetchFlakefinderReport(yesterday)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch flakefinder report: %w", err)
	}

	log.Infof("fetched flakefinder report: %s to %s (%d tests)", report.StartOfReport, report.EndOfReport, len(report.Tests))

	var entries []TestRateEntry
	for _, testName := range failedTests {
		entry := computeTestRate(testName, report)
		entries = append(entries, entry)
	}

	return &TestRateResult{
		ProwJobURL:      prowJobURL,
		JobName:         jobBuildErrors.JobName,
		ReportPeriodEnd: yesterday.Format(time.DateOnly),
		FailedTests:     entries,
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
			if name == "" || seen[name] {
				continue
			}
			seen[name] = true
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
