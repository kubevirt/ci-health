package main

import (
	"bytes"
	"html"
	"strings"
	"testing"

	cifailures "github.com/kubevirt/ci-health/pkg/ci-failures"
)

func TestReportPreservesPrimaryAndSecondaryEvidence(t *testing.T) {
	job := &cifailures.JobBuildError{
		JobURL: "https://example.com/job", Category: "infrastructure", CategoryReason: "provisioning",
		BuildLogErrorSnippets: []*cifailures.BuildLogErrorSnippet{
			{Role: "primary", ErrorText: "provisioning <failed>", LinkToLogLine: "https://example.com/log#10", Context: "primary <context>"},
			{Role: "secondary", ErrorText: "connection <failed>", LinkToLogLine: "https://example.com/log#20", Context: "secondary <context>"},
			{Role: "diagnostic", ErrorText: "diagnostic <failed>", LinkToLogLine: "https://example.com/log#30", Context: "diagnostic <context>"},
			{Role: "cleanup", ErrorText: "cleanup <failed>", LinkToLogLine: "https://example.com/log#40", Context: "cleanup <context>"},
		},
	}
	failures := Failures{CategoryName: "test failures", Anchor: "test"}
	failures.AddError("test", job)
	for _, name := range []string{"failures", "categorized-failures"} {
		t.Run(name, func(t *testing.T) {
			var output bytes.Buffer
			if err := ciFailuresTemplate.ExecuteTemplate(&output, name, &failures); err != nil {
				t.Fatal(err)
			}
			report := output.String()
			headline, evidence, found := strings.Cut(report, "<summary>secondary evidence</summary>")
			if !found {
				t.Fatalf("secondary evidence is not expandable:\n%s", report)
			}
			if !strings.Contains(headline, "<code>provisioning &lt;failed&gt;</code> [build-log](https://example.com/log#10)") || strings.Contains(headline, "cleanup") {
				t.Errorf("headline must show primary, not cleanup:\n%s", headline)
			}
			for _, snippet := range job.BuildLogErrorSnippets[1:] {
				for _, want := range []string{
					"<strong>" + snippet.Role + "</strong>",
					"<code>" + html.EscapeString(snippet.ErrorText) + "</code>",
					"[build-log](" + snippet.LinkToLogLine + ")",
					"<pre>" + html.EscapeString(snippet.Context) + "</pre>",
				} {
					if !strings.Contains(evidence, want) {
						t.Errorf("missing secondary evidence %q:\n%s", want, evidence)
					}
				}
			}
			if name == "categorized-failures" && !strings.Contains(report, "<pre>primary &lt;context&gt;</pre>") {
				t.Error("existing primary context was lost")
			}
			if strings.Contains(report, "<failed>") || strings.Contains(report, "<context>") {
				t.Error("snippet content must remain HTML-escaped")
			}
		})
	}
}

func TestReportWithoutSecondaryEvidence(t *testing.T) {
	for _, snippets := range [][]*cifailures.BuildLogErrorSnippet{
		nil,
		{{ErrorText: "legacy primary", LinkToLogLine: "https://example.com/log#10"}},
	} {
		failures := Failures{}
		failures.AddError("test", &cifailures.JobBuildError{JobURL: "https://example.com/job", BuildLogErrorSnippets: snippets})
		for _, name := range []string{"failures", "categorized-failures"} {
			var output bytes.Buffer
			if err := ciFailuresTemplate.ExecuteTemplate(&output, name, &failures); err != nil {
				t.Fatal(err)
			}
			report := output.String()
			if strings.Contains(report, "secondary evidence") {
				t.Errorf("unexpected empty secondary evidence:\n%s", report)
			}
			want := "[build-log](https://example.com/job)"
			if len(snippets) > 0 {
				want = "<code>legacy primary</code> [build-log](https://example.com/log#10)"
			}
			if !strings.Contains(report, want) {
				t.Errorf("missing existing fallback or headline %q:\n%s", want, report)
			}
		}
	}
}
