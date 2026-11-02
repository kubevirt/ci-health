package cifailures

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	cleanupStart      = regexp.MustCompile(`^\++ (ret=[1-9][0-9]*|collect_debug_logs\b)|^Cleaning up after podman`)
	operationProgress = regexp.MustCompile(`INFO: Build completed successfully|Done \./hack/|^\+ make (cluster-up|cluster-sync|functest|bazel-build-[^ ]+)\b`)
	diagnosticError   = regexp.MustCompile(`(?i)\bwarning:|\blevel=warning\b|Body was not decodable|\bretrying\b|kube-apiserver check failed`)
	failureWrapper    = regexp.MustCompile(`\bmake: \*\*\*|An error occurred during the fetch of repository|ERROR: (Analysis of target .*failed; build aborted|Build (failed\. Not running target|did NOT complete successfully))`)
	cleanupError      = regexp.MustCompile(`cannot remove container|failed to delete cluster|volume .+ is being used|cluster-down\] Error`)
)

// extractErrorSnippets keeps all matched evidence, with the selected failure first
// for the report and clustering consumers. Selection is deliberately bounded by
// shell failure handling and successful operations, rather than external ownership.
func extractErrorSnippets(content, jobURL string) []*BuildLogErrorSnippet {
	lines := strings.Split(content, "\n")
	end := len(lines)
	hasFailure := false
	for i, line := range lines {
		if cleanupStart.MatchString(line) || (hasFailure && strings.HasPrefix(line, "+ make cluster-down")) {
			end = i
			break
		}
		if operationProgress.MatchString(line) {
			hasFailure = false
		} else if buildLogLineMatchesExpression.MatchString(line) && !diagnosticError.MatchString(line) {
			hasFailure = true
		}
	}
	start := 0
	for i := 0; i < end; i++ {
		if operationProgress.MatchString(lines[i]) {
			start = i + 1
		}
	}

	var snippets []*BuildLogErrorSnippet
	primary := -1
	fallback := -1
	for i, line := range lines {
		if !buildLogLineMatchesExpression.MatchString(line) {
			continue
		}
		role := "secondary"
		switch {
		case i >= end:
			role = "cleanup"
		case i < start || diagnosticError.MatchString(line):
			role = "diagnostic"
		default:
			if fallback < 0 {
				fallback = len(snippets)
			}
			if primary < 0 && !failureWrapper.MatchString(line) && !cleanupError.MatchString(line) {
				primary = len(snippets)
			}
		}
		snippet := &BuildLogErrorSnippet{
			Role:          role,
			StartLine:     max(0, i-3) + 1,
			EndLine:       min(len(lines)-1, i+3) + 1,
			ErrorLine:     i + 1,
			ErrorText:     strings.TrimSpace(line),
			LinkToLogLine: fmt.Sprintf("%s#1:build-log.txt%%3A%d", jobURL, i+1),
		}
		var context []string
		for j := snippet.StartLine - 1; j < snippet.EndLine; j++ {
			context = append(context, strings.TrimSpace(lines[j]))
		}
		snippet.Context = strings.Join(context, "\n")
		snippets = append(snippets, snippet)
	}
	if primary < 0 {
		primary = fallback
	}
	if primary >= 0 {
		snippet := snippets[primary]
		snippet.Role = "primary"
		copy(snippets[1:primary+1], snippets[:primary])
		snippets[0] = snippet
	}
	return snippets
}
