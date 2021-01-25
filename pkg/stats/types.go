package stats

type DataPoint struct {
	Value float64
}

// RunningAverageDataItem contains data information in the form of a running average.
// It contains the actual average value and the maximum and minimum values in the data
// set.
type RunningAverageDataItem struct {
	Name       string
	Value      float64
	DataPoints []DataPoint
}

// Results represents the data obtained from GitHub. It includes the source repo from which
// the data was obtained and the number of days back from the execution time included in the
// data.
type Results struct {
	ExecutionDate string `yaml:"executionDate"`
	DataDays      int    `yaml:"dataDays"`
	Source        string
	Data          []*RunningAverageDataItem
}
