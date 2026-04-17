---
name: analyze-pr-builds
description: >
    Analyze all failed builds for a GitHub pull request.
    Use when user provides a GitHub PR URL and wants to understand why builds are failing.
allowed-tools: [Read, Glob, Bash(go run:*)]
---

## Overview

This skill analyzes all currently-failing builds for a GitHub pull request, identifies root causes, and presents a concise summary.

It works for any repo served by prow.ci.kubevirt.io (currently the `kubevirt` GitHub org).

## Data Generation

Run the `analyze-pr` command with the GitHub PR URL to fetch and analyze all failed builds:

```bash
$ go run ./cmd/ci-failures analyze-pr https://github.com/kubevirt/kubevirt/pull/17287
```

This command:
1. Fetches the PR history from prow
2. Finds all jobs for the latest commit
3. Keeps only the latest run per job (jobs that eventually passed are excluded)
4. Downloads build logs and extracts error snippets for each failure
5. Writes one YAML file per failed job to `output/tmp/{job-name}.yaml`

## Analysis

After data generation:

1. Use Glob to find all `output/tmp/*.yaml` files produced by the command
2. Read each YAML file. The structure is:
   - `job_name`: the Prow job name
   - `build_errors`: list of build errors, each containing:
     - `job_url`: link to the Prow job UI
     - `build_id`: the build number
     - `started` / `finished`: timestamps
     - `category`: error category (external, internal, pr-build, needs-investigation)
     - `category_reason`: explanation of the categorization
     - `build_log_error_snippets`: list of error matches with `error_text`, `context`, and `link_to_log_line`
3. For each failure, examine the `error_text` and `context` fields to determine the root cause
4. Group failures with the same root cause together

## Output

Present a concise summary to the user:
- Group failures by root cause
- For each group: state the root cause, list affected jobs, and include one representative error snippet
- Classify each as **fixable** (CI config or code issue) or **non-fixable** (external/infra)
- Include the `link_to_log_line` for the most relevant error in each group
