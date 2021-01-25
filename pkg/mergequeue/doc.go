// Package mergequeue provides primitives for querying data about GitHub merge
// queues managed by Prow.

// For a GitHub project which CI is managed by Prow, we define the merge queue
// as the list of Pull Requests that are ready to be merged at any given date.
// They have been already reviewed, approved, don't need to be rebased and are
// not marked to be hold. The package provides several functions to query data
// about the merge queue.
package mergequeue
