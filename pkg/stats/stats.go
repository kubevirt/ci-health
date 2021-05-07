package stats

import (
	"fmt"
	"math"
	"time"

	"github.com/fgimenez/ci-health/pkg/chatops"
	"github.com/fgimenez/ci-health/pkg/constants"
	"github.com/fgimenez/ci-health/pkg/mergequeue"
	"github.com/fgimenez/ci-health/pkg/types"
)

type statsProcessor func(*Results) (*Results, error)

type HandlerOptions struct {
	Mq       *mergequeue.Handler
	Co       *chatops.Handler
	Source   string
	EndDate  time.Time
	DataDays int

	TargetMetrics []types.Metric
}

type Handler struct {
	mq            *mergequeue.Handler
	co            *chatops.Handler
	source        string
	endDate       time.Time
	dataDays      int
	targetMetrics []types.Metric
}

func NewHandler(ho *HandlerOptions) *Handler {
	return &Handler{
		mq:            ho.Mq,
		co:            ho.Co,
		source:        ho.Source,
		endDate:       ho.EndDate,
		dataDays:      ho.DataDays,
		targetMetrics: ho.TargetMetrics,
	}
}

func (h *Handler) Run() (*Results, error) {
	results := &Results{
		EndDate:  h.endDate.Format(constants.DateFormat),
		DataDays: h.dataDays,
		Source:   h.source,
		Data:     map[string]RunningAverageDataItem{},
	}
	var err error

	for _, targetMetric := range h.targetMetrics {
		var processor statsProcessor

		switch targetMetric {
		case types.MergeQueueLengthMetric:
			processor = h.mergeQueueProcessor
		case types.TimeToMergeMetric:
			processor = h.timeToMergeProcessor
		case types.RetestsToMergeMetric:
			processor = h.retestsToMergeProcessor
		default:
			return nil, fmt.Errorf("Unknown metric: %q", targetMetric)
		}

		results, err = processor(results)
		if err != nil {
			return nil, err
		}
	}
	return results, nil
}

func (h *Handler) mergeQueueProcessor(results *Results) (*Results, error) {
	currentTime := time.Now()

	dataItem := RunningAverageDataItem{
		DataPoints: []DataPoint{},
	}

	values := []float64{}
	for i := 0; i < results.DataDays; i++ {
		queryDate := currentTime.AddDate(0, 0, -1*i)
		queueLength, prs, err := h.mq.LengthAt(queryDate)
		if err != nil {
			return nil, err
		}
		values = append(values, float64(queueLength))
		dataItem.DataPoints = append(dataItem.DataPoints,
			DataPoint{
				Value: float64(queueLength),
				PRs:   prs,
				Date:  &queryDate,
			})
	}

	dataItem.Avg = Average(values)
	dataItem.Std = Std(values)

	results.Data[constants.MergeQueueLengthName] = dataItem

	return results, nil
}

func (h *Handler) timeToMergeProcessor(results *Results) (*Results, error) {
	currentTime := time.Now()

	dataItem := RunningAverageDataItem{
		DataPoints: []DataPoint{},
	}

	timesToMerge, err := h.mq.TimesToMerge(currentTime.AddDate(0, 0, -1*results.DataDays), currentTime)
	if err != nil {
		return nil, err
	}

	values := []float64{}

	for prNumber, timeToMerge := range timesToMerge {
		days := timeToMerge.Hours() / 24
		value := round(days)

		values = append(values, value)

		dataItem.DataPoints = append(dataItem.DataPoints,
			DataPoint{
				Value: value,
				PRs:   []int{prNumber},
			})
	}

	dataItem.Avg = Average(values)
	dataItem.Std = Std(values)

	results.Data[constants.TimeToMergeName] = dataItem

	return results, nil
}

func (h *Handler) retestsToMergeProcessor(results *Results) (*Results, error) {
	currentTime := time.Now()

	dataItem := RunningAverageDataItem{
		DataPoints: []DataPoint{},
	}

	retestsToMerge, err := h.co.RetestsToMerge(currentTime.AddDate(0, 0, -1*results.DataDays), currentTime)
	if err != nil {
		return nil, err
	}

	values := []float64{}

	for prNumber, retestsToMerge := range retestsToMerge {
		value := float64(retestsToMerge)

		values = append(values, value)

		dataItem.DataPoints = append(dataItem.DataPoints,
			DataPoint{
				Value: value,
				PRs:   []int{prNumber},
			})
	}

	dataItem.Avg = Average(values)
	dataItem.Std = Std(values)

	results.Data[constants.RetestsToMergeName] = dataItem

	return results, nil
}

// Average returns the average of the given floats.
func Average(xs []float64) float64 {
	if len(xs) == 0 {
		return 0
	}
	total := 0.0
	for _, v := range xs {
		total += v
	}
	result := total / float64(len(xs))
	return round(result)
}

// Std returns the standard deviation of the given floats.
func Std(xs []float64) float64 {
	if len(xs) == 0 {
		return 0
	}
	avg := Average(xs)
	total := 0.0
	for _, v := range xs {
		total += math.Pow((v - avg), 2)
	}
	variance := total / float64(len(xs))
	result := math.Sqrt(variance)
	return round(result)
}

func round(value float64) float64 {
	return math.Round(value*100) / 100
}
