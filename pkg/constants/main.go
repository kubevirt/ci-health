package constants

const (
	DateFormat = "2006-01-02T15:04:05Z"

	LGTMLabel        = "lgtm"
	ApprovedLabel    = "approved"
	HoldLabel        = "do-not-merge/hold"
	NeedsRebaseLabel = "needs-rebase"

	MergeQueueLengthName = "AverageMergeQueueLength"
	TimeToMergeName      = "AverageTimeToMerge"

	DefaultPath                        = "/tmp/test"
	DefaultTokenPath                   = ""
	DefaultSource                      = "kubevirt/kubevirt"
	DefaultDataDays                    = 7
	DefaultLogLevel                    = "info"
	DefaultTimeToMergeYellowLevel      = 3.0
	DefaultTimeToMergeGreenLevel       = 2.0
	DefaultMergeQueueLengthYellowLevel = 15.0
	DefaultMergeQueueLengthGreenLevel  = 10.0

	TimeToMergeBadgeFileName      = "time-to-merge.svg"
	MergeQueueLengthBadgeFileName = "merge-queue-length.svg"
	TimeToMergeBadgeName          = "avg time to merge"
	MergeQueueLengthBadgeName     = "avg merge queue length"
)
