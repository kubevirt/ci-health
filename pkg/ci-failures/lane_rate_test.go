package cifailures

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestParseTestGridURL(t *testing.T) {
	tests := []struct {
		name      string
		url       string
		dashboard string
		tab       string
		wantErr   bool
	}{
		{
			name:      "standard URL",
			url:       "https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.36-sig-storage",
			dashboard: "kubevirt-periodics",
			tab:       "periodic-kubevirt-e2e-k8s-1.36-sig-storage",
		},
		{
			name:      "URL with width parameter",
			url:       "https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.36-sig-storage&width=20",
			dashboard: "kubevirt-periodics",
			tab:       "periodic-kubevirt-e2e-k8s-1.36-sig-storage",
		},
		{
			name:      "URL with trailing slash",
			url:       "https://testgrid.k8s.io/kubevirt-periodics/#periodic-kubevirt-e2e-k8s-1.36-sig-storage",
			dashboard: "kubevirt-periodics",
			tab:       "periodic-kubevirt-e2e-k8s-1.36-sig-storage",
		},
		{
			name:    "missing fragment",
			url:     "https://testgrid.k8s.io/kubevirt-periodics",
			wantErr: true,
		},
		{
			name:    "missing path",
			url:     "https://testgrid.k8s.io/#some-tab",
			wantErr: true,
		},
		{
			name:    "empty tab after stripping params",
			url:     "https://testgrid.k8s.io/kubevirt-periodics#&width=20",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dashboard, tab, err := ParseTestGridURL(tt.url)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if dashboard != tt.dashboard {
				t.Errorf("dashboard = %q, want %q", dashboard, tt.dashboard)
			}
			if tab != tt.tab {
				t.Errorf("tab = %q, want %q", tab, tt.tab)
			}
		})
	}
}

func TestExpandStatuses(t *testing.T) {
	tests := []struct {
		name     string
		statuses []testGridStatus
		want     []int
	}{
		{
			name:     "single entry",
			statuses: []testGridStatus{{Count: 5, Value: 1}},
			want:     []int{1, 1, 1, 1, 1},
		},
		{
			name: "mixed",
			statuses: []testGridStatus{
				{Count: 2, Value: 1},
				{Count: 1, Value: 12},
				{Count: 3, Value: 0},
			},
			want: []int{1, 1, 12, 0, 0, 0},
		},
		{
			name:     "empty",
			statuses: nil,
			want:     nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := expandStatuses(tt.statuses)
			if len(got) != len(tt.want) {
				t.Fatalf("len = %d, want %d", len(got), len(tt.want))
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("index %d = %d, want %d", i, got[i], tt.want[i])
				}
			}
		})
	}
}

func TestFilterColumnsInRange(t *testing.T) {
	now := time.Now().UTC()

	t.Run("all in range", func(t *testing.T) {
		timestamps := []int64{
			now.Add(-1 * time.Hour).UnixMilli(),
			now.Add(-2 * time.Hour).UnixMilli(),
			now.Add(-3 * time.Hour).UnixMilli(),
		}
		indices, _, _ := filterColumnsInRange(timestamps, 1)
		if len(indices) != 3 {
			t.Errorf("got %d indices, want 3", len(indices))
		}
	})

	t.Run("some out of range", func(t *testing.T) {
		timestamps := []int64{
			now.Add(-1 * time.Hour).UnixMilli(),
			now.Add(-48 * time.Hour).UnixMilli(),
			now.Add(-72 * time.Hour).UnixMilli(),
		}
		indices, _, _ := filterColumnsInRange(timestamps, 1)
		if len(indices) != 1 {
			t.Errorf("got %d indices, want 1", len(indices))
		}
	})

	t.Run("empty", func(t *testing.T) {
		indices, _, _ := filterColumnsInRange(nil, 7)
		if len(indices) != 0 {
			t.Errorf("got %d indices, want 0", len(indices))
		}
	})
}

