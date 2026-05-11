package cifailures

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

var githubAPIBaseURL = "https://api.github.com"

type ProwJobPR struct {
	Org      string
	Repo     string
	PRNumber int
}

type ChangeRelevanceResult struct {
	ProwJobURL   string                 `yaml:"prow_job_url"`
	JobName      string                 `yaml:"job_name"`
	Org          string                 `yaml:"org"`
	Repo         string                 `yaml:"repo"`
	PRNumber     int                    `yaml:"pr_number"`
	ChangedFiles []string               `yaml:"changed_files"`
	FailedTests  []ChangeRelevanceEntry `yaml:"failed_tests"`
	Summary      RelevanceSummary       `yaml:"summary"`
}

type ChangeRelevanceEntry struct {
	TestName     string   `yaml:"test_name"`
	SIG          string   `yaml:"sig,omitempty"`
	CodeAreas    []string `yaml:"code_areas,omitempty"`
	OverlapFiles []string `yaml:"overlap_files,omitempty"`
	Relevance    string   `yaml:"relevance"`
	Reason       string   `yaml:"reason"`
}

type RelevanceSummary struct {
	TotalFailedTests int `yaml:"total_failed_tests"`
	Related          int `yaml:"related"`
	PossiblyRelated  int `yaml:"possibly_related"`
	Unrelated        int `yaml:"unrelated"`
	Unknown          int `yaml:"unknown"`
}

var sigCodeAreas = map[string][]string{
	"sig-compute": {
		"pkg/virt-handler/",
		"pkg/virt-launcher/",
		"pkg/virt-controller/",
		"pkg/virt-api/",
		"pkg/libvmi/",
		"pkg/executor/",
		"pkg/hypervisor/",
		"cmd/virt-handler/",
		"cmd/virt-launcher/",
		"cmd/virt-controller/",
		"cmd/virt-api/",
		"tests/compute/",
	},
	"sig-network": {
		"pkg/network/",
		"tests/network/",
		"tests/libnet/",
	},
	"sig-storage": {
		"pkg/storage/",
		"pkg/container-disk/",
		"pkg/hotplug-disk/",
		"tests/storage/",
		"tests/libstorage/",
	},
	"sig-compute-migrations": {
		"pkg/virt-handler/migration/",
		"pkg/virt-controller/watch/migration",
		"tests/migration/",
		"tests/libmigration/",
	},
	"sig-operator": {
		"pkg/virt-operator/",
		"pkg/virt-controller/",
		"pkg/config/",
		"tests/operator/",
	},
	"sig-monitoring": {
		"pkg/monitoring/",
		"tests/monitoring/",
		"tests/libmonitoring/",
	},
}

var broadChangeAreas = []string{
	"staging/src/kubevirt.io/api/",
	"staging/src/kubevirt.io/client-go/",
	"pkg/util/",
	"tests/framework/",
	"tests/libvmi/",
	"tests/testsuite/",
	"hack/",
	"Makefile",
	"go.mod",
	"go.sum",
}

const prowPRLogsMarker = "/pr-logs/pull/"

func ParseProwJobPR(prowJobURL string) (*ProwJobPR, error) {
	jobURL := normalizeJobURL(prowJobURL)

	idx := strings.Index(jobURL, prowPRLogsMarker)
	if idx < 0 {
		return nil, fmt.Errorf("URL is not a PR build (no pr-logs/pull/ segment): %s", jobURL)
	}

	remainder := jobURL[idx+len(prowPRLogsMarker):]
	parts := strings.Split(remainder, "/")
	if len(parts) < 2 {
		return nil, fmt.Errorf("malformed Prow PR URL, expected {org}_{repo}/{pr}/...: %s", jobURL)
	}

	orgRepo := parts[0]
	sepIdx := strings.Index(orgRepo, "_")
	if sepIdx < 1 || sepIdx >= len(orgRepo)-1 {
		return nil, fmt.Errorf("malformed org_repo segment %q in URL: %s", orgRepo, jobURL)
	}

	prNumber, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, fmt.Errorf("invalid PR number %q in URL: %s", parts[1], jobURL)
	}

	return &ProwJobPR{
		Org:      orgRepo[:sepIdx],
		Repo:     orgRepo[sepIdx+1:],
		PRNumber: prNumber,
	}, nil
}

type ghPRFile struct {
	Filename string `json:"filename"`
}

func FetchPRChangedFiles(org, repo string, prNumber int) ([]string, error) {
	var allFiles []string
	page := 1
	client := &http.Client{Timeout: 30 * time.Second}

	for {
		params := url.Values{}
		params.Set("per_page", "100")
		params.Set("page", strconv.Itoa(page))
		apiURL := fmt.Sprintf("%s/repos/%s/%s/pulls/%d/files?%s",
			githubAPIBaseURL, org, repo, prNumber, params.Encode())

		req, err := http.NewRequest(http.MethodGet, apiURL, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to create request: %w", err)
		}
		req.Header.Set("Accept", "application/vnd.github.v3+json")

		if token := os.Getenv("GITHUB_TOKEN"); token != "" {
			req.Header.Set("Authorization", "Bearer "+token)
		}

		resp, err := client.Do(req)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch PR files: %w", err)
		}

		if resp.StatusCode != http.StatusOK {
			resp.Body.Close()
			return nil, fmt.Errorf("GitHub API returned status %d for %s", resp.StatusCode, apiURL)
		}

		var files []ghPRFile
		if err := json.NewDecoder(resp.Body).Decode(&files); err != nil {
			resp.Body.Close()
			return nil, fmt.Errorf("failed to decode GitHub API response: %w", err)
		}
		resp.Body.Close()

		for _, f := range files {
			allFiles = append(allFiles, f.Filename)
		}

		if len(files) < 100 {
			break
		}
		page++
	}

	return allFiles, nil
}

