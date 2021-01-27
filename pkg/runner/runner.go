package runner

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	log "github.com/sirupsen/logrus"

	"github.com/fgimenez/ci-health/pkg/gh"
	"github.com/fgimenez/ci-health/pkg/mergequeue"
	"github.com/fgimenez/ci-health/pkg/output"
	"github.com/fgimenez/ci-health/pkg/stats"
	"github.com/fgimenez/ci-health/pkg/types"
)

func Run(o *types.Options) (*stats.Results, error) {
	if o.LogLevel == "debug" {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}

	if o.TokenPath == "" {
		return nil, fmt.Errorf("You need to specify the GitHub token path with --gh-token")
	}

	ghClient, err := gh.NewClient(o.TokenPath, o.Source)
	if err != nil {
		return nil, err
	}

	mqHandler := mergequeue.NewHandler(ghClient)

	statsHandler := stats.NewHandler(mqHandler, o.Source, o.DataDays)

	results, err := statsHandler.Run()
	if err != nil {
		return nil, err
	}
	resultsJSON, err := json.Marshal(results)
	if err != nil {
		return nil, err
	}
	log.Infof("Results: %s", string(resultsJSON))

	if o.Path == "" {
		dir, err := ioutil.TempDir("", "ci-health")
		if err != nil {
			return nil, err
		}
		o.Path = dir
	}

	badgeOptions := &output.BadgeOptions{
		Path: o.Path,
		TimeToMergeLevels: &output.Levels{
			Red:    o.TimeToMergeRedLevel,
			Yellow: o.TimeToMergeYellowLevel,
			Green:  o.TimeToMergeGreenLevel,
		},
		MergeQueueLengthLevels: &output.Levels{
			Red:    o.MergeQueueLengthRedLevel,
			Yellow: o.MergeQueueLengthYellowLevel,
			Green:  o.MergeQueueLengthGreenLevel,
		},
	}
	badgeHandler := output.NewBadgeHandler(badgeOptions)

	err = badgeHandler.Write(results)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func setLogLevel(logLevel string) {
	if logLevel == "debug" {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}
}
