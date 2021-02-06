package constants

const (
	DateFormat = "2006-01-02T15:04:05Z"

	LGTMLabel     = "lgtm"
	ApprovedLabel = "approved"

	DoNotMergeLabelPattern = "do-not-merge/*"
	NeedsRebaseLabel       = "needs-rebase"

	MergeQueueLengthName = "AverageMergeQueueLength"
	TimeToMergeName      = "AverageTimeToMerge"

	DefaultPath                        = "/tmp/test"
	DefaultTokenPath                   = ""
	DefaultSource                      = "kubevirt/kubevirt"
	DefaultDataDays                    = 7
	DefaultLogLevel                    = "info"
	DefaultTimeToMergeYellowLevel      = 1.0
	DefaultTimeToMergeRedLevel         = 3.0
	DefaultMergeQueueLengthYellowLevel = 8.0
	DefaultMergeQueueLengthRedLevel    = 15.0

	TimeToMergeBadgeFileName      = "time-to-merge.svg"
	MergeQueueLengthBadgeFileName = "merge-queue-length.svg"
	JSONResultsFileName           = "results.json"
	MetricsFileName               = "metrics"

	TimeToMergeBadgeName      = "avg time to merge"
	MergeQueueLengthBadgeName = "avg merge queue length"
	BadgeDataFormat           = "%.2f ± std %.2f"

	AvgMergeQueueLengthMetricName = "cihealth_avg_merge_queue_lenght_total"
	AvgTimeToMergeMetricName      = "cihealth_avg_time_to_merge_days"
)

func DoNotMergeLabels() []string {
	return []string{
		NeedsRebaseLabel,
		DoNotMergeLabelPattern,
	}
}

func RequiredForMergeLabels() []string {
	return []string{
		LGTMLabel,
		ApprovedLabel,
	}
}
