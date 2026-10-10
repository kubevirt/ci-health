package stats

import (
	"errors"
	"testing"
	"time"

	cifailures "github.com/kubevirt/ci-health/pkg/ci-failures"
	"github.com/kubevirt/ci-health/pkg/sigretests"
	"github.com/kubevirt/ci-health/pkg/types"
)

func TestCauseFromCIAnalysis(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		result     *cifailures.JobBuildErrors
		err        error
		wantCause  string
		wantReason string
	}{
		{
			name:       "analysis error is ci",
			err:        errors.New("download failed"),
			wantCause:  types.FailureCauseCI,
			wantReason: "failed to analyze build; treating as ci",
		},
		{
			name:       "nil result is ci",
			wantCause:  types.FailureCauseCI,
			wantReason: "no build errors; treating as ci",
		},
		{
			name:       "empty errors is ci",
			result:     &cifailures.JobBuildErrors{},
			wantCause:  types.FailureCauseCI,
			wantReason: "no build errors; treating as ci",
		},
		{
			name: "external",
			result: &cifailures.JobBuildErrors{
				BuildErrors: []*cifailures.JobBuildError{{
					Category:       string(cifailures.CategoryExternal),
					CategoryReason: "container image pull failure",
				}},
			},
			wantCause:  types.FailureCauseExternal,
			wantReason: "container image pull failure",
		},
		{
			name: "pr-build is unlabeled",
			result: &cifailures.JobBuildErrors{
				BuildErrors: []*cifailures.JobBuildError{{
					Category:       string(cifailures.CategoryPRBuild),
					CategoryReason: "bazel build failure",
				}},
			},
			wantCause:  "",
			wantReason: "bazel build failure",
		},
		{
			name: "internal is ci",
			result: &cifailures.JobBuildErrors{
				BuildErrors: []*cifailures.JobBuildError{{
					Category:       string(cifailures.CategoryInternal),
					CategoryReason: "kind cluster creation failure",
				}},
			},
			wantCause:  types.FailureCauseCI,
			wantReason: "kind cluster creation failure",
		},
		{
			name: "needs-investigation is ci",
			result: &cifailures.JobBuildErrors{
				BuildErrors: []*cifailures.JobBuildError{{
					Category: string(cifailures.CategoryNeedsInvestigation),
				}},
			},
			wantCause:  types.FailureCauseCI,
			wantReason: string(cifailures.CategoryNeedsInvestigation),
		},
		{
			name: "kube-apiserver livez noise is ci",
			result: &cifailures.JobBuildErrors{
				BuildErrors: []*cifailures.JobBuildError{{
					Category:       string(cifailures.CategoryExternal),
					CategoryReason: "transient kube-apiserver body decode noise (from secondary snippet)",
				}},
			},
			wantCause:  types.FailureCauseCI,
			wantReason: "transient kube-apiserver body decode noise (from secondary snippet)",
		},
		{
			name: "podman teardown timeout is ci",
			result: &cifailures.JobBuildErrors{
				BuildErrors: []*cifailures.JobBuildError{{
					Category:       string(cifailures.CategoryExternal),
					CategoryReason: "podman container removal timeout",
				}},
			},
			wantCause:  types.FailureCauseCI,
			wantReason: "podman container removal timeout",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			gotCause, gotReason := CauseFromCIAnalysis(tc.result, tc.err)
			if gotCause != tc.wantCause {
				t.Errorf("cause = %q, want %q", gotCause, tc.wantCause)
			}
			if gotReason != tc.wantReason {
				t.Errorf("reason = %q, want %q", gotReason, tc.wantReason)
			}
		})
	}
}

func TestCountExternalCauses(t *testing.T) {
	t.Parallel()

	classified := map[string]ciFailureCause{
		"https://example/ext":  {cause: types.FailureCauseExternal, reason: "pull"},
		"https://example/code": {cause: "", reason: "bazel"},
		"https://example/ci":   {cause: types.FailureCauseCI, reason: "kind"},
	}
	if got := countExternalCauses(classified); got != 1 {
		t.Errorf("countExternalCauses = %d, want 1", got)
	}
}

