---
name: test-failure-rate
description: >
    Show historical success rate for tests that failed in a build.
    Use when the user wants to know if a test failure is flaky or PR-related.
allowed-tools: [Read, Glob, Bash(go run:*)]
---

## Overview

This skill looks up the historical success rate for each test that failed in a Prow build, using kubevirt flakefinder 7-day (168h) reports. Use `--days` to extend the analysis window up to 28 days by fetching and merging multiple weekly reports.

## Data generation

Run the `test-rate` subcommand with the Prow job URL:

```bash
$ go run ./cmd/ci-failures test-rate <prow-job-url>
```

To cover a longer period (up to 28 days):

```bash
$ go run ./cmd/ci-failures test-rate --days 14 <prow-job-url>
```

Example:

```bash
$ go run ./cmd/ci-failures test-rate --days 21 https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17690/pull-kubevirt-e2e-k8s-1.35-sig-compute-migrations/2053739485539078144
```

This produces `output/tmp/test-rate-{job-name}.yaml`.

## Analysis

After data generation:

1. Use Glob to find the `output/tmp/test-rate-*.yaml` file
2. Read the YAML. The structure is:
   - `prow_job_url`: the analyzed build URL
   - `job_name`: the Prow job name
   - `report_days`: number of days covered
   - `report_period_start`: start date of the merged flakefinder window
   - `report_period_end`: end date of the merged flakefinder window
   - `failed_tests`: list of failed tests, each with:
     - `test_name`: original test name from the build log
     - `matched_name`: the matching test name in the flakefinder report (empty if no match)
     - `total_succeeded`: total successes across all jobs in the covered period
     - `total_failed`: total failures across all jobs
     - `total_skipped`: total skips
     - `success_rate`: percentage (0-100)
     - `severity`: classification of the failure

## Severity classification

- **likely-pr-related** (success rate >= 95%): the test almost always passes, so this failure is probably caused by the PR's changes
- **inconclusive** (success rate >= 80%): moderate flakiness, can't determine if PR-related
- **likely-flaky** (success rate < 80%): the test fails frequently across all jobs, likely a known flake
- **unknown**: the test was not found in the flakefinder report

## Presenting results

For each failed test, report:
1. The test name
2. The success rate and total runs (succeeded + failed)
3. The severity classification
4. If severity is "unknown", note that the test may be new or not covered by flakefinder
