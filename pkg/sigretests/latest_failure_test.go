package sigretests

import (
	"testing"
)

func TestLatestFailurePerJob(t *testing.T) {
	t.Parallel()

	jobs := []annotatedJob{
		{job: job{jobName: "e2e-storage", buildNumber: "10", failure: true, buildURL: "https://prow/old-fail"}},
		{job: job{jobName: "e2e-storage", buildNumber: "11", failure: false, buildURL: "https://prow/old-pass"}},
		{job: job{jobName: "e2e-compute", buildNumber: "20", failure: true, buildURL: "https://prow/head-fail"}},
		{job: job{jobName: "e2e-network", buildNumber: "21", failure: false, buildURL: "https://prow/head-pass"}},
		{job: job{jobName: "e2e-compute", buildNumber: "19", failure: true, buildURL: "https://prow/older-compute"}},
	}

	got := latestFailurePerJob(jobs)
	if len(got) != 2 {
		t.Fatalf("got %d jobs, want 2 (failed jobs only, unique by name)", len(got))
	}
	if got[0].jobName != "e2e-compute" || got[0].buildURL != "https://prow/head-fail" {
		t.Fatalf("compute = %+v, want latest failure head-fail", got[0])
	}
	if got[1].jobName != "e2e-storage" || got[1].buildURL != "https://prow/old-fail" {
		t.Fatalf("storage = %+v, want failure even though a later run passed", got[1])
	}
}
