package types

import (
	"errors"
	"time"
)

const (
	StatsAction Action = "stats"
	BatchAction        = "batch"

	MergeQueueLengthMetricType MetricType = "merge-queue-length"
	TimeToMergeMetricType                 = "time-to-merge"
	RetestsToMergeMetricType              = "retests-to-merge"
)

type MetricType string

func (mt MetricType) IsValid() error {
	switch mt {
	case MergeQueueLengthMetricType, TimeToMergeMetricType, RetestsToMergeMetricType:
		return nil
	}
	return errors.New("Invalid MetricType value")
}

type Action string

func (a Action) IsValid() error {
	switch a {
	case StatsAction, BatchAction:
		return nil
	}
	return errors.New("Invalid Action value")
}

type Options struct {
	Action

	Path            string
	TokenPath       string
	Source          string
	DataDays        int
	LogLevel        string
	RequestedAction Action

	// stats options
	TimeToMergeRedLevel         float64
	TimeToMergeYellowLevel      float64
	MergeQueueLengthRedLevel    float64
	MergeQueueLengthYellowLevel float64
	RetestsToMergeYellowLevel   float64
	RetestsToMergeRedLevel      float64

	// batch options
	TargetMetric MetricType
	StartDate    string
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
	BodyText  string
}

type Commit struct {
	PushedDate time.Time
}

type PullRequestCommitFragment struct {
	Commit Commit
}

type TimelineItem struct {
	LabeledEventFragment      `graphql:"... on LabeledEvent"`
	UnlabeledEventFragment    `graphql:"... on UnlabeledEvent"`
	IssueCommentFragment      `graphql:"... on IssueComment"`
	PullRequestCommitFragment `graphql:"... on PullRequestCommit"`
}

type TimelineItems struct {
	Nodes []TimelineItem
}

type ChatopsPullRequestFragment struct {
	Number        int
	CreatedAt     time.Time
	MergedAt      time.Time
	TimelineItems `graphql:"timelineItems(first:100, itemTypes:[PULL_REQUEST_COMMIT, ISSUE_COMMENT])"`
}

type MergeQueuePullRequestFragment struct {
	Number        int
	CreatedAt     time.Time
	MergedAt      time.Time
	TimelineItems `graphql:"timelineItems(first:100, itemTypes:[LABELED_EVENT, UNLABELED_EVENT])"`
}

type BarePullRequestFragment struct {
	Number    int
	CreatedAt time.Time
	MergedAt  time.Time
}

type BarePRList []struct {
	BarePullRequestFragment `graphql:"... on PullRequest"`
}

type MergeQueuePRList []struct {
	MergeQueuePullRequestFragment `graphql:"... on PullRequest"`
}

type ChatopsPRList []struct {
	ChatopsPullRequestFragment `graphql:"... on PullRequest"`
}
