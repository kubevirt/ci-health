package mergequeue_test

import (
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	"github.com/fgimenez/ci-health/pkg/constants"
	"github.com/fgimenez/ci-health/pkg/mergequeue"
	"github.com/fgimenez/ci-health/pkg/types"
)

const (
	HoldLabel           = "do-not-merge/hold"
	WorkInProgressLabel = "do-not-merge/work-in-progress"
)

func TestMergeQueue(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "mergequeue Suite")
}

var (
	queryDate = time.Date(2021, time.Month(1), 22, 15, 11, 26, 0, time.UTC)
	zeroDate  = time.Time{}
)

var _ = Describe("DatePREntered", func() {
	table.DescribeTable("should return the date a PR entered the merge queue",
		func(pr *types.PullRequestFragment, date time.Time, expected time.Time) {
			actual := mergequeue.DatePREntered(pr, date)
			Expect(actual).To(Equal(expected))
		},
		table.Entry("PR in merge queue: approved first",
			&types.PullRequestFragment{
				TimelineItems: types.TimelineItems{
					Nodes: []types.TimelineItem{
						{
							LabeledEventFragment: types.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -2),
								AddedLabel: types.Label{
									Name: constants.ApprovedLabel,
								},
							},
						},
						{
							LabeledEventFragment: types.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -1),
								AddedLabel: types.Label{
									Name: constants.LGTMLabel,
								},
							},
						},
					},
				},
			},
			queryDate,
			queryDate.AddDate(0, 0, -1)),
		table.Entry("PR in merge queue: lgtm first",
			&types.PullRequestFragment{
				TimelineItems: types.TimelineItems{
					Nodes: []types.TimelineItem{
						{
							LabeledEventFragment: types.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -2),
								AddedLabel: types.Label{
									Name: constants.LGTMLabel,
								},
							},
						},
						{
							LabeledEventFragment: types.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -1),
								AddedLabel: types.Label{
									Name: constants.ApprovedLabel,
								},
							},
						},
					},
				},
			},
			queryDate,
			queryDate.AddDate(0, 0, -1)),
		table.Entry("PR in merge queue: reentered after leaving",
			&types.PullRequestFragment{
				TimelineItems: types.TimelineItems{
					Nodes: []types.TimelineItem{
						{
							LabeledEventFragment: types.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -7),
								AddedLabel: types.Label{
									Name: constants.ApprovedLabel,
								},
							},
						},
						{
							LabeledEventFragment: types.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -6),
								AddedLabel: types.Label{
									Name: constants.LGTMLabel,
								},
							},
						},
						{
							UnlabeledEventFragment: types.UnlabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -5),
								RemovedLabel: types.Label{
									Name: constants.LGTMLabel,
								},
							},
						},
						{
							LabeledEventFragment: types.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -4),
								AddedLabel: types.Label{
									Name: constants.LGTMLabel,
								},
							},
						},
					},
				},
			},
			queryDate,
			queryDate.AddDate(0, 0, -4)),
		table.Entry("PR in merge queue: hold and unhold",
			&types.PullRequestFragment{
				TimelineItems: types.TimelineItems{
					Nodes: []types.TimelineItem{
						{
							LabeledEventFragment: types.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -7),
								AddedLabel: types.Label{
									Name: constants.ApprovedLabel,
								},
							},
						},
						{
							LabeledEventFragment: types.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -6),
								AddedLabel: types.Label{
									Name: HoldLabel,
								},
							},
						},
						{
							LabeledEventFragment: types.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -5),
								AddedLabel: types.Label{
									Name: constants.LGTMLabel,
								},
							},
						},
						{
							UnlabeledEventFragment: types.UnlabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -4),
								RemovedLabel: types.Label{
									Name: HoldLabel,
								},
							},
						},
					},
				},
			},
			queryDate,
			queryDate.AddDate(0, 0, -4)),
		table.Entry("PR in merge queue: work-in-progress and removed work-in-progress",
			&types.PullRequestFragment{
				TimelineItems: types.TimelineItems{
					Nodes: []types.TimelineItem{
						{
							LabeledEventFragment: types.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -7),
								AddedLabel: types.Label{
									Name: constants.ApprovedLabel,
								},
							},
						},
						{
							LabeledEventFragment: types.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -6),
								AddedLabel: types.Label{
									Name: WorkInProgressLabel,
								},
							},
						},
						{
							LabeledEventFragment: types.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -5),
								AddedLabel: types.Label{
									Name: constants.LGTMLabel,
								},
							},
						},
						{
							UnlabeledEventFragment: types.UnlabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -4),
								RemovedLabel: types.Label{
									Name: WorkInProgressLabel,
								},
							},
						},
					},
				},
			},
			queryDate,
			queryDate.AddDate(0, 0, -4)),
		table.Entry("PR in merge queue: any do-not-merge prefixed label and removed any do-not-merge prefixed label",
			&types.PullRequestFragment{
				TimelineItems: types.TimelineItems{
					Nodes: []types.TimelineItem{
						{
							LabeledEventFragment: types.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -7),
								AddedLabel: types.Label{
									Name: constants.ApprovedLabel,
								},
							},
						},
						{
							LabeledEventFragment: types.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -6),
								AddedLabel: types.Label{
									Name: "do-not-merge/for-whatever-reason",
								},
							},
						},
						{
							LabeledEventFragment: types.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -5),
								AddedLabel: types.Label{
									Name: constants.LGTMLabel,
								},
							},
						},
						{
							UnlabeledEventFragment: types.UnlabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -4),
								RemovedLabel: types.Label{
									Name: "do-not-merge/for-whatever-reason",
								},
							},
						},
					},
				},
			},
			queryDate,
			queryDate.AddDate(0, 0, -4)),
		table.Entry("PR in merge queue: needs-rebase and rebase",
			&types.PullRequestFragment{
				TimelineItems: types.TimelineItems{
					Nodes: []types.TimelineItem{
						{
							LabeledEventFragment: types.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -7),
								AddedLabel: types.Label{
									Name: constants.ApprovedLabel,
								},
							},
						},
						{
							LabeledEventFragment: types.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -6),
								AddedLabel: types.Label{
									Name: "needs-rebase",
								},
							},
						},
						{
							LabeledEventFragment: types.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -5),
								AddedLabel: types.Label{
									Name: constants.LGTMLabel,
								},
							},
						},
						{
							UnlabeledEventFragment: types.UnlabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -4),
								RemovedLabel: types.Label{
									Name: "needs-rebase",
								},
							},
						},
					},
				},
			},
			queryDate,
			queryDate.AddDate(0, 0, -4)),
		table.Entry("PR in merge queue: needs-whatever and removed needs-whatever",
			&types.PullRequestFragment{
				TimelineItems: types.TimelineItems{
					Nodes: []types.TimelineItem{
						{
							LabeledEventFragment: types.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -7),
								AddedLabel: types.Label{
									Name: constants.ApprovedLabel,
								},
							},
						},
						{
							LabeledEventFragment: types.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -6),
								AddedLabel: types.Label{
									Name: "needs-whatever",
								},
							},
						},
						{
							LabeledEventFragment: types.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -5),
								AddedLabel: types.Label{
									Name: constants.LGTMLabel,
								},
							},
						},
						{
							UnlabeledEventFragment: types.UnlabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -4),
								RemovedLabel: types.Label{
									Name: "needs-whatever",
								},
							},
						},
					},
				},
			},
			queryDate,
			queryDate.AddDate(0, 0, -4)),
		table.Entry("PR in merge queue: hold after queryDate",
			&types.PullRequestFragment{
				TimelineItems: types.TimelineItems{
					Nodes: []types.TimelineItem{
						{
							LabeledEventFragment: types.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -2),
								AddedLabel: types.Label{
									Name: constants.ApprovedLabel,
								},
							},
						},
						{
							LabeledEventFragment: types.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -1),
								AddedLabel: types.Label{
									Name: constants.LGTMLabel,
								},
							},
						},
						{
							LabeledEventFragment: types.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, 1),
								AddedLabel: types.Label{
									Name: HoldLabel,
								},
							},
						},
					},
				},
			},
			queryDate,
			queryDate.AddDate(0, 0, -1)),
		table.Entry("PR in merge queue: work-in-progress after queryDate",
			&types.PullRequestFragment{
				TimelineItems: types.TimelineItems{
					Nodes: []types.TimelineItem{
						{
							LabeledEventFragment: types.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -2),
								AddedLabel: types.Label{
									Name: constants.ApprovedLabel,
								},
							},
						},
						{
							LabeledEventFragment: types.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -1),
								AddedLabel: types.Label{
									Name: constants.LGTMLabel,
								},
							},
						},
						{
							LabeledEventFragment: types.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, 1),
								AddedLabel: types.Label{
									Name: WorkInProgressLabel,
								},
							},
						},
					},
				},
			},
			queryDate,
			queryDate.AddDate(0, 0, -1)),
		table.Entry("PR in merge queue: needs-rebase after queryDate",
			&types.PullRequestFragment{
				TimelineItems: types.TimelineItems{
					Nodes: []types.TimelineItem{
						{
							LabeledEventFragment: types.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -2),
								AddedLabel: types.Label{
									Name: constants.ApprovedLabel,
								},
							},
						},
						{
							LabeledEventFragment: types.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -1),
								AddedLabel: types.Label{
									Name: constants.LGTMLabel,
								},
							},
						},
						{
							LabeledEventFragment: types.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, 1),
								AddedLabel: types.Label{
									Name: "needs-rebase",
								},
							},
						},
					},
				},
			},
			queryDate,
			queryDate.AddDate(0, 0, -1)),
		table.Entry("PR in merge queue: needs-whatever after queryDate",
			&types.PullRequestFragment{
				TimelineItems: types.TimelineItems{
					Nodes: []types.TimelineItem{
						{
							LabeledEventFragment: types.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -2),
								AddedLabel: types.Label{
									Name: constants.ApprovedLabel,
								},
							},
						},
						{
							LabeledEventFragment: types.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -1),
								AddedLabel: types.Label{
									Name: constants.LGTMLabel,
								},
							},
						},
						{
							LabeledEventFragment: types.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, 1),
								AddedLabel: types.Label{
									Name: "needs-whatever",
								},
							},
						},
					},
				},
			},
			queryDate,
			queryDate.AddDate(0, 0, -1)),
		table.Entry("PR in merge queue: handle empty events",
			&types.PullRequestFragment{
				TimelineItems: types.TimelineItems{
					Nodes: []types.TimelineItem{
						{
							LabeledEventFragment: types.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -4),
								AddedLabel: types.Label{
									Name: constants.ApprovedLabel,
								},
							},
						},
						{
							LabeledEventFragment: types.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -3),
								AddedLabel: types.Label{
									Name: "needs-rebase",
								},
							},
						},
						{
							LabeledEventFragment: types.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -2),
								AddedLabel: types.Label{
									Name: "",
								},
							},
							UnlabeledEventFragment: types.UnlabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -2),
								RemovedLabel: types.Label{
									Name: "needs-rebase",
								},
							},
						},
						{
							LabeledEventFragment: types.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -1),
								AddedLabel: types.Label{
									Name: constants.LGTMLabel,
								},
							},
						},
					},
				},
			},
			queryDate,
			queryDate.AddDate(0, 0, -1)),
		table.Entry("PR not in merge queue: missing all labels",
			&types.PullRequestFragment{
				TimelineItems: types.TimelineItems{
					Nodes: []types.TimelineItem{},
				},
			},
			queryDate,
			zeroDate),
		table.Entry("PR not in merge queue: missing lgtm label",
			&types.PullRequestFragment{
				TimelineItems: types.TimelineItems{
					Nodes: []types.TimelineItem{
						{
							LabeledEventFragment: types.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -1),
								AddedLabel: types.Label{
									Name: constants.ApprovedLabel,
								},
							},
						},
					},
				},
			},
			queryDate,
			zeroDate),
		table.Entry("PR not in merge queue: missing approved label",
			&types.PullRequestFragment{
				TimelineItems: types.TimelineItems{
					Nodes: []types.TimelineItem{
						{
							LabeledEventFragment: types.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -1),
								AddedLabel: types.Label{
									Name: constants.LGTMLabel,
								},
							},
						},
					},
				},
			},
			queryDate,
			zeroDate),
		table.Entry("PR not in merge queue: approved label after query date",
			&types.PullRequestFragment{
				TimelineItems: types.TimelineItems{
					Nodes: []types.TimelineItem{
						{
							LabeledEventFragment: types.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -1),
								AddedLabel: types.Label{
									Name: constants.LGTMLabel,
								},
							},
						},
						{
							LabeledEventFragment: types.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, 1),
								AddedLabel: types.Label{
									Name: constants.ApprovedLabel,
								},
							},
						},
					},
				},
			},
			queryDate,
			zeroDate),
		table.Entry("PR not in merge queue: lgtm label after query date",
			&types.PullRequestFragment{
				TimelineItems: types.TimelineItems{
					Nodes: []types.TimelineItem{
						{
							LabeledEventFragment: types.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -1),
								AddedLabel: types.Label{
									Name: constants.ApprovedLabel,
								},
							},
						},
						{
							LabeledEventFragment: types.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, 1),
								AddedLabel: types.Label{
									Name: constants.LGTMLabel,
								},
							},
						},
					},
				},
			},
			queryDate,
			zeroDate),
		table.Entry("PR not in merge queue: with needs-rebase label",
			&types.PullRequestFragment{
				TimelineItems: types.TimelineItems{
					Nodes: []types.TimelineItem{
						{
							LabeledEventFragment: types.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -3),
								AddedLabel: types.Label{
									Name: constants.ApprovedLabel,
								},
							},
						},
						{
							LabeledEventFragment: types.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -2),
								AddedLabel: types.Label{
									Name: constants.LGTMLabel,
								},
							},
						},
						{
							LabeledEventFragment: types.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -1),
								AddedLabel: types.Label{
									Name: "needs-rebase",
								},
							},
						},
					},
				},
			},
			queryDate,
			zeroDate),
		table.Entry("PR not in merge queue: with needs-whatever label",
			&types.PullRequestFragment{
				TimelineItems: types.TimelineItems{
					Nodes: []types.TimelineItem{
						{
							LabeledEventFragment: types.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -3),
								AddedLabel: types.Label{
									Name: constants.ApprovedLabel,
								},
							},
						},
						{
							LabeledEventFragment: types.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -2),
								AddedLabel: types.Label{
									Name: constants.LGTMLabel,
								},
							},
						},
						{
							LabeledEventFragment: types.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -1),
								AddedLabel: types.Label{
									Name: "need-whatever",
								},
							},
						},
					},
				},
			},
			queryDate,
			zeroDate),
		table.Entry("PR not in merge queue: with hold label",
			&types.PullRequestFragment{
				TimelineItems: types.TimelineItems{
					Nodes: []types.TimelineItem{
						{
							LabeledEventFragment: types.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -3),
								AddedLabel: types.Label{
									Name: constants.ApprovedLabel,
								},
							},
						},
						{
							LabeledEventFragment: types.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -2),
								AddedLabel: types.Label{
									Name: constants.LGTMLabel,
								},
							},
						},
						{
							LabeledEventFragment: types.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -1),
								AddedLabel: types.Label{
									Name: HoldLabel,
								},
							},
						},
					},
				},
			},
			queryDate,
			zeroDate),
		table.Entry("PR not in merge queue: with work-in-progress label",
			&types.PullRequestFragment{
				TimelineItems: types.TimelineItems{
					Nodes: []types.TimelineItem{
						{
							LabeledEventFragment: types.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -3),
								AddedLabel: types.Label{
									Name: constants.ApprovedLabel,
								},
							},
						},
						{
							LabeledEventFragment: types.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -2),
								AddedLabel: types.Label{
									Name: constants.LGTMLabel,
								},
							},
						},
						{
							LabeledEventFragment: types.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -1),
								AddedLabel: types.Label{
									Name: WorkInProgressLabel,
								},
							},
						},
					},
				},
			},
			queryDate,
			zeroDate),
		table.Entry("PR not in merge queue: with any do-not-merge prefixed label",
			&types.PullRequestFragment{
				TimelineItems: types.TimelineItems{
					Nodes: []types.TimelineItem{
						{
							LabeledEventFragment: types.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -3),
								AddedLabel: types.Label{
									Name: constants.ApprovedLabel,
								},
							},
						},
						{
							LabeledEventFragment: types.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -2),
								AddedLabel: types.Label{
									Name: constants.LGTMLabel,
								},
							},
						},
						{
							LabeledEventFragment: types.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -1),
								AddedLabel: types.Label{
									Name: "do-not-merge/for-whatever-reason",
								},
							},
						},
					},
				},
			},
			queryDate,
			zeroDate),
		table.Entry("PR not in merge queue: was previously in and lost lgtm label",
			&types.PullRequestFragment{
				TimelineItems: types.TimelineItems{
					Nodes: []types.TimelineItem{
						{
							LabeledEventFragment: types.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -3),
								AddedLabel: types.Label{
									Name: constants.ApprovedLabel,
								},
							},
						},
						{
							LabeledEventFragment: types.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -2),
								AddedLabel: types.Label{
									Name: constants.LGTMLabel,
								},
							},
						},
						{
							UnlabeledEventFragment: types.UnlabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -1),
								RemovedLabel: types.Label{
									Name: constants.LGTMLabel,
								},
							},
						},
					},
				},
			},
			queryDate,
			zeroDate),
		table.Entry("PR not in merge queue: was previously in and lost approved label",
			&types.PullRequestFragment{
				TimelineItems: types.TimelineItems{
					Nodes: []types.TimelineItem{
						{
							LabeledEventFragment: types.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -3),
								AddedLabel: types.Label{
									Name: constants.ApprovedLabel,
								},
							},
						},
						{
							LabeledEventFragment: types.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -2),
								AddedLabel: types.Label{
									Name: constants.LGTMLabel,
								},
							},
						},
						{
							UnlabeledEventFragment: types.UnlabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -1),
								RemovedLabel: types.Label{
									Name: constants.ApprovedLabel,
								},
							},
						},
					},
				},
			},
			queryDate,
			zeroDate),
		table.Entry("PR not in merge queue: lgtm after queryDate",
			&types.PullRequestFragment{
				TimelineItems: types.TimelineItems{
					Nodes: []types.TimelineItem{
						{
							LabeledEventFragment: types.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -1),
								AddedLabel: types.Label{
									Name: constants.ApprovedLabel,
								},
							},
						},
						{
							LabeledEventFragment: types.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, 1),
								AddedLabel: types.Label{
									Name: constants.LGTMLabel,
								},
							},
						},
					},
				},
			},
			queryDate,
			zeroDate),
		table.Entry("PR not in merge queue: approve after queryDate",
			&types.PullRequestFragment{
				TimelineItems: types.TimelineItems{
					Nodes: []types.TimelineItem{
						{
							LabeledEventFragment: types.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -1),
								AddedLabel: types.Label{
									Name: constants.LGTMLabel,
								},
							},
						},
						{
							LabeledEventFragment: types.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, 1),
								AddedLabel: types.Label{
									Name: constants.ApprovedLabel,
								},
							},
						},
					},
				},
			},
			queryDate,
			zeroDate),
		table.Entry("PR not in merge queue: lgtm, approve and hold at the same time",
			&types.PullRequestFragment{
				TimelineItems: types.TimelineItems{
					Nodes: []types.TimelineItem{
						{
							LabeledEventFragment: types.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -1),
								AddedLabel: types.Label{
									Name: constants.LGTMLabel,
								},
							},
						},
						{
							LabeledEventFragment: types.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -1),
								AddedLabel: types.Label{
									Name: constants.ApprovedLabel,
								},
							},
						},
						{
							LabeledEventFragment: types.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -1),
								AddedLabel: types.Label{
									Name: HoldLabel,
								},
							},
						},
					},
				},
			},
			queryDate,
			zeroDate),
		table.Entry("PR not in merge queue: lgtm, approve and work-in-progress at the same time",
			&types.PullRequestFragment{
				TimelineItems: types.TimelineItems{
					Nodes: []types.TimelineItem{
						{
							LabeledEventFragment: types.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -1),
								AddedLabel: types.Label{
									Name: constants.LGTMLabel,
								},
							},
						},
						{
							LabeledEventFragment: types.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -1),
								AddedLabel: types.Label{
									Name: constants.ApprovedLabel,
								},
							},
						},
						{
							LabeledEventFragment: types.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -1),
								AddedLabel: types.Label{
									Name: WorkInProgressLabel,
								},
							},
						},
					},
				},
			},
			queryDate,
			zeroDate),
		table.Entry("PR not in merge queue: handle empty events",
			&types.PullRequestFragment{
				TimelineItems: types.TimelineItems{
					Nodes: []types.TimelineItem{
						{
							LabeledEventFragment: types.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -3),
								AddedLabel: types.Label{
									Name: constants.ApprovedLabel,
								},
							},
						},
						{
							LabeledEventFragment: types.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -2),
								AddedLabel: types.Label{
									Name: constants.LGTMLabel,
								},
							},
						},
						{
							LabeledEventFragment: types.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -1),
								AddedLabel: types.Label{
									Name: "",
								},
							},
							UnlabeledEventFragment: types.UnlabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -1),
								RemovedLabel: types.Label{
									Name: constants.LGTMLabel,
								},
							},
						},
					},
				},
			},
			queryDate,
			zeroDate),
	)
})
