package cifailures

import (
	"fmt"
	"strings"

	"github.com/joshdk/go-junit"

	"github.com/kubevirt/ci-health/pkg/types"
)

const (
	severityLikelyFlaky  = "likely-flaky"
	massFailureThreshold = 5
	junitFetchCIReason   = "job failed with junit but no named test failures"
)

var suiteHooks = map[string]bool{
	"AfterSuite":        true,
	"BeforeSuite":       true,
	"BeforeEach":        true,
	"AfterEach":         true,
	"JustBeforeEach":    true,
	"JustAfterEach":     true,
	"ReportAfterSuite":  true,
	"ReportBeforeSuite": true,
}

// FailedTest is a named junit failure used to classify an e2e job.
type FailedTest struct {
	Name string
}

var readGCSFn = retrieveFileContentFromGCS

func junitArtifactURL(prowJobURL string) string {
	u := normalizeJobURL(prowJobURL)
	u = strings.Replace(u, "https://prow.ci.kubevirt.io/view/gs/", "https://storage.googleapis.com/", 1)
	return strings.TrimRight(u, "/") + "/artifacts/junit.functest.xml"
}

// FetchFailedJunitTests loads junit.functest.xml for a Prow job and returns failed tests.
func FetchFailedJunitTests(prowJobURL string) ([]FailedTest, error) {
	raw, err := readGCSFn(junitArtifactURL(prowJobURL))
	if err != nil {
		return nil, err
	}
	return parseFailedTestsFromJunit(raw)
}

func parseFailedTestsFromJunit(raw []byte) ([]FailedTest, error) {
	suites, err := junit.Ingest(raw)
	if err != nil {
		return nil, fmt.Errorf("failed to parse junit: %w", err)
	}

	seen := map[string]bool{}
	var tests []FailedTest
	for _, suite := range suites {
		for _, test := range suite.Tests {
			if test.Status != junit.StatusFailed && test.Status != junit.StatusError {
				continue
			}
			name := strings.TrimSpace(test.Name)
			if name == "" || seen[name] || suiteHooks[name] {
				continue
			}
			seen[name] = true
			tests = append(tests, FailedTest{Name: name})
		}
	}
	return tests, nil
}

// ClassifyE2ECause labels a SIG-lane failure as flake, ci, or unlabeled.
// All named failures likely-flaky → flake. No named junit failures → ci.
// Many named failures that are not all flaky → ci (suite collapse).
// Everything else is left unlabeled: proving the PR caused a test is not
// reliable, and unlabeled leftovers are the remainder after flake and CI.
func ClassifyE2ECause(tests []FailedTest, report *FlakefinderReport) (cause, reason string) {
	if len(tests) == 0 {
		return types.FailureCauseCI, junitFetchCIReason
	}

	var flakeReasons, otherReasons []string
	for _, test := range tests {
		rate := LookupTestRate(test.Name, report)
		label := shortTestName(test.Name)
		if rate.Severity == severityLikelyFlaky {
			flakeReasons = append(flakeReasons, fmt.Sprintf("%s: likely-flaky (%.0f%% success)", label, rate.SuccessRate))
			continue
		}
		otherReasons = append(otherReasons, fmt.Sprintf("%s: e2e test failure", label))
	}

	if len(otherReasons) == 0 {
		return types.FailureCauseFlake, joinReasons(flakeReasons)
	}
	if len(tests) >= massFailureThreshold {
		return types.FailureCauseCI, fmt.Sprintf("%d named test failures in one job; treating as ci", len(tests))
	}
	return "", joinReasons(append(append([]string{}, flakeReasons...), otherReasons...))
}

func shortTestName(name string) string {
	n := normalizeTestName(name)
	if len(n) > 100 {
		return n[:97] + "..."
	}
	return n
}

func joinReasons(reasons []string) string {
	const max = 3
	if len(reasons) > max {
		reasons = append(reasons[:max], fmt.Sprintf("(+%d more)", len(reasons)-max))
	}
	return strings.Join(reasons, "; ")
}
