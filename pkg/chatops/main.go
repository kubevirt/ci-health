// Package chatops provides primitives for querying data about actions triggered
// by comments in issues of GitHub repos managed by Prow.

package chatops

import (
	"strings"
	"time"

	log "github.com/sirupsen/logrus"

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

// RetestsToMerge returns a map with the number of retest calls it took to land
// each merged PR in the time frame between the given start and end dates.
func (co *Handler) RetestsToMerge(startDate, endDate time.Time) (map[int]int, error) {

	items, err := co.client.MergedPRsBetween(startDate, endDate)
	if err != nil {
		return nil, err
	}

	result := map[int]int{}
	for _, pr := range items {
		log.Debugf("RetestsToMerge: calculating date of last push for PR num %d merged at %s", pr.Number, pr.MergedAt)
		result[pr.Number] = RetestComments(&pr.PullRequestFragment)
	}
	return result, nil
}

// RetestComments returns the number of /retest or /test comments a PR received
// after the given date.
func RetestComments(pr *types.PullRequestFragment) int {
	var total int
	lastPush := determineLastPush(pr)

	for _, timelineItem := range pr.TimelineItems.Nodes {
		if isRetestCommentAfterLastPush(timelineItem, lastPush) {
			total += 1
		}
	}
	return total
}

func determineLastPush(pr *types.PullRequestFragment) time.Time {
	lastPush := zeroDate

	for _, timelineItem := range pr.TimelineItems.Nodes {
		if isCommit(timelineItem) {
			if timelineItem.CommitFragment.PushedDate.After(lastPush) {
				lastPush = timelineItem.CommitFragment.PushedDate
			}
		}
	}
	return lastPush
}

func isCommit(timelineItem types.TimelineItem) bool {
	return timelineItem.CommitFragment != types.CommitFragment{}
}

func isRetestCommentAfterLastPush(timelineItem types.TimelineItem, lastPush time.Time) bool {
	return timelineItem.IssueCommentFragment != types.IssueCommentFragment{} &&
		timelineItem.IssueCommentFragment.CreatedAt.After(lastPush) &&
		(strings.HasPrefix(timelineItem.IssueCommentFragment.BodyText, "/retest") ||
			strings.HasPrefix(timelineItem.IssueCommentFragment.BodyText, "/test"))
}
