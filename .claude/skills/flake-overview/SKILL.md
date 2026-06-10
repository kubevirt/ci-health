---
name: flake-overview
description: >
    Unified flake analysis combining PR-based flakefinder data with periodic lane failure rates.
    Use when the user wants an overview of test flakiness across the project.
allowed-tools: [Read, Write, WebFetch, Bash(go run:*), Bash(grep:*), Bash(find:*), Bash(ls:*), Bash(test:*), Bash(gh search code:*), Bash(gh search issues:*), Bash(gh issue list:*), Bash(mkdir:*), Bash(curl:*)]
---

## Overview

This skill runs a single command that combines two complementary views of test flakiness:

1. **flakefinder data** — aggregates 24h flakefinder reports from GCS, showing which lanes have the most test failures across PRs
2. **lane-rate data** — for each lane, fetches testgrid data directly and computes per-test failure rates

Together they answer: "which tests are flaky, is it a PR-interaction problem or a baseline issue, and is it getting better or worse?"

## Step 1: Locate the ci-health repository and run flake-overview twice

The `cmd/ci-failures` tool lives in the `kubevirt.io/ci-health` repository. If not already inside the `ci-health` repository, `cd` into it first (e.g. `cd kubevirt.io/ci-health` from the `github.com` directory).

Run the command twice — once for a 28-day window and once for a 7-day window — to enable trend detection:

```bash
go run ./cmd/ci-failures flake-overview --days 28
```

Read the output YAML, then run:

```bash
go run ./cmd/ci-failures flake-overview --days 7
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

**Always search online** — do not look for or use a local kubevirt repository clone.

1. **Extract the `It("...")` description** from each test name. The test name is a concatenation of Ginkgo `Describe`/`Context`/`It` blocks separated by spaces. The `It` text is the last meaningful segment — typically everything after the last known context boundary. For example:
   - Full name: `KubeVirt Tests Suite.[sig-compute]VSOCK Live migration should retain the CID for migration target`
   - Search string: `should retain the CID for migration target`
2. **Search on GitHub** using the `gh` CLI to find the file and line number:
   ```bash
   gh search code --repo kubevirt/kubevirt --filename '*.go' '<description>'
   ```
   This returns matching files and line numbers. If results are ambiguous, narrow by adding `-- path:tests/` or use a more specific substring.
3. **Construct a GitHub permalink**: `https://github.com/kubevirt/kubevirt/blob/main/<path>#L<line>` (e.g., `https://github.com/kubevirt/kubevirt/blob/main/tests/vmi_vsock_test.go#L121`)

Skip `.Pod` and `BeforeSuite`/`AfterSuite` entries — these are infrastructure-level, not test code.

## Step 4: Search for existing GitHub issues

For each test that qualifies as a quarantine candidate (from the quarantine prioritization criteria) or is already quarantined with a success rate below 50% (quarantine debt), search for existing open issues in `kubevirt/kubevirt` that track the flake.

1. **Extract a search-friendly snippet** from the test name — use the `It("...")` description (same as Step 3). Keep it short enough to avoid false negatives but specific enough to avoid false positives. For example:
   - Test: `KubeVirt Tests Suite.[sig-compute]HostDevices with ephemeral disk with emulated PCI devices Should successfully passthrough 2 emulated PCI devices`
   - Search snippet: `passthrough 2 emulated PCI devices`

2. **Search open issues** using the `gh` CLI:
   ```bash
   gh search issues --repo kubevirt/kubevirt --state open '<search snippet>'
   ```
   If the search returns too many results, narrow with additional keywords (e.g., `flak` or `quarantine`). If it returns nothing, try a shorter or different substring from the test name.

3. **Record the results** — for each test, note:
   - Whether an open issue exists (issue number + title + URL)
   - Whether the issue is labeled `kind/flake` or similar
   - If no issue exists, flag the test as **"no tracking issue"** — this is important information for the report

