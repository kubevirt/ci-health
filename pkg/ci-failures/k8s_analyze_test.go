package cifailures

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestParseAndDetect(t *testing.T) {
	tests := []struct {
		name          string
		artifactType  string
		json          string
		expectedCount int
	}{
		{
			name:         "pods with CrashLoopBackOff",
			artifactType: "pods",
			json: `{
				"items": [{
					"metadata": {"name": "test-pod", "namespace": "default"},
					"spec": {},
					"status": {
						"phase": "Running",
						"containerStatuses": [{
							"name": "app",
							"ready": false,
							"restartCount": 8,
							"state": {
								"waiting": {"reason": "CrashLoopBackOff"}
							}
						}]
					}
				}]
			}`,
			expectedCount: 1,
		},
		{
			name:         "nodes all healthy",
			artifactType: "nodes",
			json: `{
				"items": [{
					"metadata": {"name": "node1"},
					"status": {
						"conditions": [
							{"type": "Ready", "status": "True"},
							{"type": "DiskPressure", "status": "False"}
						]
					}
				}]
			}`,
			expectedCount: 0,
		},
		{
			name:         "events with warnings",
			artifactType: "events",
			json: `{
				"items": [{
					"metadata": {"name": "ev1"},
					"involvedObject": {"kind": "Pod", "name": "test", "namespace": "ns"},
					"type": "Warning",
					"reason": "FailedScheduling",
					"message": "0/3 nodes available",
					"count": 3
				}]
			}`,
			expectedCount: 1,
		},
		{
			name:         "vmis with failed phase",
			artifactType: "vmis",
			json: `{
				"items": [{
					"metadata": {"name": "vmi1", "namespace": "ns"},
					"status": {"phase": "Failed"}
				}]
			}`,
			expectedCount: 1,
		},
		{
			name:          "vmims empty",
			artifactType:  "vmims",
			json:          `{"items": []}`,
			expectedCount: 0,
		},
		{
			name:         "etcd-storage-profile with tmpfs exhaustion",
			artifactType: "etcd-storage-profile",
			json: `{
				"collectedAt": "2026-03-05T18:30:39Z",
				"totalSpecs": 2,
				"finalDBSizeBytes": 10000000,
				"finalTmpfsUsedBytes": 480000000,
				"finalTmpfsTotalBytes": 512000000,
				"finalWALSizeBytes": 50000000,
				"finalSnapSizeBytes": 10000000,
				"peakTmpfsUsedBytes": 480000000,
				"peakDBSizeBytes": 10000000,
				"records": []
			}`,
			expectedCount: 1,
		},
		{
			name:          "unknown type returns nil",
			artifactType:  "unknown",
			json:          `{}`,
			expectedCount: 0,
		},
		{
			name:          "invalid JSON returns nil",
			artifactType:  "pods",
			json:          `{invalid`,
			expectedCount: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			findings := parseAndDetect([]byte(tt.json), tt.artifactType)
			if len(findings) != tt.expectedCount {
				t.Errorf("expected %d findings, got %d: %+v", tt.expectedCount, len(findings), findings)
			}
		})
	}
}

func TestDiscoverAllSnapshotsFlatPath(t *testing.T) {
	mux := http.NewServeMux()
	// Non-parallel run: artifacts directly under k8s-reporter/
	mux.HandleFunc("/artifacts/k8s-reporter/1_pods.log", ok)
	mux.HandleFunc("/artifacts/k8s-reporter/1_nodes.log", ok)
	mux.HandleFunc("/", notFound)

	server := httptest.NewServer(mux)
	defer server.Close()

	snapshots := discoverAllSnapshots(server.URL)

	if len(snapshots) != 1 {
		t.Fatalf("expected 1 snapshot, got %d", len(snapshots))
	}
	if snapshots[0].snapshot.Process != "flat" {
		t.Errorf("expected process 'flat', got %q", snapshots[0].snapshot.Process)
	}
	if snapshots[0].snapshot.FailureCount != 1 {
		t.Errorf("expected failureCount 1, got %d", snapshots[0].snapshot.FailureCount)
	}
	if len(snapshots[0].urls) != 2 {
		t.Errorf("expected 2 artifact URLs, got %d", len(snapshots[0].urls))
	}
}

