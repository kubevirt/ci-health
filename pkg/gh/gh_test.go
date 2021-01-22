package gh_test

import (
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	"github.com/fgimenez/cihealth/pkg/constants"
	"github.com/fgimenez/cihealth/pkg/gh"
)

func TestGH(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "gh Suite")
}

var (
	queryDate = time.Date(2021, time.Month(1), 22, 15, 11, 26, 0, time.UTC)
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
								CreatedAt: queryDate.AddDate(0, 0, -1),
								AddedLabel: gh.Label{
									Name: constants.LGTMLabel,
								},
							},
						},
						{
							LabeledEventFragment: gh.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -2),
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
								CreatedAt: queryDate.AddDate(0, 0, -4),
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
								CreatedAt: queryDate.AddDate(0, 0, -6),
								AddedLabel: gh.Label{
									Name: constants.LGTMLabel,
								},
							},
						},
						{
							LabeledEventFragment: gh.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -7),
								AddedLabel: gh.Label{
									Name: constants.ApprovedLabel,
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
							UnlabeledEventFragment: gh.UnlabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -4),
								RemovedLabel: gh.Label{
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
							LabeledEventFragment: gh.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -6),
								AddedLabel: gh.Label{
									Name: constants.HoldLabel,
								},
							},
						},
						{
							LabeledEventFragment: gh.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -7),
								AddedLabel: gh.Label{
									Name: constants.ApprovedLabel,
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
							UnlabeledEventFragment: gh.UnlabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -4),
								RemovedLabel: gh.Label{
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
							LabeledEventFragment: gh.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -6),
								AddedLabel: gh.Label{
									Name: constants.NeedsRebaseLabel,
								},
							},
						},
						{
							LabeledEventFragment: gh.LabeledEventFragment{
								CreatedAt: queryDate.AddDate(0, 0, -7),
								AddedLabel: gh.Label{
									Name: constants.ApprovedLabel,
								},
							},
						},
					},
				},
			},
			queryDate,
			queryDate.AddDate(0, 0, -4)),

		/*
				table.Entry("PR not in merge queue: missing all labels"),
				table.Entry("PR not in merge queue: missing lgtm label"),
				table.Entry("PR not in merge queue: missing approved label"),
				table.Entry("PR not in merge queue: with needs-rebase label"),
			  table.Entry("PR not in merge queue: with hold label"),
				table.Entry("PR not in merge queue: was previously and lost lgtm label"),
				table.Entry("PR not in merge queue: was previously and lost approved label"),
		*/
	)
})
