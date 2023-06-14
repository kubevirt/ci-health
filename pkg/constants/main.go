package constants

const (
	DateFormat          = "2006-01-02T15:04:05Z"
	BatchDataDateFormat = "2006-01-02"

	LGTMLabel     = "lgtm"
	ApprovedLabel = "approved"

	DoNotMergeLabelPattern = "do-not-merge/*"
	NeedsLabelPattern      = "needs-*"

	MergeQueueLengthName = "AverageMergeQueueLength"
	TimeToMergeName      = "AverageTimeToMerge"
	RetestsToMergeName   = "AverageRetestsToMerge"
	MergedPRsName        = "AverageMergedPRs"
	MergedPRsNoRetest    = "PRsMergedWithNoRetest"

	DefaultPath                         = "/tmp/test"
	DefaultTokenPath                    = ""
	DefaultSource                       = "kubevirt/kubevirt"
	DefaultDataDays                     = 7
	DefaultLogLevel                     = "info"
	DefaultTimeToMergeYellowLevel       = 1.0
	DefaultTimeToMergeRedLevel          = 3.0
	DefaultMergeQueueLengthYellowLevel  = 8.0
	DefaultMergeQueueLengthRedLevel     = 15.0
	DefaultRetestsToMergeYellowLevel    = 3.0
	DefaultRetestsToMergeRedLevel       = 6.0
	DefaultMergedPRsYellowLevel         = 10.0
	DefaultMergedPRsRedLevel            = 7.0
	DefaultMergedPRsNoRetestYellowLevel = 0.5
	DefaultMergedPRsNoRetestRedLevel    = 0.25
	DefaultBatchStartDate               = "2019-05-06"

	TimeToMergeBadgeFileName       = "time-to-merge.svg"
	MergeQueueLengthBadgeFileName  = "merge-queue-length.svg"
	RetestsToMergeBadgeFileName    = "retests-to-merge.svg"
	MergedPRsBadgeFileName         = "merged-prs.svg"
	MergedPRsNoRetestBadgeFileName = "merged-prs-no-retest.svg"
	JSONResultsFileName            = "results.json"
	PlotFileName                   = "plot.png"
	MetricsFileName                = "metrics"

	TimeToMergeBadgeName       = "days to merge"
	MergeQueueLengthBadgeName  = "avg merge queue length"
	RetestsToMergeBadgeName    = "retests to merge"
	MergedPRsBadgeName         = "merged PRs"
	MergedPRsNoRetestBadgeName = "Merged PRs with 0 retest vs. Merged PRs"
	BadgeDataFormat            = "%.2f ± std %.2f"
	NoRetestBadgeDataFormat    = "%.0f / %.0f"

	AvgMergeQueueLengthMetricName = "cihealth_avg_merge_queue_lenght_total"
	AvgTimeToMergeMetricName      = "cihealth_avg_time_to_merge_days"
	AvgRetestsToMergeMetricName   = "cihealth_avg_retests_to_merge_total"
	AvgMergedPRsMetricName        = "cihealth_avg_merged_prs_total"
	StdMergeQueueLengthMetricName = "cihealth_std_merge_queue_lenght_total"
	StdTimeToMergeMetricName      = "cihealth_std_time_to_merge_days"
	StdRetestsToMergeMetricName   = "cihealth_std_retests_to_merge_total"
	StdMergedPRsMetricName        = "cihealth_std_merged_prs_total"

	DefaultBatchBaseOutputPath = "batch"
	DefaultBatchDataOutputPath = "data"
	DefaultBatchPlotOutputPath = "plot"
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
