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

// CategorizationResult holds the category and the reason it was assigned.
type CategorizationResult struct {
	Category ErrorCategory
	Reason   string
}

type categoryRule struct {
	category ErrorCategory
	reason   string
	pattern  *regexp.Regexp
}

// rules are evaluated in order; first match wins.
// Rules that match on error_text alone.
var rules = []categoryRule{
	// --- External: service/infra failures outside CI team's control ---

	{CategoryExternal, "download failure from external URL", regexp.MustCompile(`Error downloading \[https?://`)},
	{CategoryExternal, "bazel repository fetch failure", regexp.MustCompile(`ERROR: An error occurred during the fetch of repository`)},
	{CategoryExternal, "bazel WORKSPACE fetch rule failure", regexp.MustCompile(`fetching (http_file|http_archive|go_repository|rpm|regctl_repositories|oci_pull|oci_alias|jq_platform_repo|zstd_binary_repo) rule //external:`)},
	{CategoryExternal, "build target depends on failed external fetch", regexp.MustCompile(`which failed to fetch\. no such package`)},
	{CategoryExternal, "container image pull failure", regexp.MustCompile(`unable to copy from source docker://`)},
	{CategoryExternal, "container registry HTTP error", regexp.MustCompile(`(pinging container registry|Requesting bearer token|initializing source docker://).*received unexpected HTTP status:`)},
	{CategoryExternal, "DNS resolution failure", regexp.MustCompile(`dial tcp: lookup .+: (Temporary failure in name resolution|no such host)`)},
	{CategoryExternal, "git clone failure to GitHub", regexp.MustCompile(`fatal: unable to access 'https://github\.com/`)},
	{CategoryExternal, "podman container removal timeout", regexp.MustCompile(`cannot remove container .+ as it could not be stopped`)},
	{CategoryExternal, "podman storage file timeout", regexp.MustCompile(`timed out waiting for file /var/lib/containers/storage/`)},
	{CategoryExternal, "kind cluster deletion failure", regexp.MustCompile(`failed to delete cluster .+: failed to delete nodes`)},
	{CategoryExternal, "transient kube-apiserver body decode noise", regexp.MustCompile(`Body was not decodable \(unable to check for Status\)`)},
	{CategoryExternal, "API rate limiter timeout", regexp.MustCompile(`client rate limiter Wait returned an error: context deadline exceeded`)},
	{CategoryExternal, "skopeo/podman retry with HTTP error", regexp.MustCompile(`Failed, retrying in \d+s .* Error:`)},
	{CategoryExternal, "RPM download failure from mirror", regexp.MustCompile(`Error downloading \[http://mirror\.`)},
	// Bazel remote cache blob fetch failure
	{CategoryExternal, "bazel remote cache blob fetch failure", regexp.MustCompile(`Failed to fetch blobs because they do not exist remotely`)},
	// Bazel remote cache connection timeout (Java/Netty)
	{CategoryExternal, "bazel remote cache connection timeout", regexp.MustCompile(`io\.netty\.channel\.ConnectTimeoutException: connection timed out: bazel-cache`)},
	// Bazel remote cache download timeout
	{CategoryExternal, "bazel remote cache download timeout", regexp.MustCompile(`Exec failed due to IOException:.*Download of .+ timed out`)},
	// Bazel remote cache DNS failure
	{CategoryExternal, "bazel remote cache DNS failure", regexp.MustCompile(`Exec failed due to IOException:.*UnknownHostException: bazel-cache`)},
	// Bazel remote cache IOException (general)
	{CategoryExternal, "bazel remote cache IO failure", regexp.MustCompile(`Exec failed due to IOException:`)},

	// --- PR Build: code/build failures caused by the PR itself ---

	{CategoryPRBuild, "bazel build failure", regexp.MustCompile(`ERROR: Build failed\. Not running target`)},
	{CategoryPRBuild, "bazel analysis failure", regexp.MustCompile(`ERROR: Analysis of target '.+' failed; build aborted`)},
	// Test panic detected in output
	{CategoryPRBuild, "panic detected in test output", regexp.MustCompile(`ERROR: Found panic in test output`)},

	// --- Internal: CI configuration errors fixable by sig-ci ---

	{CategoryInternal, "taint not found during cluster setup", regexp.MustCompile(`error: taint "node-role\.kubernetes\.io/(control-plane|master):NoSchedule" not found`)},
	{CategoryInternal, "make cluster lifecycle target failure", regexp.MustCompile(`make: \*\*\* \[Makefile:\d+: cluster-(up|sync|down)\] Error`)},
	{CategoryInternal, "make bazel-build target failure", regexp.MustCompile(`make: \*\*\* \[Makefile:\d+: bazel-build-(images|functests)\] Error`)},
	{CategoryInternal, "KubeVirt deployment timeout", regexp.MustCompile(`timed out waiting for the condition on kubevirts/kubevirt`)},
	{CategoryInternal, "deployment timeout", regexp.MustCompile(`timed out waiting for the condition on deployments/`)},
	{CategoryInternal, "cluster bootstrap JWS token issue", regexp.MustCompile(`could not find a JWS signature in the cluster-info ConfigMap`)},
	// Control plane failed to start during kind cluster creation
	{CategoryInternal, "control plane startup failure", regexp.MustCompile(`error.*execution phase wait-control-plane: failed while waiting for the control plane to start`)},
	// Kind cluster creation failure (log line pattern not found)
	{CategoryInternal, "kind cluster creation failure", regexp.MustCompile(`ERROR: failed to create cluster: could not find a log line that matches`)},
}

