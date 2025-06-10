// Package chatops provides primitives for querying data about actions triggered
// by comments in issues of GitHub repos managed by Prow.

package chatops

import (
	"strings"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/kubevirt/ci-health/pkg/constants"
	"github.com/kubevirt/ci-health/pkg/gh"
	"github.com/kubevirt/ci-health/pkg/types"
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

// RetestsToMerge returns a map with the number of retest calls it took to land
// each merged PR in the time frame between the given start and end dates.
func (co *Handler) RetestsToMerge(startDate, endDate time.Time) (map[types.PR]int, error) {

	items, err := co.client.ChatopsMergedPRsBetween(startDate, endDate)
	if err != nil {
		return nil, err
	}

	result := map[types.PR]int{}
	for _, prItem := range items {
		log.Debugf("RetestsToMerge: calculating date of last commit or force push for PR num %d merged at %s", prItem.Number, prItem.MergedAt)
		pr := types.PR{
			Number:   prItem.Number,
			MergedAt: prItem.MergedAt.Format(constants.DateFormat),
		}
		result[pr] = RetestComments(&prItem.ChatopsPullRequestFragment)
	}
	return result, nil
}

// RetestComments returns the number of /retest or /test comments a PR received
// after the last commit or force push.
func RetestComments(pr *types.ChatopsPullRequestFragment) int {
	var total int
	const phase2Intro = "Required labels detected, running phase 2 presubmits:"

	lastPush := determineLastPush(pr)

	for _, timelineItem := range pr.TimelineItems.Nodes {
		if strings.Contains(timelineItem.BodyText, phase2Intro) {
			continue
		}
		if isRetestCommentAfterLastPush(timelineItem, lastPush) {
			total += 1
		}
	}
	return total
}

func determineLastPush(pr *types.ChatopsPullRequestFragment) time.Time {
	lastPush := zeroDate

	var itemDate time.Time
	for _, timelineItem := range pr.TimelineItems.Nodes {
		if isCommit(timelineItem) {
			itemDate = timelineItem.PullRequestCommitFragment.Commit.CommittedDate
		} else if isHeadRefForcePush(timelineItem) {
			itemDate = timelineItem.HeadRefForcePushFragment.CreatedAt
		} else if isBaseRefForcePush(timelineItem) {
			itemDate = timelineItem.BaseRefForcePushFragment.CreatedAt
		}
		if itemDate.After(lastPush) {
			lastPush = itemDate
		}
	}
	return lastPush
}

func isCommit(timelineItem types.TimelineItem) bool {
	return timelineItem.PullRequestCommitFragment != types.PullRequestCommitFragment{}
}

func isHeadRefForcePush(timelineItem types.TimelineItem) bool {
	return timelineItem.HeadRefForcePushFragment.Actor.Login != ""
}

func isBaseRefForcePush(timelineItem types.TimelineItem) bool {
	return timelineItem.BaseRefForcePushFragment.Actor.Login != ""

}

func isRetestCommentAfterLastPush(timelineItem types.TimelineItem, lastPush time.Time) bool {
	return timelineItem.IssueCommentFragment != types.IssueCommentFragment{} &&
		timelineItem.IssueCommentFragment.CreatedAt.After(lastPush) &&
		(strings.HasPrefix(timelineItem.IssueCommentFragment.BodyText, "/retest"))
}
