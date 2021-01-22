package gh

import "time"

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
	TimelineItems `graphql:"timelineItems(first:100, itemTypes:[LABELED_EVENT, UNLABELED_EVENT])"`
}
