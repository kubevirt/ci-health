package mergequeue

import (
	"time"

	"github.com/fgimenez/ci-health/pkg/constants"
	"github.com/fgimenez/ci-health/pkg/types"
)

var (
	zeroDate = time.Time{}
)

// DatePREntered returns when a PR entered the merge queue before a
// given date, zero value date if it was not in the merge queue by that date.
func DatePREntered(pr *types.PullRequestFragment, date time.Time) time.Time {
	labelsAdded := make(map[string]time.Time)
	labelsRemoved := make(map[string]time.Time)

	for _, timelineItem := range pr.TimelineItems.Nodes {
		if (timelineItem.LabeledEventFragment != types.LabeledEventFragment{}) {
			labelsAdded[timelineItem.LabeledEventFragment.AddedLabel.Name] = timelineItem.LabeledEventFragment.CreatedAt
			labelsRemoved[timelineItem.LabeledEventFragment.AddedLabel.Name] = zeroDate
		} else if (timelineItem.UnlabeledEventFragment != types.UnlabeledEventFragment{}) {
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
