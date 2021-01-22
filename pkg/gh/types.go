package gh

import "time"

type LabeledEventFragment struct {
	CreatedAt  time.Time
	AddedLabel struct {
		Name string
	} `graphql:"addedLabel:label"`
}

type UnlabeledEventFragment struct {
	CreatedAt    time.Time
	RemovedLabel struct {
		Name string
	} `graphql:"removedLabel:label"`
}

type TimelineItem struct {
	LabeledEventFragment   `graphql:"... on LabeledEvent"`
	UnlabeledEventFragment `graphql:"... on UnlabeledEvent"`
}

type PullRequestFragment struct {
	Number        int
	TimelineItems struct {
		Nodes []TimelineItem
	} `graphql:"timelineItems(first:100, itemTypes:[LABELED_EVENT, UNLABELED_EVENT])"`
}
