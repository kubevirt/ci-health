---
name: flake-overview
description: >
    Unified flake analysis combining PR-based flakefinder data with periodic lane failure rates.
    Use when the user wants an overview of test flakiness across the project.
allowed-tools: [Read, Bash(go run:*), Bash(grep:*), Bash(find:*), Bash(ls:*), Bash(test:*), Write]
---

## Overview

This skill runs a single command that combines two complementary views of test flakiness:

1. **flakefinder data** — aggregates 24h flakefinder reports from GCS, showing which lanes have the most test failures across PRs
2. **lane-rate data** — for each lane, fetches testgrid data directly and computes per-test failure rates

Together they answer: "which tests are flaky, is it a PR-interaction problem or a baseline issue, and is it getting better or worse?"

## Step 1: Locate the ci-health repository and run flake-overview twice

The `cmd/ci-failures` tool lives in the `kubevirt.io/ci-health` repository. The skill may be executed from either the top-level `github.com` directory or directly inside the `ci-health` repository. Determine the correct path to the tool based on the current working directory:
- From `github.com/`: `go run ./kubevirt.io/ci-health/cmd/ci-failures ...`
- From `kubevirt.io/ci-health/`: `go run ./cmd/ci-failures ...`

Run the command twice — once for a 28-day window and once for a 7-day window — to enable trend detection:

```bash
go run ./kubevirt.io/ci-health/cmd/ci-failures flake-overview --days 28
```

Read the output YAML, then run:

```bash
go run ./kubevirt.io/ci-health/cmd/ci-failures flake-overview --days 7
```

Read that output YAML too. Both runs produce a file at `output/tmp/flake-overview.yaml` (or `output/tmp/sessions/{session-id}/flake-overview.yaml` when a session ID is set). The second run overwrites the first, so read the first output before running the second.

The command:
1. Fetches flakefinder 24h reports from GCS for the requested window
2. Aggregates per-lane failure counts (filtering out `.*-root$` lanes by default)
3. For each lane, fetches testgrid data and computes per-test success rates (6 lanes in parallel)
4. Combines everything into a single YAML

## Step 2: Read and cross-reference both outputs

Use the 28-day data as the baseline and the 7-day data for trend detection. For each test, compare:

- **7d rate similar to 28d**: long-standing flake — quarantine is warranted
- **7d rate much worse than 28d**: flakiness is worsening — possible recent regression, investigate recent merges
- **7d rate much better than 28d**: flakiness is improving — a fix may have landed recently, quarantine may not be needed

### YAML structure

```yaml
report_date: "2026-05-13"
analysis_days: 14
total_failures: 1207
lanes:
  - name: periodic-kubevirt-e2e-k8s-1.36-sig-compute
    url: https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.36-sig-compute
    failures: 348
    share_percent: 28.83
    total_builds: 56
    failed_tests:
      - test_name: "full test description"
        succeeded: 0
        failed: 55
        skipped: 1
        total_runs: 55
        success_rate: 0.0
        severity: likely-flaky
        recent_failures:
          - build_id: "2054504557802688512"
            timestamp: "2026-05-13T10:11:03Z"
            message: "error message excerpt..."
```

Key fields:
- `total_failures` — total failures across all lanes in the window
- `lanes[*].failures` — flakefinder failure count for this lane
- `lanes[*].share_percent` — percentage of total failures
- `lanes[*].total_builds` — number of testgrid builds analyzed
- `lanes[*].failed_tests` — per-test failure data from testgrid
- `failed_tests[*].success_rate` — percentage (0-100)
- `failed_tests[*].severity` — likely-pr-related (>=95%), inconclusive (>=80%), likely-flaky (<80%)

## Deeper flakiness analysis

The severity classification and success rates are a starting point. Apply these techniques to the data for more accurate diagnosis.

### Flip-rate analysis from recent_failures

Examine the `recent_failures` list for each test to detect state transition patterns:

- **High flip rate** (failures interleaved with passes, spread across the time window): genuine flake — the test is nondeterministic
- **Low flip rate** (failures clustered consecutively): the test was broken for a period then fixed, or is still broken — this is a regression, not a flake
- **Key heuristic**: 3 or more consecutive failures should be treated as deterministic breakage, not flakiness. A test at 60% success rate with alternating pass/fail is very different from one that failed for a solid week then passed for a solid week — same rate, completely different diagnosis.

### Cross-lane dispersion

Collect all failing test names across all lanes and categorize by how many lanes they appear in:

- **Dispersed** (fails in 3+ lanes): genuine cross-environment flake — strong quarantine candidate regardless of success rate
- **Concentrated** (1–2 lanes only): lane-specific issue — check whether those lanes share a k8s version, provider, or SIG. The test itself may be fine; the environment is the problem.
- **Universal** (fails in every or nearly every lane): likely a regression on main, not a flake — check if it was recently introduced by looking at whether recent_failures timestamps cluster near the present

### Infra-correlated failure detection

Within each lane, compare `build_id`s across different failing tests:

- When multiple unrelated tests share the same `build_id`s in their `recent_failures`, they failed together in the same builds — this is an infrastructure flake at the lane level, not individual test problems
- If a small number of builds account for most failures across many tests in a lane, the lane experienced infra instability during those builds
- Flag these as **infra-correlated** — these tests should NOT be quarantined; the lane's infrastructure should be investigated instead

