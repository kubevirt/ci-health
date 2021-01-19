package stats

// RunningAverageDataItem contains data information in the form of a running average.
// It contains the actual average value, the maximum and minimum values in the data set
// and the number of days for which the average was calculated.
type RunningAverageDataItem struct {
	Name            string
	Value, Max, Min float64
	DataDays        int
}

// Results represents the data obtained from GitHub.
type Results struct {
	ExecutionDate string
	Source        string
	Data          []RunningAverageDataItem
}
