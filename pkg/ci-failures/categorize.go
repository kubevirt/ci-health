package cifailures

import "regexp"

// ErrorCategory represents the classification of a CI build failure.
type ErrorCategory string

const (
	// CategoryExternal represents failures caused by external services outside our control.
	// These must NOT be counted against the sig-ci failure rate.
	CategoryExternal ErrorCategory = "external"

	// CategoryInternal represents CI configuration errors fixable by sig-ci.
	// These ARE counted against the sig-ci failure rate.
	CategoryInternal ErrorCategory = "internal"

	// CategoryPRBuild represents build failures caused by the PR code itself.
	// These are engineering errors, not CI errors.
	CategoryPRBuild ErrorCategory = "pr-build"

	// CategoryNeedsInvestigation represents errors that don't match any known pattern.
	// These require manual review to create a new pattern.
	CategoryNeedsInvestigation ErrorCategory = "needs-investigation"
)

type categoryRule struct {
	category ErrorCategory
	pattern  *regexp.Regexp
}

// rules are evaluated in order; first match wins.
var rules = []categoryRule{
	// --- External: service/infra failures outside CI team's control ---

	// GitHub/external download failures (404, 502, etc.)
	{CategoryExternal, regexp.MustCompile(`Error downloading \[https://`)},
	// Podman container removal timeout
	{CategoryExternal, regexp.MustCompile(`cannot remove container .+ as it could not be stopped`)},
	// Podman storage file timeout
	{CategoryExternal, regexp.MustCompile(`timed out waiting for file /var/lib/containers/storage/`)},
	// kind cluster deletion failure via podman
	{CategoryExternal, regexp.MustCompile(`failed to delete cluster .+: failed to delete nodes`)},
	// Transient kube-apiserver body decode noise (not a real failure)
	{CategoryExternal, regexp.MustCompile(`Body was not decodable \(unable to check for Status\)`)},
	// API rate limiter timeout
	{CategoryExternal, regexp.MustCompile(`client rate limiter Wait returned an error: context deadline exceeded`)},

	// --- PR Build: code/build failures caused by the PR itself ---

	// Bazel analysis failure
	{CategoryPRBuild, regexp.MustCompile(`ERROR: Analysis of target '.+' failed; build aborted`)},
	// Bazel build failure
	{CategoryPRBuild, regexp.MustCompile(`ERROR: Build failed\. Not running target`)},
	// Bazel build make target failures
	{CategoryPRBuild, regexp.MustCompile(`make: \*\*\* \[Makefile:\d+: bazel-build-(images|functests)\] Error`)},

	// --- Internal: CI configuration errors fixable by sig-ci ---

	// Taint not found during cluster setup
	{CategoryInternal, regexp.MustCompile(`error: taint "node-role\.kubernetes\.io/(control-plane|master):NoSchedule" not found`)},
	// make target failures for CI lifecycle (cluster-up, cluster-sync, cluster-down)
	{CategoryInternal, regexp.MustCompile(`make: \*\*\* \[Makefile:\d+: cluster-(up|sync|down)\] Error`)},
	// KubeVirt deployment timeout
	{CategoryInternal, regexp.MustCompile(`timed out waiting for the condition on kubevirts/kubevirt`)},
	// Cluster bootstrap JWS token issue
	{CategoryInternal, regexp.MustCompile(`could not find a JWS signature in the cluster-info ConfigMap`)},
}

// CategorizeError classifies a single error text into one of the known categories.
func CategorizeError(errorText string) ErrorCategory {
	for _, rule := range rules {
		if rule.pattern.MatchString(errorText) {
			return rule.category
		}
	}
	return CategoryNeedsInvestigation
}

// CategorizeJobBuildError classifies a JobBuildError using its first (most prominent) error snippet.
func CategorizeJobBuildError(err *JobBuildError) ErrorCategory {
	if len(err.BuildLogErrorSnippets) == 0 {
		return CategoryNeedsInvestigation
	}
	return CategorizeError(err.BuildLogErrorSnippets[0].ErrorText)
}
