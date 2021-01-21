package stats

import (
	"github.com/fgimenez/cihealth/pkg/gh"
)

type statsProcessor func(*Results, *gh.Client) (*Results, error)

var processors []statsProcessor

func init() {
	processors = []statsProcessor{mergeQueueProcessor, timeToMergeProcessor}
}

func Run(client *gh.Client) (*Results, error) {
	results := &Results{
		DataDays: client.DataDays(),
		Source:   client.Source(),
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

	return results, nil
}

func timeToMergeProcessor(results *Results, client *gh.Client) (*Results, error) {
	return results, nil
}
