package cifailures

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestStripTimestamp(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"13:34:47: [FAIL] some test", "[FAIL] some test"},
		{"  13:34:47:  [FAIL] some test", "[FAIL] some test"},
		{"[FAIL] no timestamp", "[FAIL] no timestamp"},
		{"", ""},
	}
	for _, tt := range tests {
		got := stripTimestamp(tt.input)
		if got != tt.expected {
			t.Errorf("stripTimestamp(%q) = %q, want %q", tt.input, got, tt.expected)
		}
	}
}

func TestNormalizeTestName(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "strips [It] marker",
			input:    "[sig-compute] Live Migrations [It] with a live-migrate eviction strategy set",
			expected: "[sig-compute] Live Migrations with a live-migrate eviction strategy set",
		},
		{
			name:     "strips trailing labels",
			input:    "[sig-compute] Live Migrations with a strategy set [sig-compute, sig-compute-migrations, Serial]",
			expected: "[sig-compute] Live Migrations with a strategy set",
		},
		{
			name:     "strips both [It] and trailing labels",
			input:    "[sig-compute] Live Migrations [It] with a strategy set [sig-compute, sig-compute-migrations]",
			expected: "[sig-compute] Live Migrations with a strategy set",
		},
		{
			name:     "no-op for clean name",
			input:    "[sig-compute] Live Migrations with a strategy set",
			expected: "[sig-compute] Live Migrations with a strategy set",
		},
		{
			name:     "does not strip single-element brackets",
			input:    "[sig-compute] Live Migrations [rfe_id:393] with a strategy set",
			expected: "[sig-compute] Live Migrations [rfe_id:393] with a strategy set",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := normalizeTestName(tt.input)
			if got != tt.expected {
				t.Errorf("normalizeTestName() = %q, want %q", got, tt.expected)
			}
		})
	}
}

func TestFindMatchingTest(t *testing.T) {
	report := &FlakefinderReport{
		Tests: []string{
			"[sig-compute] Live Migrations with a live-migrate eviction strategy set",
			"[sig-network] VMI connectivity should allow regular network traffic",
		},
	}

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "exact match",
			input:    "[sig-compute] Live Migrations with a live-migrate eviction strategy set",
			expected: "[sig-compute] Live Migrations with a live-migrate eviction strategy set",
		},
		{
			name:     "match after normalization (strip [It])",
			input:    "[sig-compute] Live Migrations [It] with a live-migrate eviction strategy set",
			expected: "[sig-compute] Live Migrations with a live-migrate eviction strategy set",
		},
		{
			name:     "match after normalization (strip trailing labels)",
			input:    "[sig-compute] Live Migrations with a live-migrate eviction strategy set [sig-compute, Serial]",
			expected: "[sig-compute] Live Migrations with a live-migrate eviction strategy set",
		},
		{
			name:     "no match",
			input:    "[sig-storage] something completely different",
			expected: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findMatchingTest(tt.input, report)
			if got != tt.expected {
				t.Errorf("findMatchingTest() = %q, want %q", got, tt.expected)
			}
		})
	}
}

