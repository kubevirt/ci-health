---
name: analyze-ci-failures
description: >
    Analyze root causes for all the recent ci-failures.
    Use when user mentions to analyze recent ci failures
allowed-tools: [Read, Write, Grep, Bash(go run:*), Bash(stat:*), Bash(ls:*), Bash(git fetch:*), Bash(git diff:*)]
---

## Overview

This skill helps the user to find the root cause and mitigations for all ci failures aka failing prowjobs that have not reached the testing stage.

## Analysis data generation

As a base for analyzing the failure, the go tool `cmd/ci-failures` from the `kubevirt.io/ci-health` repository is to be used. As a base for the data the file
`output/kubevirt/kubevirt/results.json` is used.

If not already inside the `ci-health` repository, `cd` into it first (e.g. `cd kubevirt.io/ci-health` from the `github.com` directory).

Example how to generate the data files for the prowjobs failing ci:

```bash
$ go run ./cmd/ci-failures generate yaml
```

This produces:
- `output/tmp/errors-{sig}/*.yaml` — per-job error extractions with snippets, line links, and context
- `output/tmp/build-logs/{buildID}.yaml` — full cached build logs with prow URL, GCS URL, timestamps, and complete log content

At the end of execution, the command prints a YAML list of all generated error files to stdout:
```
generated_files:
  - output/tmp/errors-sig-compute/pull-kubevirt-e2e-k8s-1.35-sig-compute.yaml
  - output/tmp/errors-sig-network/pull-kubevirt-e2e-k8s-1.35-sig-network.yaml
  ...
```
Parse this list to know exactly which files to read — no separate file discovery step (find/glob) is needed.

### Stale data
Before you proceed, 
1) check whether there's updates on `output/kubevirt/kubevirt/results.json` available on the main branch - in that case first update the current branch with latest changes by rebasing on main.
2) compare `output/tmp` file time stamps with `output/kubevirt/kubevirt/results.json` - if the files in `output/tmp` are older than the latter, they are stale and need to get deleted.

To compare timestamps, use `ls -lt` to list both files together — the newer file appears first:
```bash
$ ls -lt output/kubevirt/kubevirt/results.json output/tmp/errors-*/*.yaml
```
If `results.json` appears above (newer than) the error YAML files, the generated data is stale.

To delete stale data, remove the entire `output/tmp` directory:
```bash
$ rm -rf output/tmp/*
```

## Data Available in the YAML Files

### Error YAML files (`output/tmp/errors-{sig}/*.yaml`)

Each file contains structured data — most metadata can be read directly without parsing URLs:

- **SIG**: encoded in the parent directory name (e.g., `errors-sig-compute/`, `errors-sig-network/`, `errors-sig-storage/`, `errors-sig-monitoring/`, `errors-sig-performance/`)
- **Job name** (`job_name` field): the Prow job name, e.g., `pull-kubevirt-e2e-k8s-1.35-sig-network`
- **Branch**: deducible from `job_name` — if it ends with a version suffix like `-1.6`, it targets `release-1.6`; no suffix means `main`
- **Kubernetes version**: deducible from `job_name` — the `k8s-X.Y` or `kind-X.Y` segment gives the K8s version (e.g., `1.35`)
- **PR number**: deducible from `job_url` — the segment after `pull/kubevirt_kubevirt/` (e.g., `17129`)
- **Build errors** (`build_errors` list), each containing:
    - `job_url`: direct link to the Prow job UI
    - `build_id`: the Prow build number (also usable to find the full build log)
    - `started` / `finished`: timestamps
    - `build_log_error_snippets`: list of extracted error matches, each with:
        - `error_text`: the matched error line
        - `context`: surrounding lines (typically 3 before and after) — this is the primary source for root cause analysis
        - `link_to_log_line`: direct link to the error line in Prow UI
        - `start_line`, `error_line`, `end_line`: line numbers in the build log

### Full build log files (`output/tmp/build-logs/{buildID}.yaml`)

Only needed when error snippets are insufficient for root cause analysis. Use the `show-log` command to read decoded build logs instead of reading the YAML files directly (they contain base64-encoded binary content):

```bash
# Show last 30 lines of a build log
$ go run ./cmd/ci-failures show-log <build-id> --tail 30

# Search for error patterns (case-insensitive)
$ go run ./cmd/ci-failures show-log <build-id> --grep "error" --tail 20
```

Flags:
- `--tail N` — print only the last N lines (default: all)
- `--grep pattern` — filter lines containing the pattern (case-insensitive); output includes line numbers

## Procedure

1. **Use generated file list**: Parse the `generated_files:` YAML list printed by the `generate yaml` command to get all error YAML file paths — no separate discovery step needed
2. **Read each error YAML file**: Extract the job name, SIG (from directory), and all build errors with their snippets
3. **Analyze error snippets**: For each build error, examine `error_text` and `context` fields to determine root cause. In most cases the snippets contain enough information
4. **Deep-dive when needed**: If snippets are ambiguous or empty, use `go run ./cmd/ci-failures show-log <build_id> --grep "error" --tail 20` to search the full build log
5. **Classify** each failure as fixable or non-fixable
6. **Group similar failures** across jobs — errors with the same root cause should be combined into a single group even if they appear in different SIGs or branches

## Flakiness and infrastructure event detection

After grouping failures by root cause, apply these additional analyses:

### Cross-PR flake detection

When the same test fails across multiple *different* PRs in the data, it is almost certainly a systemic flake, not caused by any individual PR. To identify these:

1. After reading all error YAML files, collect failing test names across all jobs and PRs (extract PR number from `job_url`)
2. Tests that appear in failures from 3+ different PRs are **quarantine candidates** — flag them prominently
3. For these tests, the mitigation is quarantining the test, not fixing individual PRs

### Infrastructure event detection

When many errors from the same time window share infrastructure signatures, group them as a single infra event rather than analyzing each individually:

1. Compare `started`/`finished` timestamps across build errors from different jobs
2. If multiple unrelated jobs (different SIGs, different PRs) all failed within the same **2-4 hour window** with similar error patterns (e.g., registry pull failures, DNS timeouts, node provisioning failures), they likely hit the same infra event
3. Group these as **"infra event at {time range}"** and describe the common infrastructure signature
4. Common infra signatures to look for:
   - Image pull errors (`ErrImagePull`, `ImagePullBackOff`) — registry or network issue
   - Node provisioning failures — cloud provider issue
   - DNS resolution failures — cluster DNS issue
   - Timeout errors across many unrelated tests — general cluster instability
   - `context deadline exceeded` in cluster setup — provisioning infra overloaded

### Prioritization

When presenting results, prioritize by actionability:
1. **Cross-PR flakes** (quarantine candidates) — highest signal, immediately actionable
2. **Infrastructure events** — not fixable by code changes, useful for incident tracking
3. **Individual fixable failures** — specific to a PR or code change
4. **Non-fixable/external failures** — lowest priority

## Goal for analysis

The output is to be looked at as described above, then reasons and possible mitigations are to be deduced.
