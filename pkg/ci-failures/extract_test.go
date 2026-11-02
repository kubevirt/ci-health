package cifailures

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestExtractErrorSnippets(t *testing.T) {
	tests := []struct {
		name     string
		lines    []string
		primary  int
		category ErrorCategory
	}{
		{
			name: "hostname provisioning error survives decode noise and teardown",
			lines: []string{
				`12:00:00: Body was not decodable: json parse error: invalid body`,
				`12:00:01: kube-apiserver is healthy after 1s`,
				`12:00:02: Error from server (NotFound): nodes "node01" not found`,
				`make: *** [Makefile:174: cluster-up] Error 1`,
				`+ ret=2`,
				`+ make cluster-down`,
				`make: *** [Makefile:182: cluster-down] Error 1`,
			},
			primary: 3, category: CategoryInternal,
		},
		{
			name: "missing layer survives Podman cleanup timeout",
			lines: []string{
				`WARNING: Download failed: recovered mirror warning`,
				`INFO: Build completed successfully, 1 total action`,
				`Error during deploy: deploying images: layer sha256:d689b4d17b73 not found in any of the original registries`,
				`make: *** [Makefile:194: cluster-sync] Error 1`,
				`+ ret=2`,
				`Error: cannot remove container abc as it could not be stopped: given PID did not die within timeout`,
			},
			primary: 3, category: CategoryNeedsInvestigation,
		},
		{
			name: "Buildifier fetch selects underlying download rather than wrapper",
			lines: []string{
				`INFO: Build completed successfully, 21 total actions`,
				`WARNING: Download from https://github.com/buildifier failed: GET returned 504 Gateway Time-out`,
				`ERROR: /cache/http.bzl:178:33: An error occurred during the fetch of repository 'buildifier_linux_amd64':`,
				`Traceback (most recent call last):`,
				`Error in download: java.io.IOException: Error downloading [https://github.com/buildifier] to /tmp/buildifier: GET returned 504 Gateway Time-out`,
				`ERROR: Analysis of target '//:buildifier' failed; build aborted: Analysis failed`,
				`make: *** [Makefile:28: bazel-build-functests] Error 1`,
				`+ ret=2`,
				`make: *** [Makefile:182: cluster-down] Error 1`,
			},
			primary: 5, category: CategoryExternal,
		},
		{
			name: "quay EOF before diagnostics trap",
			lines: []string{
				`+ make cluster-up`,
				`Error: unable to copy from source docker://quay.io/phoracek/lspci: unexpected EOF`,
				`make: *** [Makefile:174: cluster-up] Error 125`,
				`++ collect_debug_logs`,
				`Error: diagnostic command failed`,
			},
			primary: 2, category: CategoryExternal,
		},
		{
			name: "ERR trap diagnostics cannot override initiating make failure",
			lines: []string{
				`make: *** [Makefile:174: cluster-up] Error 1`,
				`++ collect_debug_logs`,
				`INFO: Build completed successfully, 1 total action`,
				`Error: unable to copy from source docker://quay.io/debug-image: unexpected EOF`,
				`+ ret=2`,
			},
			primary: 1, category: CategoryInternal,
		},
		{
			name: "Bazel compiler failure retains PR classification",
			lines: []string{
				`ERROR: /workspace/pkg/example/BUILD.bazel:3:11: GoCompilePkg //pkg/example:example failed: (Exit 1): builder failed: error executing GoCompilePkg command`,
				`pkg/example/example.go:12: undefined: missingSymbol`,
				`ERROR: Build failed. Not running target`,
				`make: *** [Makefile:28: bazel-build-functests] Error 1`,
				`+ ret=2`,
			},
			primary: 1, category: CategoryPRBuild,
		},
		{
			name: "Bazel summary without compiler diagnostics retains PR classification",
			lines: []string{
				`ERROR: Build did NOT complete successfully`,
				`ERROR: Build failed. Not running target`,
				`make: *** [Makefile:28: bazel-build-functests] Error 1`,
				`+ ret=2`,
			},
			primary: 1, category: CategoryPRBuild,
		},
		{
			name: "successful operation invalidates earlier external error",
			lines: []string{
				`ERROR: Error downloading [https://example.com/retried]`,
				`INFO: Build completed successfully, 1 total action`,
				`Error from server (NotFound): nodes "node02" not found`,
				`make: *** [Makefile:174: cluster-up] Error 1`,
			},
			primary: 3, category: CategoryInternal,
		},
		{
			name: "recovered retry does not override PR failure",
			lines: []string{
				`level=warning msg="Failed, retrying in 1s ... Error: unable to copy from source docker://quay.io/image"`,
				`ERROR: Build failed. Not running target`,
				`+ ret=2`,
			},
			primary: 2, category: CategoryPRBuild,
		},
		{
			name: "unknown error retained rather than recognized cleanup",
			lines: []string{
				`Error: an unfamiliar deployment failure`,
				`make: *** [Makefile:194: cluster-sync] Error 1`,
				`+ ret=2`,
				`Error: cannot remove container abc as it could not be stopped`,
			},
			primary: 1, category: CategoryNeedsInvestigation,
		},
		{
			name: "cleanup failure without an earlier operation failure",
			lines: []string{
				`+ make cluster-down`,
				`Error: cannot remove container abc as it could not be stopped`,
				`make: *** [Makefile:182: cluster-down] Error 1`,
			},
			primary: 2, category: CategoryInternal,
		},
		{
			name: "cleanup without saved shell status cannot replace primary",
			lines: []string{
				`Error from server (NotFound): nodes "node01" not found`,
				`+ make cluster-down`,
				`Error: unable to copy from source docker://quay.io/cleanup-image`,
			},
			primary: 1, category: CategoryInternal,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			const url = "https://example.com/build"
			snippets := extractErrorSnippets(strings.Join(tt.lines, "\n"), url)
			if len(snippets) == 0 {
				t.Fatal("no snippets")
			}
			primary := snippets[0]
			if primary.Role != "primary" || primary.ErrorLine != tt.primary || primary.ErrorText != tt.lines[tt.primary-1] {
				t.Fatalf("primary = %+v, want line %d", primary, tt.primary)
			}
			if primary.LinkToLogLine != fmt.Sprintf("%s#1:build-log.txt%%3A%d", url, tt.primary) {
				t.Fatalf("wrong primary link: %s", primary.LinkToLogLine)
			}
			job := &JobBuildError{BuildLogErrorSnippets: snippets}
			if got := CategorizeJobBuildError(job); got != tt.category {
				t.Fatalf("category = %q (%s), want %q", got, job.CategoryReason, tt.category)
			}
			seen := map[int]bool{}
			for _, s := range snippets {
				if seen[s.ErrorLine] {
					t.Fatalf("duplicate line %d", s.ErrorLine)
				}
				seen[s.ErrorLine] = true
				start, end := max(1, s.ErrorLine-3), min(len(tt.lines), s.ErrorLine+3)
				if s.StartLine != start || s.EndLine != end || s.Context != strings.Join(tt.lines[start-1:end], "\n") {
					t.Fatalf("incorrect context bounds: %+v", s)
				}
			}
			for i, line := range tt.lines {
				if buildLogLineMatchesExpression.MatchString(line) && !seen[i+1] {
					t.Fatalf("lost evidence at line %d", i+1)
				}
			}
		})
	}
}