var sigLabelRegex = regexp.MustCompile(`\[sig-([a-zA-Z0-9-]+)\]`)

func ExtractSIGFromTestName(testName string) string {
	matches := sigLabelRegex.FindStringSubmatch(testName)
	if matches == nil {
		return ""
	}
	return "sig-" + matches[1]
}

func ExtractSIGFromJobName(jobName string) string {
	bestMatch := ""
	for sig := range sigCodeAreas {
		if strings.Contains(jobName, sig) && len(sig) > len(bestMatch) {
			bestMatch = sig
		}
	}
	return bestMatch
}

func CheckFileOverlap(changedFiles, prefixes []string) []string {
	var overlap []string
	seen := map[string]bool{}
	for _, f := range changedFiles {
		for _, p := range prefixes {
			if strings.HasPrefix(f, p) || f == p {
				if !seen[f] {
					seen[f] = true
					overlap = append(overlap, f)
				}
				break
			}
		}
	}
	return overlap
}

func ClassifyRelevance(sigOverlap, broadOverlap []string, sig string) (relevance, reason string) {
	if sig == "" {
		return "unknown", "no SIG label found in test name or job name"
	}
	if _, ok := sigCodeAreas[sig]; !ok {
		return "unknown", fmt.Sprintf("SIG %q not in code area mapping", sig)
	}
	if len(sigOverlap) > 0 {
		return "related", fmt.Sprintf("PR changes %d file(s) in %s code areas", len(sigOverlap), sig)
	}
	if len(broadOverlap) > 0 {
		return "possibly-related", fmt.Sprintf("PR changes %d file(s) in broad-impact areas (API, shared utils, test framework)", len(broadOverlap))
	}
	return "unrelated", fmt.Sprintf("PR changes do not overlap with %s code areas", sig)
}

func AnalyzeChangeRelevance(prowJobURL string) (*ChangeRelevanceResult, error) {
	prowJobURL = normalizeJobURL(prowJobURL)

	prInfo, err := ParseProwJobPR(prowJobURL)
	if err != nil {
		return nil, err
	}

	log.Infof("PR: %s/%s#%d", prInfo.Org, prInfo.Repo, prInfo.PRNumber)

	jobBuildErrors, err := AnalyzeBuild(prowJobURL)
	if err != nil {
		return nil, fmt.Errorf("failed to analyze build: %w", err)
	}

	failedTests := extractFailedTestNames(jobBuildErrors)
	if len(failedTests) == 0 {
		return nil, fmt.Errorf("no failed tests found in build log")
	}

	log.Infof("found %d failed test(s) in build log", len(failedTests))

	changedFiles, err := FetchPRChangedFiles(prInfo.Org, prInfo.Repo, prInfo.PRNumber)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch PR changed files: %w", err)
	}

	log.Infof("PR has %d changed file(s)", len(changedFiles))

	var entries []ChangeRelevanceEntry
	var summary RelevanceSummary

	broadOverlap := CheckFileOverlap(changedFiles, broadChangeAreas)
	for _, testName := range failedTests {
		sig := ExtractSIGFromTestName(testName)
		if sig == "" {
			sig = ExtractSIGFromJobName(jobBuildErrors.JobName)
		}

		var codeAreas []string
		if areas, ok := sigCodeAreas[sig]; ok {
			codeAreas = areas
		}

		sigOverlap := CheckFileOverlap(changedFiles, codeAreas)
		relevance, reason := ClassifyRelevance(sigOverlap, broadOverlap, sig)

		entry := ChangeRelevanceEntry{
			TestName:  testName,
			SIG:       sig,
			CodeAreas: codeAreas,
			Relevance: relevance,
			Reason:    reason,
		}
		if len(sigOverlap) > 0 {
			entry.OverlapFiles = sigOverlap
		} else if len(broadOverlap) > 0 {
			entry.OverlapFiles = broadOverlap
		}

		entries = append(entries, entry)

		summary.TotalFailedTests++
		switch relevance {
		case "related":
			summary.Related++
		case "possibly-related":
			summary.PossiblyRelated++
		case "unrelated":
			summary.Unrelated++
		case "unknown":
			summary.Unknown++
		}
	}

	return &ChangeRelevanceResult{
		ProwJobURL:   prowJobURL,
		JobName:      jobBuildErrors.JobName,
		Org:          prInfo.Org,
		Repo:         prInfo.Repo,
		PRNumber:     prInfo.PRNumber,
		ChangedFiles: changedFiles,
		FailedTests:  entries,
		Summary:      summary,
	}, nil
}

func WriteChangeRelevanceResultYAML(outputPath string, result *ChangeRelevanceResult) error {
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