func TestJobFailuresForPR(t *testing.T) {
	t.Parallel()

	ciURL := "https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/1/pull-kubevirt-e2e-k8s-1.37-sig-compute/111"
	codeURL := "https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/1/pull-kubevirt-e2e-k8s-1.37-sig-storage/222"
	details := []sigretests.FailedJobDetail{
		{JobName: "pull-kubevirt-e2e-k8s-1.37-sig-compute", URL: ciURL, SIG: "ci"},
		{JobName: "pull-kubevirt-e2e-k8s-1.37-sig-storage", URL: codeURL, SIG: "storage"},
	}
	ciCauses := map[string]ciFailureCause{
		ciURL: {cause: types.FailureCauseExternal, reason: "container image pull failure"},
	}

	got := jobFailuresForPR(1, details, ciCauses, nil)
	if len(got) != 2 {
		t.Fatalf("got %d failures, want 2", len(got))
	}
	if got[0].JobName != "pull-kubevirt-e2e-k8s-1.37-sig-compute" {
		t.Fatalf("sorted first job = %q", got[0].JobName)
	}
	if got[0].Cause != types.FailureCauseExternal || got[0].SIG != "ci" {
		t.Errorf("ci job cause=%q sig=%q, want external/ci", got[0].Cause, got[0].SIG)
	}
	if got[1].Cause != "" || got[1].SIG != "storage" {
		t.Errorf("storage job cause=%q sig=%q, want unlabeled/storage", got[1].Cause, got[1].SIG)
	}
	if got[1].Reason != "e2e test failure" {
		t.Errorf("storage reason = %q", got[1].Reason)
	}
}

func TestSourceOrgRepo(t *testing.T) {
	t.Parallel()
	tests := []struct {
		source   string
		wantOrg  string
		wantRepo string
	}{
		{source: "kubevirt/kubevirt", wantOrg: "kubevirt", wantRepo: "kubevirt"},
		{source: "foo/bar", wantOrg: "foo", wantRepo: "bar"},
		{source: "invalid", wantOrg: "kubevirt", wantRepo: "kubevirt"},
		{source: "", wantOrg: "kubevirt", wantRepo: "kubevirt"},
	}
	for _, tc := range tests {
		h := &Handler{source: tc.source}
		org, repo := h.sourceOrgRepo()
		if org != tc.wantOrg || repo != tc.wantRepo {
			t.Errorf("source %q: got %s/%s, want %s/%s", tc.source, org, repo, tc.wantOrg, tc.wantRepo)
		}
	}
}

func TestCollectOpenPRRetestReportNilChatops(t *testing.T) {
	t.Parallel()
	h := &Handler{}
	got, n, err := h.collectOpenPRRetestReport(time.Time{}, nil)
	if err != nil {
		t.Fatalf("err = %v", err)
	}
	if n != 0 || len(got) != 0 {
		t.Fatalf("got %d summaries, considered %d, want none", len(got), n)
	}
}

func TestBuildPRRetestReport(t *testing.T) {
	t.Parallel()

	prs := []types.PR{
		{Number: 10, MergedAt: "2026-09-17T00:00:00Z"},
		{Number: 20, MergedAt: "2026-09-16T00:00:00Z"},
		{Number: 30, MergedAt: "2026-09-15T00:00:00Z"},
	}
	retestCounts := map[int]int{20: 3, 10: 3}
	failuresByPR := map[int][]types.PRJobFailure{
		20: {{JobName: "pull-kubevirt-e2e-k8s-1.37-sig-compute", Cause: types.FailureCauseFlake, SIG: "compute"}},
	}

	got := BuildPRRetestReport(prs, retestCounts, failuresByPR)
	if len(got) != 2 {
		t.Fatalf("got %d summaries, want 2 (zero-retest PR 30 omitted)", len(got))
	}
	// Highest retest count first; ties broken by higher PR number.
	if got[0].Number != 20 || got[0].RetestCount != 3 {
		t.Errorf("got[0] = PR %d retests %d, want 20/3", got[0].Number, got[0].RetestCount)
	}
	if got[1].Number != 10 || got[1].RetestCount != 3 {
		t.Errorf("got[1] = PR %d retests %d, want 10/3", got[1].Number, got[1].RetestCount)
	}
	if len(got[0].Failures) != 1 {
		t.Errorf("PR 20 failures = %d, want 1", len(got[0].Failures))
	}
}

