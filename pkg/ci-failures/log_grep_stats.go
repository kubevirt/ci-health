package cifailures

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"sort"
	"strings"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

type LogGrepStatsResult struct {
	Pattern         string              `yaml:"pattern"`
	Since           string              `yaml:"since"`
	CaseInsensitive bool                `yaml:"case_insensitive"`
	TotalURLs       int                 `yaml:"total_urls"`
	CIFailureURLs   int                 `yaml:"ci_failure_urls"`
	TotalMatched    int                 `yaml:"total_matched"`
	Months          []LogGrepMonthStats `yaml:"months"`
}

type LogGrepMonthStats struct {
	Month   string              `yaml:"month"`
	Matched int                 `yaml:"matched"`
	Builds  []LogGrepBuildMatch `yaml:"builds"`
}

type LogGrepBuildMatch struct {
	BuildID     int       `yaml:"build_id"`
	ProwJobURL  string    `yaml:"prow_job_url"`
	Started     time.Time `yaml:"started"`
	Occurrences int       `yaml:"occurrences"`
}

func AnalyzeLogGrepStats(repoDir, pattern, since string, caseInsensitive bool, concurrency int) (*LogGrepStatsResult, error) {
	if concurrency < 1 {
		concurrency = 10
	}

	urls, err := collectFailureURLsFromGitHistory(repoDir, since)
	if err != nil {
		return nil, fmt.Errorf("failed to collect failure URLs from git history: %w", err)
	}

	log.Infof("collected %d unique failure URLs from git history since %s", len(urls), since)

	ciFailureURLs, err := filterCIFailureURLs(urls, concurrency)
	if err != nil {
		return nil, fmt.Errorf("failed to filter CI failure URLs: %w", err)
	}

	log.Infof("filtered to %d CI failure URLs (no junit.functest.xml)", len(ciFailureURLs))

	if err := os.MkdirAll(logsDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create build logs directory: %w", err)
	}

	logPaths, err := downloadBuildLogsForURLs(ciFailureURLs, concurrency)
	if err != nil {
		return nil, fmt.Errorf("failed to download build logs: %w", err)
	}

	log.Infof("downloaded/cached %d build logs", len(logPaths))

	result := grepBuildLogs(logPaths, pattern, caseInsensitive)
	result.Since = since
	result.TotalURLs = len(urls)
	result.CIFailureURLs = len(ciFailureURLs)

	return result, nil
}

func collectFailureURLsFromGitHistory(repoDir, since string) ([]string, error) {
	cmd := exec.Command("git", "log", "--format=%H", fmt.Sprintf("--since=%s", since), "--", "output/kubevirt/kubevirt/results.json")
	cmd.Dir = repoDir
	out, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("git log failed: %w", err)
	}

	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	if len(lines) == 0 || (len(lines) == 1 && lines[0] == "") {
		return nil, nil
	}

	seen := make(map[string]struct{})
	for _, commitHash := range lines {
		commitHash = strings.TrimSpace(commitHash)
		if commitHash == "" {
			continue
		}

		showCmd := exec.Command("git", "show", fmt.Sprintf("%s:output/kubevirt/kubevirt/results.json", commitHash))
		showCmd.Dir = repoDir
		jsonBytes, err := showCmd.Output()
		if err != nil {
			log.Warnf("skipping commit %s: git show failed: %v", commitHash, err)
			continue
		}

		urls, err := extractFailureURLsFromJSON(jsonBytes)
		if err != nil {
			log.Warnf("skipping commit %s: failed to parse results.json: %v", commitHash, err)
			continue
		}

		for _, u := range urls {
			seen[normalizeJobURL(u)] = struct{}{}
		}
	}

	result := make([]string, 0, len(seen))
	for u := range seen {
		result = append(result, u)
	}
	sort.Strings(result)
	return result, nil
}

func extractFailureURLsFromJSON(data []byte) ([]string, error) {
	var result struct {
		Data struct {
			SIGRetests struct {
				FailedJobLeaderBoard []struct {
					FailureURLs []string `json:"FailureURLs"`
				} `json:"FailedJobLeaderBoard"`
			} `json:"SIGRetests"`
		} `json:"Data"`
	}

	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	var urls []string
	for _, board := range result.Data.SIGRetests.FailedJobLeaderBoard {
		urls = append(urls, board.FailureURLs...)
	}
	return urls, nil
}

