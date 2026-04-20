---
name: analyze-ci-failures
description: >
    Analyze root causes for all the recent ci-failures.
    Use when user mentions to analyze recent ci failures
allowed-tools: [Read, Write, Glob, Grep, Bash(go run:*), Bash(stat:*), Bash(ls:*), Bash(git fetch:*), Bash(git diff:*)]
---

## Overview

This skill helps the user to find the root cause and mitigations for all ci failures aka failing prowjobs that have not reached the testing stage.

## Analysis data generation

As a base for analyzing the failure, the go tool `cmd/ci-failures` from this project is to be used. As a base for the data the file
`output/kubevirt/kubevirt/results.json` is used.

Example how to generate the data files for the prowjobs failing ci:

```bash
$ go run ./cmd/ci-failures generate yaml
```

This produces:
- `output/tmp/errors-{sig}/*.yaml` — per-job error extractions with snippets, line links, and context
- `output/tmp/build-logs/{buildID}.yaml` — full cached build logs with prow URL, GCS URL, timestamps, and complete log content

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

Only needed when error snippets are insufficient for root cause analysis. Structure:
- `prow_job_url`: link to the Prow job
- `build_log_url`: GCS URL for the raw build log
- `log_content`: the complete build log text
- `started` / `finished`: timestamps

These files can be very large. When reading them, start searching from the end since failures appear at the bottom.

## Procedure

1. **Discover error files**: Use Glob to find all `output/tmp/errors-*/*.yaml` files
2. **Read each error YAML file**: Extract the job name, SIG (from directory), and all build errors with their snippets
3. **Analyze error snippets**: For each build error, examine `error_text` and `context` fields to determine root cause. In most cases the snippets contain enough information
4. **Deep-dive when needed**: If snippets are ambiguous, read the full build log from `output/tmp/build-logs/{build_id}.yaml`
5. **Classify** each failure as fixable or non-fixable
6. **Group similar failures** across jobs — errors with the same root cause should be combined into a single group even if they appear in different SIGs or branches

## Goal for analysis

The output is to be looked at as described above, then reasons and possible mitigations are to be deduced.
