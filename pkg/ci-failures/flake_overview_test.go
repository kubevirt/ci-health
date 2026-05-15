package cifailures

import (
	"regexp"
	"testing"
	"time"
)

func TestFlakefinder24hReportURL(t *testing.T) {
	date := time.Date(2026, 5, 12, 0, 0, 0, 0, time.UTC)
	got := flakefinder24hReportURL("kubevirt", "kubevirt", date)
	want := "https://storage.googleapis.com/kubevirt-prow/reports/flakefinder/kubevirt/kubevirt/flakefinder-2026-05-12-024h.json"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestTestgridURLForJob(t *testing.T) {
	tests := []struct {
		job, want string
	}{
		{"pull-kubevirt-e2e-k8s-1.35-sig-compute", "https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.35-sig-compute"},
		{"periodic-kubevirt-e2e-k8s-1.36-sig-storage", "https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.36-sig-storage"},
		{"unknown-job", ""},
	}
	for _, tc := range tests {
		got := testgridURLForJob(tc.job)
		if got != tc.want {
			t.Errorf("testgridURLForJob(%q) = %q, want %q", tc.job, got, tc.want)
		}
	}
}

func TestAggregateLaneFailures(t *testing.T) {
	reports := []*FlakefinderReport{
		{
			Tests: []string{"test-a", "test-b"},
			Data: map[string]map[string]*TestDetails{
				"test-a": {
					"periodic-kubevirt-e2e-k8s-1.36-sig-compute":      {Failed: 3},
					"pull-kubevirt-e2e-k8s-1.35-sig-compute":          {Failed: 1},
				},
				"test-b": {
					"periodic-kubevirt-e2e-k8s-1.36-sig-compute":      {Failed: 2},
					"periodic-kubevirt-e2e-k8s-1.36-sig-compute-root": {Failed: 5},
				},
			},
		},
		{
			Tests: []string{"test-a"},
			Data: map[string]map[string]*TestDetails{
				"test-a": {
					"periodic-kubevirt-e2e-k8s-1.36-sig-compute": {Failed: 1},
				},
			},
		},
	}

	filterRegex := regexp.MustCompile(`.*-root$`)

	lanes, total := aggregateLaneFailures(reports, filterRegex)

	if total != 7 {
		t.Errorf("total = %d, want 7", total)
	}
	if len(lanes) != 2 {
		t.Fatalf("len(lanes) = %d, want 2", len(lanes))
	}
	if lanes[0].name != "periodic-kubevirt-e2e-k8s-1.36-sig-compute" || lanes[0].failures != 6 {
		t.Errorf("lanes[0] = {%s, %d}, want {periodic-kubevirt-e2e-k8s-1.36-sig-compute, 6}", lanes[0].name, lanes[0].failures)
	}
	if lanes[1].name != "pull-kubevirt-e2e-k8s-1.35-sig-compute" || lanes[1].failures != 1 {
		t.Errorf("lanes[1] = {%s, %d}, want {pull-kubevirt-e2e-k8s-1.35-sig-compute, 1}", lanes[1].name, lanes[1].failures)
	}
}

func TestAggregateLaneFailures_NoFilter(t *testing.T) {
	reports := []*FlakefinderReport{
		{
			Tests: []string{"test-a"},
			Data: map[string]map[string]*TestDetails{
				"test-a": {
					"periodic-job": {Failed: 3},
					"pull-job":     {Failed: 2},
				},
			},
		},
	}

	lanes, total := aggregateLaneFailures(reports, nil)

	if total != 5 {
		t.Errorf("total = %d, want 5", total)
	}
	if len(lanes) != 2 {
		t.Fatalf("len(lanes) = %d, want 2", len(lanes))
	}
}

func TestAggregateLaneFailures_SkipsZeroFailures(t *testing.T) {
	reports := []*FlakefinderReport{
		{
			Tests: []string{"test-a"},
			Data: map[string]map[string]*TestDetails{
				"test-a": {
					"periodic-job": {Failed: 0},
				},
			},
		},
	}

	lanes, total := aggregateLaneFailures(reports, nil)

	if total != 0 {
		t.Errorf("total = %d, want 0", total)
	}
	if len(lanes) != 0 {
		t.Errorf("len(lanes) = %d, want 0", len(lanes))
	}
}
