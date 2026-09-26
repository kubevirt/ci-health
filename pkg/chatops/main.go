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

// OpenPRRetests returns /retest counts in [updatedSince, until] for currently
// open PRs updated at or after updatedSince. considered is the number of open
// PRs matched by the search, including those whose timeline could not be fetched.
func (co *Handler) OpenPRRetests(updatedSince, until time.Time) (counts map[types.PR]int, considered int, err error) {
	numbers, err := co.client.OpenPRNumbers(updatedSince)
	if err != nil {
		return nil, 0, err
	}
	counts = map[types.PR]int{}
	for _, number := range numbers {
		frag, fetchErr := co.client.FetchChatopsPR(number)
		if fetchErr != nil {
			log.WithError(fetchErr).Warnf("OpenPRRetests: failed to fetch timeline for PR %d, skipping", number)
			continue
		}
		counts[types.PR{Number: number}] = AllRetestCommentsBetween(frag, updatedSince, until)
	}
	return counts, len(numbers), nil
}

// AllRetestsBetween returns the number of /retest comments on each merged PR
// whose comment time falls in [startDate, endDate].
func (co *Handler) AllRetestsBetween(startDate, endDate time.Time) (map[int]int, error) {
	items, err := co.client.ChatopsMergedPRsBetween(startDate, endDate)
	if err != nil {
		return nil, err
	}
	result := map[int]int{}
	for _, prItem := range items {
		result[prItem.Number] = AllRetestCommentsBetween(&prItem.ChatopsPullRequestFragment, startDate, endDate)
	}
	return result, nil
}

// AllRetestCommentsBetween counts /retest comments with CreatedAt in [start, end].
// A zero start or end bound is unbounded. Phase-2 bot comments are excluded.
func AllRetestCommentsBetween(pr *types.ChatopsPullRequestFragment, start, end time.Time) int {
	var total int
	const phase2Intro = "Required labels detected, running phase 2 presubmits:"
	for _, timelineItem := range pr.Nodes {
		if timelineItem.IssueCommentFragment == (types.IssueCommentFragment{}) {
			continue
		}
		if strings.Contains(timelineItem.BodyText, phase2Intro) {
			continue
		}
		if !strings.HasPrefix(timelineItem.BodyText, "/retest") {
			continue
		}
		created := timelineItem.IssueCommentFragment.CreatedAt
		if !start.IsZero() && created.Before(start) {
			continue
		}
		if !end.IsZero() && created.After(end) {
			continue
		}
		total++
	}
	return total
}

// RetestComments returns the number of /retest or /test comments a PR received
// after the last commit or force push.
func RetestComments(pr *types.ChatopsPullRequestFragment) int {
	var total int
	const phase2Intro = "Required labels detected, running phase 2 presubmits:"

	lastPush := determineLastPush(pr)

	for _, timelineItem := range pr.Nodes {
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
	for _, timelineItem := range pr.Nodes {
		if isCommit(timelineItem) {
			itemDate = timelineItem.Commit.CommittedDate
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
		(strings.HasPrefix(timelineItem.BodyText, "/retest"))
}
