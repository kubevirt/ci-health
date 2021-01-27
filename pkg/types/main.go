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
	TimeToMergeGreenLevel       float64
	MergeQueueLengthRedLevel    float64
	MergeQueueLengthYellowLevel float64
	MergeQueueLengthGreenLevel  float64
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

type TimelineItem struct {
	LabeledEventFragment   `graphql:"... on LabeledEvent"`
	UnlabeledEventFragment `graphql:"... on UnlabeledEvent"`
}

type TimelineItems struct {
	Nodes []TimelineItem
}

type PullRequestFragment struct {
	Number        int
	CreatedAt     time.Time
	MergedAt      time.Time
	TimelineItems `graphql:"timelineItems(first:100, itemTypes:[LABELED_EVENT, UNLABELED_EVENT])"`
}
