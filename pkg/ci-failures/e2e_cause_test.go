package cifailures

import (
	"strings"
	"testing"

	"github.com/kubevirt/ci-health/pkg/types"
)

const tlsVirtTemplateTest = `[sig-compute] Infrastructure tls configuration [It] should enforce TLS configuration on virt-template components [sig-compute, Serial]`

func TestClassifyE2ECause(t *testing.T) {
	t.Parallel()

	tlsTest := FailedTest{Name: tlsVirtTemplateTest}
	stableReport := reportWithRate(tlsVirtTemplateTest, 100, 0)
	flakyReport := reportWithRate(tlsVirtTemplateTest, 10, 40)

	tests := []struct {
		name   string
		tests  []FailedTest
		report *FlakefinderReport
		want   string
	}{
		{
			name:   "unknown rate is unlabeled",
			tests:  []FailedTest{tlsTest},
			report: &FlakefinderReport{},
			want:   "",
		},
		{
			name:   "stable e2e is unlabeled",
			tests:  []FailedTest{tlsTest},
			report: stableReport,
			want:   "",
		},
		{
			name:   "likely-flaky is flake",
			tests:  []FailedTest{tlsTest},
			report: flakyReport,
			want:   types.FailureCauseFlake,
		},
		{
			name: "mixed flaky and stable stays unlabeled",
			tests: []FailedTest{
				tlsTest,
				{Name: `[sig-network] bridge binding should work [It]`},
			},
			report: flakyReport,
			want:   "",
		},
		{
			name: "no named tests is ci",
			want: types.FailureCauseCI,
		},
		{
			name: "mass unlabeled failures are ci",
			tests: []FailedTest{
				{Name: `[sig-operator] test one [It]`},
				{Name: `[sig-operator] test two [It]`},
				{Name: `[sig-operator] test three [It]`},
				{Name: `[sig-operator] test four [It]`},
				{Name: `[sig-operator] test five [It]`},
			},
			report: &FlakefinderReport{},
			want:   types.FailureCauseCI,
		},
		{
			name: "four unlabeled failures stay unlabeled",
			tests: []FailedTest{
				{Name: `[sig-operator] test one [It]`},
				{Name: `[sig-operator] test two [It]`},
				{Name: `[sig-operator] test three [It]`},
				{Name: `[sig-operator] test four [It]`},
			},
			report: &FlakefinderReport{},
			want:   "",
		},
		{
			name: "mass failures that are all flaky are flake",
			tests: []FailedTest{
				{Name: `[sig-operator] test one [It]`},
				{Name: `[sig-operator] test two [It]`},
				{Name: `[sig-operator] test three [It]`},
				{Name: `[sig-operator] test four [It]`},
				{Name: `[sig-operator] test five [It]`},
			},
			report: massFlakyReport(),
			want:   types.FailureCauseFlake,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got, reason := ClassifyE2ECause(tc.tests, tc.report)
			if got != tc.want {
				t.Fatalf("cause = %q (%s), want %q", got, reason, tc.want)
			}
			if got == types.FailureCauseFlake && reason == "" {
				t.Fatal("flake reason is empty")
			}
		})
	}
}

func TestParseFailedTestsFromJunit(t *testing.T) {
	t.Parallel()

	xml := `<?xml version="1.0" encoding="UTF-8"?>
<testsuites>
  <testsuite name="kubevirt" tests="2" failures="1">
    <testcase name="[sig-compute] Infrastructure tls configuration [It] should enforce TLS configuration on virt-template components [sig-compute, Serial]">
      <failure type="Failure">Expected success
tests/infrastructure/tls-configuration.go:114</failure>
    </testcase>
    <testcase name="passing test"/>
    <testcase name="AfterSuite">
      <failure type="Failure">suite cleanup failed</failure>
    </testcase>
  </testsuite>
</testsuites>`

	got, err := parseFailedTestsFromJunit([]byte(xml))
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	if len(got) != 1 {
		t.Fatalf("got %d failed tests, want 1", len(got))
	}
	if !strings.Contains(got[0].Name, "virt-template") {
		t.Errorf("name = %q", got[0].Name)
	}
}

func TestJunitArtifactURL(t *testing.T) {
	t.Parallel()
	in := "https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19055/pull-kubevirt-e2e-k8s-1.34-sig-compute/123"
	want := "https://storage.googleapis.com/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19055/pull-kubevirt-e2e-k8s-1.34-sig-compute/123/artifacts/junit.functest.xml"
	if got := junitArtifactURL(in); got != want {
		t.Errorf("junitArtifactURL = %q, want %q", got, want)
	}
}

func TestLookupTestRateNilReport(t *testing.T) {
	t.Parallel()
	got := LookupTestRate("anything", nil)
	if got.Severity != "unknown" {
		t.Errorf("severity = %q, want unknown", got.Severity)
	}
}

func reportWithRate(testName string, succeeded, failed int) *FlakefinderReport {
	return &FlakefinderReport{
		Tests: []string{testName},
		Data: map[string]map[string]*TestDetails{
			testName: {
				"pull-kubevirt-e2e-k8s-1.34-sig-compute": {
					Succeeded: succeeded,
					Failed:    failed,
				},
			},
		},
	}
}

func massFlakyReport() *FlakefinderReport {
	names := []string{
		`[sig-operator] test one [It]`,
		`[sig-operator] test two [It]`,
		`[sig-operator] test three [It]`,
		`[sig-operator] test four [It]`,
		`[sig-operator] test five [It]`,
	}
	data := map[string]map[string]*TestDetails{}
	for _, name := range names {
		data[name] = map[string]*TestDetails{
			"lane": {Succeeded: 10, Failed: 40},
		}
	}
	return &FlakefinderReport{Tests: names, Data: data}
}
