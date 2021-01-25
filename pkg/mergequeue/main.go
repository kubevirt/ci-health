package mergequeue

import (
	"time"

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
		if DatePREntered(&pr.PullRequestFragment, date) != zeroDate {
			result++
		}
	}

	return result, nil
}

// DatePREntered returns when a PR entered the merge queue before a
// given date, zero value date if it was not in the merge queue by that date.
func DatePREntered(pr *types.PullRequestFragment, date time.Time) time.Time {
	labelsAdded := make(map[string]time.Time)
	labelsRemoved := make(map[string]time.Time)

	for _, timelineItem := range pr.TimelineItems.Nodes {
		if (timelineItem.LabeledEventFragment != types.LabeledEventFragment{} &&
			date.After(timelineItem.LabeledEventFragment.CreatedAt)) {

			labelsAdded[timelineItem.LabeledEventFragment.AddedLabel.Name] = timelineItem.LabeledEventFragment.CreatedAt
			labelsRemoved[timelineItem.LabeledEventFragment.AddedLabel.Name] = zeroDate

		} else if (timelineItem.UnlabeledEventFragment != types.UnlabeledEventFragment{} &&
			date.After(timelineItem.UnlabeledEventFragment.CreatedAt)) {

			labelsAdded[timelineItem.UnlabeledEventFragment.RemovedLabel.Name] = zeroDate
			labelsRemoved[timelineItem.UnlabeledEventFragment.RemovedLabel.Name] = timelineItem.UnlabeledEventFragment.CreatedAt

		}
	}

	if labelsAdded[constants.LGTMLabel].After(labelsAdded[constants.HoldLabel]) &&
		labelsAdded[constants.LGTMLabel].After(labelsAdded[constants.NeedsRebaseLabel]) &&
		labelsAdded[constants.LGTMLabel].Before(date) &&
		labelsAdded[constants.ApprovedLabel].After(labelsAdded[constants.HoldLabel]) &&
		labelsAdded[constants.ApprovedLabel].After(labelsAdded[constants.NeedsRebaseLabel]) &&
		labelsAdded[constants.ApprovedLabel].Before(date) {

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
