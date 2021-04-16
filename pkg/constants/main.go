package constants

const (
	DateFormat = "2006-01-02T15:04:05Z"

	LGTMLabel     = "lgtm"
	ApprovedLabel = "approved"

	DoNotMergeLabelPattern = "do-not-merge/*"
	NeedsLabelPattern      = "needs-*"

	MergeQueueLengthName = "AverageMergeQueueLength"
	TimeToMergeName      = "AverageTimeToMerge"
	RetestsToMergeName   = "AverageRetestsToMerge"

	DefaultPath                        = "/tmp/test"
	DefaultTokenPath                   = ""
	DefaultSource                      = "kubevirt/kubevirt"
	DefaultDataDays                    = 7
	DefaultLogLevel                    = "info"
	DefaultTimeToMergeYellowLevel      = 1.0
	DefaultTimeToMergeRedLevel         = 3.0
	DefaultMergeQueueLengthYellowLevel = 8.0
	DefaultMergeQueueLengthRedLevel    = 15.0
	DefaultRetestsToMergeYellowLevel   = 2.0
	DefaultRetestsToMergeRedLevel      = 4.0

	TimeToMergeBadgeFileName      = "time-to-merge.svg"
	MergeQueueLengthBadgeFileName = "merge-queue-length.svg"
	RetestsToMergeBadgeFileName   = "retests-to-merge.svg"
	JSONResultsFileName           = "results.json"
	MetricsFileName               = "metrics"

	TimeToMergeBadgeName      = "days to merge"
	MergeQueueLengthBadgeName = "avg merge queue length"
	RetestsToMergeBadgeName   = "retests to merge"
	BadgeDataFormat           = "%.2f ± std %.2f"

	AvgMergeQueueLengthMetricName = "cihealth_avg_merge_queue_lenght_total"
	AvgTimeToMergeMetricName      = "cihealth_avg_time_to_merge_days"
	AvgRetestsToMergeMetricName   = "cihealth_avg_retests_to_merge_total"
	StdMergeQueueLengthMetricName = "cihealth_std_merge_queue_lenght_total"
	StdTimeToMergeMetricName      = "cihealth_std_time_to_merge_days"
	StdRetestsToMergeMetricName   = "cihealth_std_retests_to_merge_total"
)

func DoNotMergeLabels() []string {
	return []string{
		NeedsLabelPattern,
		DoNotMergeLabelPattern,
	}
}

func RequiredForMergeLabels() []string {
	return []string{
		LGTMLabel,
		ApprovedLabel,
	}
}
