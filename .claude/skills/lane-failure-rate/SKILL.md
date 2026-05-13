---
name: lane-failure-rate
description: >
    Show failure rates for all tests in a testgrid lane over the last N days.
    Use when the user provides a testgrid URL and wants to see which tests are failing in that lane.
allowed-tools: [Read, Glob, Bash(go run:*)]
---

## Overview

This skill analyzes a testgrid lane and calculates failure rates for every test that has at least one failure in the requested time window. Unlike test-failure-rate (which starts from a single build's failed tests and looks them up in flakefinder), this skill operates at the lane level — it fetches all test results directly from testgrid and computes rates from raw pass/fail data.

## Data generation

Run the `lane-rate` subcommand with a testgrid URL:

```bash
$ go run ./cmd/ci-failures lane-rate <testgrid-url>
```

To control the analysis window (default 14 days):

```bash
$ go run ./cmd/ci-failures lane-rate --days 21 <testgrid-url>
```

Example:

```bash
$ go run ./cmd/ci-failures lane-rate --days 14 "https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.36-sig-storage"
```

This produces `output/tmp/lane-rate-{tab-name}.yaml`.

## Analysis

After data generation:

1. Use Glob to find the `output/tmp/lane-rate-*.yaml` file
2. Read the YAML. The structure is:
   - `testgrid_url`: the original testgrid URL analyzed
   - `dashboard`: the testgrid dashboard name (e.g., "kubevirt-periodics")
   - `tab`: the testgrid tab/lane name (e.g., "periodic-kubevirt-e2e-k8s-1.36-sig-storage")
   - `report_days`: number of days covered
   - `report_period_start`: start of the analysis window (ISO 8601)
   - `report_period_end`: end of the analysis window (ISO 8601)
   - `total_builds`: number of builds in the analysis window
   - `failed_tests`: list of tests with at least one failure, sorted by success rate ascending (worst first), each with:
     - `test_name`: full test name from testgrid
     - `succeeded`: number of passing runs
     - `failed`: number of failing runs (includes flaky runs)
     - `skipped`: number of skipped/not-run entries
     - `total_runs`: succeeded + failed (excludes skipped)
     - `success_rate`: percentage (0-100)
     - `severity`: classification of the failure pattern
     - `recent_failures`: up to 10 most recent failures, each with:
       - `build_id`: the build/column identifier
       - `timestamp`: ISO 8601 timestamp of the build
       - `message`: failure message from testgrid (may be empty)

## Severity classification

- **likely-pr-related** (success rate >= 95%): rare failure, likely a one-off
- **inconclusive** (success rate >= 80%): moderate failure rate, needs investigation
- **likely-flaky** (success rate < 80%): frequent failures, likely a known flake
- **unknown**: no pass+fail data available

## Presenting results

1. Lead with a summary: total builds analyzed, number of tests with failures
2. Group tests by severity, starting with the worst (likely-flaky)
3. For each test, report the test name, success rate, total runs, and severity
4. For tests with low success rates, include the most recent failure messages to help diagnose the issue
5. Note the total builds count to contextualize rates — a test that failed 1/2 is different from 1/42
6. Tests tagged `[QUARANTINE]` are already known flakes; note this when presenting

## Combining with test-failure-rate

The `lane-rate` command gives a lane-wide overview, while `test-failure-rate` gives a cross-lane view for specific tests. Use `lane-rate` to find which tests are problematic in a lane, then use `test-failure-rate` on a specific build to see if those tests also fail in other lanes.
