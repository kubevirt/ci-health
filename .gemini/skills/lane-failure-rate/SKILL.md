---
name: lane-failure-rate
description: Analyze failure rates, flip-rates, and flakiness patterns for all tests in a testgrid lane.
argument-hint: <testgrid-url>
allowed-tools: Read, Glob, Bash(go run *), Bash(gh *)
---

## Overview

This skill analyzes a testgrid lane and calculates failure rates for every test that has at least one failure in the requested time window. Unlike test-failure-rate (which starts from a single build's failed tests and looks them up in flakefinder), this skill operates at the lane level — it fetches all test results directly from testgrid and computes rates from raw pass/fail data.

## Lane discovery

When the user asks to analyze lanes for a Kubernetes version (e.g. "kubevirt 1.36 lanes") without providing specific testgrid URLs, use the `discover-lanes` subcommand to find all matching lanes across both periodic and presubmit dashboards.

```bash
$ go run ./cmd/ci-failures discover-lanes $ARGUMENTS
```

Example:

```bash
$ go run ./cmd/ci-failures discover-lanes $ARGUMENTS
```

This queries the testgrid summary API for both `kubevirt-periodics` and `kubevirt-presubmits` dashboards and produces `output/tmp/discover-lanes-{version}.yaml`.

### Discovery output YAML structure

```yaml
version_filter: "1.36"
dashboards:
  - name: kubevirt-periodics
    lanes:
      - tab: periodic-kubevirt-e2e-k8s-1.36-sig-compute
        overall_status: FAILING
        testgrid_url: "https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.36-sig-compute"
  - name: kubevirt-presubmits
    lanes:
      - tab: pull-kubevirt-e2e-k8s-1.36-sig-compute
        overall_status: FLAKY
        testgrid_url: "https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.36-sig-compute"
```

Note: some SIG lanes only exist on one dashboard (e.g. `sig-compute-serial` is presubmit-only).

### Specialized presubmit lanes

Presubmits often include specialized lanes beyond the standard SIG lanes — for specific architectures (e.g. `sig-compute-arm64`), subsets (e.g. `sig-compute-windows`), or serial execution (`sig-compute-serial`). These lanes typically do **not** have periodic counterparts. However, the tests they run usually also exist in the standard SIG lane of the same type (e.g. an arm64-specific test failure may also appear in the regular `sig-compute` periodic lane). When cross-referencing failures, check the standard SIG lane for the same test name even if the specialized lane has no periodic equivalent.

### Workflow for multi-lane analysis

1. Run `discover-lanes` to find all matching lanes
2. Read the discovery YAML to get testgrid URLs
3. Run `lane-rate` for each discovered lane (can run in parallel), using the `testgrid_url` from the discovery output
4. Analyze and present results across all lanes, noting cross-lane patterns

### Interpreting periodic vs presubmit data

- **Periodic lanes** run on a fixed schedule against the latest merged code. Failures here indicate problems in mainline — they affect everyone.
- **Presubmit lanes** run against unmerged PR code. Failures may be caused by individual PR changes, so the baseline failure rate is noisier. However, tests that fail consistently in presubmits with high volume are strong flake signals.
- When a test fails in **both** periodic and presubmit lanes, it's almost certainly a real issue in mainline (not PR-caused).
- When a test fails **only** in presubmits, it may be a genuine flake that's easier to trigger under presubmit conditions (different concurrency, resource pressure, timing).

## Data generation

Run the `lane-rate` subcommand from the `kubevirt.io/ci-health` repository with a testgrid URL.

If not already inside the `ci-health` repository, `cd` into it first (e.g. `cd kubevirt.io/ci-health` from the `github.com` directory).

```bash
$ go run ./cmd/ci-failures lane-rate $ARGUMENTS
```

To control the analysis window (default 14 days):

```bash
$ go run ./cmd/ci-failures lane-rate --days 21 $ARGUMENTS
```

To filter out one-off failures and only keep tests at or below a success rate threshold (useful for noisy presubmit lanes):

```bash
$ go run ./cmd/ci-failures lane-rate --max-success-rate 95 $ARGUMENTS
```

Example:

```bash
$ go run ./cmd/ci-failures lane-rate $ARGUMENTS
```

This produces `output/tmp/lane-rate-{tab-name}.yaml`.

When analyzing presubmit lanes (which typically have many one-off test failures from individual PRs), use `--max-success-rate 95` to focus on tests with meaningful failure patterns.

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
7. **Quarantine transitions**: a test may be quarantined or unquarantined during the observation period. When this happens, the results show two entries for the same underlying test — one with `[QUARANTINE]` in the name and one without — each covering only the portion of the window where that variant was active. Treat these as a single test: combine their failure counts mentally, and note the transition (e.g. "quarantined mid-window, 0/45 before + 0/6 after"). A high skip count on one variant and low skip count on the other is the telltale sign of a mid-window quarantine change.

### Summary table

After per-test details, provide a summary:

| Test | Rate | Runs | Pattern | Assessment |
|------|------|------|---------|------------|

Where **Pattern** is one of:
- **flaky** — high flip rate, failures spread across time (genuine flake)
- **burst** — failures clustered in a short time window (transient event)
- **consecutive** — 3+ failures in a row (deterministic breakage)
- **infra-correlated** — fails alongside other unrelated tests in same builds

## Issue tracking

After presenting the analysis and top priorities, check GitHub for existing tracker issues and offer to create missing ones.

### Searching for existing issues

For each top-priority test failure, search the upstream repository (typically `kubevirt/kubevirt`) for open issues that reference the test name or a distinctive keyword from the error message:

```bash
gh issue list -R kubevirt/kubevirt --state open --search "virtiofs ServiceAccount token" --limit 5
```

Use a short, distinctive substring of the test name — not the full name, which is too long for search. Try the most specific part (e.g. `virtiofs ServiceAccount`, `should pause VMI on IO error`, `volume migrate block to filesystem`). If no results, try the error signature (e.g. `expired after 200 seconds config.go:173`).

### Presenting results

For each top-priority item, report whether a tracking issue exists:
- **Found**: link the issue (number + title)
- **Not found**: note that no existing issue was found

After the search results, offer to create GitHub issues for any top-priority failures that lack a tracker. Do not create issues automatically — ask the user first. When creating, include:
- Test name and lane(s) affected
- Success rate and time window
- Failure pattern (consecutive, flaky, infra-correlated)
- Error message snippet
- Testgrid link(s)

## Opaque lanes (missing JUnit data)

Some lanes — typically unit test lanes on non-Bazel architectures (e.g. s390x) — do not produce JUnit XML artifacts. When this happens, testgrid only has two entries:

```
{lane-name}.Overall
{lane-name}.Pod
```

### Detecting opaque lanes

After running `lane-rate`, check whether the `failed_tests` list contains only entries ending in `.Overall` or `.Pod` (or is empty despite the lane having a non-zero failure rate in testgrid). If so, the lane is opaque — all failures are collapsed into a single `.Pod` entry, and per-test analysis is meaningless.

### What to do with opaque lanes

1. **Flag the limitation**: tell the user that this lane lacks per-test JUnit data, so lane-rate analysis cannot distinguish individual test failures. The aggregate failure rate is still valid, but per-test breakdown is not.
2. **Suggest manual build analysis**: use the `analyze-build` skill on individual failing builds to identify specific test failures from the raw build log.
3. **Root cause**: KubeVirt unit tests use Ginkgo with a shared `KubeVirtTestSuiteSetup` function (`staging/src/kubevirt.io/client-go/testutils/setup.go`) that already supports JUnit XML via `v1reporter.NewV1JUnitReporter`, but the reporter is only activated under Bazel environment variables (`GO_TEST_WRAP`, `XML_OUTPUT_FILE`). Non-Bazel test paths (like `make go-test` used by s390x) skip JUnit output because these env vars are absent. The fix is to enable the Ginkgo JUnit reporter when the `ARTIFACTS` env var is set (which Prow provides).
4. **Link to tracking issue**: https://github.com/kubevirt/kubevirt/issues/18225

## Combining with test-failure-rate

The `lane-rate` command gives a lane-wide overview, while `test-failure-rate` gives a cross-lane view for specific tests. Use `lane-rate` to find which tests are problematic in a lane, then use `test-failure-rate` on a specific build to see if those tests also fail in other lanes.