func TestComputeTestRate(t *testing.T) {
	report := &FlakefinderReport{
		Tests: []string{
			"[sig-compute] Live Migrations with a strategy set",
		},
		Data: map[string]map[string]*TestDetails{
			"[sig-compute] Live Migrations with a strategy set": {
				"pull-kubevirt-e2e-k8s-1.35-sig-compute":   {Succeeded: 80, Failed: 10, Skipped: 5},
				"pull-kubevirt-e2e-k8s-1.36-sig-compute":   {Succeeded: 70, Failed: 20, Skipped: 3},
				"pull-kubevirt-e2e-kind-1.35-sig-compute":  {Succeeded: 50, Failed: 5, Skipped: 1},
			},
		},
	}

	entry := computeTestRate("[sig-compute] Live Migrations with a strategy set", report)

	if entry.TotalSucceeded != 200 {
		t.Errorf("TotalSucceeded = %d, want 200", entry.TotalSucceeded)
	}
	if entry.TotalFailed != 35 {
		t.Errorf("TotalFailed = %d, want 35", entry.TotalFailed)
	}
	if entry.TotalSkipped != 9 {
		t.Errorf("TotalSkipped = %d, want 9", entry.TotalSkipped)
	}

	expectedRate := float64(200) / float64(235) * 100.0
	if entry.SuccessRate != expectedRate {
		t.Errorf("SuccessRate = %f, want %f", entry.SuccessRate, expectedRate)
	}
	if entry.Severity != "inconclusive" {
		t.Errorf("Severity = %q, want %q", entry.Severity, "inconclusive")
	}

	// Two k8s versions: 1.35 (two lanes aggregated) and 1.36 (one lane)
	if len(entry.K8sVersions) != 2 {
		t.Fatalf("len(K8sVersions) = %d, want 2", len(entry.K8sVersions))
	}
	// Sorted by success rate ascending: 1.36 (77.8%) before 1.35 (89.7%)
	if entry.K8sVersions[0].Version != "1.36" {
		t.Errorf("K8sVersions[0].Version = %q, want %q", entry.K8sVersions[0].Version, "1.36")
	}
	if entry.K8sVersions[0].Succeeded != 70 || entry.K8sVersions[0].Failed != 20 {
		t.Errorf("K8sVersions[0] = %d/%d, want 70/20", entry.K8sVersions[0].Succeeded, entry.K8sVersions[0].Failed)
	}
	if entry.K8sVersions[0].Severity != "likely-flaky" {
		t.Errorf("K8sVersions[0].Severity = %q, want %q", entry.K8sVersions[0].Severity, "likely-flaky")
	}
	if len(entry.K8sVersions[0].Lanes) != 1 {
		t.Fatalf("K8sVersions[0] lane count = %d, want 1", len(entry.K8sVersions[0].Lanes))
	}

	if entry.K8sVersions[1].Version != "1.35" {
		t.Errorf("K8sVersions[1].Version = %q, want %q", entry.K8sVersions[1].Version, "1.35")
	}
	// 1.35 aggregates k8s-1.35 (80/10) + kind-1.35 (50/5) = 130/15
	if entry.K8sVersions[1].Succeeded != 130 || entry.K8sVersions[1].Failed != 15 {
		t.Errorf("K8sVersions[1] = %d/%d, want 130/15", entry.K8sVersions[1].Succeeded, entry.K8sVersions[1].Failed)
	}
	if len(entry.K8sVersions[1].Lanes) != 2 {
		t.Fatalf("K8sVersions[1] lane count = %d, want 2", len(entry.K8sVersions[1].Lanes))
	}
}

func TestExtractK8sVersion(t *testing.T) {
	tests := []struct {
		lane    string
		version string
	}{
		{"pull-kubevirt-e2e-k8s-1.35-sig-compute", "1.35"},
		{"periodic-kubevirt-e2e-k8s-1.36-sig-network", "1.36"},
		{"periodic-kubevirt-e2e-kind-1.34-sev", "1.34"},
		{"pull-kubevirt-e2e-kind-1.35-sig-compute-arm64", "1.35"},
		{"periodic-kubevirt-e2e-test-S390X", "unknown"},
		{"some-job-without-version", "unknown"},
	}
	for _, tt := range tests {
		t.Run(tt.lane, func(t *testing.T) {
			got := extractK8sVersion(tt.lane)
			if got != tt.version {
				t.Errorf("extractK8sVersion(%q) = %q, want %q", tt.lane, got, tt.version)
			}
		})
	}
}

func TestComputeTestRateNoMatch(t *testing.T) {
	report := &FlakefinderReport{
		Tests: []string{"[sig-compute] something else"},
		Data:  map[string]map[string]*TestDetails{},
	}

	entry := computeTestRate("[sig-storage] unknown test", report)

	if entry.MatchedName != "" {
		t.Errorf("MatchedName = %q, want empty", entry.MatchedName)
	}
	if entry.Severity != "unknown" {
		t.Errorf("Severity = %q, want %q", entry.Severity, "unknown")
	}
}

