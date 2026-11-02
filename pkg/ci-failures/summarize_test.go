package cifailures

import (
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestReadJobFilePreservesEvidence(t *testing.T) {
	for _, role := range []string{"primary", "secondary", "diagnostic", "cleanup", ""} {
		t.Run(role, func(t *testing.T) {
			path := filepath.Join(t.TempDir(), "job.yaml")
			data := `job_name: test-job
build_errors:
  - job_url: https://example.com/job
    build_id: 123
    result: FAILURE
    category: infrastructure
    category_reason: provisioning
    build_log_error_snippets:
      - error_text: provisioning failed
        link_to_log_line: https://example.com/log#10
        role: ` + role + `
        context: provisioning context
        start_line: 8
        error_line: 10
        end_line: 12
      - error_text: cleanup failed
        link_to_log_line: https://example.com/log#20
        role: cleanup
`
			if err := os.WriteFile(path, []byte(data), 0o600); err != nil {
				t.Fatal(err)
			}
			job, err := readJobFile(path)
			if err != nil {
				t.Fatal(err)
			}
			want := SummaryJob{
				JobName: "test-job", JobURL: "https://example.com/job", BuildID: 123,
				Result: "FAILURE", Category: "infrastructure", CategoryReason: "provisioning",
				ErrorSnippets: []SummaryErrorSnippet{
					{ErrorText: "provisioning failed", Link: "https://example.com/log#10", Role: role,
						Context: "provisioning context", StartLine: 8, ErrorLine: 10, EndLine: 12},
					{ErrorText: "cleanup failed", Link: "https://example.com/log#20", Role: "cleanup"},
				},
			}
			if !reflect.DeepEqual(job, want) {
				t.Fatalf("readJobFile() = %#v, want %#v", job, want)
			}
		})
	}
}

func TestSummarizeSessionSerializesEvidence(t *testing.T) {
	dir := t.TempDir()
	input := JobBuildErrors{
		JobName: "test-job",
		BuildErrors: []*JobBuildError{{
			BuildLogErrorSnippets: []*BuildLogErrorSnippet{
				{ErrorText: "primary failure", LinkToLogLine: "https://example.com/log#10", Role: "primary",
					Context: "before\nprimary failure\nafter", StartLine: 8, ErrorLine: 10, EndLine: 12},
				{ErrorText: "cleanup failure", LinkToLogLine: "https://example.com/log#20", Role: "cleanup",
					Context: "cleanup context", StartLine: 18, ErrorLine: 20, EndLine: 22},
				{ErrorText: "legacy failure", LinkToLogLine: "https://example.com/log#30"},
			},
		}},
	}
	data, err := yaml.Marshal(input)
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "job.yaml"), data, 0o600); err != nil {
		t.Fatal(err)
	}
	summary, err := SummarizeSession(dir)
	if err != nil {
		t.Fatal(err)
	}
	data, err = yaml.Marshal(summary)
	if err != nil {
		t.Fatal(err)
	}
	var serialized SessionSummary
	if err := yaml.Unmarshal(data, &serialized); err != nil {
		t.Fatal(err)
	}
	want := []SummaryErrorSnippet{
		{ErrorText: "primary failure", Link: "https://example.com/log#10", Role: "primary",
			Context: "before\nprimary failure\nafter", StartLine: 8, ErrorLine: 10, EndLine: 12},
		{ErrorText: "cleanup failure", Link: "https://example.com/log#20", Role: "cleanup",
			Context: "cleanup context", StartLine: 18, ErrorLine: 20, EndLine: 22},
		{ErrorText: "legacy failure", Link: "https://example.com/log#30"},
	}
	if len(serialized.Jobs) != 1 || !reflect.DeepEqual(serialized.Jobs[0].ErrorSnippets, want) {
		t.Fatalf("serialized summary lost evidence:\n%s", data)
	}
	_, legacy, found := strings.Cut(string(data), "- error_text: legacy failure")
	if !found {
		t.Fatalf("missing legacy snippet:\n%s", data)
	}
	for _, field := range []string{"role:", "context:", "start_line:", "error_line:", "end_line:"} {
		if strings.Contains(legacy, field) {
			t.Errorf("empty field %q must be omitted:\n%s", field, legacy)
		}
	}
}
