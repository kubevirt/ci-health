# AGENTS.md

This file provides guidance to AI agents when working with code in this repository.

## Project Purpose

The `ci-health` project is a Go-based tool designed to analyze and report on the performance of the KubeVirt organization's CI/CD system, which is built on Prow. It calculates key metrics to provide insights into the health and efficiency of the development and testing pipeline.

The primary metrics tracked are:
- **Merge Queue Length**: The number of pull requests ready to be merged at any given time.
- **Time to Merge**: The duration it takes for a PR to be merged after it enters the merge queue.
- **Retests to Merge**: The number of `/test` or `/retest` commands issued on a PR before it merges.
- **Failures per SIG**: Tracks test failures attributed to different Special Interest Groups (SIGs).
- **Quarantined Tests**: Monitors the number of tests that are temporarily disabled due to flakiness.

The output of this tool is a set of SVG badges and PNG plots that are embedded in the `README.md` to provide a real-time dashboard of CI health.

## Project Structure

The project is organized into two main directories: `cmd` and `pkg`.

### `cmd/` - Command-Line Applications

This directory contains the entry points for the executable tools:

- **`stats/`**: A tool that gathers data from the most recent period (e.g., the last 7 days), calculates metrics, and generates SVG badges. This is used for the near real-time status in the README.
- **`batch/`**: A tool for processing historical data. It can fetch data for a specified date range and then generate plots (PNG images) to visualize trends over time.
- **`html-report/`**: A tool to generate detailed HTML reports of test failures, categorized by SIG.
- **`ci-failures/`**: A tool that extracts information about jobs that have failed before the test phase for a lane has started. It downloads build logs, extracts error snippets, and generates markdown reports to drill down on these failures.

### `pkg/` - Core Libraries

This directory contains the packages that implement the core logic of the `ci-health` tool:

- **`chatops/`**: Provides primitives for querying data about actions triggered by comments in GitHub repos managed by Prow, such as analyzing `/retest` and `/test` comments for merge time and retry count calculations.
- **`ci-failures/`**: Provides the core logic for fetching, downloading, and processing CI failures, including clustering similar errors.
- **`constants/`**: Defines application-wide constants including date formats, metric names, badge file names, label patterns, and default configuration values.
- **`gh/`**: Handles interactions with the GitHub API to fetch data about pull requests, labels, and comments.
- **`htmlreport/`**: Generates HTML reports of CI failures and test metrics using embedded Go templates, including parsing JUnit XML test reports for SIG-specific failure analysis.
- **`mergequeue/`**, **`metrics/`**, **`stats/`**: Contain the logic for calculating the defined CI metrics.
- **`output/`**: Manages the writing of results, such as JSON data, SVG badges, and plots, to the filesystem.
- **`plot/`**: Implements the functionality to generate PNG plots from historical data.
- **`prow/`**: Contains data type definitions for Prow job metadata including started/finished job events with timestamps, pull request information, and test results.
- **`quarantine/`**: Logic specific to tracking and reporting on quarantined tests.
- **`runner/`**: Orchestrates the execution of batch and stats operations by coordinating between GitHub client, merge queue handler, and chatops handler.
- **`sigretests/`**: Logic for attributing retests and failures to the responsible SIGs.
- **`timeutils/`**: Provides utility functions for date/time manipulation, such as finding the closest Monday date.
- **`types/`**: Defines core application types such as Metric enum, Actions, Options, Results, and PR structures used throughout the codebase.

## Development Workflow

### Draft Pull Requests

**IMPORTANT**: When creating pull requests, ALWAYS create them as drafts using `gh pr create --draft`.

This applies to all PRs regardless of size or scope. The author will mark them as ready for review when appropriate.

## GitHub Actions Workflows

The repository uses GitHub Actions to automate the process of data collection, report generation, and testing.

### `test.yaml`

- **Trigger**: Runs on pull requests that modify Go source files (`cmd/`, `pkg/`, `e2e/`), Go module files (`go.*`), or the workflows themselves (`.github/`).
- **Purpose**: To ensure the integrity of the codebase.
- **Action**: It executes `go test ./...` to run all the unit tests in the project. This prevents regressions and ensures that new code is tested.

### `badges-update.yaml`

- **Trigger**: Runs on a schedule (every 3 hours) and can also be dispatched manually.
- **Purpose**: To keep the CI health status badges in the `README.md` up-to-date with the latest data.
- **Action**:
    1. Checks out the repository.
    2. Runs the `stats` command to fetch data for the last 7 days for the `kubevirt/kubevirt` repository.
    3. The command overwrites the SVG badges and JSON results in the `output/` directory.
    4. A subsequent step commits the updated files back to the repository.

### `ci-failures.yml`

- **Trigger**: Runs on pushes to `main` that modify `output/kubevirt/kubevirt/results.json`.
- **Purpose**: To automatically regenerate the CI failures summary report whenever new results data is committed.
- **Action**:
    1. Checks out the repository.
    2. Runs `go run ./cmd/ci-failures generate report` to produce an updated markdown summary of CI failures.
    3. Commits the updated output files back to the repository.

### `historical-update.yaml`

- **Trigger**: Runs on a schedule (weekly, every Monday at 00:10 UTC).
- **Purpose**: To track and visualize long-term CI health trends.
- **Action**:
    1. Checks out the repository.
    2. Runs the `batch` command in `fetch` and then `plot` mode for several key metrics: `retests-to-merge`, `time-to-merge`, `merge-queue-length`, `merged-prs`, and `quarantined-tests`.
    3. This process updates the historical data files and regenerates the trend plots (PNG images).
    4. A final step commits the updated data and plots back to the repository.
