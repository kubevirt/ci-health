package types

import "time"

type Options struct {
	Path                        string
	TokenPath                   string
	Source                      string
	DataDays                    int
	LogLevel                    string
	TimeToMergeRedLevel         float64
	TimeToMergeYellowLevel      float64
	MergeQueueLengthRedLevel    float64
	MergeQueueLengthYellowLevel float64
	RetestsToMergeYellowLevel   float64
	RetestsToMergeRedLevel      float64
}

type Label struct {
	Name string
}

type LabeledEventFragment struct {
	CreatedAt  time.Time
	AddedLabel Label `graphql:"addedLabel:label"`
}

type UnlabeledEventFragment struct {
	CreatedAt    time.Time
	RemovedLabel Label `graphql:"removedLabel:label"`
}

type IssueCommentFragment struct {
	CreatedAt time.Time
	BodyText  string `graphql:"bodyText:string"`
}

type CommitFragment struct {
	PushedDate time.Time
}

type TimelineItem struct {
	LabeledEventFragment   `graphql:"... on LabeledEvent"`
	UnlabeledEventFragment `graphql:"... on UnlabeledEvent"`
	IssueCommentFragment   `graphql:"... on IssueComment"`
	CommitFragment         `graphql:"... on Commit"`
}

type TimelineItems struct {
	Nodes []TimelineItem
}

type PullRequestFragment struct {
	Number        int
	CreatedAt     time.Time
	MergedAt      time.Time
	TimelineItems `graphql:"timelineItems(first:100, itemTypes:[LABELED_EVENT, UNLABELED_EVENT, COMMIT, ISSUE_COMMENT])"`
}
