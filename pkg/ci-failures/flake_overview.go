package cifailures

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"os"
	"regexp"
	"sort"
	"strings"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

type FlakeOverviewResult struct {
	ReportDate    string              `yaml:"report_date"`
	AnalysisDays  int                 `yaml:"analysis_days"`
	TotalFailures int                 `yaml:"total_failures"`
	Lanes         []FlakeOverviewLane `yaml:"lanes"`
}

type FlakeOverviewLane struct {
	Name         string                   `yaml:"name"`
	URL          string                   `yaml:"url"`
	Failures     int                      `yaml:"failures"`
	SharePercent float64                  `yaml:"share_percent"`
	TotalBuilds  int                      `yaml:"total_builds"`
	FailedTests  []FlakeOverviewTestEntry `yaml:"failed_tests,omitempty"`
}

type FlakeOverviewTestEntry struct {
	TestName    string  `yaml:"test_name"`
	Failed      int     `yaml:"failed"`
	TotalRuns   int     `yaml:"total_runs"`
	SuccessRate float64 `yaml:"success_rate"`
	Severity    string  `yaml:"severity"`
}

type laneSummary struct {
	name     string
	url      string
	failures int
}

func flakefinder24hReportURL(org, repo string, date time.Time) string {
	return fmt.Sprintf(
		"https://storage.googleapis.com/kubevirt-prow/reports/flakefinder/%s/%s/flakefinder-%s-024h.json",
		org, repo, date.Format(time.DateOnly),
	)
}

func fetchFlakefinder24hReport(reportURL string) (*FlakefinderReport, error) {
	resp, err := httpClient.Get(reportURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch %s: %w", reportURL, err)
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.WithError(err).Warn("failed closing response body")
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP %d fetching %s", resp.StatusCode, reportURL)
	}

	var report FlakefinderReport
	if err := json.NewDecoder(resp.Body).Decode(&report); err != nil {
		return nil, fmt.Errorf("failed to decode flakefinder JSON from %s: %w", reportURL, err)
	}
	return &report, nil
}

func testgridURLForJob(jobName string) string {
	switch {
	case strings.HasPrefix(jobName, "pull"):
		return fmt.Sprintf("https://testgrid.k8s.io/kubevirt-presubmits#%s", jobName)
	case strings.HasPrefix(jobName, "periodic"):
		return fmt.Sprintf("https://testgrid.k8s.io/kubevirt-periodics#%s", jobName)
	default:
		return ""
	}
}

func aggregateLaneFailures(reports []*FlakefinderReport, filterLaneRegex *regexp.Regexp) ([]laneSummary, int) {
	laneFailures := map[string]int{}

	for _, report := range reports {
		for _, testName := range report.Tests {
			for jobName, details := range report.Data[testName] {
				if details.Failed == 0 {
					continue
				}
				if filterLaneRegex != nil && filterLaneRegex.MatchString(jobName) {
					continue
				}
				laneFailures[jobName] += details.Failed
			}
		}
	}

	totalFailures := 0
	lanes := make([]laneSummary, 0, len(laneFailures))
	for name, failures := range laneFailures {
		totalFailures += failures
		lanes = append(lanes, laneSummary{
			name:     name,
			url:      testgridURLForJob(name),
			failures: failures,
		})
	}

	sort.Slice(lanes, func(i, j int) bool {
		if lanes[i].failures != lanes[j].failures {
			return lanes[i].failures > lanes[j].failures
		}
		return lanes[i].name < lanes[j].name
	})

	return lanes, totalFailures
}

func AnalyzeFlakeOverview(days int, org, repo, filterLaneRegexStr string, concurrency int) (*FlakeOverviewResult, error) {
	var filterLaneRegex *regexp.Regexp
	if filterLaneRegexStr != "" {
		var err error
		filterLaneRegex, err = regexp.Compile(filterLaneRegexStr)
		if err != nil {
			return nil, fmt.Errorf("invalid filter-lane-regex %q: %w", filterLaneRegexStr, err)
		}
	}

	reports, err := fetch24hReportsForDays(days, org, repo)
	if err != nil {
		return nil, err
	}

	lanes, totalFailures := aggregateLaneFailures(reports, filterLaneRegex)
	log.Infof("aggregated %d total failures across %d lanes", totalFailures, len(lanes))

	overviewLanes := fetchLaneRates(lanes, totalFailures, days, concurrency)

	return &FlakeOverviewResult{
		ReportDate:    time.Now().UTC().Format(time.DateOnly),
		AnalysisDays:  days,
		TotalFailures: totalFailures,
		Lanes:         overviewLanes,
	}, nil
}

func fetch24hReportsForDays(days int, org, repo string) ([]*FlakefinderReport, error) {
	reports := make([]*FlakefinderReport, 0, days)
	date := time.Now().UTC().AddDate(0, 0, -1)

	for i := 0; i < days; i++ {
		reportURL := flakefinder24hReportURL(org, repo, date)
		log.Infof("fetching flakefinder report: %s", reportURL)

		report, err := fetchFlakefinder24hReport(reportURL)
		if err != nil {
			return nil, fmt.Errorf("day %s: %w", date.Format(time.DateOnly), err)
		}
		reports = append(reports, report)
		date = date.AddDate(0, 0, -1)
	}

	return reports, nil
}

func fetchLaneRates(lanes []laneSummary, totalFailures, days, concurrency int) []FlakeOverviewLane {
	overviewLanes := make([]FlakeOverviewLane, len(lanes))
	if concurrency < 1 {
		concurrency = 1
	}
	sem := make(chan struct{}, concurrency)
	var wg sync.WaitGroup

	for i, lane := range lanes {
		sharePercent := 0.0
		if totalFailures > 0 {
			sharePercent = math.Round(float64(lane.failures)/float64(totalFailures)*10000) / 100
		}
		overviewLanes[i] = FlakeOverviewLane{
			Name:         lane.name,
			URL:          lane.url,
			Failures:     lane.failures,
			SharePercent: sharePercent,
		}

		if lane.url == "" {
			continue
		}

		wg.Add(1)
		go func(idx int, l laneSummary) {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()

			result, err := AnalyzeLaneRate(l.url, days)
			if err != nil {
				log.WithError(err).Warnf("lane-rate failed for %s, skipping", l.name)
				return
			}
			overviewLanes[idx].TotalBuilds = result.TotalBuilds
			overviewLanes[idx].FailedTests = compactTestEntries(result.FailedTests)
		}(i, lane)
	}

	wg.Wait()
	return overviewLanes
}

func compactTestEntries(full []LaneRateTestEntry) []FlakeOverviewTestEntry {
	compact := make([]FlakeOverviewTestEntry, 0, len(full))
	for _, e := range full {
		if e.Severity == "likely-pr-related" {
			continue
		}
		compact = append(compact, FlakeOverviewTestEntry{
			TestName:    e.TestName,
			Failed:      e.Failed,
			TotalRuns:   e.TotalRuns,
			SuccessRate: e.SuccessRate,
			Severity:    e.Severity,
		})
	}
	return compact
}

func WriteFlakeOverviewResultYAML(outputPath string, result *FlakeOverviewResult) error {
	f, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create %s: %w", outputPath, err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.WithError(err).Warn("failed closing output file")
		}
	}()

	encoder := yaml.NewEncoder(f)
	if err := encoder.Encode(result); err != nil {
		return fmt.Errorf("failed to encode YAML: %w", err)
	}
	return nil
}
