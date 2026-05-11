package cifailures

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"slices"
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

	containerLogFiles := fetchContainerLogs(gcsBaseURL, cacheDir)

	result := &K8sAnalysisResult{
		ProwJobURL:        prowJobURL,
		JobName:           jobName,
		BuildID:           buildID,
		Started:           started,
		Finished:          finished,
		Snapshots:         snapshotList,
		Findings:          allFindings,
		Summary:           buildSummary(allFindings),
		EtcdProfile:       etcdProfile,
		ContainerLogFiles: containerLogFiles,
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

var containerLogComponents = []string{
	"virt-controller",
	"virt-handler",
	"virt-api",
	"virt-operator",
}

func fetchContainerLogs(gcsBaseURL, cacheDir string) []ContainerLogFile {
	k8sReporterBase := gcsBaseURL + "/artifacts/k8s-reporter"

	// Container logs are in suite snapshots: {n}_{namespace}_{podName}-{containerName}.log
	// Try suite first, then flat layout.
	bases := []string{
		k8sReporterBase + "/suite",
		k8sReporterBase,
	}

	for _, base := range bases {
		files := discoverContainerLogFiles(base)
		if len(files) == 0 {
			continue
		}

		logCacheDir := filepath.Join(cacheDir, "container-logs")
		if err := os.MkdirAll(logCacheDir, 0755); err != nil {
			log.WithError(err).Warn("failed to create container-logs cache dir")
			return nil
		}

		var result []ContainerLogFile
		for _, f := range files {
			cached, size, err := downloadContainerLog(f.url, logCacheDir, f.fileName)
			if err != nil {
				log.WithError(err).Warnf("failed to download container log %s", f.fileName)
				continue
			}
			result = append(result, ContainerLogFile{
				PodName:       f.podName,
				ContainerName: f.containerName,
				Namespace:     f.namespace,
				CachedPath:    cached,
				SizeBytes:     size,
			})
			log.Infof("cached container log %s (%d bytes)", f.fileName, size)
		}
		return result
	}

	log.Info("no container log files found for kubevirt components")
	return nil
}

type containerLogMeta struct {
	fileName      string
	namespace     string
	podName       string
	containerName string
	url           string
}

// gcsListResponse is the JSON response from the GCS Storage JSON API list objects endpoint.
type gcsListResponse struct {
	Items []struct {
		Name string `json:"name"`
		Size string `json:"size"`
	} `json:"items"`
	NextPageToken string `json:"nextPageToken"`
}

// gcsBaseToListURL converts a GCS storage URL to the corresponding JSON API listing URL.
// Returns the listing URL and the download base URL, or empty strings if parsing fails.
func gcsBaseToListURL(gcsBasePath string) (listURL, downloadBase string) {
	const storagePrefix = "https://storage.googleapis.com/"
	if !strings.HasPrefix(gcsBasePath, storagePrefix) {
		return "", ""
	}

	remainder := strings.TrimPrefix(gcsBasePath, storagePrefix)
	slashIdx := strings.Index(remainder, "/")
	if slashIdx < 0 {
		return "", ""
	}
	bucket := remainder[:slashIdx]
	prefix := remainder[slashIdx+1:]
	if !strings.HasSuffix(prefix, "/") {
		prefix += "/"
	}

	v := url.Values{}
	v.Set("prefix", prefix)
	v.Set("maxResults", "500")

	return fmt.Sprintf("%sstorage/v1/b/%s/o?%s", storagePrefix, bucket, v.Encode()),
		storagePrefix + bucket + "/"
}

func discoverContainerLogFiles(gcsBasePath string) []containerLogMeta {
	listURL, downloadBase := gcsBaseToListURL(gcsBasePath)
	if listURL == "" {
		return nil
	}
	return listAndFilterContainerLogs(listURL, downloadBase)
}

func listAndFilterContainerLogs(listURL, downloadBase string) []containerLogMeta {
	var result []containerLogMeta
	pageURL := listURL

	for pageURL != "" {
		resp, err := http.Get(pageURL)
		if err != nil {
			log.WithError(err).Warn("failed to list GCS objects for container logs")
			return nil
		}

		if resp.StatusCode != http.StatusOK {
			closeAndLogErr(resp.Body)
			log.Warnf("GCS listing returned status %d", resp.StatusCode)
			return nil
		}

		var listing gcsListResponse
		if err := json.NewDecoder(resp.Body).Decode(&listing); err != nil {
			closeAndLogErr(resp.Body)
			log.WithError(err).Warn("failed to decode GCS listing response")
			return nil
		}
		closeAndLogErr(resp.Body)

		for _, item := range listing.Items {
			fileName := filepath.Base(item.Name)
			meta, ok := parseContainerLogFileName(fileName)
			if !ok {
				continue
			}
			if !isKubevirtComponent(meta.containerName) {
				continue
			}
			meta.url = downloadBase + item.Name
			result = append(result, meta)
		}

		if listing.NextPageToken == "" {
			break
		}
		pageURL = listURL + "&pageToken=" + url.QueryEscape(listing.NextPageToken)
	}

	return result
}

// parseContainerLogFileName parses filenames like "0_kubevirt_virt-controller-65cc5dc497-75jjr-virt-controller.log"
// into namespace, pod name, and container name components.
func parseContainerLogFileName(fileName string) (containerLogMeta, bool) {
	if !strings.HasSuffix(fileName, ".log") {
		return containerLogMeta{}, false
	}

	name := strings.TrimSuffix(fileName, ".log")

	// Format: {n}_{namespace}_{podName}-{containerName}
	// First segment before _ is the failure count number
	firstUnderscore := strings.Index(name, "_")
	if firstUnderscore < 0 {
		return containerLogMeta{}, false
	}

	rest := name[firstUnderscore+1:]

	// Second _ separates namespace from podName-containerName
	secondUnderscore := strings.Index(rest, "_")
	if secondUnderscore < 0 {
		return containerLogMeta{}, false
	}

	namespace := rest[:secondUnderscore]
	podContainer := rest[secondUnderscore+1:]

	// The container name is the last hyphen-separated segment that matches a known
	// component name. Since pod names can contain hyphens (e.g., "virt-controller-65cc5dc497-75jjr"),
	// we find the container name by checking known suffixes.
	for _, component := range containerLogComponents {
		suffix := "-" + component
		if strings.HasSuffix(podContainer, suffix) {
			podName := podContainer[:len(podContainer)-len(suffix)]
			return containerLogMeta{
				fileName:      fileName,
				namespace:     namespace,
				podName:       podName,
				containerName: component,
			}, true
		}
	}

	// Also match files where the pod name prefix equals the container name
	// (e.g., "virt-handler" container in "virt-handler-n64dk" pod has suffix "-virt-handler")
	// Already covered above. For "_previous" logs, skip them.
	if strings.HasSuffix(podContainer, "_previous") {
		return containerLogMeta{}, false
	}

	return containerLogMeta{}, false
}

func isKubevirtComponent(containerName string) bool {
	return slices.Contains(containerLogComponents, containerName)
}

func downloadContainerLog(gcsURL, cacheDir, fileName string) (string, int64, error) {
	cachePath := filepath.Join(cacheDir, fileName)

	info, err := os.Stat(cachePath)
	if err == nil {
		return cachePath, info.Size(), nil
	}
	if !errors.Is(err, fs.ErrNotExist) {
		return "", 0, err
	}

	resp, err := http.Get(gcsURL)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", 0, fmt.Errorf("failed to download log from GCS: status %d", resp.StatusCode)
	}

	out, err := os.Create(cachePath)
	if err != nil {
		return "", 0, err
	}
	defer out.Close()

	size, err := io.Copy(out, resp.Body)
	if err != nil {
		return "", 0, err
	}

	return cachePath, size, nil
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
