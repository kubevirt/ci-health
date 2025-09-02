# Hack Scripts

This directory contains various scripts for development and CI purposes.

## `show-ci-failure-jobs.sh`

**One-liner:** Identifies and lists Prow job URLs that failed before the testing stage by checking for the absence of test result artifacts.

This script identifies CI jobs that have failed for reasons other than test failures. It reads the `results.json` file to get a list of all failed jobs from the last 7 days. For each failed job, it checks if a `junit.functest.xml` artifact was generated. If this file is missing, it indicates the job failed before the main testing phase could run. The script then prints the URLs of these jobs, which are likely victims of infrastructure or build script errors rather than code-related test failures.

## `extract-errors.sh`

**One-liner:** Downloads build logs from a list of Prow job URLs, extracts error-related lines, and saves them to grouped output files.

This script takes a file containing Prow job URLs as input and downloads the `build-log.txt` for each. It then searches these logs for keywords like `ERROR`, `panic`, or `failure`. The matching lines, along with a few lines of context, are saved into separate output files in the `output/tmp/` directory. The jobs are grouped by SIGs (e.g., `sig-compute`, `sig-network`), and the script outputs the names of the files it creates, making it easier to analyze failures for specific areas.

## `download-build-logs.sh`

**One-liner:** Downloads the full `build-log.txt` for each Prow job URL provided in a file.

This script automates the process of fetching full build logs for analysis. It reads a list of Prow job URLs from `output/tmp/ci-failure-jobs.txt`. For each URL, it constructs the corresponding Google Cloud Storage URL for the `build-log.txt` file. It then downloads each log, saving it to the `output/tmp/build-logs/` directory with a filename corresponding to the job's build ID. Each downloaded file also includes a header containing the original Prow job URL for easy reference.
