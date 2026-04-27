---
name: analyze-build
description: >
    Analyze the root cause inside a failed prowjob.
    Use when user mentions to analyze a prowjob failure or adds a prowjob url
allowed-tools: [Read, Glob, Bash(go run:*)]
---

## Overview

This skill helps the user to find the root cause and mitigations for failing prowjobs.

It combines two analyses:
1. **Build log analysis** — extracts error snippets from the build log and categorizes them
2. **Kubernetes cluster state analysis** — downloads k8s-reporter artifacts (pods, nodes, events, VMIs) and detects infrastructure-level failures like CrashLoopBackOff, OOMKilled, NotReady nodes, failed migrations, and warning events

## Analysis data generation

As a base for analyzing the failure, the go tool `cmd/ci-failures` from this project is to be used.

Example how to generate the data files for a specific prowjob, given is the url to the prowjob build:

```bash
$ go run ./cmd/ci-failures analyze-build https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16885/pull-kubevirt-e2e-k8s-1.34-windows2016/2034979182789791744
```

This produces two output files:
- `output/tmp/{job-name}.yaml` — build log error analysis
- `output/tmp/k8s-analysis-{build-id}.yaml` — kubernetes cluster state analysis (if k8s-reporter artifacts exist)

## Analysis

After data generation:

1. Use Glob to find all `output/tmp/*.yaml` files produced by the command
2. Read the build log analysis YAML (`{job-name}.yaml`). The structure is:
   - `job_name`: the Prow job name
   - `build_errors`: list of build errors with `category`, `category_reason`, and `build_log_error_snippets`
3. Read the k8s analysis YAML (`k8s-analysis-*.yaml`) if present. The structure is:
   - `snapshots`: list of cluster state capture points (process + failure count)
   - `findings`: list of detected issues, each with `detector`, `severity`, `kind`, `name`, `reason`, `message`, and `snapshot`
   - `summary`: aggregate counts by kind, severity, and detector
4. Correlate both analyses: k8s findings (e.g. CrashLoopBackOff on a component pod) often explain build log errors (e.g. test timeouts)
5. Deduce the root cause and possible mitigations from the combined data
