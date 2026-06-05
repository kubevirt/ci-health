package cifailures

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

var testgridDashboards = []string{
	"kubevirt-periodics",
	"kubevirt-presubmits",
}

type DiscoverLanesResult struct {
	VersionFilter string                   `yaml:"version_filter"`
	Dashboards    []DiscoverLanesDashboard `yaml:"dashboards"`
}

type DiscoverLanesDashboard struct {
	Name  string               `yaml:"name"`
	Lanes []DiscoverLanesEntry `yaml:"lanes"`
}

type DiscoverLanesEntry struct {
	Tab           string `yaml:"tab"`
	OverallStatus string `yaml:"overall_status"`
	TestGridURL   string `yaml:"testgrid_url"`
}

type testgridSummaryEntry struct {
	OverallStatus string `json:"overall_status"`
}

func fetchDashboardSummary(dashboard string) (map[string]testgridSummaryEntry, error) {
	apiURL := fmt.Sprintf("%s/%s/summary", testgridBaseURL, dashboard)
	log.Infof("fetching testgrid summary from %s", apiURL)

	resp, err := httpClient.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch testgrid summary: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("testgrid returned status %d for %s", resp.StatusCode, apiURL)
	}

	var summary map[string]testgridSummaryEntry
	if err := json.NewDecoder(resp.Body).Decode(&summary); err != nil {
		return nil, fmt.Errorf("failed to decode testgrid summary: %w", err)
	}

	return summary, nil
}

func DiscoverLanes(version string) (*DiscoverLanesResult, error) {
	result := &DiscoverLanesResult{
		VersionFilter: version,
	}

	for _, dashboard := range testgridDashboards {
		summary, err := fetchDashboardSummary(dashboard)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch summary for %s: %w", dashboard, err)
		}

		var entries []DiscoverLanesEntry
		for tab, entry := range summary {
			if !strings.Contains(tab, version) {
				continue
			}
			entries = append(entries, DiscoverLanesEntry{
				Tab:           tab,
				OverallStatus: entry.OverallStatus,
				TestGridURL:   fmt.Sprintf("%s/%s#%s", testgridBaseURL, dashboard, tab),
			})
		}

		sort.Slice(entries, func(i, j int) bool {
			return entries[i].Tab < entries[j].Tab
		})

		result.Dashboards = append(result.Dashboards, DiscoverLanesDashboard{
			Name:  dashboard,
			Lanes: entries,
		})
	}

	return result, nil
}

func WriteDiscoverLanesResultYAML(outputPath string, result *DiscoverLanesResult) error {
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
