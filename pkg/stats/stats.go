package stats

import (
	"math"
	"time"

	"github.com/fgimenez/ci-health/pkg/constants"
	"github.com/fgimenez/ci-health/pkg/mergequeue"
)

type statsProcessor func(*Results) (*Results, error)

type Handler struct {
	mq       *mergequeue.Handler
	source   string
	dataDays int
}

func NewHandler(mq *mergequeue.Handler, source string, dataDays int) *Handler {
	return &Handler{
		mq,
		source,
		dataDays,
	}
}

func (h *Handler) Run() (*Results, error) {
	results := &Results{
		ExecutionDate: time.Now().Format(constants.DateFormat),
		DataDays:      h.dataDays,
		Source:        h.source,
		Data:          map[string]RunningAverageDataItem{},
	}
	var err error

	for _, processor := range []statsProcessor{h.mergeQueueProcessor, h.timeToMergeProcessor} {
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
		queueLength, err := h.mq.LengthAt(currentTime.AddDate(0, 0, -1*i))
		if err != nil {
			return nil, err
		}
		values = append(values, float64(queueLength))
		dataItem.DataPoints = append(dataItem.DataPoints, DataPoint{Value: float64(queueLength)})
	}

	dataItem.Value = Average(values)

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
	for _, timeToMerge := range timesToMerge {
		days := timeToMerge.Hours() / 24
		value := round(days)

		values = append(values, value)
		dataItem.DataPoints = append(dataItem.DataPoints, DataPoint{Value: value})
	}

	dataItem.Value = Average(values)

	results.Data[constants.TimeToMergeName] = dataItem

	return results, nil
}

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

func round(value float64) float64 {
	return math.Round(value*100) / 100
}