4. **Rate-limit fallback** — GitHub search API has rate limits. If `gh search issues` returns a 403, fall back to the GitHub web search URL via WebFetch:
   ```
   https://github.com/kubevirt/kubevirt/issues?q=is%3Aissue+is%3Aopen+<url-encoded+search+snippet>
   ```
   This uses a different rate-limit bucket and is usually available when the API is exhausted. Parse the fetched page for issue titles and URLs. If both approaches fail, note which tests could not be checked.

## Step 5: Detect quarantine status changes

Compare the current run's quarantine data against the **most recent prior report** to surface what changed since the last analysis.

1. **Find the prior report**: list files in `output/kubevirt/kubevirt/flake-overview/` sorted by date descending. Take the most recent file whose date is **before** today's date. If no prior report exists (first run), skip this step entirely.

2. **Read the prior report** and extract three sets:
   - **Prior quarantined tests**: all test names that appeared with `[QUARANTINE]` in any per-lane table
   - **Prior quarantine candidates**: all tests listed in the "Quarantine Candidate Summary" table (non-quarantined tests flagged for quarantine)
   - **Prior quarantine debt**: all tests listed in the "Quarantine Debt" section

3. **Build the current sets** from the current YAML data and analysis:
   - **Current quarantined tests**: all test names containing `[QUARANTINE]` in the YAML `test_name` fields
   - **Current quarantine candidates**: tests identified by the quarantine prioritization criteria in this run
   - **Current quarantine debt**: quarantined tests with <50% success rate in this run

4. **Compute the diff** across five categories:

   - **Newly quarantined**: tests that have `[QUARANTINE]` in the current data but were listed as candidates or appeared without the tag in the prior report. This means the quarantine action was taken.
   - **Removed from quarantine**: tests that had `[QUARANTINE]` in the prior report but no longer have the tag in the current data. Either the tag was removed (test fixed) or the test was deleted.
   - **Stale quarantine candidates**: tests that were flagged as quarantine candidates in the prior report and are still not quarantined in the current data. Action was recommended but not taken.
   - **New quarantine debt**: tests that are quarantined with <50% success rate now but were not in the prior report's debt section (either newly quarantined and immediately broken, or a previously-healthy quarantined test degraded below 50%).
   - **Resolved quarantine debt**: tests that were in the prior report's debt section but are no longer (either fixed above 50%, unquarantined, or removed entirely).

   For matching tests across reports, normalize test names by stripping markdown link syntax, trimming whitespace, and comparing the core test description (the `It("...")` text). Exact name matching is fragile because reports may abbreviate differently — match on the distinctive substring (e.g., "USB passthrough 1 device", "IO error pause VMI").

## Presenting results

Always write the report to `output/kubevirt/kubevirt/flake-overview/flake-overview-YYYY-MM-DD.md`. The reports are long and structured — a file is easier to read, share, and diff against future runs than inline output.

The report must be **grouped by SIG** (derived from lane names: `sig-compute`, `sig-storage`, `sig-network`, `sig-operator`, `sig-monitoring`, `sig-performance`). Lanes that don't match a SIG (e.g., S390X, arm64) go under a "Platform-specific" group.

### Report structure

#### 1. Summary section (top of report)

Lead with a concise summary covering:

- Report date and analysis windows (28d and 7d)
- Total failures and number of lanes for each window
- Overall trend (improving / stable / worsening) — compute by comparing the 7d daily rate against the **non-overlapping** prior-21d rate: `prior_21d_rate = (28d_total - 7d_total) / 21`. If 7d rate < prior 21d rate → improving; if roughly equal (±15%) → stable; if higher → worsening. Do **not** compare 7d against the raw 28d average (which includes the 7d data and biases toward "stable").
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
- [Quarantine Status Changes](#quarantine-status-changes)
- [Quarantine Candidate Summary](#quarantine-candidate-summary)
```

**Anchor stability**: The "Quarantine Status Changes" heading includes a date suffix (`(since YYYY-MM-DD)`), which changes the generated anchor. To keep the TOC link stable, add an explicit HTML anchor before the heading: `<a id="quarantine-status-changes"></a>`. Do the same for any heading whose text varies between reports.

#### 2. Infrastructure section

Before the SIG sections, flag infra-correlated lanes:

- Presubmit Pod-level failure rates (all lanes with `.Pod` failures)
- Lanes with complete setup failures (BeforeSuite/AfterSuite at 0%)
- Include 28d-vs-7d trend for infra health

#### 3. Per-SIG sections

For each SIG, create a section with:

- **SIG header** with total failures and failure share across all lanes in that SIG
- **h3 lane groupings** — always group lanes under `### Periodic lanes` and/or `### Presubmit lanes` headings, even when there is only one lane in the group. This maintains consistent h2→h3→h4 heading hierarchy across all SIG sections.
- **Lanes table**: list all lanes for this SIG with failures, share, builds, and 28d-vs-7d trend
- **Failing tests table** (h4 per lane) for each lane, sorted by success rate (worst first). For each test include:
  - Test name as a **markdown link to the source code** (GitHub permalink from Step 3). If no source link was resolved, use the plain test name.
  - Success rate (28d and 7d), total runs, and severity
  - **Trend**: arrow or label indicating direction (28d→7d)
  - **Dispersion**: how many lanes this test fails in (dispersed / concentrated / universal)
  - **Pattern**: flaky (high flip rate), burst (clustered timestamps), consecutive (deterministic), or infra-correlated
  - `[QUARANTINE]` flag if present in the test name

#### 4. Quarantine status changes (before the candidate summary)

If a prior report exists (from Step 5), add a section tracking what changed:

```markdown
## Quarantine Status Changes (since YYYY-MM-DD)
```

Include only the subsections that have entries — omit empty categories. If no prior report exists, omit this entire section.

**Newly quarantined** — tests that gained the `[QUARANTINE]` tag since the prior report:

| Test | SIG | Prior Status |
|------|-----|--------------|

Where **Prior Status** is "quarantine candidate (priority)" or "not flagged".

**Removed from quarantine** — tests that lost the `[QUARANTINE]` tag:

| Test | SIG | Prior Rate | Current Rate | Notes |
|------|-----|-----------|--------------|-------|

Where **Notes** explains whether the test was fixed (now passing), deleted, or still failing without the tag.

**Stale quarantine candidates** — tests flagged as candidates in the prior report that are still not quarantined:

| Test | SIG | First Flagged | Current Rate | Issue |
|------|-----|--------------|--------------|-------|

Where **First Flagged** is the date of the prior report that first recommended quarantine, and **Issue** links to any tracking issue (from Step 4).

**Quarantine debt resolved** — tests that were in the prior report's debt section but are no longer:

| Test | SIG | Prior Rate | Resolution |
|------|-----|-----------|------------|

Where **Resolution** is "fixed (now X%)", "unquarantined", or "test removed".

**New quarantine debt** — tests newly fallen below 50% success since the prior report:

| Test | SIG | Prior Rate | Current Rate |
|------|-----|-----------|--------------|

#### 5. Quarantine candidate summary (bottom of report)

After all SIG sections, provide a cross-SIG quarantine candidate table:

| Test | SIG | Lanes | Worst Rate (28d) | 7d Trend | Pattern | Dispersion | Priority | Issue |
|------|-----|-------|-------------------|----------|---------|------------|----------|-------|

Where:
- **Test** is a markdown link to the source code (GitHub permalink from Step 3)
- **Priority** is high/medium/low/not-candidate per the quarantine prioritization criteria
- **7d Trend** is improving/stable/worsening with the rate delta
- **Issue** is a link to the existing open GitHub issue (from Step 4), or "none" if no tracking issue exists
- Only include non-quarantined tests (exclude those already tagged `[QUARANTINE]`)

Then list already-quarantined tests that remain severely broken (success rate <50%) as a "quarantine debt" reminder — these have been quarantined but never fixed. These should also link to source code and include the **Issue** column.