func TestDiscoverAllSnapshotsParallelSingleFailure(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/artifacts/k8s-reporter/1/1_pods.log", ok)
	mux.HandleFunc("/artifacts/k8s-reporter/1/1_nodes.log", ok)
	mux.HandleFunc("/artifacts/k8s-reporter/1/1_events.log", ok)
	mux.HandleFunc("/", notFound)

	server := httptest.NewServer(mux)
	defer server.Close()

	snapshots := discoverAllSnapshots(server.URL)

	if len(snapshots) != 1 {
		t.Fatalf("expected 1 snapshot, got %d", len(snapshots))
	}
	if snapshots[0].snapshot.Process != "1" {
		t.Errorf("expected process '1', got %q", snapshots[0].snapshot.Process)
	}
	if len(snapshots[0].urls) != 3 {
		t.Errorf("expected 3 artifacts, got %d", len(snapshots[0].urls))
	}
}

func TestDiscoverAllSnapshotsMultipleFailures(t *testing.T) {
	mux := http.NewServeMux()
	// Process 1, failures 1-3
	mux.HandleFunc("/artifacts/k8s-reporter/1/1_pods.log", ok)
	mux.HandleFunc("/artifacts/k8s-reporter/1/1_nodes.log", ok)
	mux.HandleFunc("/artifacts/k8s-reporter/1/2_pods.log", ok)
	mux.HandleFunc("/artifacts/k8s-reporter/1/2_nodes.log", ok)
	mux.HandleFunc("/artifacts/k8s-reporter/1/3_pods.log", ok)
	mux.HandleFunc("/", notFound)

	server := httptest.NewServer(mux)
	defer server.Close()

	snapshots := discoverAllSnapshots(server.URL)

	if len(snapshots) != 3 {
		t.Fatalf("expected 3 snapshots, got %d", len(snapshots))
	}
	for i, expected := range []int{1, 2, 3} {
		if snapshots[i].snapshot.FailureCount != expected {
			t.Errorf("snapshot %d: expected failureCount %d, got %d", i, expected, snapshots[i].snapshot.FailureCount)
		}
	}
}

func TestDiscoverAllSnapshotsMultipleProcesses(t *testing.T) {
	mux := http.NewServeMux()
	// Process 1: 2 failures
	mux.HandleFunc("/artifacts/k8s-reporter/1/1_pods.log", ok)
	mux.HandleFunc("/artifacts/k8s-reporter/1/2_pods.log", ok)
	// Process 4: 1 failure
	mux.HandleFunc("/artifacts/k8s-reporter/4/1_pods.log", ok)
	mux.HandleFunc("/", notFound)

	server := httptest.NewServer(mux)
	defer server.Close()

	snapshots := discoverAllSnapshots(server.URL)

	if len(snapshots) != 3 {
		t.Fatalf("expected 3 snapshots, got %d", len(snapshots))
	}

	// Process 1 should come first (lower process number)
	if snapshots[0].snapshot.Process != "1" || snapshots[0].snapshot.FailureCount != 1 {
		t.Errorf("expected process=1/failure=1, got %s", snapshots[0].snapshot)
	}
	if snapshots[1].snapshot.Process != "1" || snapshots[1].snapshot.FailureCount != 2 {
		t.Errorf("expected process=1/failure=2, got %s", snapshots[1].snapshot)
	}
	if snapshots[2].snapshot.Process != "4" || snapshots[2].snapshot.FailureCount != 1 {
		t.Errorf("expected process=4/failure=1, got %s", snapshots[2].snapshot)
	}
}

func TestDiscoverAllSnapshotsSuiteAndParallel(t *testing.T) {
	mux := http.NewServeMux()
	// Process 1
	mux.HandleFunc("/artifacts/k8s-reporter/1/1_pods.log", ok)
	// Suite (uses failureCount 0)
	mux.HandleFunc("/artifacts/k8s-reporter/suite/0_pods.log", ok)
	mux.HandleFunc("/", notFound)

	server := httptest.NewServer(mux)
	defer server.Close()

	snapshots := discoverAllSnapshots(server.URL)

	if len(snapshots) != 2 {
		t.Fatalf("expected 2 snapshots, got %d", len(snapshots))
	}

	processes := map[string]bool{}
	for _, s := range snapshots {
		processes[s.snapshot.Process] = true
	}
	if !processes["1"] {
		t.Error("expected process '1'")
	}
	if !processes["suite"] {
		t.Error("expected process 'suite'")
	}
}

