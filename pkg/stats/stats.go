package stats

import (
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
	queueLength, err := h.mq.LengthAt(time.Now())
	if err != nil {
		return nil, err
	}

	dataItem := &RunningAverageDataItem{
		Name:  "MergeQueueLength",
		Value: float64(queueLength),
	}

	results.Data = append(results.Data, dataItem)

	return results, nil
}

func (h *Handler) timeToMergeProcessor(results *Results) (*Results, error) {
	return results, nil
}
