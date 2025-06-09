package chatops_test

import (
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	"github.com/kubevirt/ci-health/pkg/chatops"
	"github.com/kubevirt/ci-health/pkg/types"
)

const (
	retestComment = "/retest"
)

func TestChatops(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "chatops Suite")
}

var (
	queryDate = time.Date(2021, time.Month(1), 22, 15, 11, 26, 0, time.UTC)
	zeroDate  = time.Time{}
)

var _ = Describe("RetestComments", func() {
	table.DescribeTable("should return the number of retest comments for a PR",
		func(pr *types.ChatopsPullRequestFragment, expected int) {
			actual := chatops.RetestComments(pr)
			Expect(actual).To(Equal(expected))
		},
		table.Entry("No retests after single commit",
			&types.ChatopsPullRequestFragment{
				TimelineItems: types.TimelineItems{
					Nodes: []types.TimelineItem{
						{
							PullRequestCommitFragment: types.PullRequestCommitFragment{
								Commit: types.Commit{
									CommittedDate: queryDate.AddDate(0, 0, -2),
								},
							},
						},
					},
				},
			},
			0),
		table.Entry("No retests after single HeadRef force push",
			&types.ChatopsPullRequestFragment{
				TimelineItems: types.TimelineItems{
					Nodes: []types.TimelineItem{
						{
							HeadRefForcePushFragment: types.HeadRefForcePushFragment{
								CreatedAt: queryDate.AddDate(0, 0, -2),
								Actor: types.Actor{
									Login: "user1",
								},
							},
						},
					},
				},
			},
			0),
		table.Entry("No retests after single BaseRef force push",
			&types.ChatopsPullRequestFragment{
				TimelineItems: types.TimelineItems{
					Nodes: []types.TimelineItem{
						{
							BaseRefForcePushFragment: types.BaseRefForcePushFragment{
								CreatedAt: queryDate.AddDate(0, 0, -2),
								Actor: types.Actor{
									Login: "user1",
								},
							},
						},
					},
				},
			},
			0),
		table.Entry("One retest call after last commit",
			&types.ChatopsPullRequestFragment{
				TimelineItems: types.TimelineItems{
					Nodes: []types.TimelineItem{
						{
							PullRequestCommitFragment: types.PullRequestCommitFragment{
								Commit: types.Commit{
									CommittedDate: queryDate.AddDate(0, 0, -2),
								},
							},
						},
						{
							IssueCommentFragment: types.IssueCommentFragment{
								CreatedAt: queryDate.AddDate(0, 0, -1),
								BodyText:  "/retest",
							},
						},
					},
				},
			},
			1),
		table.Entry("One retest call after last HeadRef force push",
			&types.ChatopsPullRequestFragment{
				TimelineItems: types.TimelineItems{
					Nodes: []types.TimelineItem{
						{
							HeadRefForcePushFragment: types.HeadRefForcePushFragment{
								CreatedAt: queryDate.AddDate(0, 0, -2),
								Actor: types.Actor{
									Login: "user1",
								},
							},
						},
						{
							IssueCommentFragment: types.IssueCommentFragment{
								CreatedAt: queryDate.AddDate(0, 0, -1),
								BodyText:  "/retest",
							},
						},
					},
				},
			},
			1),
		table.Entry("One retest call after last BaseRef force push",
			&types.ChatopsPullRequestFragment{
				TimelineItems: types.TimelineItems{
					Nodes: []types.TimelineItem{
						{
							BaseRefForcePushFragment: types.BaseRefForcePushFragment{
								CreatedAt: queryDate.AddDate(0, 0, -2),
								Actor: types.Actor{
									Login: "user1",
								},
							},
						},
						{
							IssueCommentFragment: types.IssueCommentFragment{
								CreatedAt: queryDate.AddDate(0, 0, -1),
								BodyText:  "/retest",
							},
						},
					},
				},
			},
			1),
		table.Entry("HeadRef force push without actor.login is ignored",
			&types.ChatopsPullRequestFragment{
				TimelineItems: types.TimelineItems{
					Nodes: []types.TimelineItem{
						{
							IssueCommentFragment: types.IssueCommentFragment{
								CreatedAt: queryDate.AddDate(0, 0, -3),
								BodyText:  "/retest",
							},
						},
						{
							HeadRefForcePushFragment: types.HeadRefForcePushFragment{
								CreatedAt: queryDate.AddDate(0, 0, -2),
							},
						},
						{
							IssueCommentFragment: types.IssueCommentFragment{
								CreatedAt: queryDate.AddDate(0, 0, -1),
								BodyText:  "/retest",
							},
						},
					},
				},
			},
			2),
		table.Entry("BaseRef force push without actor.login is ignored",
			&types.ChatopsPullRequestFragment{
				TimelineItems: types.TimelineItems{
					Nodes: []types.TimelineItem{
						{
							IssueCommentFragment: types.IssueCommentFragment{
								CreatedAt: queryDate.AddDate(0, 0, -3),
								BodyText:  "/retest",
							},
						},
						{
							BaseRefForcePushFragment: types.BaseRefForcePushFragment{
								CreatedAt: queryDate.AddDate(0, 0, -2),
							},
						},
						{
							IssueCommentFragment: types.IssueCommentFragment{
								CreatedAt: queryDate.AddDate(0, 0, -1),
								BodyText:  "/retest",
							},
						},
					},
				},
			},
			2),
		table.Entry("Multiple retest calls after last commit",
			&types.ChatopsPullRequestFragment{
				TimelineItems: types.TimelineItems{
					Nodes: []types.TimelineItem{
						{
							PullRequestCommitFragment: types.PullRequestCommitFragment{
								Commit: types.Commit{
									CommittedDate: queryDate.AddDate(0, 0, -5),
								},
							},
						},
						{
							IssueCommentFragment: types.IssueCommentFragment{
								CreatedAt: queryDate.AddDate(0, 0, -4),
								BodyText:  "/retest",
							},
						},
						{
							IssueCommentFragment: types.IssueCommentFragment{
								CreatedAt: queryDate.AddDate(0, 0, -3),
								BodyText:  "/retest",
							},
						},
						{
							IssueCommentFragment: types.IssueCommentFragment{
								CreatedAt: queryDate.AddDate(0, 0, -2),
								BodyText:  "/retest",
							},
						},
					},
				},
			},
			3),
		table.Entry("Multiple retest calls after last HeadRef force push",
			&types.ChatopsPullRequestFragment{
				TimelineItems: types.TimelineItems{
					Nodes: []types.TimelineItem{
						{
							HeadRefForcePushFragment: types.HeadRefForcePushFragment{
								CreatedAt: queryDate.AddDate(0, 0, -5),
								Actor: types.Actor{
									Login: "user1",
								},
							},
						},
						{
							IssueCommentFragment: types.IssueCommentFragment{
								CreatedAt: queryDate.AddDate(0, 0, -4),
								BodyText:  "/retest",
							},
						},
						{
							IssueCommentFragment: types.IssueCommentFragment{
								CreatedAt: queryDate.AddDate(0, 0, -3),
								BodyText:  "/retest",
							},
						},
						{
							IssueCommentFragment: types.IssueCommentFragment{
								CreatedAt: queryDate.AddDate(0, 0, -2),
								BodyText:  "/retest",
							},
						},
					},
				},
			},
			3),
		table.Entry("Multiple retest calls after last BaseRef force push",
			&types.ChatopsPullRequestFragment{
				TimelineItems: types.TimelineItems{
					Nodes: []types.TimelineItem{
						{
							BaseRefForcePushFragment: types.BaseRefForcePushFragment{
								CreatedAt: queryDate.AddDate(0, 0, -5),
								Actor: types.Actor{
									Login: "user1",
								},
							},
						},
						{
							IssueCommentFragment: types.IssueCommentFragment{
								CreatedAt: queryDate.AddDate(0, 0, -4),
								BodyText:  "/retest",
							},
						},
						{
							IssueCommentFragment: types.IssueCommentFragment{
								CreatedAt: queryDate.AddDate(0, 0, -3),
								BodyText:  "/retest",
							},
						},
						{
							IssueCommentFragment: types.IssueCommentFragment{
								CreatedAt: queryDate.AddDate(0, 0, -2),
								BodyText:  "/retest",
							},
						},
					},
				},
			},
			3),
		table.Entry("Retest calls before last commit are ignored",
			&types.ChatopsPullRequestFragment{
				TimelineItems: types.TimelineItems{
					Nodes: []types.TimelineItem{
						{
							IssueCommentFragment: types.IssueCommentFragment{
								CreatedAt: queryDate.AddDate(0, 0, -4),
								BodyText:  "/retest",
							},
						},
						{
							PullRequestCommitFragment: types.PullRequestCommitFragment{
								Commit: types.Commit{
									CommittedDate: queryDate.AddDate(0, 0, -3),
								},
							},
						},
						{
							IssueCommentFragment: types.IssueCommentFragment{
								CreatedAt: queryDate.AddDate(0, 0, -2),
								BodyText:  "/retest",
							},
						},
						{
							IssueCommentFragment: types.IssueCommentFragment{
								CreatedAt: queryDate.AddDate(0, 0, -1),
								BodyText:  "/retest",
							},
						},
					},
				},
			},
			2),
		table.Entry("Retest calls before last HeadRef force push are ignored",
			&types.ChatopsPullRequestFragment{
				TimelineItems: types.TimelineItems{
					Nodes: []types.TimelineItem{
						{
							IssueCommentFragment: types.IssueCommentFragment{
								CreatedAt: queryDate.AddDate(0, 0, -4),
								BodyText:  "/retest",
							},
						},
						{
							HeadRefForcePushFragment: types.HeadRefForcePushFragment{
								CreatedAt: queryDate.AddDate(0, 0, -3),
								Actor: types.Actor{
									Login: "user1",
								},
							},
						},
						{
							IssueCommentFragment: types.IssueCommentFragment{
								CreatedAt: queryDate.AddDate(0, 0, -2),
								BodyText:  "/retest",
							},
						},
						{
							IssueCommentFragment: types.IssueCommentFragment{
								CreatedAt: queryDate.AddDate(0, 0, -1),
								BodyText:  "/retest",
							},
						},
					},
				},
			},
			2),
		table.Entry("Retest calls before last BaseRef force push are ignored",
			&types.ChatopsPullRequestFragment{
				TimelineItems: types.TimelineItems{
					Nodes: []types.TimelineItem{
						{
							IssueCommentFragment: types.IssueCommentFragment{
								CreatedAt: queryDate.AddDate(0, 0, -4),
								BodyText:  "/retest",
							},
						},
						{
							BaseRefForcePushFragment: types.BaseRefForcePushFragment{
								CreatedAt: queryDate.AddDate(0, 0, -3),
								Actor: types.Actor{
									Login: "user1",
								},
							},
						},
						{
							IssueCommentFragment: types.IssueCommentFragment{
								CreatedAt: queryDate.AddDate(0, 0, -2),
								BodyText:  "/retest",
							},
						},
						{
							IssueCommentFragment: types.IssueCommentFragment{
								CreatedAt: queryDate.AddDate(0, 0, -1),
								BodyText:  "/retest",
							},
						},
					},
				},
			},
			2),
		table.Entry("Without last commit or force push all retests are counted",
			&types.ChatopsPullRequestFragment{
				TimelineItems: types.TimelineItems{
					Nodes: []types.TimelineItem{
						{
							IssueCommentFragment: types.IssueCommentFragment{
								CreatedAt: queryDate.AddDate(0, 0, -4),
								BodyText:  "/retest",
							},
						},
						{
							IssueCommentFragment: types.IssueCommentFragment{
								CreatedAt: queryDate.AddDate(0, 0, -2),
								BodyText:  "/retest",
							},
						},
						{
							IssueCommentFragment: types.IssueCommentFragment{
								CreatedAt: queryDate.AddDate(0, 0, -1),
								BodyText:  "/retest",
							},
						},
					},
				},
			},
			3),
		table.Entry("Retest in the body of the comment are ignored",
			&types.ChatopsPullRequestFragment{
				TimelineItems: types.TimelineItems{
					Nodes: []types.TimelineItem{
						{
							PullRequestCommitFragment: types.PullRequestCommitFragment{
								Commit: types.Commit{
									CommittedDate: queryDate.AddDate(0, 0, -3),
								},
							},
						},
						{
							IssueCommentFragment: types.IssueCommentFragment{
								CreatedAt: queryDate.AddDate(0, 0, -2),
								BodyText:  "some other text /retest blabla",
							},
						},
						{
							IssueCommentFragment: types.IssueCommentFragment{
								CreatedAt: queryDate.AddDate(0, 0, -1),
								BodyText:  "/retest",
							},
						},
					},
				},
			},
			1),
		table.Entry("Additional lines in retest comment body are allowed",
			&types.ChatopsPullRequestFragment{
				TimelineItems: types.TimelineItems{
					Nodes: []types.TimelineItem{
						{
							PullRequestCommitFragment: types.PullRequestCommitFragment{
								Commit: types.Commit{
									CommittedDate: queryDate.AddDate(0, 0, -3),
								},
							},
						},
						{
							IssueCommentFragment: types.IssueCommentFragment{
								CreatedAt: queryDate.AddDate(0, 0, -1),
								BodyText: `/retest

another line
even more lines
`,
							},
						},
					},
				},
			},
			1),
		table.Entry("One retest call after last HeadRef force push, previous after commit/BaseRef force push are ignored",
			&types.ChatopsPullRequestFragment{
				TimelineItems: types.TimelineItems{
					Nodes: []types.TimelineItem{
						{
							PullRequestCommitFragment: types.PullRequestCommitFragment{
								Commit: types.Commit{
									CommittedDate: queryDate.AddDate(0, 0, -6),
								},
							},
						},
						{
							IssueCommentFragment: types.IssueCommentFragment{
								CreatedAt: queryDate.AddDate(0, 0, -5),
								BodyText:  "/retest",
							},
						},
						{
							BaseRefForcePushFragment: types.BaseRefForcePushFragment{
								CreatedAt: queryDate.AddDate(0, 0, -4),
								Actor: types.Actor{
									Login: "user1",
								},
							},
						},
						{
							IssueCommentFragment: types.IssueCommentFragment{
								CreatedAt: queryDate.AddDate(0, 0, -3),
								BodyText:  "/retest",
							},
						},
						{
							HeadRefForcePushFragment: types.HeadRefForcePushFragment{
								CreatedAt: queryDate.AddDate(0, 0, -2),
								Actor: types.Actor{
									Login: "user1",
								},
							},
						},
						{
							IssueCommentFragment: types.IssueCommentFragment{
								CreatedAt: queryDate.AddDate(0, 0, -1),
								BodyText:  "/retest",
							},
						},
					},
				},
			},
			1),
		table.Entry("One retest call after last BaseRef force push, previous after commit/HeadRef force push are ignored",
			&types.ChatopsPullRequestFragment{
				TimelineItems: types.TimelineItems{
					Nodes: []types.TimelineItem{
						{
							PullRequestCommitFragment: types.PullRequestCommitFragment{
								Commit: types.Commit{
									CommittedDate: queryDate.AddDate(0, 0, -6),
								},
							},
						},
						{
							IssueCommentFragment: types.IssueCommentFragment{
								CreatedAt: queryDate.AddDate(0, 0, -5),
								BodyText:  "/retest",
							},
						},
						{
							HeadRefForcePushFragment: types.HeadRefForcePushFragment{
								CreatedAt: queryDate.AddDate(0, 0, -4),
								Actor: types.Actor{
									Login: "user1",
								},
							},
						},
						{
							IssueCommentFragment: types.IssueCommentFragment{
								CreatedAt: queryDate.AddDate(0, 0, -3),
								BodyText:  "/retest",
							},
						},
						{
							BaseRefForcePushFragment: types.BaseRefForcePushFragment{
								CreatedAt: queryDate.AddDate(0, 0, -2),
								Actor: types.Actor{
									Login: "user1",
								},
							},
						},
						{
							IssueCommentFragment: types.IssueCommentFragment{
								CreatedAt: queryDate.AddDate(0, 0, -1),
								BodyText:  "/retest",
							},
						},
					},
				},
			},
			1),
	)
})
