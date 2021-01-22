package stats

import (
	"time"

	"github.com/fgimenez/cihealth/pkg/gh"
)

type statsProcessor func(*Results, *gh.Client) (*Results, error)

var processors []statsProcessor

func init() {
	processors = []statsProcessor{mergeQueueProcessor, timeToMergeProcessor}
}

func Run(client *gh.Client) (*Results, error) {
	results := &Results{
		ExecutionDate: time.Now().Format("2006-01-02T15:04:05Z"),
		DataDays:      client.DataDays(),
		Source:        client.Source(),
	}
	var err error

	for _, processor := range processors {
		results, err = processor(results, client)
		if err != nil {
			return nil, err
		}
	}
	return results, nil
}

func mergeQueueProcessor(results *Results, client *gh.Client) (*Results, error) {
	queueLength, err := client.GetOpenApprovedPRsByDate(time.Now())
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

func timeToMergeProcessor(results *Results, client *gh.Client) (*Results, error) {
	return results, nil
}
