---
name: change-relevance
description: >
    Determine whether a PR's code changes are related to tests that failed in a Prow build.
    Use when user wants to know if test failures are caused by their PR changes.
allowed-tools: [Read, Glob, Bash(go run:*)]
---

## Overview

This skill checks whether a PR's changed files overlap with the code areas that failing tests cover. It extracts the PR number from the Prow job URL, fetches the changed files from GitHub, maps each failed test to source directories via its SIG label, and reports whether the changes are related.

## Data generation

Run the `change-relevance` subcommand with a Prow job URL for a **PR build**:

```bash
$ go run ./cmd/ci-failures change-relevance <prow-job-url>
```

Example:

```bash
$ go run ./cmd/ci-failures change-relevance https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17760/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2053843644594524160
```

This produces `output/tmp/change-relevance-{job-name}.yaml`.

Only PR builds (URLs containing `pr-logs/pull/`) are supported. Periodic builds will return an error.

## Analysis

After data generation:

1. Use Glob to find the `output/tmp/change-relevance-*.yaml` file
2. Read the YAML. The structure is:
   - `prow_job_url`: the analyzed build URL
   - `job_name`: the Prow job name
   - `org`, `repo`, `pr_number`: the GitHub PR identity
   - `changed_files`: list of files changed in the PR
   - `failed_tests`: list of failed tests, each with:
     - `test_name`: original test name from the build log
     - `sig`: the SIG label extracted from the test name (or inferred from job name)
     - `code_areas`: source directory prefixes mapped to that SIG
     - `overlap_files`: PR files that overlap with the test's code areas (empty if none)
     - `relevance`: classification of the relationship
     - `reason`: human-readable explanation
   - `summary`: aggregate counts of each relevance level

## Relevance classification

- **related**: The PR directly changes files in the SIG's code areas. The test failure is likely caused by the PR.
- **possibly-related**: The PR changes broad-impact files (API definitions, shared utilities, test framework) that could affect any test. Cannot rule out the PR as the cause.
- **unrelated**: The PR does not change any files in the test's code areas or in broad-impact areas. The failure is likely a pre-existing flake.
- **unknown**: The test has no SIG label and the SIG could not be inferred from the job name, or the SIG is not in the mapping.

## Presenting results

For each failed test, report:
1. The test name
2. The SIG and relevance classification
3. If related or possibly-related, list the overlapping files
4. If unrelated, note that the failure is likely a flake

Summarize with the aggregate counts from the `summary` section.

## Combining with test-failure-rate

For the most complete picture, run both `change-relevance` and `test-rate` on the same Prow job URL. Together they answer:
- **test-rate**: Is this test historically flaky?
- **change-relevance**: Did this PR touch the code the test covers?

A test that is "likely-pr-related" by success rate AND "related" by change analysis is very likely caused by the PR.