func TestDownloadK8sArtifactCaching(t *testing.T) {
	callCount := 0
	podJSON := `{"items": []}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		callCount++
		fmt.Fprint(w, podJSON) //nolint:errcheck
	}))
	defer server.Close()

	cacheDir := t.TempDir()

	raw1, err := downloadK8sArtifact(server.URL+"/pods.log", cacheDir, "pods")
	if err != nil {
		t.Fatalf("first download failed: %v", err)
	}
	if callCount != 1 {
		t.Fatalf("expected 1 HTTP call, got %d", callCount)
	}

	raw2, err := downloadK8sArtifact(server.URL+"/pods.log", cacheDir, "pods")
	if err != nil {
		t.Fatalf("second download failed: %v", err)
	}
	if callCount != 1 {
		t.Fatalf("expected cache hit (still 1 HTTP call), got %d", callCount)
	}

	if string(raw1) != string(raw2) {
		t.Error("cached content differs from original")
	}

	cached, err := os.ReadFile(filepath.Join(cacheDir, "pods.json"))
	if err != nil {
		t.Fatalf("cache file not found: %v", err)
	}
	if string(cached) != podJSON {
		t.Errorf("cache file content mismatch: got %q", string(cached))
	}
}

func TestBuildSummary(t *testing.T) {
	findings := []*K8sFinding{
		{Detector: "pod-crash-loop", Severity: "error", Kind: "Pod"},
		{Detector: "pod-crash-loop", Severity: "error", Kind: "Pod"},
		{Detector: "node-not-ready", Severity: "error", Kind: "Node"},
		{Detector: "event-warning", Severity: "warning", Kind: "Event"},
	}

	s := buildSummary(findings)

	if s.TotalFindings != 4 {
		t.Errorf("expected 4 total, got %d", s.TotalFindings)
	}
	if s.ByKind["Pod"] != 2 {
		t.Errorf("expected 2 pods, got %d", s.ByKind["Pod"])
	}
	if s.ByKind["Node"] != 1 {
		t.Errorf("expected 1 node, got %d", s.ByKind["Node"])
	}
	if s.BySeverity["error"] != 3 {
		t.Errorf("expected 3 errors, got %d", s.BySeverity["error"])
	}
	if s.BySeverity["warning"] != 1 {
		t.Errorf("expected 1 warning, got %d", s.BySeverity["warning"])
	}
	if s.ByDetector["pod-crash-loop"] != 2 {
		t.Errorf("expected 2 pod-crash-loop, got %d", s.ByDetector["pod-crash-loop"])
	}
}

func TestWriteK8sAnalysisResultYAML(t *testing.T) {
	result := &K8sAnalysisResult{
		ProwJobURL: "https://prow.ci.kubevirt.io/view/gs/test",
		JobName:    "test-job",
		BuildID:    12345,
		Snapshots:  []K8sSnapshot{{Process: "1", FailureCount: 1}},
		Findings: []*K8sFinding{{
			Detector: "pod-crash-loop",
			Severity: "error",
			Kind:     "Pod",
			Name:     "test-pod",
			Reason:   "CrashLoopBackOff",
			Message:  "Container app in CrashLoopBackOff",
			Snapshot: K8sSnapshot{Process: "1", FailureCount: 1},
		}},
		Summary: buildSummary([]*K8sFinding{{
			Detector: "pod-crash-loop",
			Severity: "error",
			Kind:     "Pod",
		}}),
	}

	outputPath := filepath.Join(t.TempDir(), "test-output.yaml")
	if err := WriteK8sAnalysisResultYAML(outputPath, result); err != nil {
		t.Fatalf("failed to write YAML: %v", err)
	}

	content, err := os.ReadFile(outputPath)
	if err != nil {
		t.Fatalf("failed to read output: %v", err)
	}

	yamlStr := string(content)
	for _, expected := range []string{"prow_job_url:", "pod-crash-loop", "CrashLoopBackOff", "test-pod", "process:", "failure_count:"} {
		if !strings.Contains(yamlStr, expected) {
			t.Errorf("YAML output missing %q", expected)
		}
	}
}

func TestFetchJobTimestamps(t *testing.T) {
	started := map[string]any{"timestamp": 1713520103}
	finished := map[string]any{"timestamp": 1713521809, "passed": false, "result": "FAILURE"}

	startedJSON, _ := json.Marshal(started)
	finishedJSON, _ := json.Marshal(finished)

	mux := http.NewServeMux()
	mux.HandleFunc("/started.json", func(w http.ResponseWriter, r *http.Request) {
		w.Write(startedJSON) //nolint:errcheck
	})
	mux.HandleFunc("/finished.json", func(w http.ResponseWriter, r *http.Request) {
		w.Write(finishedJSON) //nolint:errcheck
	})

	server := httptest.NewServer(mux)
	defer server.Close()

	s, f, err := fetchJobTimestamps(server.URL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if s.Unix() != 1713520103 {
		t.Errorf("started timestamp mismatch: got %d", s.Unix())
	}
	if f.Unix() != 1713521809 {
		t.Errorf("finished timestamp mismatch: got %d", f.Unix())
	}
}

func TestK8sSnapshotString(t *testing.T) {
	s := K8sSnapshot{Process: "1", FailureCount: 3}
	if s.String() != "process=1/failure=3" {
		t.Errorf("unexpected String(): %q", s.String())
	}
}

func TestFetchEtcdProfile(t *testing.T) {
	profileJSON := `{
		"collectedAt": "2026-03-05T18:30:39Z",
		"totalSpecs": 5,
		"finalDBSizeBytes": 10000000,
		"finalTmpfsUsedBytes": 200000000,
		"finalTmpfsTotalBytes": 536870912,
		"finalWALSizeBytes": 50000000,
		"finalSnapSizeBytes": 10000000,
		"peakTmpfsUsedBytes": 450000000,
		"peakDBSizeBytes": 12000000,
		"records": []
	}`

	mux := http.NewServeMux()
	mux.HandleFunc("/artifacts/etcd-profiler/etcd-storage-profile.json", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodHead {
			w.WriteHeader(http.StatusOK)
			return
		}
		fmt.Fprint(w, profileJSON) //nolint:errcheck
	})
	mux.HandleFunc("/", notFound)

	server := httptest.NewServer(mux)
	defer server.Close()

	cacheDir := t.TempDir()
	profile := fetchEtcdProfile(server.URL, cacheDir)

	if profile == nil {
		t.Fatal("expected non-nil profile")
	}
	if profile.TotalSpecs != 5 {
		t.Errorf("expected 5 specs, got %d", profile.TotalSpecs)
	}
	if profile.PeakTmpfsUsed != 450000000 {
		t.Errorf("expected peak tmpfs 450000000, got %d", profile.PeakTmpfsUsed)
	}
	if profile.FinalTmpfsTotal != 536870912 {
		t.Errorf("expected tmpfs total 536870912, got %d", profile.FinalTmpfsTotal)
	}
}

func TestFetchEtcdProfileNotPresent(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", notFound)

	server := httptest.NewServer(mux)
	defer server.Close()

	cacheDir := t.TempDir()
	profile := fetchEtcdProfile(server.URL, cacheDir)

	if profile != nil {
		t.Errorf("expected nil profile when not present, got %+v", profile)
	}
}

func TestParseContainerLogFileName(t *testing.T) {
	tests := []struct {
		name       string
		fileName   string
		expectOK   bool
		expectNS   string
		expectPod  string
		expectCont string
	}{
		{
			name:       "virt-controller log",
			fileName:   "0_kubevirt_virt-controller-65cc5dc497-75jjr-virt-controller.log",
			expectOK:   true,
			expectNS:   "kubevirt",
			expectPod:  "virt-controller-65cc5dc497-75jjr",
			expectCont: "virt-controller",
		},
		{
			name:       "virt-handler log",
			fileName:   "0_kubevirt_virt-handler-7cdkk-virt-handler.log",
			expectOK:   true,
			expectNS:   "kubevirt",
			expectPod:  "virt-handler-7cdkk",
			expectCont: "virt-handler",
		},
		{
			name:       "virt-api log",
			fileName:   "0_kubevirt_virt-api-5bf5dd6df9-6nxcf-virt-api.log",
			expectOK:   true,
			expectNS:   "kubevirt",
			expectPod:  "virt-api-5bf5dd6df9-6nxcf",
			expectCont: "virt-api",
		},
		{
			name:       "virt-operator log",
			fileName:   "0_kubevirt_virt-operator-7b646765b4-4b2ks-virt-operator.log",
			expectOK:   true,
			expectNS:   "kubevirt",
			expectPod:  "virt-operator-7b646765b4-4b2ks",
			expectCont: "virt-operator",
		},
		{
			name:     "previous log is skipped",
			fileName: "0_kubevirt_virt-controller-65cc5dc497-75jjr-virt-controller_previous.log",
			expectOK: false,
		},
		{
			name:     "non-kubevirt component is skipped by caller",
			fileName: "0_kube-system_coredns-64b977d65c-gkhvz-coredns.log",
			expectOK: false,
		},
		{
			name:     "object dump is not a container log",
			fileName: "0_pods.log",
			expectOK: false,
		},
		{
			name:     "non-log file",
			fileName: "0_pods.json",
			expectOK: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			meta, ok := parseContainerLogFileName(tt.fileName)
			if ok != tt.expectOK {
				t.Fatalf("expected ok=%v, got %v", tt.expectOK, ok)
			}
			if !ok {
				return
			}
			if meta.namespace != tt.expectNS {
				t.Errorf("namespace: expected %q, got %q", tt.expectNS, meta.namespace)
			}
			if meta.podName != tt.expectPod {
				t.Errorf("podName: expected %q, got %q", tt.expectPod, meta.podName)
			}
			if meta.containerName != tt.expectCont {
				t.Errorf("containerName: expected %q, got %q", tt.expectCont, meta.containerName)
			}
		})
	}
}

func TestListAndFilterContainerLogs(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/list", func(w http.ResponseWriter, r *http.Request) {
		resp := gcsListResponse{
			Items: []struct {
				Name string `json:"name"`
				Size string `json:"size"`
			}{
				{Name: "prefix/suite/0_kubevirt_virt-controller-abc123-virt-controller.log", Size: "1000"},
				{Name: "prefix/suite/0_kubevirt_virt-handler-xyz-virt-handler.log", Size: "2000"},
				{Name: "prefix/suite/0_pods.log", Size: "5000"},
				{Name: "prefix/suite/0_kube-system_coredns-abc-coredns.log", Size: "3000"},
				{Name: "prefix/suite/0_kubevirt_virt-controller-abc123-virt-controller_previous.log", Size: "900"},
			},
		}
		json.NewEncoder(w).Encode(resp) //nolint:errcheck
	})

	server := httptest.NewServer(mux)
	defer server.Close()

	files := listAndFilterContainerLogs(server.URL+"/list", server.URL+"/")

	if len(files) != 2 {
		t.Fatalf("expected 2 container log files, got %d", len(files))
	}

	foundController := false
	foundHandler := false
	for _, f := range files {
		if f.containerName == "virt-controller" {
			foundController = true
		}
		if f.containerName == "virt-handler" {
			foundHandler = true
		}
	}
	if !foundController {
		t.Error("expected virt-controller log")
	}
	if !foundHandler {
		t.Error("expected virt-handler log")
	}
}

func TestListAndFilterContainerLogsPagination(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/storage/v1/b/bucket/o", func(w http.ResponseWriter, r *http.Request) {
		pageToken := r.URL.Query().Get("pageToken")
		var resp gcsListResponse
		switch pageToken {
		case "":
			resp = gcsListResponse{
				Items: []struct {
					Name string `json:"name"`
					Size string `json:"size"`
				}{
					{Name: "suite/0_kubevirt_virt-controller-abc-virt-controller.log", Size: "1000"},
				},
				NextPageToken: "page2",
			}
		case "page2":
			resp = gcsListResponse{
				Items: []struct {
					Name string `json:"name"`
					Size string `json:"size"`
				}{
					{Name: "suite/0_kubevirt_virt-handler-xyz-virt-handler.log", Size: "2000"},
				},
			}
		}
		json.NewEncoder(w).Encode(resp) //nolint:errcheck
	})

	server := httptest.NewServer(mux)
	defer server.Close()

	listURL := server.URL + "/storage/v1/b/bucket/o?prefix=suite%2F&maxResults=500"
	files := listAndFilterContainerLogs(listURL, server.URL+"/")

	if len(files) != 2 {
		t.Fatalf("expected 2 container log files across pages, got %d", len(files))
	}
	if files[0].containerName != "virt-controller" {
		t.Errorf("expected virt-controller first, got %q", files[0].containerName)
	}
	if files[1].containerName != "virt-handler" {
		t.Errorf("expected virt-handler second, got %q", files[1].containerName)
	}
}

func TestGcsBaseToListURL(t *testing.T) {
	listURL, downloadBase := gcsBaseToListURL("https://storage.googleapis.com/kubevirt-prow/logs/job/123/artifacts/k8s-reporter/suite")
	if listURL == "" {
		t.Fatal("expected non-empty listURL")
	}
	if !strings.Contains(listURL, "storage/v1/b/kubevirt-prow/o") {
		t.Errorf("unexpected listURL: %s", listURL)
	}
	if !strings.Contains(listURL, "maxResults=500") {
		t.Errorf("expected maxResults parameter in listURL: %s", listURL)
	}
	if !strings.Contains(listURL, "prefix=") {
		t.Errorf("expected prefix parameter in listURL: %s", listURL)
	}
	if downloadBase != "https://storage.googleapis.com/kubevirt-prow/" {
		t.Errorf("unexpected downloadBase: %s", downloadBase)
	}

	listURL, _ = gcsBaseToListURL("http://localhost:8080/test")
	if listURL != "" {
		t.Error("expected empty listURL for non-GCS URL")
	}
}

func TestDownloadContainerLogCaching(t *testing.T) {
	callCount := 0
	logContent := "level=error msg=\"test error\"\n"

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		callCount++
		fmt.Fprint(w, logContent) //nolint:errcheck
	}))
	defer server.Close()

	cacheDir := t.TempDir()
	fileName := "0_kubevirt_virt-controller-abc-virt-controller.log"

	path1, size1, err := downloadContainerLog(server.URL+"/log", cacheDir, fileName)
	if err != nil {
		t.Fatalf("first download failed: %v", err)
	}
	if callCount != 1 {
		t.Fatalf("expected 1 HTTP call, got %d", callCount)
	}
	if size1 != int64(len(logContent)) {
		t.Errorf("expected size %d, got %d", len(logContent), size1)
	}

	path2, _, err := downloadContainerLog(server.URL+"/log", cacheDir, fileName)
	if err != nil {
		t.Fatalf("second download failed: %v", err)
	}
	if callCount != 1 {
		t.Errorf("expected cache hit (still 1 call), got %d", callCount)
	}
	if path1 != path2 {
		t.Error("cached path should be the same")
	}
}

func TestFetchContainerLogsEndToEnd(t *testing.T) {
	logContent := "level=error msg=\"schema validation failed\"\n"

	mux := http.NewServeMux()
	mux.HandleFunc("/list", func(w http.ResponseWriter, r *http.Request) {
		resp := gcsListResponse{
			Items: []struct {
				Name string `json:"name"`
				Size string `json:"size"`
			}{
				{Name: "suite/0_kubevirt_virt-controller-abc-virt-controller.log", Size: "100"},
			},
		}
		json.NewEncoder(w).Encode(resp) //nolint:errcheck
	})
	mux.HandleFunc("/suite/0_kubevirt_virt-controller-abc-virt-controller.log", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, logContent) //nolint:errcheck
	})
	mux.HandleFunc("/", notFound)

	server := httptest.NewServer(mux)
	defer server.Close()

	cacheDir := t.TempDir()
	logCacheDir := filepath.Join(cacheDir, "container-logs")
	if err := os.MkdirAll(logCacheDir, 0755); err != nil {
		t.Fatal(err)
	}

	// Use listAndFilterContainerLogs directly, then download
	files := listAndFilterContainerLogs(server.URL+"/list", server.URL+"/")
	if len(files) != 1 {
		t.Fatalf("expected 1 file from listing, got %d", len(files))
	}

	f := files[0]
	cached, size, err := downloadContainerLog(f.url, logCacheDir, f.fileName)
	if err != nil {
		t.Fatalf("download failed: %v", err)
	}

	if f.containerName != "virt-controller" {
		t.Errorf("expected container name 'virt-controller', got %q", f.containerName)
	}
	if f.namespace != "kubevirt" {
		t.Errorf("expected namespace 'kubevirt', got %q", f.namespace)
	}
	if size != int64(len(logContent)) {
		t.Errorf("expected size %d, got %d", len(logContent), size)
	}

	content, err := os.ReadFile(cached)
	if err != nil {
		t.Fatalf("failed to read cached file: %v", err)
	}
	if string(content) != logContent {
		t.Errorf("cached content mismatch: got %q", string(content))
	}
}

func ok(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func notFound(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}