func TestComputeTestRateSeverityBuckets(t *testing.T) {
	tests := []struct {
		name             string
		succeeded        int
		failed           int
		expectedSeverity string
	}{
		{"high success rate", 98, 2, "likely-pr-related"},
		{"exactly 95%", 95, 5, "likely-pr-related"},
		{"moderate success rate", 85, 15, "inconclusive"},
		{"exactly 80%", 80, 20, "inconclusive"},
		{"low success rate", 60, 40, "likely-flaky"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testName := "[sig-compute] test"
			report := &FlakefinderReport{
				Tests: []string{testName},
				Data: map[string]map[string]*TestDetails{
					testName: {
						"job-a": {Succeeded: tt.succeeded, Failed: tt.failed},
					},
				},
			}

			entry := computeTestRate(testName, report)
			if entry.Severity != tt.expectedSeverity {
				t.Errorf("Severity = %q, want %q (rate=%.1f%%)", entry.Severity, tt.expectedSeverity, entry.SuccessRate)
			}
		})
	}
}

func TestExtractFailedTestNames(t *testing.T) {
	jobBuildErrors := &JobBuildErrors{
		BuildErrors: []*JobBuildError{
			{
				BuildLogErrorSnippets: []*BuildLogErrorSnippet{
					{ErrorText: "13:34:47: [FAIL] [sig-compute] Live Migrations [It] with a strategy set [sig-compute, Serial]"},
					{ErrorText: "some other error without FAIL marker"},
					{ErrorText: "[FAIL] [sig-network] VMI connectivity [It] should work [sig-network]"},
					{ErrorText: "13:34:47: [FAIL] [sig-compute] Live Migrations [It] with a strategy set [sig-compute, Serial]"},
				},
			},
		},
	}

	result := extractFailedTestNames(jobBuildErrors)

	if len(result) != 2 {
		t.Fatalf("got %d test names, want 2", len(result))
	}
	if result[0] != "[sig-compute] Live Migrations [It] with a strategy set [sig-compute, Serial]" {
		t.Errorf("result[0] = %q, unexpected", result[0])
	}
	if result[1] != "[sig-network] VMI connectivity [It] should work [sig-network]" {
		t.Errorf("result[1] = %q, unexpected", result[1])
	}
}

func TestMergeReports(t *testing.T) {
	r1 := &FlakefinderReport{
		StartOfReport: "2026-04-26",
		EndOfReport:   "2026-05-03",
		Headers:       []string{"job-a", "job-b"},
		Tests:         []string{"[sig-compute] test one", "[sig-compute] test two"},
		Data: map[string]map[string]*TestDetails{
			"[sig-compute] test one": {
				"job-a": {Succeeded: 50, Failed: 10, Skipped: 2},
			},
			"[sig-compute] test two": {
				"job-b": {Succeeded: 30, Failed: 5, Skipped: 1},
			},
		},
	}
	r2 := &FlakefinderReport{
		StartOfReport: "2026-05-03",
		EndOfReport:   "2026-05-10",
		Headers:       []string{"job-a", "job-c"},
		Tests:         []string{"[sig-compute] test one", "[sig-network] test three"},
		Data: map[string]map[string]*TestDetails{
			"[sig-compute] test one": {
				"job-a": {Succeeded: 40, Failed: 8, Skipped: 3},
			},
			"[sig-network] test three": {
				"job-c": {Succeeded: 20, Failed: 2, Skipped: 0},
			},
		},
	}

	merged := mergeReports([]*FlakefinderReport{r1, r2})

	if merged.StartOfReport != "2026-04-26" {
		t.Errorf("StartOfReport = %q, want %q", merged.StartOfReport, "2026-04-26")
	}
	if merged.EndOfReport != "2026-05-10" {
		t.Errorf("EndOfReport = %q, want %q", merged.EndOfReport, "2026-05-10")
	}
	if len(merged.Tests) != 3 {
		t.Fatalf("len(Tests) = %d, want 3", len(merged.Tests))
	}

	d := merged.Data["[sig-compute] test one"]["job-a"]
	if d == nil {
		t.Fatal("missing data for test one / job-a")
	}
	if d.Succeeded != 90 {
		t.Errorf("test one job-a Succeeded = %d, want 90", d.Succeeded)
	}
	if d.Failed != 18 {
		t.Errorf("test one job-a Failed = %d, want 18", d.Failed)
	}
	if d.Skipped != 5 {
		t.Errorf("test one job-a Skipped = %d, want 5", d.Skipped)
	}

	if _, ok := merged.Data["[sig-compute] test two"]["job-b"]; !ok {
		t.Error("missing data for test two / job-b")
	}
	if _, ok := merged.Data["[sig-network] test three"]["job-c"]; !ok {
		t.Error("missing data for test three / job-c")
	}
}

