// Package mergequeue provides primitives for querying data about GitHub merge
// queues managed by Prow.

// For a GitHub project which CI is managed by Prow, we define the merge queue
// as the list of Pull Requests that are ready to be merged at any given date.
// They have been already reviewed, approved, don't need to be rebased and are
// not marked to be hold. The package provides several functions to query data
// about the merge queue.
package mergequeue

import (
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/fgimenez/ci-health/pkg/constants"
	"github.com/fgimenez/ci-health/pkg/gh"
	"github.com/fgimenez/ci-health/pkg/types"
)

var (
	zeroDate = time.Time{}
)

type Handler struct {
	client *gh.Client
}

func NewHandler(client *gh.Client) *Handler {
	return &Handler{
		client,
	}
}

// LengthAt returns the merge queue size for a given date.
func (mq *Handler) LengthAt(date time.Time) (int, error) {
	prs, err := mq.client.OpenPRsAt(date)
	if err != nil {
		return 0, err
	}

	result := 0
	for _, pr := range prs {
		log.Debugf("LengthAt: calculating mq date entered for PR num %d from %s", pr.Number, date)
		if DatePREntered(&pr.PullRequestFragment, date) != zeroDate {
			result++
		}
	}

	return result, nil
}

// TimesToMerge returns the duration each merged PR took to land in the time
// frame between the given start and end dates.
func (mq *Handler) TimesToMerge(startDate, endDate time.Time) ([]time.Duration, error) {
	prs, err := mq.client.MergedPRsBetween(startDate, endDate)
	if err != nil {
		return nil, err
	}

	result := []time.Duration{}
	for _, pr := range prs {
		log.Debugf("TimesToMerge: calculating mq date entered for PR num %d merged at %s", pr.Number, pr.MergedAt)
		mqStart := DatePREntered(&pr.PullRequestFragment, pr.MergedAt)
		if mqStart.Equal(zeroDate) {
			return nil, fmt.Errorf("Merge queue enter date not found for PR %d", pr.Number)
		}
		result = append(result, pr.MergedAt.Sub(mqStart))
	}

	return result, nil
}

// DatePREntered returns when a PR entered the merge queue before a
// given date, zero value date if it was not in the merge queue by that date.
func DatePREntered(pr *types.PullRequestFragment, date time.Time) time.Time {
	labelsAdded := make(map[string]time.Time)
	labelsRemoved := make(map[string]time.Time)
	for _, timelineItem := range pr.TimelineItems.Nodes {
		if isLabeledEvent(timelineItem, date) {

			labelsAdded[timelineItem.LabeledEventFragment.AddedLabel.Name] = timelineItem.LabeledEventFragment.CreatedAt
			labelsRemoved[timelineItem.LabeledEventFragment.AddedLabel.Name] = zeroDate

		} else if isUnlabeledEvent(timelineItem, date) {

			labelsAdded[timelineItem.UnlabeledEventFragment.RemovedLabel.Name] = zeroDate
			labelsRemoved[timelineItem.UnlabeledEventFragment.RemovedLabel.Name] = timelineItem.UnlabeledEventFragment.CreatedAt

		}
	}

	if isPRInMergeQueue(labelsAdded, date) {

		if labelsRemoved[constants.HoldLabel].After(labelsAdded[constants.LGTMLabel]) &&
			labelsRemoved[constants.HoldLabel].After(labelsAdded[constants.ApprovedLabel]) {

			return labelsRemoved[constants.HoldLabel]
		}
		if labelsRemoved[constants.NeedsRebaseLabel].After(labelsAdded[constants.LGTMLabel]) &&
			labelsRemoved[constants.NeedsRebaseLabel].After(labelsAdded[constants.ApprovedLabel]) {

			return labelsRemoved[constants.NeedsRebaseLabel]
		}

		if labelsAdded[constants.ApprovedLabel].After(labelsAdded[constants.LGTMLabel]) {

			return labelsAdded[constants.ApprovedLabel]
		}
		return labelsAdded[constants.LGTMLabel]
	}
	return zeroDate
}

func isLabeledEvent(timelineItem types.TimelineItem, date time.Time) bool {
	return timelineItem.LabeledEventFragment != types.LabeledEventFragment{} &&
		timelineItem.LabeledEventFragment.AddedLabel.Name != "" &&
		date.After(timelineItem.LabeledEventFragment.CreatedAt)
}

func isUnlabeledEvent(timelineItem types.TimelineItem, date time.Time) bool {
	return timelineItem.UnlabeledEventFragment != types.UnlabeledEventFragment{} &&
		timelineItem.UnlabeledEventFragment.RemovedLabel.Name != "" &&
		date.After(timelineItem.UnlabeledEventFragment.CreatedAt)
}

func isPRInMergeQueue(labelsAdded map[string]time.Time, date time.Time) bool {
	return labelsAdded[constants.LGTMLabel].After(labelsAdded[constants.HoldLabel]) &&
		labelsAdded[constants.LGTMLabel].After(labelsAdded[constants.NeedsRebaseLabel]) &&
		labelsAdded[constants.LGTMLabel].Before(date) &&
		labelsAdded[constants.ApprovedLabel].After(labelsAdded[constants.HoldLabel]) &&
		labelsAdded[constants.ApprovedLabel].After(labelsAdded[constants.NeedsRebaseLabel]) &&
		labelsAdded[constants.ApprovedLabel].Before(date)
}