type ciFilterResult struct {
	url  string
	pass bool
}

func filterCIFailureURLs(urls []string, concurrency int) ([]string, error) {
	results := make(chan ciFilterResult, len(urls))
	sem := make(chan struct{}, concurrency)
	var wg sync.WaitGroup

	for _, u := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()

			exists, err := checkJunitFuncTestXMLExists(url)
			if err != nil {
				log.Warnf("junit check failed for %s: %v, including URL", url, err)
				results <- ciFilterResult{url: url, pass: true}
				return
			}
			results <- ciFilterResult{url: url, pass: !exists}
		}(u)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	var filtered []string
	for r := range results {
		if r.pass {
			filtered = append(filtered, r.url)
		}
	}
	sort.Strings(filtered)
	return filtered, nil
}

type downloadResult struct {
	url  string
	path string
	err  error
}

func downloadBuildLogsForURLs(urls []string, concurrency int) ([]string, error) {
	results := make(chan downloadResult, len(urls))
	sem := make(chan struct{}, concurrency)
	var wg sync.WaitGroup

	for _, u := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()

			path, err := storeBuildLogIfNotExists(url)
			results <- downloadResult{url: url, path: path, err: err}
		}(u)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	var paths []string
	for r := range results {
		if r.err != nil {
			log.Warnf("failed to download build log for %s: %v, skipping", r.url, r.err)
			continue
		}
		paths = append(paths, r.path)
	}
	sort.Strings(paths)
	return paths, nil
}

func grepBuildLogs(logPaths []string, pattern string, caseInsensitive bool) *LogGrepStatsResult {
	searchPattern := pattern
	if caseInsensitive {
		searchPattern = strings.ToLower(pattern)
	}

	monthMap := make(map[string]*LogGrepMonthStats)

	for _, logPath := range logPaths {
		bl, err := readBuildLog(logPath)
		if err != nil {
			log.Warnf("failed to read build log %s: %v, skipping", logPath, err)
			continue
		}

		content := bl.LogContent
		if caseInsensitive {
			content = strings.ToLower(content)
		}

		count := strings.Count(content, searchPattern)
		if count == 0 {
			continue
		}

		buildID, err := buildIDFromJobURL(bl.ProwJobURL)
		if err != nil {
			log.Warnf("failed to extract build ID from %s: %v, skipping", bl.ProwJobURL, err)
			continue
		}

		monthKey := bl.Started.Format("2006-01")
		ms, exists := monthMap[monthKey]
		if !exists {
			ms = &LogGrepMonthStats{Month: monthKey}
			monthMap[monthKey] = ms
		}
		ms.Matched++
		ms.Builds = append(ms.Builds, LogGrepBuildMatch{
			BuildID:     buildID,
			ProwJobURL:  bl.ProwJobURL,
			Started:     bl.Started,
			Occurrences: count,
		})
	}

	var months []LogGrepMonthStats
	for _, ms := range monthMap {
		sort.Slice(ms.Builds, func(i, j int) bool {
			return ms.Builds[i].Started.Before(ms.Builds[j].Started)
		})
		months = append(months, *ms)
	}
	sort.Slice(months, func(i, j int) bool {
		return months[i].Month < months[j].Month
	})

	totalMatched := 0
	for _, m := range months {
		totalMatched += m.Matched
	}

	return &LogGrepStatsResult{
		Pattern:         pattern,
		CaseInsensitive: caseInsensitive,
		TotalMatched:    totalMatched,
		Months:          months,
	}
}

func WriteLogGrepStatsResultYAML(outputPath string, result *LogGrepStatsResult) error {
	outputFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file %s: %w", outputPath, err)
	}
	defer func() {
		if err := outputFile.Close(); err != nil {
			log.WithError(err).Warn("failed closing output file")
		}
	}()

	encoder := yaml.NewEncoder(outputFile)
	if err := encoder.Encode(result); err != nil {
		return fmt.Errorf("failed to encode YAML: %w", err)
	}
	return nil
}
