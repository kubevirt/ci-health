package gh_test

import (
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	"github.com/fgimenez/ci-health/pkg/constants"
	"github.com/fgimenez/ci-health/pkg/gh"
)

func TestGH(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "gh Suite")
}

var (
	queryDate = time.Date(2021, time.Month(1), 22, 15, 11, 26, 0, time.UTC)
	zeroDate  = time.Time{}
)

var _ = Describe("DatePREnteredMergeQueue", func() {
	table.DescribeTable("should return the right values for different timeline events",
		func(pr *gh.PullRequestFragment, date time.Time, expected time.Time) {
			actual := gh.DatePREnteredMergeQueue(pr, date)
			Expect(actual).To(Equal(expected))
		},
		table.Entry("PR in merge queue: approved first",
			&gh.PullRequestFragment{
				TimelineItems: gh.TimelineItems{
					Nodes: []gh.TimelineItem{
						{
							LabeledEventFragment: gh.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -2),
								AddedLabel: gh.Label{
									Name: constants.ApprovedLabel,
								},
							},
						},
						{
							LabeledEventFragment: gh.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -1),
								AddedLabel: gh.Label{
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
			&gh.PullRequestFragment{
				TimelineItems: gh.TimelineItems{
					Nodes: []gh.TimelineItem{
						{
							LabeledEventFragment: gh.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -2),
								AddedLabel: gh.Label{
									Name: constants.LGTMLabel,
								},
							},
						},
						{
							LabeledEventFragment: gh.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -1),
								AddedLabel: gh.Label{
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
			&gh.PullRequestFragment{
				TimelineItems: gh.TimelineItems{
					Nodes: []gh.TimelineItem{
						{
							LabeledEventFragment: gh.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -7),
								AddedLabel: gh.Label{
									Name: constants.ApprovedLabel,
								},
							},
						},
						{
							LabeledEventFragment: gh.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -6),
								AddedLabel: gh.Label{
									Name: constants.LGTMLabel,
								},
							},
						},
						{
							UnlabeledEventFragment: gh.UnlabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -5),
								RemovedLabel: gh.Label{
									Name: constants.LGTMLabel,
								},
							},
						},
						{
							LabeledEventFragment: gh.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -4),
								AddedLabel: gh.Label{
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
			&gh.PullRequestFragment{
				TimelineItems: gh.TimelineItems{
					Nodes: []gh.TimelineItem{
						{
							LabeledEventFragment: gh.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -7),
								AddedLabel: gh.Label{
									Name: constants.ApprovedLabel,
								},
							},
						},
						{
							LabeledEventFragment: gh.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -6),
								AddedLabel: gh.Label{
									Name: constants.HoldLabel,
								},
							},
						},
						{
							LabeledEventFragment: gh.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -5),
								AddedLabel: gh.Label{
									Name: constants.LGTMLabel,
								},
							},
						},
						{
							UnlabeledEventFragment: gh.UnlabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -4),
								RemovedLabel: gh.Label{
									Name: constants.HoldLabel,
								},
							},
						},
					},
				},
			},
			queryDate,
			queryDate.AddDate(0, 0, -4)),
		table.Entry("PR in merge queue: needs-rebase and rebase",
			&gh.PullRequestFragment{
				TimelineItems: gh.TimelineItems{
					Nodes: []gh.TimelineItem{
						{
							LabeledEventFragment: gh.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -7),
								AddedLabel: gh.Label{
									Name: constants.ApprovedLabel,
								},
							},
						},
						{
							LabeledEventFragment: gh.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -6),
								AddedLabel: gh.Label{
									Name: constants.NeedsRebaseLabel,
								},
							},
						},
						{
							LabeledEventFragment: gh.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -5),
								AddedLabel: gh.Label{
									Name: constants.LGTMLabel,
								},
							},
						},
						{
							UnlabeledEventFragment: gh.UnlabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -4),
								RemovedLabel: gh.Label{
									Name: constants.NeedsRebaseLabel,
								},
							},
						},
					},
				},
			},
			queryDate,
			queryDate.AddDate(0, 0, -4)),
		table.Entry("PR not in merge queue: missing all labels",
			&gh.PullRequestFragment{
				TimelineItems: gh.TimelineItems{
					Nodes: []gh.TimelineItem{},
				},
			},
			queryDate,
			zeroDate),
		table.Entry("PR not in merge queue: missing lgtm label",
			&gh.PullRequestFragment{
				TimelineItems: gh.TimelineItems{
					Nodes: []gh.TimelineItem{
						{
							LabeledEventFragment: gh.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -1),
								AddedLabel: gh.Label{
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
			&gh.PullRequestFragment{
				TimelineItems: gh.TimelineItems{
					Nodes: []gh.TimelineItem{
						{
							LabeledEventFragment: gh.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -1),
								AddedLabel: gh.Label{
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
			&gh.PullRequestFragment{
				TimelineItems: gh.TimelineItems{
					Nodes: []gh.TimelineItem{
						{
							LabeledEventFragment: gh.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -1),
								AddedLabel: gh.Label{
									Name: constants.LGTMLabel,
								},
							},
						},
						{
							LabeledEventFragment: gh.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, 1),
								AddedLabel: gh.Label{
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
			&gh.PullRequestFragment{
				TimelineItems: gh.TimelineItems{
					Nodes: []gh.TimelineItem{
						{
							LabeledEventFragment: gh.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -1),
								AddedLabel: gh.Label{
									Name: constants.ApprovedLabel,
								},
							},
						},
						{
							LabeledEventFragment: gh.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, 1),
								AddedLabel: gh.Label{
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
			&gh.PullRequestFragment{
				TimelineItems: gh.TimelineItems{
					Nodes: []gh.TimelineItem{
						{
							LabeledEventFragment: gh.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -3),
								AddedLabel: gh.Label{
									Name: constants.ApprovedLabel,
								},
							},
						},
						{
							LabeledEventFragment: gh.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -2),
								AddedLabel: gh.Label{
									Name: constants.LGTMLabel,
								},
							},
						},
						{
							LabeledEventFragment: gh.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -1),
								AddedLabel: gh.Label{
									Name: constants.NeedsRebaseLabel,
								},
							},
						},
					},
				},
			},
			queryDate,
			zeroDate),
		table.Entry("PR not in merge queue: with hold label",
			&gh.PullRequestFragment{
				TimelineItems: gh.TimelineItems{
					Nodes: []gh.TimelineItem{
						{
							LabeledEventFragment: gh.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -3),
								AddedLabel: gh.Label{
									Name: constants.ApprovedLabel,
								},
							},
						},
						{
							LabeledEventFragment: gh.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -2),
								AddedLabel: gh.Label{
									Name: constants.LGTMLabel,
								},
							},
						},
						{
							LabeledEventFragment: gh.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -1),
								AddedLabel: gh.Label{
									Name: constants.HoldLabel,
								},
							},
						},
					},
				},
			},
			queryDate,
			zeroDate),
		table.Entry("PR not in merge queue: was previously in and lost lgtm label",
			&gh.PullRequestFragment{
				TimelineItems: gh.TimelineItems{
					Nodes: []gh.TimelineItem{
						{
							LabeledEventFragment: gh.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -3),
								AddedLabel: gh.Label{
									Name: constants.ApprovedLabel,
								},
							},
						},
						{
							LabeledEventFragment: gh.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -2),
								AddedLabel: gh.Label{
									Name: constants.LGTMLabel,
								},
							},
						},
						{
							UnlabeledEventFragment: gh.UnlabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -1),
								RemovedLabel: gh.Label{
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
			&gh.PullRequestFragment{
				TimelineItems: gh.TimelineItems{
					Nodes: []gh.TimelineItem{
						{
							LabeledEventFragment: gh.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -3),
								AddedLabel: gh.Label{
									Name: constants.ApprovedLabel,
								},
							},
						},
						{
							LabeledEventFragment: gh.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -2),
								AddedLabel: gh.Label{
									Name: constants.LGTMLabel,
								},
							},
						},
						{
							UnlabeledEventFragment: gh.UnlabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -1),
								RemovedLabel: gh.Label{
									Name: constants.ApprovedLabel,
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