func TestBuildPRRetestReportOmitsZeroCount(t *testing.T) {
	t.Parallel()

	prs := []types.PR{{Number: 17974}, {Number: 42}}
	got := BuildPRRetestReport(prs, map[int]int{17974: 1, 42: 0}, nil)
	if len(got) != 1 || got[0].Number != 17974 || got[0].RetestCount != 1 {
		t.Fatalf("got %+v, want only PR 17974 with 1 /retest", got)
	}
}

func TestClassifyCIFailuresUsesAnalyzeBuild(t *testing.T) {
	orig := analyzeBuildFn
	t.Cleanup(func() { analyzeBuildFn = orig })

	analyzeBuildFn = func(url string) (*cifailures.JobBuildErrors, error) {
		switch url {
		case "ext":
			return &cifailures.JobBuildErrors{
				BuildErrors: []*cifailures.JobBuildError{{
					Category:       string(cifailures.CategoryExternal),
					CategoryReason: "dns",
				}},
			}, nil
		case "pr-build":
			return &cifailures.JobBuildErrors{
				BuildErrors: []*cifailures.JobBuildError{{
					Category:       string(cifailures.CategoryPRBuild),
					CategoryReason: "bazel",
				}},
			}, nil
		default:
			return nil, errors.New("boom")
		}
	}

	got := classifyCIFailures([]string{"ext", "pr-build", "bad"})
	if got["ext"].cause != types.FailureCauseExternal {
		t.Errorf("ext cause = %q, want external", got["ext"].cause)
	}
	if got["pr-build"].cause != "" {
		t.Errorf("pr-build cause = %q, want unlabeled", got["pr-build"].cause)
	}
	if got["bad"].cause != types.FailureCauseCI {
		t.Errorf("bad cause = %q, want ci", got["bad"].cause)
	}
	if countExternalCauses(got) != 1 {
		t.Errorf("external count = %d, want 1", countExternalCauses(got))
	}
}

func TestJobFailuresForPRClassifiesE2EFlake(t *testing.T) {
	origTests, origFF := fetchFailedTestsFn, loadFlakefinderFn
	t.Cleanup(func() {
		fetchFailedTestsFn = origTests
		loadFlakefinderFn = origFF
	})

	testName := `[sig-compute] Infrastructure tls configuration [It] should enforce TLS configuration on virt-template components [sig-compute, Serial]`
	fetchFailedTestsFn = func(url string) ([]cifailures.FailedTest, error) {
		return []cifailures.FailedTest{{Name: testName}}, nil
	}
	loadFlakefinderFn = func(days int) (*cifailures.FlakefinderReport, error) {
		if days != 7 {
			t.Errorf("flakefinder days = %d, want 7", days)
		}
		return &cifailures.FlakefinderReport{
			Tests: []string{testName},
			Data: map[string]map[string]*cifailures.TestDetails{
				testName: {
					"lane": {Succeeded: 10, Failed: 40},
				},
			},
		}, nil
	}

	clf := newE2EClassifier(7)
	details := []sigretests.FailedJobDetail{{
		JobName: "pull-kubevirt-e2e-k8s-1.34-sig-compute",
		URL:     "https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19055/pull-kubevirt-e2e-k8s-1.34-sig-compute/1",
		SIG:     "compute",
	}}
	got := jobFailuresForPR(19055, details, nil, clf)
	if len(got) != 1 {
		t.Fatalf("got %d failures, want 1", len(got))
	}
	if got[0].Cause != types.FailureCauseFlake {
		t.Fatalf("cause = %q (%s), want flake", got[0].Cause, got[0].Reason)
	}
}

