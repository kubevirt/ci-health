---
name: trigger-pr-jobs
description: >
    Trigger missing or failed required status checks for a GitHub pull request.
    Use when user provides a GitHub PR URL and wants to retrigger failed or missing CI jobs.
argument-hint: <PR-URL>
allowed-tools: [Bash(gh api:*), Bash(gh pr checks:*), Bash(gh pr comment:*), Bash(gh pr view:*)]
---

## Overview

This skill identifies required status checks that are either missing (no result) or failed on a GitHub pull request, then posts `/test` commands to retrigger them.

It works for any repo in the `kubevirt` GitHub org that uses Prow.

## Steps

### 1. Extract PR details

Parse the PR URL from the arguments to get `owner`, `repo`, and `pr_number`. Validate that all three values were successfully extracted and are non-empty. If the URL is malformed or any component is missing, report the error to the user and stop.

### 2. Get the target branch

```bash
gh api repos/{owner}/{repo}/pulls/{pr_number} --jq '.base.ref'
```

### 3. Get required status checks

```bash
gh api repos/{owner}/{repo}/branches/{branch}/protection/required_status_checks --jq '.contexts[]'
```

If this returns a 404 (branch protection not configured), treat it as having no required checks. Tell the user that no required status checks are configured for the target branch and stop.

### 4. Get current check results

Get the status of all checks reported on the PR:

```bash
gh pr checks {pr_number} --repo {owner}/{repo} --json name,state,bucket
```

This returns a JSON array of objects containing the check name and its current state (e.g., `SUCCESS`, `FAILURE`, `PENDING`).

### 5. Find checks to trigger

Compare the required checks against the checks present on the commit. A check needs triggering if it is:
- **Missing**: appears in required checks but has no status/check-run at all on the latest commit
- **Failed**: appears in the PR checks with `fail` status

**Only trigger Prow jobs.** Filter to job names starting with `pull-`. Non-Prow checks cannot be triggered via `/test` and must be excluded:
- `dco`, `tide` — Prow tide/plugin statuses, not jobs
- `coverage/coveralls` — third-party service
- GitHub Actions checks (e.g. `check-vep`) — triggered by GitHub, not Prow
- Other third-party checks (e.g. `Sourcery review`)

If there are no checks to trigger, tell the user and stop.

### 6. Post /test commands

Post a **single comment** on the PR with one `/test {job_name}` line per check that needs triggering:

```bash
gh pr comment {pr_number} --repo {owner}/{repo} --body "/test job-1
/test job-2
/test job-3"
```

Always combine all `/test` lines into one comment to avoid noise on the PR.

## Smart retrigger guidance

Before posting `/test` commands, consider whether retriggering is likely to help:

- **Known flakes** (test has < 80% success rate in flakefinder): retriggering will likely succeed — worth doing
- **Likely PR-related** (test has > 95% success rate): retriggering will likely fail again — suggest the user investigate with `analyze-build` or `test-failure-rate` first instead of wasting CI resources
- **Infra failures** (many unrelated tests failed together): retriggering may help if the infra event was transient

When there are many failures, suggest running `test-failure-rate` on one of the failed build URLs first to quickly assess whether the failures are flaky. If all failures are known flakes, retrigger confidently. If most are likely PR-related, warn the user before retriggering.

## Output

Present a summary table to the user showing:
- Which checks were missing (no result)
- Which checks were failed
- Confirmation that the `/test` commands were posted, with a link to the comment
