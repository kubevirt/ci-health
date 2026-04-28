package cifailures

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/kubevirt/ci-health/pkg/prow"
	"github.com/kubevirt/ci-health/pkg/sigretests"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

const k8sArtifactsCacheDir = "output/tmp/k8s-artifacts"

var k8sArtifactTypes = []string{
	"pods",
	"nodes",
	"events",
	"vmis",
	"vmims",
}

// artifactSnapshot groups the GCS URLs for one cluster state snapshot.
type artifactSnapshot struct {
	snapshot K8sSnapshot
	urls     map[string]string // artifactType -> GCS URL
}

// AnalyzeK8s downloads k8s-reporter artifacts for a prow job URL,
// parses them, runs failure detectors, and returns the analysis result.
func AnalyzeK8s(prowJobURL string) (*K8sAnalysisResult, error) {
	prowJobURL = normalizeJobURL(prowJobURL)

	gcsBaseURL := strings.Replace(prowJobURL,
		"https://prow.ci.kubevirt.io/view/gs/",
		"https://storage.googleapis.com/", 1)

	jobName, err := jobNameFromURL(prowJobURL)
	if err != nil {
		return nil, err
	}

	buildID, err := buildIDFromJobURL(prowJobURL)
	if err != nil {
		return nil, fmt.Errorf("failed to extract build id: %v", err)
	}

	started, finished, err := fetchJobTimestamps(gcsBaseURL)
	if err != nil {
		log.WithError(err).Warn("failed to fetch job timestamps, continuing without them")
	}

	cacheDir := filepath.Join(k8sArtifactsCacheDir, fmt.Sprintf("%d", buildID))
	if err = os.MkdirAll(cacheDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create cache directory: %w", err)
	}

	snapshots := discoverAllSnapshots(gcsBaseURL)
	if len(snapshots) == 0 {
		return nil, fmt.Errorf("no k8s-reporter artifacts found at %s/artifacts/k8s-reporter/", gcsBaseURL)
	}

	var allFindings []*K8sFinding
	var snapshotList []K8sSnapshot

	for _, snap := range snapshots {
		snapshotList = append(snapshotList, snap.snapshot)
		snapCacheDir := filepath.Join(cacheDir, snap.snapshot.Process, fmt.Sprintf("%d", snap.snapshot.FailureCount))
		if mkErr := os.MkdirAll(snapCacheDir, 0755); mkErr != nil {
			log.WithError(mkErr).Warnf("failed to create snapshot cache dir, skipping %s", snap.snapshot)
			continue
		}

		for artifactType, artifactURL := range snap.urls {
			raw, dlErr := downloadK8sArtifact(artifactURL, snapCacheDir, artifactType)
			if dlErr != nil {
				log.WithError(dlErr).Warnf("failed to download %s for %s, skipping", artifactType, snap.snapshot)
				continue
			}

			findings := parseAndDetect(raw, artifactType)
			for _, f := range findings {
				f.Snapshot = snap.snapshot
			}
			allFindings = append(allFindings, findings...)
		}
	}

	etcdProfile := fetchEtcdProfile(gcsBaseURL, cacheDir)
	if etcdProfile != nil {
		etcdFindings := detectEtcdIssues(etcdProfile)
		allFindings = append(allFindings, etcdFindings...)
	}

	result := &K8sAnalysisResult{
		ProwJobURL:  prowJobURL,
		JobName:     jobName,
		BuildID:     buildID,
		Started:     started,
		Finished:    finished,
		Snapshots:   snapshotList,
		Findings:    allFindings,
		Summary:     buildSummary(allFindings),
		EtcdProfile: etcdProfile,
	}

	return result, nil
}

func fetchJobTimestamps(gcsBaseURL string) (time.Time, time.Time, error) {
	var zero time.Time

	startedContent, err := retrieveFileContentFromGCS(gcsBaseURL + "/started.json")
	if err != nil {
		return zero, zero, fmt.Errorf("failed to fetch started.json: %v", err)
	}
	var started prow.Started
	if err = json.Unmarshal(startedContent, &started); err != nil {
		return zero, zero, fmt.Errorf("failed to decode started.json: %v", err)
	}

	finishedContent, err := retrieveFileContentFromGCS(gcsBaseURL + "/finished.json")
	if err != nil {
		return zero, zero, fmt.Errorf("failed to fetch finished.json: %v", err)
	}
	var finished prow.Finished
	if err = json.Unmarshal(finishedContent, &finished); err != nil {
		return zero, zero, fmt.Errorf("failed to decode finished.json: %v", err)
	}

	return started.Time(), finished.Time(), nil
}

