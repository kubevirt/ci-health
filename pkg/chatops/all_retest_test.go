package chatops_test

import (
	"testing"
	"time"

	"github.com/kubevirt/ci-health/pkg/chatops"
	"github.com/kubevirt/ci-health/pkg/types"
)

func TestAllRetestCommentsIncludesBeforeLastPush(t *testing.T) {
	t.Parallel()

	push := time.Date(2026, 9, 12, 0, 0, 0, 0, time.UTC)
	pr := &types.ChatopsPullRequestFragment{
		TimelineItems: types.TimelineItems{
			Nodes: []types.TimelineItem{
				{
					IssueCommentFragment: types.IssueCommentFragment{
						CreatedAt: push.Add(-24 * time.Hour),
						BodyText:  "/retest",
					},
				},
				{
					PullRequestCommitFragment: types.PullRequestCommitFragment{
						Commit: types.Commit{CommittedDate: push},
					},
				},
				{
					IssueCommentFragment: types.IssueCommentFragment{
						CreatedAt: push.Add(time.Hour),
						BodyText:  "/retest",
					},
				},
				{
					IssueCommentFragment: types.IssueCommentFragment{
						CreatedAt: push.Add(2 * time.Hour),
						BodyText:  "Required labels detected, running phase 2 presubmits:",
					},
				},
			},
		},
	}
	if got := chatops.RetestComments(pr); got != 1 {
		t.Errorf("RetestComments = %d, want 1 (after last push only)", got)
	}
	if got := chatops.AllRetestCommentsBetween(pr, time.Time{}, time.Time{}); got != 2 {
		t.Errorf("AllRetestCommentsBetween unbounded = %d, want 2 (before and after last push)", got)
	}
	start := push.Add(-12 * time.Hour)
	end := push.Add(12 * time.Hour)
	if got := chatops.AllRetestCommentsBetween(pr, start, end); got != 1 {
		t.Errorf("AllRetestCommentsBetween = %d, want 1 (only comments in the window)", got)
	}
}
