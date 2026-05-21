---
name: lane-failure-rate
description: >
    Analyze failure rates and flakiness patterns for all tests in a testgrid lane.
    Uses flip-rate analysis, burst detection, and cross-test correlation in addition
    to success rates. Use when the user provides a testgrid URL and wants to see
    which tests are failing in that lane.
allowed-tools: [Read, Glob, Bash(go run:*)]
---

## Overview

This skill analyzes a testgrid lane and calculates failure rates for every test that has at least one failure in the requested time window. Unlike test-failure-rate (which starts from a single build's failed tests and looks them up in flakefinder), this skill operates at the lane level — it fetches all test results directly from testgrid and computes rates from raw pass/fail data.

## Data generation

Run the `lane-rate` subcommand from the `kubevirt.io/ci-health` repository with a testgrid URL.

If not already inside the `ci-health` repository, `cd` into it first (e.g. `cd kubevirt.io/ci-health` from the `github.com` directory).

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

## Deeper flakiness analysis

The severity classification above is based on aggregate success rate. Use the `recent_failures` data to perform richer pattern analysis that better distinguishes genuine flakes from deterministic breakage and infrastructure events.

### Flip-rate analysis

Examine the `recent_failures` list to detect state transitions (pass→fail→pass patterns):

- **High flip rate** (failures interleaved with passes, spread across the window): the test is nondeterministic — a genuine flake
- **Low flip rate** (failures clustered consecutively): the test was broken for a period then fixed, or is still broken — not a flake but a regression
- **Key heuristic**: 3 or more consecutive failures should be treated as deterministic breakage, not flakiness (per Google testing research). When all `recent_failures` timestamps are consecutive builds, the test has a real bug regardless of what the aggregate success rate says

A test with 75% success rate could be either: flaky (alternating pass/fail across 40 runs) or broken-then-fixed (failed 10 runs in a row then passed 30). The flip rate distinguishes the two — same success rate, completely different diagnosis.

### Burst detection

Examine the timestamps in `recent_failures`:

- **Clustered failures** (many failures in a short time window, e.g. several within the same day): likely an infrastructure event, a temporary regression, or a bad merge that was later reverted
- **Spread failures** (failures distributed evenly across the analysis window): genuine flaky behavior — the test fails randomly regardless of external factors
- **Single burst then clean**: a transient event — the test is healthy now despite the lower success rate

### Cross-test infrastructure flake detection

Compare `recent_failures` timestamps across different tests in the results:

- When multiple unrelated tests share the same failure timestamps (same `build_id`s), they failed together in the same build — this is an infrastructure flake signal, not individual test problems
- Group tests that share failure timestamps and flag them as **infra-correlated**
- Common infra patterns: node instability, storage backend issues, registry pull failures, network flakes — these cause unrelated tests to fail simultaneously
- If a small number of builds account for most failures across many tests, the lane experienced infra instability during those builds

## Presenting results

1. Lead with a summary: total builds analyzed, number of tests with failures
2. **Flag infra-correlated failures first**: if multiple tests share failure timestamps, group them and note the affected build IDs — these are lane-level problems, not test-level
3. Group remaining tests by severity, starting with the worst (likely-flaky)
4. For each test, report:
   - Test name, success rate, total runs, and severity
   - **Failure pattern**: flaky (high flip rate, spread), burst (clustered), consecutive (deterministic), or infra-correlated
   - Most recent failure messages for tests with low success rates
5. Note the total builds count to contextualize rates — a test that failed 1/2 is different from 1/42
6. Tests tagged `[QUARANTINE]` are already known flakes; note this when presenting

### Summary table

After per-test details, provide a summary:

| Test | Rate | Runs | Pattern | Assessment |
|------|------|------|---------|------------|

Where **Pattern** is one of:
- **flaky** — high flip rate, failures spread across time (genuine flake)
- **burst** — failures clustered in a short time window (transient event)
- **consecutive** — 3+ failures in a row (deterministic breakage)
- **infra-correlated** — fails alongside other unrelated tests in same builds

## Combining with test-failure-rate

The `lane-rate` command gives a lane-wide overview, while `test-failure-rate` gives a cross-lane view for specific tests. Use `lane-rate` to find which tests are problematic in a lane, then use `test-failure-rate` on a specific build to see if those tests also fail in other lanes.
