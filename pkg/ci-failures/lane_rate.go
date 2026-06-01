package cifailures

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

var testgridBaseURL = "https://testgrid.k8s.io"

var httpClient = &http.Client{
	Timeout: 30 * time.Second,
}

type testGridTableResponse struct {
	Timestamps []int64         `json:"timestamps"`
	Tests      []testGridTest  `json:"tests"`
	ColumnIDs  []string        `json:"column_ids"`
}

type testGridTest struct {
	Name     string             `json:"name"`
	Statuses []testGridStatus   `json:"statuses"`
	Messages []string           `json:"messages"`
}

type testGridStatus struct {
	Count int `json:"count"`
	Value int `json:"value"`
}

const (
	tgStatusPass    = 1
	tgStatusFail    = 12
	tgStatusFlaky   = 3
)

type LaneRateResult struct {
	TestGridURL       string              `yaml:"testgrid_url"`
	Dashboard         string              `yaml:"dashboard"`
	Tab               string              `yaml:"tab"`
	ReportDays        int                 `yaml:"report_days"`
	ReportPeriodStart string              `yaml:"report_period_start"`
	ReportPeriodEnd   string              `yaml:"report_period_end"`
	TotalBuilds       int                 `yaml:"total_builds"`
	FailedTests       []LaneRateTestEntry `yaml:"failed_tests"`
}

type LaneRateTestEntry struct {
	TestName       string            `yaml:"test_name"`
	Succeeded      int               `yaml:"succeeded"`
	Failed         int               `yaml:"failed"`
	Skipped        int               `yaml:"skipped"`
	TotalRuns      int               `yaml:"total_runs"`
	SuccessRate    float64           `yaml:"success_rate"`
	Severity       string            `yaml:"severity"`
	RecentFailures []LaneRateFailure `yaml:"recent_failures,omitempty"`
}

type LaneRateFailure struct {
	BuildID   string `yaml:"build_id"`
	Timestamp string `yaml:"timestamp"`
	Message   string `yaml:"message,omitempty"`
}

func ParseTestGridURL(rawURL string) (dashboard, tab string, err error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return "", "", fmt.Errorf("invalid URL: %w", err)
	}

	dashboard = strings.Trim(u.Path, "/")
	if dashboard == "" {
		return "", "", fmt.Errorf("no dashboard in URL path: %s", rawURL)
	}

	fragment := u.Fragment
	if fragment == "" {
		return "", "", fmt.Errorf("no tab in URL fragment: %s", rawURL)
	}

	if idx := strings.Index(fragment, "&"); idx >= 0 {
		fragment = fragment[:idx]
	}
	tab = fragment
	if tab == "" {
		return "", "", fmt.Errorf("no tab name in URL fragment: %s", rawURL)
	}

	return dashboard, tab, nil
}

func fetchTestGridTable(dashboard, tab string) (*testGridTableResponse, error) {
	params := url.Values{}
	params.Add("tab", tab)
	apiURL := fmt.Sprintf("%s/%s/table?%s", testgridBaseURL, dashboard, params.Encode())
	log.Infof("fetching testgrid table from %s", apiURL)

	resp, err := httpClient.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch testgrid table: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("testgrid returned status %d for %s", resp.StatusCode, apiURL)
	}

	var table testGridTableResponse
	if err := json.NewDecoder(resp.Body).Decode(&table); err != nil {
		return nil, fmt.Errorf("failed to decode testgrid response: %w", err)
	}

	return &table, nil
}

func expandStatuses(statuses []testGridStatus) []int {
	var total int
	for _, s := range statuses {
		total += s.Count
	}
	result := make([]int, 0, total)
	for _, s := range statuses {
		for range s.Count {
			result = append(result, s.Value)
		}
	}
	return result
}