func TestExtractErrorSnippetsWithoutInitiatingFailure(t *testing.T) {
	for _, content := range []string{"", "everything succeeded", "Body was not decodable: json parse error: invalid body", "+ ret=2\nError: cleanup failed"} {
		job := &JobBuildError{BuildLogErrorSnippets: extractErrorSnippets(content, "https://example.com")}
		if got := CategorizeJobBuildError(job); got != CategoryNeedsInvestigation {
			t.Errorf("%q: category = %s, want needs-investigation", content, got)
		}
	}
}

func TestAnalyzeBuildAndExtractErrorsUseSameEvidence(t *testing.T) {
	previousLogsDir := logsDir
	logsDir = t.TempDir()
	t.Cleanup(func() { logsDir = previousLogsDir })
	const url = "https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19492/pull-kubevirt-e2e-k8s-1.36-sig-storage/123"
	cached := &buildLog{
		ProwJobURL: url, Result: "FAILURE",
		LogContent: "Error from server (NotFound): nodes \"node01\" not found\nmake: *** [Makefile:174: cluster-up] Error 1\n+ ret=2\nmake: *** [Makefile:182: cluster-down] Error 1",
	}
	data, err := yaml.Marshal(cached)
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(logsDir, "123.yaml"), data, 0o600); err != nil {
		t.Fatal(err)
	}
	single, err := AnalyzeBuild(url)
	if err != nil {
		t.Fatal(err)
	}
	paths, err := ExtractErrors([]string{url}, t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	if len(paths) != 1 {
		t.Fatalf("output paths = %v, want one", paths)
	}
	data, err = os.ReadFile(paths[0])
	if err != nil {
		t.Fatal(err)
	}
	var batch JobBuildErrors
	if err := yaml.Unmarshal(data, &batch); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(single, &batch) {
		t.Fatalf("single and batch results differ: %+v, %+v", single, batch)
	}
	if single.BuildErrors[0].Category != string(CategoryInternal) || single.BuildErrors[0].BuildLogErrorSnippets[0].ErrorLine != 1 {
		t.Fatalf("wrong primary/category: %+v", single.BuildErrors[0])
	}
}

func TestClusterErrorsUsesInitiatingFailure(t *testing.T) {
	first := &JobBuildError{BuildLogErrorSnippets: extractErrorSnippets("Error: first causal failure\n+ ret=2\nError: cleanup A", "first")}
	second := &JobBuildError{BuildLogErrorSnippets: extractErrorSnippets("Error: first causal failure\n+ ret=2\nError: cleanup B", "second")}
	third := &JobBuildError{BuildLogErrorSnippets: extractErrorSnippets("Error: unrelated causal failure\n+ ret=2\nError: cleanup A", "third")}
	clusters := ClusterErrors([]*JobBuildError{first, second, third}, 1)
	if len(clusters) != 2 {
		t.Fatalf("clusters = %d, want 2", len(clusters))
	}
	if len(clusters[0].Errors) != 2 {
		t.Fatalf("same causes did not cluster: %+v", clusters)
	}
}

func TestHistoricalMissingLayersClusterTogether(t *testing.T) {
	const cause = "Error during deploy: deploying images: layer sha256:d689b4d17b73dba36b3168c448be39ccb0e84fd1ad8914f378c78f8e20b51fa9 not found in any of the original registries"
	operator := &JobBuildError{BuildLogErrorSnippets: extractErrorSnippets("15:40:36: "+cause+"\n+ ret=2\nmake: *** [Makefile:182: cluster-down] Error 1", "operator")}
	storage := &JobBuildError{BuildLogErrorSnippets: extractErrorSnippets("15:59:16: "+cause+"\n+ ret=2\nError: cannot remove container abc as it could not be stopped", "storage")}
	clusters := ClusterErrors([]*JobBuildError{operator, storage}, 0.9)
	if len(clusters) != 1 || len(clusters[0].Errors) != 2 {
		t.Fatalf("historical missing-layer failures did not cluster: %+v", clusters)
	}
}
