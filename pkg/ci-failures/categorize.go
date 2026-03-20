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
// Rules that match on error_text alone.
var rules = []categoryRule{
	// --- External: service/infra failures outside CI team's control ---

	// GitHub/external download failures (404, 502, etc.) via bazel
	{CategoryExternal, regexp.MustCompile(`Error downloading \[https?://`)},
	// Bazel repository fetch failures (covers all "An error occurred during the fetch of repository" errors)
	{CategoryExternal, regexp.MustCompile(`ERROR: An error occurred during the fetch of repository`)},
	// Bazel WORKSPACE fetch rule failures (traceback from fetching rules)
	{CategoryExternal, regexp.MustCompile(`fetching (http_file|http_archive|go_repository|rpm|regctl_repositories|oci_pull|oci_alias|jq_platform_repo|zstd_binary_repo) rule //external:`)},
	// Bazel build target depends on failed external fetch
	{CategoryExternal, regexp.MustCompile(`which failed to fetch\. no such package`)},
	// Container image pull failures from quay.io (500, 502, 504, DNS, etc.)
	{CategoryExternal, regexp.MustCompile(`unable to copy from source docker://`)},
	// Container registry bearer token / ping failures
	{CategoryExternal, regexp.MustCompile(`(pinging container registry|Requesting bearer token|initializing source docker://).*received unexpected HTTP status:`)},
	// DNS resolution failures for container registries
	{CategoryExternal, regexp.MustCompile(`dial tcp: lookup .+: (Temporary failure in name resolution|no such host)`)},
	// git clone failures to GitHub (HTTP2 framing, 500, 502, etc.)
	{CategoryExternal, regexp.MustCompile(`fatal: unable to access 'https://github\.com/`)},
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
	// Skopeo/podman retry warnings with HTTP errors
	{CategoryExternal, regexp.MustCompile(`Failed, retrying in \d+s .* Error:`)},
	// RPM download failures from mirrors (centos, fedora, etc.)
	{CategoryExternal, regexp.MustCompile(`Error downloading \[http://mirror\.`)},

	// --- PR Build: code/build failures caused by the PR itself ---

	// Bazel build failure (only when NOT caused by external fetch — those are caught above)
	{CategoryPRBuild, regexp.MustCompile(`ERROR: Build failed\. Not running target`)},
	// Bazel analysis failure (only when NOT caused by external fetch — those are caught above)
	{CategoryPRBuild, regexp.MustCompile(`ERROR: Analysis of target '.+' failed; build aborted`)},

	// --- Internal: CI configuration errors fixable by sig-ci ---

	// Taint not found during cluster setup
	{CategoryInternal, regexp.MustCompile(`error: taint "node-role\.kubernetes\.io/(control-plane|master):NoSchedule" not found`)},
	// make target failures for CI lifecycle (cluster-up, cluster-sync, cluster-down)
	{CategoryInternal, regexp.MustCompile(`make: \*\*\* \[Makefile:\d+: cluster-(up|sync|down)\] Error`)},
	// make target failures for bazel build (generic, after more specific external patterns are tried)
	{CategoryInternal, regexp.MustCompile(`make: \*\*\* \[Makefile:\d+: bazel-build-(images|functests)\] Error`)},
	// KubeVirt deployment timeout
	{CategoryInternal, regexp.MustCompile(`timed out waiting for the condition on kubevirts/kubevirt`)},
	// Other deployment timeouts (e.g. CNAO)
	{CategoryInternal, regexp.MustCompile(`timed out waiting for the condition on deployments/`)},
	// Cluster bootstrap JWS token issue
	{CategoryInternal, regexp.MustCompile(`could not find a JWS signature in the cluster-info ConfigMap`)},
}

// contextExternalPatterns are checked on the context field to reclassify
// errors that would otherwise be internal/pr-build but are actually caused
// by external service failures visible only in the surrounding log lines.
var contextExternalPatterns = []*regexp.Regexp{
	regexp.MustCompile(`Error downloading \[https?://`),
	regexp.MustCompile(`unable to copy from source docker://`),
	regexp.MustCompile(`fatal: unable to access 'https://github\.com/`),
	regexp.MustCompile(`received unexpected HTTP status:`),
	regexp.MustCompile(`dial tcp: lookup .+: (Temporary failure in name resolution|no such host)`),
	regexp.MustCompile(`which failed to fetch\. no such package`),
	regexp.MustCompile(`Error in download`),
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

// CategorizeJobBuildError classifies a JobBuildError by examining all its error snippets.
// It uses the first snippet's error_text for primary classification, then checks
// the context of all snippets to detect external failures that aren't apparent
// from the error_text alone.
func CategorizeJobBuildError(err *JobBuildError) ErrorCategory {
	if len(err.BuildLogErrorSnippets) == 0 {
		return CategoryNeedsInvestigation
	}

	primary := CategorizeError(err.BuildLogErrorSnippets[0].ErrorText)

	// If already external, no need to refine
	if primary == CategoryExternal {
		return CategoryExternal
	}

	// For non-external categories, check if any snippet's context reveals
	// an external root cause (e.g. a make error caused by a download failure)
	for _, snippet := range err.BuildLogErrorSnippets {
		// Check context
		for _, pat := range contextExternalPatterns {
			if pat.MatchString(snippet.Context) {
				return CategoryExternal
			}
		}
		// Also check error_text of non-primary snippets
		if CategorizeError(snippet.ErrorText) == CategoryExternal {
			return CategoryExternal
		}
	}

	return primary
}