func TestJobFailuresForPRLeavesStableE2EUnlabeled(t *testing.T) {
	origTests, origFF := fetchFailedTestsFn, loadFlakefinderFn
	t.Cleanup(func() {
		fetchFailedTestsFn = origTests
		loadFlakefinderFn = origFF
	})

	testName := `[sig-compute] virt-template should render [It]`
	fetchFailedTestsFn = func(url string) ([]cifailures.FailedTest, error) {
		return []cifailures.FailedTest{{Name: testName}}, nil
	}
	loadFlakefinderFn = func(days int) (*cifailures.FlakefinderReport, error) {
		return &cifailures.FlakefinderReport{
			Tests: []string{testName},
			Data: map[string]map[string]*cifailures.TestDetails{
				testName: {
					"lane": {Succeeded: 100, Failed: 0},
				},
			},
		}, nil
	}

	clf := newE2EClassifier(7)
	details := []sigretests.FailedJobDetail{{
		JobName: "pull-kubevirt-e2e-k8s-1.34-sig-compute",
		URL:     "https://prow.example/job",
		SIG:     "compute",
	}}
	got := jobFailuresForPR(42, details, nil, clf)
	if got[0].Cause != "" {
		t.Fatalf("cause = %q (%s), want unlabeled", got[0].Cause, got[0].Reason)
	}
}

func TestJobFailuresForPRClassifiesEmptyJunitAsCI(t *testing.T) {
	origTests, origFF := fetchFailedTestsFn, loadFlakefinderFn
	t.Cleanup(func() {
		fetchFailedTestsFn = origTests
		loadFlakefinderFn = origFF
	})

	fetchFailedTestsFn = func(url string) ([]cifailures.FailedTest, error) {
		return nil, nil
	}
	loadFlakefinderFn = func(days int) (*cifailures.FlakefinderReport, error) {
		return &cifailures.FlakefinderReport{}, nil
	}

	clf := newE2EClassifier(7)
	details := []sigretests.FailedJobDetail{{
		JobName: "pull-kubevirt-e2e-k8s-1.34-sig-monitoring-1.9",
		URL:     "https://prow.example/monitoring",
		SIG:     "monitoring",
	}}
	got := jobFailuresForPR(19140, details, nil, clf)
	if got[0].Cause != types.FailureCauseCI {
		t.Fatalf("cause = %q (%s), want ci", got[0].Cause, got[0].Reason)
	}
	if got[0].Reason != "job failed with junit but no named test failures" {
		t.Fatalf("reason = %q", got[0].Reason)
	}
}

func TestJobFailuresForPRClassifiesJunitFetchErrorAsCI(t *testing.T) {
	origTests, origFF := fetchFailedTestsFn, loadFlakefinderFn
	t.Cleanup(func() {
		fetchFailedTestsFn = origTests
		loadFlakefinderFn = origFF
	})

	fetchFailedTestsFn = func(url string) ([]cifailures.FailedTest, error) {
		return nil, errors.New("gcs 500")
	}
	loadFlakefinderFn = func(days int) (*cifailures.FlakefinderReport, error) {
		return &cifailures.FlakefinderReport{}, nil
	}

	clf := newE2EClassifier(7)
	details := []sigretests.FailedJobDetail{{
		JobName: "pull-kubevirt-e2e-k8s-1.34-sig-compute",
		URL:     "https://prow.example/compute",
		SIG:     "compute",
	}}
	got := jobFailuresForPR(1, details, nil, clf)
	if got[0].Cause != types.FailureCauseCI {
		t.Fatalf("cause = %q (%s), want ci", got[0].Cause, got[0].Reason)
	}
	if got[0].Reason != "failed to fetch junit; treating as ci" {
		t.Fatalf("reason = %q", got[0].Reason)
	}
}