func TestMergeReportsSingle(t *testing.T) {
	r := &FlakefinderReport{
		StartOfReport: "2026-05-03",
		EndOfReport:   "2026-05-10",
		Tests:         []string{"test"},
		Data:          map[string]map[string]*TestDetails{},
	}
	merged := mergeReports([]*FlakefinderReport{r})
	if merged != r {
		t.Error("single report merge should return the same pointer")
	}
}

func TestFetchFlakefinderReports(t *testing.T) {
	reportsByDate := map[string]*FlakefinderReport{}

	now := time.Now().UTC()
	week1Date := now.AddDate(0, 0, -1).Format(time.DateOnly)
	reportsByDate[week1Date] = &FlakefinderReport{
		StartOfReport: now.AddDate(0, 0, -8).Format(time.DateOnly),
		EndOfReport:   week1Date,
		Tests:         []string{"t1"},
		Data:          map[string]map[string]*TestDetails{},
	}
	week2Date := now.AddDate(0, 0, -8).Format(time.DateOnly)
	reportsByDate[week2Date] = &FlakefinderReport{
		StartOfReport: now.AddDate(0, 0, -15).Format(time.DateOnly),
		EndOfReport:   week2Date,
		Tests:         []string{"t2"},
		Data:          map[string]map[string]*TestDetails{},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for date, report := range reportsByDate {
			expected := fmt.Sprintf("/flakefinder-%s-168h.json", date)
			if r.URL.Path == expected {
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(report)
				return
			}
		}
		http.NotFound(w, r)
	}))
	defer server.Close()

	origBase := flakefinderBaseURL
	defer func() { flakefinderBaseURL = origBase }()
	flakefinderBaseURL = server.URL

	result, err := fetchFlakefinderReports(14)
	if err != nil {
		t.Fatalf("fetchFlakefinderReports(14) error: %v", err)
	}
	if len(result) != 2 {
		t.Errorf("got %d reports, want 2", len(result))
	}
}

func TestFetchFlakefinderReport(t *testing.T) {
	report := FlakefinderReport{
		StartOfReport: "2026-05-03",
		EndOfReport:   "2026-05-10",
		Headers:       []string{"job-a", "job-b"},
		Tests:         []string{"[sig-compute] test one"},
		Data: map[string]map[string]*TestDetails{
			"[sig-compute] test one": {
				"job-a": {Succeeded: 90, Failed: 10, Skipped: 0, Severity: "yellow"},
			},
		},
	}

	reportJSON, err := json.Marshal(report)
	if err != nil {
		t.Fatal(err)
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write(reportJSON)
	}))
	defer server.Close()

	origBase := flakefinderBaseURL
	defer func() { flakefinderBaseURL = origBase }()
	flakefinderBaseURL = server.URL

	date := time.Date(2026, 5, 10, 0, 0, 0, 0, time.UTC)
	result, err := fetchFlakefinderReport(date)
	if err != nil {
		t.Fatalf("fetchFlakefinderReport() error: %v", err)
	}

	if result.StartOfReport != "2026-05-03" {
		t.Errorf("StartOfReport = %q, want %q", result.StartOfReport, "2026-05-03")
	}
	if len(result.Tests) != 1 {
		t.Errorf("len(Tests) = %d, want 1", len(result.Tests))
	}
}