func TestAnalyzeLaneRate(t *testing.T) {
	now := time.Now().UTC()

	mockResponse := testGridTableResponse{
		Timestamps: []int64{
			now.Add(-1 * time.Hour).UnixMilli(),
			now.Add(-2 * time.Hour).UnixMilli(),
			now.Add(-3 * time.Hour).UnixMilli(),
			now.Add(-4 * time.Hour).UnixMilli(),
			now.Add(-5 * time.Hour).UnixMilli(),
		},
		ColumnIDs: []string{"build-5", "build-4", "build-3", "build-2", "build-1"},
		Tests: []testGridTest{
			{
				Name: "suite.Overall",
				Statuses: []testGridStatus{
					{Count: 4, Value: 1},
					{Count: 1, Value: 12},
				},
				Messages: []string{"", "", "", "", "overall fail"},
			},
			{
				Name: "suite.[sig-storage] test-always-passes",
				Statuses: []testGridStatus{
					{Count: 5, Value: 1},
				},
				Messages: []string{"", "", "", "", ""},
			},
			{
				Name: "suite.[sig-storage] test-sometimes-fails",
				Statuses: []testGridStatus{
					{Count: 3, Value: 1},
					{Count: 2, Value: 12},
				},
				Messages: []string{"", "", "", "error msg 1", "error msg 2"},
			},
			{
				Name: "suite.[sig-storage] test-always-fails",
				Statuses: []testGridStatus{
					{Count: 5, Value: 12},
				},
				Messages: []string{"fail1", "fail2", "fail3", "fail4", "fail5"},
			},
		},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(mockResponse)
	}))
	defer server.Close()

	origURL := testgridBaseURL
	testgridBaseURL = server.URL
	defer func() { testgridBaseURL = origURL }()

	result, err := AnalyzeLaneRate(server.URL+"/test-dashboard#test-tab", 1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result.Dashboard != "test-dashboard" {
		t.Errorf("dashboard = %q, want %q", result.Dashboard, "test-dashboard")
	}
	if result.Tab != "test-tab" {
		t.Errorf("tab = %q, want %q", result.Tab, "test-tab")
	}
	if result.TotalBuilds != 5 {
		t.Errorf("total builds = %d, want 5", result.TotalBuilds)
	}

	// Should have 2 entries: sometimes-fails and always-fails
	// Overall and always-passes should be excluded
	if len(result.FailedTests) != 2 {
		t.Fatalf("got %d failed tests, want 2", len(result.FailedTests))
	}

	// Sorted by success rate ascending — always-fails (0%) first
	alwaysFails := result.FailedTests[0]
	if alwaysFails.Failed != 5 {
		t.Errorf("always-fails: failed = %d, want 5", alwaysFails.Failed)
	}
	if alwaysFails.Succeeded != 0 {
		t.Errorf("always-fails: succeeded = %d, want 0", alwaysFails.Succeeded)
	}
	if alwaysFails.Severity != "likely-flaky" {
		t.Errorf("always-fails: severity = %q, want %q", alwaysFails.Severity, "likely-flaky")
	}

	sometimesFails := result.FailedTests[1]
	if sometimesFails.Failed != 2 {
		t.Errorf("sometimes-fails: failed = %d, want 2", sometimesFails.Failed)
	}
	if sometimesFails.Succeeded != 3 {
		t.Errorf("sometimes-fails: succeeded = %d, want 3", sometimesFails.Succeeded)
	}
	if sometimesFails.SuccessRate != 60.0 {
		t.Errorf("sometimes-fails: success rate = %f, want 60.0", sometimesFails.SuccessRate)
	}

	if len(sometimesFails.RecentFailures) != 2 {
		t.Errorf("sometimes-fails: recent failures = %d, want 2", len(sometimesFails.RecentFailures))
	}
	if sometimesFails.RecentFailures[0].Message != "error msg 1" {
		t.Errorf("sometimes-fails: first failure message = %q, want %q", sometimesFails.RecentFailures[0].Message, "error msg 1")
	}
}
