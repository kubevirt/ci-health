---
name: analyze-pr-builds
description: >
    Analyze all failed builds for a GitHub pull request.
    Use when user provides a GitHub PR URL and wants to understand why builds are failing.
allowed-tools: [Read, Glob, Bash(go run:*), Bash(grep:*)]
---

## Overview

This skill analyzes all currently-failing builds for a GitHub pull request, identifies root causes, and presents a concise summary.

It works for any repo served by prow.ci.kubevirt.io (currently the `kubevirt` GitHub org).

## Data Generation

Run the `analyze-pr` command with the GitHub PR URL to fetch and analyze all failed builds:

```bash
$ go run ./cmd/ci-failures analyze-pr https://github.com/kubevirt/kubevirt/pull/17287
```

This command:
1. Fetches the PR history from prow
2. Finds all jobs for the latest commit
3. Keeps only the latest run per job (jobs that eventually passed are excluded)
4. Downloads build logs and extracts error snippets for each failure
5. Writes one YAML file per failed job to `output/tmp/{job-name}.yaml`
6. For each failed job, also downloads k8s-reporter artifacts (and etcd profiler data if available) and writes `output/tmp/k8s-analysis-{build-id}.yaml`

## Analysis

After data generation:

1. Use Glob to find all `output/tmp/*.yaml` files produced by the command
2. Read each build log YAML (`{job-name}.yaml`). The structure is:
   - `job_name`: the Prow job name
   - `build_errors`: list of build errors, each containing:
     - `job_url`: link to the Prow job UI
     - `build_id`: the build number
     - `started` / `finished`: timestamps
     - `category`: error category (external, internal, pr-build, needs-investigation)
     - `category_reason`: explanation of the categorization
     - `build_log_error_snippets`: list of error matches with `error_text`, `context`, and `link_to_log_line`
3. Read each k8s analysis YAML (`k8s-analysis-*.yaml`) if present. The structure is:
   - `snapshots`: list of cluster state capture points (process + failure count)
   - `findings`: list of detected issues with `detector`, `severity`, `kind`, `name`, `reason`, `message`, and `snapshot`
   - `summary`: aggregate counts by kind, severity, and detector
   - `etcd_profile` (optional): full etcd storage profile with peak/final tmpfs and WAL metrics, plus per-spec records. Etcd-related findings use kind `EtcdProfile` and detectors `etcd-tmpfs-exhaustion`, `etcd-tmpfs-pressure`, `etcd-large-wal`, `etcd-db-growth`
   - `container_log_files` (optional): list of cached kubevirt component container logs (virt-controller, virt-handler, virt-api, virt-operator) with `pod_name`, `container_name`, `namespace`, `cached_path`, and `size_bytes`
4. **Cross-reference failing VMI/VM names against container logs**: extract VMI/VM names from build log errors and k8s findings, then use `grep` on the cached virt-controller and virt-handler logs to find the actual rejection reason. Example: `grep -i "testvmi-xxxxx" <cached_path>`
5. For each failure, examine build log errors, k8s findings, AND container log matches to determine the root cause
6. Correlate findings across all sources — e.g. CrashLoopBackOff on a component pod often explains test timeouts; etcd tmpfs exhaustion can explain apiserver failures; virt-controller log errors (schema validation, sync failures) reveal why VMIs are stuck in non-Running phases
7. Group failures with the same root cause together

## Cross-job failure correlation

Before presenting results, analyze how failures correlate across jobs to distinguish PR-related issues from flakes and infra events:

### Same test failing across multiple jobs
When the same test fails in multiple different jobs for this PR, it strengthens the "PR-related" signal — the PR likely broke something that is exercised across environments. This is especially strong when the test passes consistently in periodic lanes but fails in all of this PR's presubmit runs.

### Different tests failing in different jobs with no overlap
When each job has different failures with no common tests, suspect independent flakes or separate infra events rather than a PR-caused issue. No single code change typically causes unrelated tests to fail in unrelated ways.

### Multiple unrelated tests failing in the same job
When many unrelated tests fail together in a single job but pass in other jobs, suspect an infrastructure flake in that specific run — node issues, network blips, storage problems. Check the k8s analysis for that build for corroborating evidence (NotReady nodes, CrashLoopBackOff, etcd issues).

### Consistent failure across retries
If the PR has been retried and the same tests fail every time, it's almost certainly PR-related. If different tests fail on each retry, it's flaky infrastructure.

## Output

Present a concise summary to the user:
- Group failures by root cause
- For each group: state the root cause, list affected jobs, and include one representative error snippet
- Include relevant k8s findings that explain or contribute to the failure
- Note the **cross-job pattern** for each group: appears in all jobs (likely PR-caused), appears in one job only (possible flake), or correlates with infra findings (infra flake)
- Classify each as **fixable** (CI config or code issue) or **non-fixable** (external/infra)
- Include the `link_to_log_line` for the most relevant error in each group

## Suggested follow-up skills

After the root-cause summary, suggest follow-up analyses when appropriate:
- **test-failure-rate**: for failures that look flaky — confirms whether the test has a history of flakiness across all lanes and k8s versions
- **change-relevance**: for failures that look PR-related — confirms whether the PR actually changed files in the test's code area
- **lane-failure-rate**: for failures concentrated in one job/lane — shows whether that lane is generally unhealthy
