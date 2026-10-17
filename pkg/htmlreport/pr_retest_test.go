package htmlreport

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/kubevirt/ci-health/pkg/constants"
	"github.com/kubevirt/ci-health/pkg/types"
)

func TestBuildPRRetestReportData(t *testing.T) {
	t.Parallel()

	results := &types.Results{
		Source:        "kubevirt/kubevirt",
		DataDays:      7,
		EndDate:       "2026-09-17T12:00:00Z",
		OpenPRCount:   5,
		MergedPRCount: 3,
		PRRetestReport: []types.PRRetestSummary{
			{
				Number:      20,
				MergedAt:    "2026-09-16T00:00:00Z",
				RetestCount: 3,
				Failures: []types.PRJobFailure{
					{JobName: "pull-kubevirt-e2e-k8s-1.37-sig-compute", SIG: "compute", Reason: "e2e test failure", URL: "https://prow.example/compute"},
					{JobName: "pull-kubevirt-e2e-k8s-1.37-sig-compute-serial", SIG: "ci", Cause: types.FailureCauseExternal, Reason: "image pull", URL: "https://prow.example/ext"},
				},
			},
			{Number: 10, MergedAt: "2026-09-15T00:00:00Z", RetestCount: 1},
			{
				Number:      30,
				MergedAt:    "2026-09-14T00:00:00Z",
				RetestCount: 0,
				Failures: []types.PRJobFailure{
					{JobName: "pull-kubevirt-e2e-k8s-1.37-sig-storage", URL: "https://prow.example/clean-code"},
				},
			},
		},
		OpenPRRetestReport: []types.PRRetestSummary{
			{
				Number:      99,
				RetestCount: 2,
				Failures: []types.PRJobFailure{
					{JobName: "pull-kubevirt-e2e-k8s-1.37-sig-network", SIG: "network", Cause: types.FailureCauseFlake, URL: "https://prow.example/open-flake"},
				},
			},
		},
	}

	got := buildPRRetestReportData(results, time.Date(2026, 9, 17, 12, 0, 0, 0, time.UTC))
	if got.Merged.HeadlineLead != "2 / 3" {
		t.Fatalf("merged headline = %q, want 2 / 3", got.Merged.HeadlineLead)
	}
	if len(got.Merged.RetestedPRs) != 1 || got.Merged.RetestedPRs[0].Number != 20 {
		t.Fatalf("merged explained = %+v, want PR 20 only", got.Merged.RetestedPRs)
	}
	if len(got.Merged.UnexplainedPRs) != 1 || got.Merged.UnexplainedPRs[0].Number != 10 {
		t.Fatalf("merged unexplained = %+v, want PR 10", got.Merged.UnexplainedPRs)
	}
	if got.Merged.CauseOther != 1 || got.Merged.CauseExternal != 1 || got.Merged.CauseCI != 0 || got.Merged.CauseFlake != 0 {
		t.Fatalf("merged causes other=%d flake=%d ci=%d external=%d", got.Merged.CauseOther, got.Merged.CauseFlake, got.Merged.CauseCI, got.Merged.CauseExternal)
	}
	if got.Open.Title != "Open PRs" || got.Merged.Title != "Merged PRs" {
		t.Fatalf("titles open=%q merged=%q", got.Open.Title, got.Merged.Title)
	}
	if got.Open.Total != 5 || len(got.Open.RetestedPRs) != 1 || got.Open.RetestedPRs[0].Number != 99 {
		t.Fatalf("open total=%d retested=%d", got.Open.Total, len(got.Open.RetestedPRs))
	}
	if got.Open.CauseFlake != 1 || got.Open.CauseOther != 0 {
		t.Fatalf("open causes flake=%d other=%d, want 1 and 0", got.Open.CauseFlake, got.Open.CauseOther)
	}
	if got.Open.HeadlineLead != "1 / 5" {
		t.Fatalf("open headline = %q, want 1 / 5", got.Open.HeadlineLead)
	}
}