type contextExternalPattern struct {
	reason  string
	pattern *regexp.Regexp
}

// contextExternalPatterns are checked on the context field to reclassify
// errors that would otherwise be internal/pr-build but are actually caused
// by external service failures visible only in the surrounding log lines.
var contextExternalPatterns = []contextExternalPattern{
	{"download failure in context", regexp.MustCompile(`Error downloading \[https?://`)},
	{"container image pull failure in context", regexp.MustCompile(`unable to copy from source docker://`)},
	{"git clone failure in context", regexp.MustCompile(`fatal: unable to access 'https://github\.com/`)},
	{"HTTP error in context", regexp.MustCompile(`received unexpected HTTP status:`)},
	{"DNS failure in context", regexp.MustCompile(`dial tcp: lookup .+: (Temporary failure in name resolution|no such host)`)},
	{"failed external fetch in context", regexp.MustCompile(`which failed to fetch\. no such package`)},
	{"download error in context", regexp.MustCompile(`Error in download`)},
}

// CategorizeError classifies a single error text into one of the known categories.
func CategorizeError(errorText string) ErrorCategory {
	return categorizeError(errorText).Category
}

func categorizeError(errorText string) CategorizationResult {
	for _, rule := range rules {
		if rule.pattern.MatchString(errorText) {
			return CategorizationResult{rule.category, rule.reason}
		}
	}
	return CategorizationResult{CategoryNeedsInvestigation, "no matching pattern"}
}

// CategorizeJobBuildError classifies a JobBuildError by examining all its error snippets.
// It uses the first snippet's error_text for primary classification, then checks
// the context of all snippets to detect external failures that aren't apparent
// from the error_text alone. It sets err.CategoryReason with the explanation.
func CategorizeJobBuildError(err *JobBuildError) ErrorCategory {
	result := categorizeJobBuildError(err)
	err.CategoryReason = result.Reason
	return result.Category
}

func categorizeJobBuildError(err *JobBuildError) CategorizationResult {
	if len(err.BuildLogErrorSnippets) == 0 {
		return CategorizationResult{CategoryNeedsInvestigation, "no error snippets"}
	}

	primary := categorizeError(err.BuildLogErrorSnippets[0].ErrorText)

	// If already external, no need to refine
	if primary.Category == CategoryExternal {
		return primary
	}

	// For non-external categories, check if any snippet's context reveals
	// an external root cause (e.g. a make error caused by a download failure)
	for _, snippet := range err.BuildLogErrorSnippets {
		// Check context
		for _, pat := range contextExternalPatterns {
			if pat.pattern.MatchString(snippet.Context) {
				return CategorizationResult{CategoryExternal, pat.reason}
			}
		}
		// Also check error_text of non-primary snippets
		if r := categorizeError(snippet.ErrorText); r.Category == CategoryExternal {
			return CategorizationResult{CategoryExternal, r.Reason + " (from secondary snippet)"}
		}
	}

	return primary
}
