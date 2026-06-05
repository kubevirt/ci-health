package cifailures

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestParseProwJobPR(t *testing.T) {
	tests := []struct {
		name      string
		url       string
		wantOrg   string
		wantRepo  string
		wantPR    int
		wantError bool
	}{
		{
			name:     "standard PR URL",
			url:      "https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17760/pull-kubevirt-e2e-k8s-1.35-sig-compute/2053843644594524160",
			wantOrg:  "kubevirt",
			wantRepo: "kubevirt",
			wantPR:   17760,
		},
		{
			name:     "repo with hyphens",
			url:      "https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_containerized-data-importer/456/some-job/123456",
			wantOrg:  "kubevirt",
			wantRepo: "containerized-data-importer",
			wantPR:   456,
		},
		{
			name:     "double-slash URL normalized",
			url:      "https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/100/job/999",
			wantOrg:  "kubevirt",
			wantRepo: "kubevirt",
			wantPR:   100,
		},
		{
			name:      "periodic build",
			url:       "https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/logs/periodic-kubevirt-e2e-k8s-1.35-sig-compute/2053843644594524160",
			wantError: true,
		},
		{
			name:      "malformed org_repo",
			url:       "https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/badformat/17760/job/123",
			wantError: true,
		},
		{
			name:      "non-numeric PR number",
			url:       "https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/abc/job/123",
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ParseProwJobPR(tt.url)
			if tt.wantError {
				if err == nil {
					t.Error("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if result.Org != tt.wantOrg {
				t.Errorf("Org = %q, want %q", result.Org, tt.wantOrg)
			}
			if result.Repo != tt.wantRepo {
				t.Errorf("Repo = %q, want %q", result.Repo, tt.wantRepo)
			}
			if result.PRNumber != tt.wantPR {
				t.Errorf("PRNumber = %d, want %d", result.PRNumber, tt.wantPR)
			}
		})
	}
}

func TestExtractSIGFromTestName(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"sig-compute", "[sig-compute] Live Migrations with a strategy set", "sig-compute"},
		{"sig-network", "[sig-network] VMI connectivity should allow traffic", "sig-network"},
		{"sig-compute-migrations", "[sig-compute-migrations] Live Migration tests", "sig-compute-migrations"},
		{"sig-storage", "[sig-storage] DataVolume tests", "sig-storage"},
		{"no SIG label", "some test without sig label", ""},
		{"first SIG wins", "[sig-compute] test [sig-network] cross-sig", "sig-compute"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExtractSIGFromTestName(tt.input)
			if got != tt.expected {
				t.Errorf("ExtractSIGFromTestName(%q) = %q, want %q", tt.input, got, tt.expected)
			}
		})
	}
}

func TestExtractSIGFromJobName(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"sig-compute job", "pull-kubevirt-e2e-k8s-1.35-sig-compute", "sig-compute"},
		{"sig-network job", "pull-kubevirt-e2e-k8s-1.35-sig-network", "sig-network"},
		{"sig-compute-migrations job", "pull-kubevirt-e2e-k8s-1.35-sig-compute-migrations", "sig-compute-migrations"},
		{"no SIG in job name", "pull-kubevirt-unit-test", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExtractSIGFromJobName(tt.input)
			if got != tt.expected {
				t.Errorf("ExtractSIGFromJobName(%q) = %q, want %q", tt.input, got, tt.expected)
			}
		})
	}
}

func TestCheckFileOverlap(t *testing.T) {
	tests := []struct {
		name         string
		changedFiles []string
		prefixes     []string
		wantCount    int
	}{
		{
			name:         "direct match",
			changedFiles: []string{"pkg/virt-handler/vm.go", "pkg/network/pod.go"},
			prefixes:     []string{"pkg/virt-handler/"},
			wantCount:    1,
		},
		{
			name:         "no overlap",
			changedFiles: []string{"pkg/network/pod.go", "tests/network/vmi_test.go"},
			prefixes:     []string{"pkg/virt-handler/", "pkg/storage/"},
			wantCount:    0,
		},
		{
			name:         "multiple matches",
			changedFiles: []string{"pkg/virt-handler/vm.go", "pkg/virt-launcher/cmd.go", "README.md"},
			prefixes:     []string{"pkg/virt-handler/", "pkg/virt-launcher/"},
			wantCount:    2,
		},
		{
			name:         "exact file match",
			changedFiles: []string{"go.mod", "go.sum", "pkg/foo/bar.go"},
			prefixes:     []string{"go.mod", "go.sum"},
			wantCount:    2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CheckFileOverlap(tt.changedFiles, tt.prefixes)
			if len(got) != tt.wantCount {
				t.Errorf("CheckFileOverlap() returned %d files, want %d: %v", len(got), tt.wantCount, got)
			}
		})
	}
}

