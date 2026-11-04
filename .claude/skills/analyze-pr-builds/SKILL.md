---
name: analyze-pr-builds
description: >
    Analyze all failed builds for a GitHub pull request.
    Use when user provides a GitHub PR URL and wants to understand why builds are failing.
allowed-tools: [Read, Bash(ci-failures:*), Bash(go install:*), Bash(go env:*), Bash(command -v:*), Bash(grep:*)]
---

## Overview

This skill analyzes all currently-failing builds for a GitHub pull request, identifies root causes, and presents a concise summary.

It works for any repo served by prow.ci.kubevirt.io (currently the `kubevirt` GitHub org).

## Data Generation

First read [CLI installation and setup](../../../docs/skills-overview.md#cli-installation-and-setup).
Use the installed binary from a writable ci-health checkout root. Choose a fresh
session ID for this analysis and pass it to both commands below. The example ID
is illustrative; do not reuse an existing session containing another analysis.

Two commands to run sequentially:

### Step 1: Fetch and analyze all failed builds

```bash
ci-failures --session-id pr-17287-review-1 analyze-pr https://github.com/kubevirt/kubevirt/pull/17287
```

This downloads build logs, k8s-reporter artifacts, and container logs for every currently-failing job in the PR. Output files go to the session directory.

### Step 2: Get a condensed summary

```bash
ci-failures --session-id pr-17287-review-1 summarize-session
```

This reads all session YAML files and emits a **single condensed YAML** to stdout
containing key fields for initial analysis. Read the individual analysis files
when snippet context, timestamps, snapshots, or the full etcd profile are needed;
these details are not included in the condensed summary. If no failed builds
were found, report that outcome rather than reading old session files.

## Analysis

After running both commands:

1. Read the `summarize-session` stdout output. The structure is:
   - `jobs`: list of failed jobs, each with:
     - `job_name`: the Prow job name
     - `job_url`: link to the Prow job UI
     - `build_id`: the build number
     - `result`: the Prow build result from `finished.json` — one of `SUCCESS`, `FAILURE`, `ABORTED`, `ERROR`
     - `passed`: boolean from `finished.json`
     - `category`: error category (external, internal, pr-build, needs-investigation, prow-aborted, prow-error)
     - `category_reason`: explanation of the categorization
     - `error_snippets`: list of `{error_text, link}` — one-line error text and link to the log line
    - `k8s_analyses` (optional): list of k8s analysis results per build, each with:
     - `build_id`, `job_name`, `total_findings`
      - `findings` (optional): list of `{detector, severity, kind, namespace, name, reason, message}`
      - `container_log_files` (optional): list of `{pod_name, container_name, namespace, cached_path, size_bytes}` for grep cross-referencing
2. **Cross-reference failing VMI/VM names against container logs**: extract VMI/VM names from error snippets and k8s findings, then use `grep` on the cached virt-controller and virt-handler logs to find the actual rejection reason. Example: `grep -i "testvmi-xxxxx" <cached_path>`
3. Correlate findings across all sources — e.g. CrashLoopBackOff on a component pod often explains test timeouts; etcd tmpfs exhaustion can explain apiserver failures; virt-controller log errors (schema validation, sync failures) reveal why VMIs are stuck in non-Running phases
4. Group failures with the same root cause together

### Handling aborted and error builds

Jobs with `category: prow-aborted` were killed by Prow before completing (e.g., superseded by a newer commit). Jobs with `category: prow-error` failed to schedule. In both cases:
- Report them separately as "aborted" or "scheduling error" — not as test failures
- **Exclude them from cross-job failure correlation** below — they inflate failure counts without providing test-level signal
- Suggest retriggering if the latest commit has no run for that job

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
- Include the `link` for the most relevant error in each group

## Suggested follow-up skills

After the root-cause summary, suggest follow-up analyses when appropriate:
- **test-failure-rate**: for failures that look flaky — confirms whether the test has a history of flakiness across all lanes and k8s versions
- **change-relevance**: for failures that look PR-related — confirms whether the PR actually changed files in the test's code area
- **lane-failure-rate**: for failures concentrated in one job/lane — shows whether that lane is generally unhealthy
