---
name: test-failure-rate
description: >
    Analyze test failures from a Prow build against historical flakefinder data.
    Classifies failures as flaky, PR-related, or infrastructure-related using
    success rates, failure dispersion across lanes, and cross-test correlation.
allowed-tools: [Read, Glob, Bash(ci-failures:*), Bash(go install:*), Bash(go env:*), Bash(command -v:*)]
---

## Overview

This skill looks up the historical success rate for each test that failed in a Prow build, using kubevirt flakefinder 7-day (168h) reports. Use `--days` to extend the analysis window up to 28 days by fetching and merging multiple weekly reports.

## Prerequisites

Read [CLI installation and setup](../../../docs/skills-overview.md#cli-installation-and-setup) before running commands. Use the installed `ci-failures` binary from a deliberately selected, writable `ci-health` repository working directory, with a fresh session per analysis as described there. `<output-dir>` below means `output/tmp` or `output/tmp/sessions/<session-id>` according to that setup.

## Data generation

Run the `test-rate` subcommand from the working directory selected in Prerequisites with the Prow job URL.

```bash
ci-failures test-rate <prow-job-url>
```

To cover a longer period (up to 28 days):

```bash
ci-failures test-rate --days 14 <prow-job-url>
```

Example:

```bash
ci-failures test-rate --days 21 https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17690/pull-kubevirt-e2e-k8s-1.35-sig-compute-migrations/2053739485539078144
```

This produces `<output-dir>/test-rate-<job-name>.yaml`.

Read the exact file path from the `"wrote test rate analysis to ..."` log line, not a broad glob over older outputs. Runs for the same job name overwrite that file within a session, even for different builds or windows; read and retain results before rerunning.

Failed test names are extracted from Ginkgo `[FAIL]` lines in build-log error snippets, not from a complete JUnit inventory. If the command reports `no failed tests found in build log`, no test-rate YAML is written; do not read a stale file or infer the build passed. Setup failures, unsupported log formats, and missing snippets may yield no names. Use build analysis for these cases.

`--days` is clamped to 28 (values below 1 default to 7) and selects whole 168h reports, not an exact rolling-day slice. Prefer 7, 14, 21, or 28 days. Inspect the logged fetched reports and `report_period_start`/`report_period_end` for actual coverage, gaps, or overlap before comparing windows; `report_days` records the requested/clamped value, not measured coverage.

## Analysis

After data generation:

1. Read the exact `<output-dir>/test-rate-<job-name>.yaml` path logged by this invocation
2. Read the YAML. The structure is:
   - `prow_job_url`: the analyzed build URL
   - `job_name`: the Prow job name
   - `report_days`: requested/clamped days; whole weekly reports determine actual coverage
   - `report_period_start`: start date of the merged flakefinder window
   - `report_period_end`: end date of the merged flakefinder window
   - `failed_tests`: list of failed tests, each with:
     - `test_name`: original test name from the build log
     - `matched_name`: the matching test name in the flakefinder report (optional; omitted if no match)
     - `total_succeeded`: total successes across all jobs in the covered period
     - `total_failed`: total failures across all jobs
     - `total_skipped`: total skips
     - `success_rate`: percentage (0-100)
     - `severity`: classification of the failure
     - `k8s_versions`: optional per-Kubernetes-version breakdown, omitted when unavailable, sorted by success rate ascending (worst first), each with:
       - `version`: Kubernetes version (e.g., "1.35", "1.36", or "unknown")
       - `succeeded`, `failed`, `skipped`: counts aggregated across all lanes for this version
       - `success_rate`: percentage (0-100)
       - `severity`: classification for this version
       - `lanes`: per-lane breakdown within this version, sorted by success rate ascending, each with:
         - `lane`: job name
         - `succeeded`, `failed`, `skipped`: counts for this lane
         - `success_rate`: percentage (0-100)
         - `severity`: classification for this lane

## Severity classification

- **likely-pr-related** (success rate >= 95%): the test almost always passes, so this failure is probably caused by the PR's changes
- **inconclusive** (success rate >= 80%): moderate flakiness, can't determine if PR-related
- **likely-flaky** (success rate < 80%): the test fails frequently across all jobs, likely a known flake
- **unknown**: the test was not found in the flakefinder report

## Deeper flakiness analysis

The severity classification above is a starting point. Success rate alone can be misleading — a test at 85% could be flaky (failing sporadically everywhere) or broken in one specific lane (failing deterministically there, passing everywhere else). Use these techniques to distinguish the two.

### Failure dispersion across lanes

Count how many lanes show failures for a given test:

- **Concentrated failures** (1–2 lanes account for most failures): likely a lane-specific or k8s-version-specific issue, not a true flake. Check whether those lanes share an environment trait (same k8s version, same provider, Serial tag).
- **Dispersed failures** (failures spread across many lanes): genuine flaky behavior — the test is nondeterministic regardless of environment.

When failures are concentrated, upgrade the severity toward "likely-pr-related" or "lane-specific issue" even if the aggregate success rate is low. When dispersed, treat it as a stronger flake signal even if the aggregate rate looks moderate.

### Periodic lanes as ground truth

**Periodic lanes** (prefixed `periodic-`) run against `main` and reflect the true baseline flakiness of a test. **Presubmit lanes** (prefixed `pull-`) run against PR code and may fail due to faulty changes.

Only draw periodic/presubmit comparisons when those lanes are present in the fetched weekly reports. Absence is not a passing result; use `ci-failures lane-rate` for the relevant TestGrid lane when weekly report coverage is missing.

- Base the flakiness assessment primarily on **periodic lane** success rates
- If a test fails frequently in periodic lanes, it is a known flake regardless of presubmit results
- If a test passes consistently in periodic lanes but fails in the presubmit lane under analysis, the failure is likely PR-related
- Presubmit lane data is supplementary — mention it but don't let it override the periodic signal

### Infrastructure flake detection

When multiple unrelated tests all fail in the **same lane(s)** during the analysis period, suspect an infrastructure problem rather than individual test flakiness:

- Look for lanes where many of the build's failed tests show low success rates simultaneously
- If a lane shows failures across tests that share no code path, that lane likely experienced infra instability (node issues, storage problems, network flakes)
- Flag these as "infra flake" and note the affected lane — the individual tests are not at fault

### Windowed trend analysis

Use `--days` to compare different time windows and detect trends:

Compare actual report bounds and fetched weekly coverage first. Requested windows can round up to whole weeks and may overlap; do not label them exact, independent rolling windows.

- **7 days vs 21 days**: if the 7-day rate is much worse than the 21-day rate, the flakiness is recent (possible regression). If similar, it's a long-standing flake.
- **Recommend quarantine** when a test shows sustained flakiness (consistent across windows) with dispersed failures

## Presenting results

### What to report

For each failed test, report:
1. The test name
2. The overall success rate and total runs (succeeded + failed)
3. The severity classification
4. The **failure pattern** — concentrated in specific lanes/versions or dispersed across many (this often matters more than the raw success rate)
5. If severity is "unknown", note that the test may be new or not covered by flakefinder
6. Per-k8s-version success rates, highlighting versions where the test is notably worse
7. Per-lane details, **leading with periodic lanes** — call out when a specific periodic lane drives the flakiness
8. Note when periodic and presubmit lanes diverge (e.g. periodic passes 100% but pull lane shows failures)
9. When multiple tests share the same failing lane pattern, group them and flag as a probable infrastructure flake

### Summary table

After the per-test details, provide a summary table:

| Test | Success Rate | Runs | Pattern | Assessment |
|------|-------------|------|---------|------------|

Where **Pattern** is one of:
- **dispersed** — failures spread across many lanes (true flake)
- **concentrated** — failures in 1–2 lanes (lane-specific or version-specific)
- **infra-correlated** — fails alongside other unrelated tests in same lane(s)
- **unknown** — not found in flakefinder

And **Assessment** refines the raw severity by incorporating the pattern analysis.

## Suggested follow-up skills

After the failure rate summary, suggest follow-up analyses when appropriate:
- **change-relevance**: for failures classified as "likely-pr-related" — confirms whether the PR actually changed files in the test's code area before blaming the PR
- **lane-failure-rate**: for failures concentrated in specific lanes — shows whether those lanes are generally unhealthy or if the concentration is specific to this test
- **analyze-build**: for infra-correlated failures or when the k8s analysis would help — deep-dives into cluster state (pods, nodes, events, etcd) to find infrastructure root causes
