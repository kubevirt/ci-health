package cifailures

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"gopkg.in/yaml.v3"
)

func TestExtractFailureURLsFromJSON(t *testing.T) {
	tests := []struct {
		name    string
		json    string
		want    int
		wantErr bool
	}{
		{
			name: "multiple boards with URLs",
			json: `{
				"Data": {
					"SIGRetests": {
						"FailedJobLeaderBoard": [
							{"FailureURLs": ["https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/111/job-a/100"]},
							{"FailureURLs": ["https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/222/job-b/200", "https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/333/job-c/300"]}
						]
					}
				}
			}`,
			want: 3,
		},
		{
			name: "empty leaderboard",
			json: `{
				"Data": {
					"SIGRetests": {
						"FailedJobLeaderBoard": []
					}
				}
			}`,
			want: 0,
		},
		{
			name: "no FailedJobLeaderBoard key",
			json: `{
				"Data": {
					"SIGRetests": {}
				}
			}`,
			want: 0,
		},
		{
			name:    "invalid JSON",
			json:    `{not valid`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			urls, err := extractFailureURLsFromJSON([]byte(tt.json))
			if tt.wantErr {
				if err == nil {
					t.Fatal("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if len(urls) != tt.want {
				t.Fatalf("expected %d URLs, got %d", tt.want, len(urls))
			}
		})
	}
}

func writeBuildLogYAML(t *testing.T, dir string, bl *buildLog) string {
	t.Helper()
	buildID, err := buildIDFromJobURL(bl.ProwJobURL)
	if err != nil {
		t.Fatalf("failed to extract build ID: %v", err)
	}
	path := filepath.Join(dir, fmt.Sprintf("%d.yaml", buildID))
	f, err := os.Create(path)
	if err != nil {
		t.Fatalf("failed to create file: %v", err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			t.Fatalf("failed to close file: %v", err)
		}
	}()
	if err := yaml.NewEncoder(f).Encode(bl); err != nil {
		t.Fatalf("failed to encode YAML: %v", err)
	}
	return path
}

func TestGrepBuildLogs(t *testing.T) {
	t.Run("basic match", func(t *testing.T) {
		dir := t.TempDir()
		p1 := writeBuildLogYAML(t, dir, &buildLog{
			ProwJobURL: "https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/111/job-a/100",
			LogContent: "line1\ncould not establish a connection to the node after a generous timeout\nline3",
			Started:    time.Date(2025, 3, 15, 10, 0, 0, 0, time.UTC),
		})
		p2 := writeBuildLogYAML(t, dir, &buildLog{
			ProwJobURL: "https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/222/job-b/200",
			LogContent: "line1\nno match here\nline3",
			Started:    time.Date(2025, 3, 20, 10, 0, 0, 0, time.UTC),
		})

		result := grepBuildLogs([]string{p1, p2}, "could not establish a connection to the node after a generous timeout", true)

		if result.TotalMatched != 1 {
			t.Fatalf("expected 1 matched, got %d", result.TotalMatched)
		}
		if len(result.Months) != 1 {
			t.Fatalf("expected 1 month, got %d", len(result.Months))
		}
		if result.Months[0].Month != "2025-03" {
			t.Fatalf("expected month 2025-03, got %s", result.Months[0].Month)
		}
		if result.Months[0].Matched != 1 {
			t.Fatalf("expected 1 matched in month, got %d", result.Months[0].Matched)
		}
	})

	t.Run("multiple occurrences", func(t *testing.T) {
		dir := t.TempDir()
		p1 := writeBuildLogYAML(t, dir, &buildLog{
			ProwJobURL: "https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/111/job-a/100",
			LogContent: "error here\nerror here\nerror here",
			Started:    time.Date(2025, 1, 10, 10, 0, 0, 0, time.UTC),
		})

		result := grepBuildLogs([]string{p1}, "error here", true)

		if result.TotalMatched != 1 {
			t.Fatalf("expected 1 matched, got %d", result.TotalMatched)
		}
		if result.Months[0].Builds[0].Occurrences != 3 {
			t.Fatalf("expected 3 occurrences, got %d", result.Months[0].Builds[0].Occurrences)
		}
	})

	t.Run("case insensitive", func(t *testing.T) {
		dir := t.TempDir()
		p1 := writeBuildLogYAML(t, dir, &buildLog{
			ProwJobURL: "https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/111/job-a/100",
			LogContent: "Could Not Establish A Connection",
			Started:    time.Date(2025, 2, 5, 10, 0, 0, 0, time.UTC),
		})

		insensitive := grepBuildLogs([]string{p1}, "could not establish a connection", true)
		if insensitive.TotalMatched != 1 {
			t.Fatalf("case-insensitive: expected 1 matched, got %d", insensitive.TotalMatched)
		}

		sensitive := grepBuildLogs([]string{p1}, "could not establish a connection", false)
		if sensitive.TotalMatched != 0 {
			t.Fatalf("case-sensitive: expected 0 matched, got %d", sensitive.TotalMatched)
		}
	})

	t.Run("month grouping", func(t *testing.T) {
		dir := t.TempDir()
		p1 := writeBuildLogYAML(t, dir, &buildLog{
			ProwJobURL: "https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/111/job-a/100",
			LogContent: "target phrase",
			Started:    time.Date(2025, 1, 10, 10, 0, 0, 0, time.UTC),
		})
		p2 := writeBuildLogYAML(t, dir, &buildLog{
			ProwJobURL: "https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/222/job-b/200",
			LogContent: "target phrase",
			Started:    time.Date(2025, 1, 20, 10, 0, 0, 0, time.UTC),
		})
		p3 := writeBuildLogYAML(t, dir, &buildLog{
			ProwJobURL: "https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/333/job-c/300",
			LogContent: "target phrase",
			Started:    time.Date(2025, 2, 5, 10, 0, 0, 0, time.UTC),
		})

		result := grepBuildLogs([]string{p1, p2, p3}, "target phrase", true)

		if len(result.Months) != 2 {
			t.Fatalf("expected 2 months, got %d", len(result.Months))
		}
		if result.Months[0].Month != "2025-01" {
			t.Fatalf("expected first month 2025-01, got %s", result.Months[0].Month)
		}
		if result.Months[0].Matched != 2 {
			t.Fatalf("expected 2 matched in Jan, got %d", result.Months[0].Matched)
		}
		if result.Months[1].Month != "2025-02" {
			t.Fatalf("expected second month 2025-02, got %s", result.Months[1].Month)
		}
		if result.Months[1].Matched != 1 {
			t.Fatalf("expected 1 matched in Feb, got %d", result.Months[1].Matched)
		}
	})

	t.Run("empty log paths", func(t *testing.T) {
		result := grepBuildLogs(nil, "anything", true)

		if result.TotalMatched != 0 {
			t.Fatalf("expected 0 matched, got %d", result.TotalMatched)
		}
		if len(result.Months) != 0 {
			t.Fatalf("expected 0 months, got %d", len(result.Months))
		}
	})

	t.Run("no matches", func(t *testing.T) {
		dir := t.TempDir()
		p1 := writeBuildLogYAML(t, dir, &buildLog{
			ProwJobURL: "https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/111/job-a/100",
			LogContent: "some log content without the pattern",
			Started:    time.Date(2025, 5, 1, 10, 0, 0, 0, time.UTC),
		})

		result := grepBuildLogs([]string{p1}, "nonexistent pattern xyz", true)

		if result.TotalMatched != 0 {
			t.Fatalf("expected 0 matched, got %d", result.TotalMatched)
		}
		if len(result.Months) != 0 {
			t.Fatalf("expected 0 months, got %d", len(result.Months))
		}
	})
}
