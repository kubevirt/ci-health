# ci-health

![ci-health-tests](https://github.com/fgimenez/ci-health/workflows/ci-health-tests/badge.svg)

This repo contains code to calculate metrics about the performance of CI systems
based on Prow.

## Definitions

* Merge queue: list of Pull Requests that are ready to be merged at any given
date. For being ready to be merged they must:

  * Have the `lgtm` label.
  * Have the `approved` label.
  * Not have the `do-not-merge/hold` label.
  * Not have the `needs-rebase` label.

* Merge queue length: number of PRs in the merge queue.
* Time to merge: for each merged PR, the time it took since it entered the merge
queue for the last time until it got finally merged.

## Status
This status is updated every 3 hours. The average values are calculated with
data from the previous 7 days since the execution time.

### kubevirt/kubevirt

![avg-merge-queue-lenght](https://fgimenez.github.io/ci-health/badges/kubevirt/kubevirt/merge-queue-length.svg)
![avg-time-to-merge](https://fgimenez.github.io/ci-health/badges/kubevirt/kubevirt/time-to-merge.svg)