// discoverAllSnapshots finds all cluster state snapshots across all parallel
// processes and failure counts. The artifact layout is:
//
//	k8s-reporter/{process}/{failureCount}_{type}.log  (parallel runs)
//	k8s-reporter/{failureCount}_{type}.log            (non-parallel runs)
//	k8s-reporter/suite/{failureCount}_{type}.log      (suite-level dump)
func discoverAllSnapshots(gcsBaseURL string) []artifactSnapshot {
	k8sReporterBase := gcsBaseURL + "/artifacts/k8s-reporter"

	// Discover which base paths have artifacts.
	type basePath struct {
		path    string
		process string // "1", "2", ..., "suite", or "flat"
	}
	var activePaths []basePath

	// Check parallel processes 1-12
	for p := 1; p <= 12; p++ {
		path := fmt.Sprintf("%s/%d", k8sReporterBase, p)
		if probeExists(fmt.Sprintf("%s/1_pods.log", path)) {
			activePaths = append(activePaths, basePath{path: path, process: fmt.Sprintf("%d", p)})
		}
	}

	// Check suite path
	suitePath := k8sReporterBase + "/suite"
	if probeExists(fmt.Sprintf("%s/0_pods.log", suitePath)) {
		activePaths = append(activePaths, basePath{path: suitePath, process: "suite"})
	}

	// Check flat path (non-parallel)
	if len(activePaths) == 0 {
		if probeExists(fmt.Sprintf("%s/1_pods.log", k8sReporterBase)) {
			activePaths = append(activePaths, basePath{path: k8sReporterBase, process: "flat"})
		}
	}

	if len(activePaths) == 0 {
		return nil
	}

	// For each active path, discover all failure counts.
	var result []artifactSnapshot

	for _, bp := range activePaths {
		startN := 1
		if bp.process == "suite" {
			startN = 0
		}

		for n := startN; n <= 20; n++ {
			urls := map[string]string{}
			for _, artifactType := range k8sArtifactTypes {
				url := fmt.Sprintf("%s/%d_%s.log", bp.path, n, artifactType)
				if probeExists(url) {
					urls[artifactType] = url
				}
			}
			if len(urls) == 0 {
				break
			}
			snap := artifactSnapshot{
				snapshot: K8sSnapshot{Process: bp.process, FailureCount: n},
				urls:     urls,
			}
			log.Infof("discovered snapshot %s with %d artifact types", snap.snapshot, len(urls))
			result = append(result, snap)
		}
	}

	return result
}

func probeExists(url string) bool {
	resp, err := sigretests.HttpHeadWithRetry(url)
	if err != nil {
		return false
	}
	_ = resp.Body.Close()
	return resp.StatusCode == http.StatusOK
}

func downloadK8sArtifact(gcsURL, cacheDir, artifactType string) ([]byte, error) {
	cachePath := filepath.Join(cacheDir, artifactType+".json")

	_, err := os.Stat(cachePath)
	if !errors.Is(err, fs.ErrNotExist) {
		return os.ReadFile(cachePath)
	}

	raw, err := retrieveFileContentFromGCS(gcsURL)
	if err != nil {
		return nil, err
	}

	if writeErr := os.WriteFile(cachePath, raw, 0644); writeErr != nil {
		log.WithError(writeErr).Warnf("failed to cache %s artifact", artifactType)
	}

	return raw, nil
}

func parseAndDetect(raw []byte, artifactType string) []*K8sFinding {
	switch artifactType {
	case "pods":
		var pods K8sPodList
		if err := json.Unmarshal(raw, &pods); err != nil {
			log.WithError(err).Warnf("failed to parse %s JSON", artifactType)
			return nil
		}
		return detectPodFailures(&pods)
	case "nodes":
		var nodes K8sNodeList
		if err := json.Unmarshal(raw, &nodes); err != nil {
			log.WithError(err).Warnf("failed to parse %s JSON", artifactType)
			return nil
		}
		return detectNodeFailures(&nodes)
	case "events":
		var events K8sEventList
		if err := json.Unmarshal(raw, &events); err != nil {
			log.WithError(err).Warnf("failed to parse %s JSON", artifactType)
			return nil
		}
		return detectEventFailures(&events)
	case "vmis":
		var vmis K8sVMIList
		if err := json.Unmarshal(raw, &vmis); err != nil {
			log.WithError(err).Warnf("failed to parse %s JSON", artifactType)
			return nil
		}
		return detectVMIFailures(&vmis)
	case "vmims":
		var vmims K8sVMIMList
		if err := json.Unmarshal(raw, &vmims); err != nil {
			log.WithError(err).Warnf("failed to parse %s JSON", artifactType)
			return nil
		}
		return detectVMIMFailures(&vmims)
	case "etcd-storage-profile":
		var profile EtcdStorageProfile
		if err := json.Unmarshal(raw, &profile); err != nil {
			log.WithError(err).Warnf("failed to parse %s JSON", artifactType)
			return nil
		}
		return detectEtcdIssues(&profile)
	default:
		log.Warnf("unknown artifact type: %s", artifactType)
		return nil
	}
}

func fetchEtcdProfile(gcsBaseURL, cacheDir string) *EtcdStorageProfile {
	etcdURL := gcsBaseURL + "/artifacts/etcd-profiler/etcd-storage-profile.json"
	if !probeExists(etcdURL) {
		log.Info("no etcd-storage-profile.json found, skipping etcd analysis")
		return nil
	}

	raw, err := downloadK8sArtifact(etcdURL, cacheDir, "etcd-storage-profile")
	if err != nil {
		log.WithError(err).Warn("failed to download etcd-storage-profile.json")
		return nil
	}

	var profile EtcdStorageProfile
	if err := json.Unmarshal(raw, &profile); err != nil {
		log.WithError(err).Warn("failed to parse etcd-storage-profile.json")
		return nil
	}

	log.Infof("loaded etcd profile: %d specs, peak tmpfs %d/%d bytes",
		profile.TotalSpecs, profile.PeakTmpfsUsed, profile.FinalTmpfsTotal)
	return &profile
}

func buildSummary(findings []*K8sFinding) K8sSummary {
	summary := K8sSummary{
		TotalFindings: len(findings),
		ByKind:        map[string]int{},
		BySeverity:    map[string]int{},
		ByDetector:    map[string]int{},
	}
	for _, f := range findings {
		summary.ByKind[f.Kind]++
		summary.BySeverity[f.Severity]++
		summary.ByDetector[f.Detector]++
	}
	return summary
}

// WriteK8sAnalysisResultYAML writes the analysis result to a file in YAML format.
func WriteK8sAnalysisResultYAML(outputPath string, result *K8sAnalysisResult) error {
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