## Quarantine prioritization

When identifying quarantine candidates, rank by these criteria (not just success rate):

1. **High priority**: dispersed across 3+ lanes + high flip rate + low success rate — genuine cross-environment flake causing widespread CI noise
2. **Medium priority**: concentrated in 1–2 lanes + low success rate — lane-specific issue, may resolve with infra changes but still wastes CI capacity
3. **Low priority**: burst pattern or consecutive failures — transient, may already be resolved; verify with trend detection before quarantining
4. **Not a quarantine candidate**: infra-correlated (many tests failing together in same builds within a lane) — quarantining individual tests won't help; the lane needs investigation

## Step 3: Resolve test source locations

For each unique failing test, find its source file and line number so the report can link directly to the code. This helps readers who are not familiar with the test codebase.

1. **Locate the kubevirt/kubevirt repo**: look for `kubevirt.io/kubevirt` (or `kubevirt/kubevirt`) relative to the ci-health repo's parent directory.
2. **Extract the `It("...")` description** from each test name. The test name is a concatenation of Ginkgo `Describe`/`Context`/`It` blocks separated by spaces. The `It` text is the last meaningful segment — typically everything after the last known context boundary. For example:
   - Full name: `KubeVirt Tests Suite.[sig-compute]VSOCK Live migration should retain the CID for migration target`
   - Search string: `should retain the CID for migration target`
3. **Grep for the description** in the `tests/` directory: `grep -rn "<description>" tests/ --include="*.go"`
4. **Construct a GitHub permalink**: `https://github.com/kubevirt/kubevirt/blob/main/<path>#L<line>` (e.g., `https://github.com/kubevirt/kubevirt/blob/main/tests/vmi_vsock_test.go#L121`)

If the kubevirt repo is not available locally, skip this step — the report will still be useful without links, just less navigable.

Skip `.Pod` and `BeforeSuite`/`AfterSuite` entries — these are infrastructure-level, not test code.

## Presenting results

Always write the report to `output/kubevirt/kubevirt/flake-overview/flake-overview-YYYY-MM-DD.md`. The reports are long and structured — a file is easier to read, share, and diff against future runs than inline output.

The report must be **grouped by SIG** (derived from lane names: `sig-compute`, `sig-storage`, `sig-network`, `sig-operator`, `sig-monitoring`, `sig-performance`). Lanes that don't match a SIG (e.g., S390X, arm64) go under a "Platform-specific" group.

### Report structure

#### 1. Summary section (top of report)

Lead with a concise summary covering:

- Report date and analysis windows (28d and 7d)
- Total failures and number of lanes for each window
- Overall trend (improving / stable / worsening) based on 7d-vs-28d proportional comparison
- Top 3 worst SIGs by total failure count
- Number of quarantine candidates identified (high / medium / low priority)
- Number of infra-unstable lanes

Add a **navigation links** section immediately after the summary with markdown anchor links to each SIG section, e.g.:

```markdown
## Navigation

- [sig-compute](#sig-compute)
- [sig-storage](#sig-storage)
- [sig-network](#sig-network)
- [sig-operator](#sig-operator)
- [sig-monitoring](#sig-monitoring)
- [sig-performance](#sig-performance)
- [Platform-specific](#platform-specific)
- [Infrastructure](#infrastructure)
- [Quarantine Candidate Summary](#quarantine-candidate-summary)
```

#### 2. Infrastructure section

Before the SIG sections, flag infra-correlated lanes:

- Presubmit Pod-level failure rates (all lanes with `.Pod` failures)
- Lanes with complete setup failures (BeforeSuite/AfterSuite at 0%)
- Include 28d-vs-7d trend for infra health

#### 3. Per-SIG sections

For each SIG, create a section with:

- **SIG header** with total failures and failure share across all lanes in that SIG
- **Lanes table**: list all lanes for this SIG with failures, share, builds, and 28d-vs-7d trend
- **Failing tests table** for each lane, sorted by success rate (worst first). For each test include:
  - Test name as a **markdown link to the source code** (GitHub permalink from Step 3). If no source link was resolved, use the plain test name.
  - Success rate (28d and 7d), total runs, and severity
  - **Trend**: arrow or label indicating direction (28d→7d)
  - **Dispersion**: how many lanes this test fails in (dispersed / concentrated / universal)
  - **Pattern**: flaky (high flip rate), burst (clustered timestamps), consecutive (deterministic), or infra-correlated
  - `[QUARANTINE]` flag if present in the test name

#### 4. Quarantine candidate summary (bottom of report)

After all SIG sections, provide a cross-SIG quarantine candidate table:

| Test | SIG | Lanes | Worst Rate (28d) | 7d Trend | Pattern | Dispersion | Priority |
|------|-----|-------|-------------------|----------|---------|------------|----------|

Where:
- **Test** is a markdown link to the source code (GitHub permalink from Step 3)
- **Priority** is high/medium/low/not-candidate per the quarantine prioritization criteria
- **7d Trend** is improving/stable/worsening with the rate delta
- Only include non-quarantined tests (exclude those already tagged `[QUARANTINE]`)

Then list already-quarantined tests that remain severely broken (success rate <50%) as a "quarantine debt" reminder — these have been quarantined but never fixed. These should also link to source code.