func filterColumnsInRange(timestamps []int64, days int) (indices []int, periodStart, periodEnd time.Time) {
	periodEnd = time.Now().UTC()
	periodStart = periodEnd.AddDate(0, 0, -days)
	cutoffMs := periodStart.UnixMilli()

	for i, ts := range timestamps {
		if ts >= cutoffMs {
			indices = append(indices, i)
		}
	}
	return indices, periodStart, periodEnd
}

func AnalyzeLaneRate(testgridURL string, days int, maxSuccessRate ...float64) (*LaneRateResult, error) {
	dashboard, tab, err := ParseTestGridURL(testgridURL)
	if err != nil {
		return nil, err
	}

	table, err := fetchTestGridTable(dashboard, tab)
	if err != nil {
		return nil, err
	}

	validIndices, periodStart, periodEnd := filterColumnsInRange(table.Timestamps, days)

	log.Infof("analyzing %d builds in %d-day window (%s to %s), %d tests total",
		len(validIndices), days,
		periodStart.Format(time.DateOnly), periodEnd.Format(time.DateOnly),
		len(table.Tests))

	var entries []LaneRateTestEntry

	for _, test := range table.Tests {
		if strings.HasSuffix(test.Name, ".Overall") {
			continue
		}

		flat := expandStatuses(test.Statuses)

		var succeeded, failed, skipped int
		var recentFailures []LaneRateFailure

		for _, idx := range validIndices {
			if idx >= len(flat) {
				skipped++
				continue
			}

			switch flat[idx] {
			case tgStatusPass:
				succeeded++
			case tgStatusFail, tgStatusFlaky:
				failed++
				if len(recentFailures) < 10 {
					f := LaneRateFailure{
						Timestamp: time.UnixMilli(table.Timestamps[idx]).UTC().Format(time.RFC3339),
					}
					if idx < len(table.ColumnIDs) {
						f.BuildID = strings.TrimLeft(table.ColumnIDs[idx], "")
					}
					if idx < len(test.Messages) && test.Messages[idx] != "" {
						f.Message = test.Messages[idx]
					}
					recentFailures = append(recentFailures, f)
				}
			default:
				skipped++
			}
		}

		if failed == 0 {
			continue
		}

		totalRuns := succeeded + failed
		var successRate float64
		if totalRuns > 0 {
			successRate = float64(succeeded) / float64(totalRuns) * 100.0
		}

		entries = append(entries, LaneRateTestEntry{
			TestName:       test.Name,
			Succeeded:      succeeded,
			Failed:         failed,
			Skipped:        skipped,
			TotalRuns:      totalRuns,
			SuccessRate:    successRate,
			Severity:       classifySeverity(successRate, totalRuns),
			RecentFailures: recentFailures,
		})
	}

	if len(maxSuccessRate) > 0 && maxSuccessRate[0] < 100 {
		threshold := maxSuccessRate[0]
		filtered := entries[:0]
		for _, e := range entries {
			if e.SuccessRate <= threshold {
				filtered = append(filtered, e)
			}
		}
		entries = filtered
	}

	sort.Slice(entries, func(i, j int) bool {
		if entries[i].SuccessRate != entries[j].SuccessRate {
			return entries[i].SuccessRate < entries[j].SuccessRate
		}
		return entries[i].TestName < entries[j].TestName
	})

	return &LaneRateResult{
		TestGridURL:       testgridURL,
		Dashboard:         dashboard,
		Tab:               tab,
		ReportDays:        days,
		ReportPeriodStart: periodStart.Format(time.RFC3339),
		ReportPeriodEnd:   periodEnd.Format(time.RFC3339),
		TotalBuilds:       len(validIndices),
		FailedTests:       entries,
	}, nil
}

func WriteLaneRateResultYAML(outputPath string, result *LaneRateResult) error {
	outputFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file %s: %w", outputPath, err)
	}
	defer outputFile.Close()

	encoder := yaml.NewEncoder(outputFile)
	if err := encoder.Encode(result); err != nil {
		return fmt.Errorf("failed to encode YAML: %w", err)
	}
	return nil
}
