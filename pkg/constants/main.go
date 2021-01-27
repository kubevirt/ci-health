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
	DefaultTimeToMergeRedLevel         = 3.0
	DefaultTimeToMergeYellowLevel      = 2.0
	DefaultTimeToMergeGreenLevel       = 1.0
	DefaultMergeQueueLengthRedLevel    = 15.0
	DefaultMergeQueueLengthYellowLevel = 10.0
	DefaultMergeQueueLengthGreenLevel  = 6.0

	TimeToMergeBadgeFileName      = "time-to-merge.svg"
	MergeQueueLenghtBadgeFileName = "merge-queue-length.svg"
)
