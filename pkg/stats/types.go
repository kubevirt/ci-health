package stats

import "time"

type DataPoint struct {
	Value string
	Date  time.Time `json:",omitempty"`
	PRs   []int     `json:",omitempty"`
}

// RunningAverageDataItem contains data information in the form of a running average.
// It contains the actual average value and the data points used to obtain it.
type RunningAverageDataItem struct {
	Value      string
	DataPoints []DataPoint
}

// Results represents the data obtained from GitHub. It includes the source repo from which
// the data was obtained and the number of days back from the execution time included in the
// data.
type Results struct {
	ExecutionDate string
	DataDays      int
	Source        string
	Data          map[string]RunningAverageDataItem
}