func TestWritePRRetestReport(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	results := &types.Results{
		Source:        "kubevirt/kubevirt",
		DataDays:      7,
		EndDate:       "2026-09-17T12:00:00Z",
		OpenPRCount:   4,
		MergedPRCount: 2,
		PRRetestReport: []types.PRRetestSummary{
			{
				Number:      19111,
				RetestCount: 2,
				Failures: []types.PRJobFailure{
					{JobName: "pull-kubevirt-e2e-k8s-1.37-sig-storage", SIG: "storage", Reason: "e2e test failure", URL: "https://prow.example/storage"},
					{JobName: "pull-kubevirt-e2e-k8s-1.37-sig-compute", SIG: "ci", Cause: types.FailureCauseCI, Reason: "kind cluster creation failure", URL: "https://prow.example/ci"},
					{JobName: "pull-kubevirt-e2e-k8s-1.34-sig-compute", SIG: "compute", Cause: types.FailureCauseFlake, Reason: "likely-flaky (20% success)", URL: "https://prow.example/flake"},
				},
			},
			{Number: 19185, RetestCount: 1},
		},
		OpenPRRetestReport: []types.PRRetestSummary{
			{
				Number:      19180,
				RetestCount: 1,
				Failures: []types.PRJobFailure{
					{JobName: "pull-kubevirt-e2e-k8s-1.37-sig-operator", SIG: "operator", Cause: types.FailureCauseCI, URL: "https://prow.example/open-ci"},
				},
			},
		},
	}

	if err := WritePRRetestReport(results, dir); err != nil {
		t.Fatalf("WritePRRetestReport: %v", err)
	}
	body, err := os.ReadFile(filepath.Join(dir, constants.PRRetestReportFileName))
	if err != nil {
		t.Fatalf("read report: %v", err)
	}
	html := string(body)
	for _, want := range []string{
		"PR retest report",
		"Open PRs",
		"Merged PRs",
		"#19111",
		"#19180",
		"#19185",
		"https://github.com/kubevirt/kubevirt/pull/19111",
		"https://github.com/kubevirt/kubevirt/pull/19180",
		"cause-other",
		"cause-ci",
		"cause-flake",
		`href="https://prow.example/flake"`,
		`href="https://prow.example/storage"`,
		`href="https://prow.example/open-ci"`,
		"pull-kubevirt-e2e-k8s-1.37-sig-storage",
		"2 / 2",
		"1 / 4",
		`id="open-prs"`,
		`id="merged-prs"`,
		"pull-kubevirt-e2e-k8s-1.37-sig-compute",
		"no failed required e2e job to classify",
	} {
		if !strings.Contains(html, want) {
			t.Errorf("report missing %q", want)
		}
	}
	for _, notWant := range []string{
		"Merged with no /retest",
		"https://prow.example/leftover",
		"commit-details",
		">HEAD<",
		"abcdef1",
		"failed, then passed",
		"cause-pill resolved",
		"SIG badges",
		"zero retests",
	} {
		if strings.Contains(html, notWant) {
			t.Errorf("report should not contain %q", notWant)
		}
	}
	if strings.Count(html, "flake 1") < 1 {
		t.Errorf("report missing merged flake cause pill")
	}
	// Cause pills belong on job cards, not duplicated in the PR title.
	titlePills := strings.Count(html, `class="pr-title"`)
	if titlePills < 2 {
		t.Errorf("expected PR titles, got %d", titlePills)
	}
}

func TestCauseDisplay(t *testing.T) {
	t.Parallel()
	if got := causeDisplay(types.FailureCauseFlake); got != "flake" {
		t.Errorf("flake = %q", got)
	}
	if got := causeDisplay(""); got != "other" {
		t.Errorf("empty = %q, want other", got)
	}
	if got := causeDisplay("code"); got != "other" {
		t.Errorf("legacy code = %q, want other", got)
	}
}

func TestPrGitHubURL(t *testing.T) {
	t.Parallel()
	got := prGitHubURL("kubevirt/kubevirt", 42)
	want := "https://github.com/kubevirt/kubevirt/pull/42"
	if got != want {
		t.Errorf("prGitHubURL = %q, want %q", got, want)
	}
}