func TestClassifyRelevance(t *testing.T) {
	tests := []struct {
		name          string
		sigOverlap    []string
		broadOverlap  []string
		sig           string
		wantRelevance string
	}{
		{
			name:          "related - direct overlap",
			sigOverlap:    []string{"pkg/virt-handler/vm.go"},
			broadOverlap:  nil,
			sig:           "sig-compute",
			wantRelevance: "related",
		},
		{
			name:          "possibly-related - broad overlap only",
			sigOverlap:    nil,
			broadOverlap:  []string{"staging/src/kubevirt.io/api/core/types.go"},
			sig:           "sig-compute",
			wantRelevance: "possibly-related",
		},
		{
			name:          "unrelated - no overlap",
			sigOverlap:    nil,
			broadOverlap:  nil,
			sig:           "sig-compute",
			wantRelevance: "unrelated",
		},
		{
			name:          "unknown - no SIG",
			sigOverlap:    nil,
			broadOverlap:  nil,
			sig:           "",
			wantRelevance: "unknown",
		},
		{
			name:          "unknown - unmapped SIG",
			sigOverlap:    nil,
			broadOverlap:  nil,
			sig:           "sig-unknown-new",
			wantRelevance: "unknown",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			relevance, _ := ClassifyRelevance(tt.sigOverlap, tt.broadOverlap, tt.sig)
			if relevance != tt.wantRelevance {
				t.Errorf("ClassifyRelevance() = %q, want %q", relevance, tt.wantRelevance)
			}
		})
	}
}

func TestFetchPRChangedFiles(t *testing.T) {
	files := []ghPRFile{
		{Filename: "pkg/virt-handler/vm.go"},
		{Filename: "pkg/virt-handler/vm_test.go"},
		{Filename: "tests/compute/vmi_lifecycle_test.go"},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		expected := "/repos/kubevirt/kubevirt/pulls/100/files"
		if r.URL.Path != expected {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(files) //nolint:errcheck
	}))
	defer server.Close()

	origBase := githubAPIBaseURL
	defer func() { githubAPIBaseURL = origBase }()
	githubAPIBaseURL = server.URL

	result, err := FetchPRChangedFiles("kubevirt", "kubevirt", 100)
	if err != nil {
		t.Fatalf("FetchPRChangedFiles() error: %v", err)
	}
	if len(result) != 3 {
		t.Errorf("got %d files, want 3", len(result))
	}
}

func TestFetchPRChangedFilesPagination(t *testing.T) {
	page1 := make([]ghPRFile, 100)
	for i := range page1 {
		page1[i] = ghPRFile{Filename: fmt.Sprintf("file%d.go", i)}
	}
	page2 := []ghPRFile{{Filename: "last.go"}}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		page := r.URL.Query().Get("page")
		w.Header().Set("Content-Type", "application/json")
		if page == "2" {
			json.NewEncoder(w).Encode(page2) //nolint:errcheck
		} else {
			json.NewEncoder(w).Encode(page1) //nolint:errcheck
		}
	}))
	defer server.Close()

	origBase := githubAPIBaseURL
	defer func() { githubAPIBaseURL = origBase }()
	githubAPIBaseURL = server.URL

	result, err := FetchPRChangedFiles("kubevirt", "kubevirt", 100)
	if err != nil {
		t.Fatalf("FetchPRChangedFiles() error: %v", err)
	}
	if len(result) != 101 {
		t.Errorf("got %d files, want 101", len(result))
	}
}

func TestFetchPRChangedFilesNotFound(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	}))
	defer server.Close()

	origBase := githubAPIBaseURL
	defer func() { githubAPIBaseURL = origBase }()
	githubAPIBaseURL = server.URL

	_, err := FetchPRChangedFiles("kubevirt", "kubevirt", 999)
	if err == nil {
		t.Error("expected error for 404, got nil")
	}
}
