---
name: analyze-build
description: >
    Analyze the root cause inside a failed prowjob.
    Use when user mentions to analyze a prowjob failure or adds a prowjob url
allowed-tools: [Read, Glob, Bash(go run:*), Bash(grep:*)]
---

## Overview

This skill helps the user to find the root cause and mitigations for failing prowjobs.

It combines four analyses:
1. **Build log analysis** — extracts error snippets from the build log and categorizes them
2. **Kubernetes cluster state analysis** — downloads k8s-reporter artifacts (pods, nodes, events, VMIs) and detects infrastructure-level failures like CrashLoopBackOff, OOMKilled, NotReady nodes, failed migrations, and warning events
3. **Etcd storage profiling** — if the job was run with `KUBEVIRT_PROFILE_ETCD=true`, downloads `etcd-storage-profile.json` from `artifacts/etcd-profiler/` and detects tmpfs exhaustion, tmpfs pressure, large WAL files, and DB growth trends
4. **Container log analysis** — downloads virt-controller, virt-handler, virt-api, and virt-operator container logs from k8s-reporter suite artifacts for interactive grep

## Analysis data generation

As a base for analyzing the failure, the go tool `cmd/ci-failures` from the `kubevirt.io/ci-health` repository is to be used.

If not already inside the `ci-health` repository, `cd` into it first (e.g. `cd kubevirt.io/ci-health` from the `github.com` directory).

Example how to generate the data files for a specific prowjob, given is the url to the prowjob build:

```bash
$ go run ./cmd/ci-failures analyze-build https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16885/pull-kubevirt-e2e-k8s-1.34-windows2016/2034979182789791744
```

This produces up to two output files in a session directory:
- `output/tmp/sessions/<session-id>/{job-name}.yaml` — build log error analysis
- `output/tmp/sessions/<session-id>/k8s-analysis-{build-id}.yaml` — kubernetes cluster state analysis (if k8s-reporter artifacts exist), includes etcd profiler findings and full etcd profile data when available

The tool logs each written file path. Extract the session directory from the first `"wrote analysis to ..."` or `"wrote k8s analysis to ..."` log line (e.g. `output/tmp/sessions/abc123/`).

## Analysis

### Infrastructure flake precheck

Before diving into individual test failures, assess whether the entire build was an infrastructure flake:

1. Count the number of distinct failed tests in the build log
2. Check the k8s analysis (if present) for infrastructure-level problems: NotReady nodes, CrashLoopBackOff on system pods (virt-controller, virt-handler, virt-api, virt-operator), etcd issues (tmpfs exhaustion, large WAL)
3. If **5 or more unrelated tests** failed AND **infrastructure findings** are present, classify the build as an **infrastructure flake** — the individual test failures are symptoms, not causes. Report the infra root cause and skip per-test analysis.
4. If only a few tests failed or no infra findings exist, proceed with per-test analysis below.

This avoids wasting time analyzing each test individually when the real problem is a bad node or a crashed component.

### Quick flake check

Consider suggesting the `test-failure-rate` skill on this build URL as a quick pre-filter. If all failing tests are already known flakes (success rate < 80% across lanes), the build failure is likely not actionable and the user can retrigger instead of investigating.

### Detailed analysis

After data generation:

1. Extract the session directory path from the tool's log output, then list all `*.yaml` files in that directory
2. Read the build log analysis YAML (`{job-name}.yaml`). The structure is:
   - `job_name`: the Prow job name
   - `build_errors`: list of build errors with `category`, `category_reason`, and `build_log_error_snippets`
3. Read the k8s analysis YAML (`k8s-analysis-*.yaml`) if present. The structure is:
   - `snapshots`: list of cluster state capture points (process + failure count)
   - `findings`: list of detected issues, each with `detector`, `severity`, `kind`, `name`, `reason`, `message`, and `snapshot`
   - `summary`: aggregate counts by kind, severity, and detector
   - `etcd_profile` (optional): full etcd storage profile with `peak_tmpfs_used_bytes`, `final_tmpfs_total_bytes`, `final_wal_size_bytes`, and per-spec `records` showing DB size deltas and tmpfs usage. Etcd-related findings use kind `EtcdProfile` and detectors `etcd-tmpfs-exhaustion`, `etcd-tmpfs-pressure`, `etcd-large-wal`, `etcd-db-growth`
4. Read the container log files list from `container_log_files` in the k8s analysis YAML. Each entry has:
   - `pod_name`, `container_name`, `namespace`: identify the component
   - `cached_path`: local file path to grep
   - `size_bytes`: file size
5. **Cross-reference failing VMI/VM names against container logs**: extract VMI/VM names from build log errors (e.g. `testvmi-xxxxx`) and k8s findings (e.g. VMIs stuck in Scheduling), then use `grep` on the cached virt-controller and virt-handler logs to find why those VMIs failed. Common patterns to look for:
   - Schema validation errors (e.g. `admission webhook` or `validation failed`)
   - Sync failures (e.g. `syncError` or `failed to sync`)
   - Device/resource issues (e.g. `device not found`, `resource not available`)
   - Example: `grep -i "testvmi-xxxxx" <cached_path>` to find all log lines related to a specific VMI
6. Correlate all analyses: k8s findings (e.g. CrashLoopBackOff on a component pod) often explain build log errors (e.g. test timeouts); etcd findings (e.g. tmpfs exhaustion) can explain node-level or apiserver issues; container log errors (e.g. schema validation failures in virt-controller) reveal the actual rejection reason when VMIs are stuck in non-Running phases
7. Deduce the root cause and possible mitigations from the combined data
