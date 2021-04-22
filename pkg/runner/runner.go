package runner

import (
	"fmt"
	"io/ioutil"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/fgimenez/ci-health/pkg/chatops"
	"github.com/fgimenez/ci-health/pkg/gh"
	"github.com/fgimenez/ci-health/pkg/mergequeue"
	"github.com/fgimenez/ci-health/pkg/metrics"
	"github.com/fgimenez/ci-health/pkg/output"
	"github.com/fgimenez/ci-health/pkg/stats"
	"github.com/fgimenez/ci-health/pkg/types"
)

func RunStats(o *types.StatsOptions) (*stats.Results, error) {
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
	coHandler := chatops.NewHandler(ghClient)

	statsHandler := stats.NewHandler(mqHandler, coHandler, o.Source, o.DataDays)

	results, err := statsHandler.Run()
	if err != nil {
		return nil, err
	}

	if o.Path == "" {
		dir, err := ioutil.TempDir("", "ci-health")
		if err != nil {
			return nil, err
		}
		o.Path = dir
	}
	if _, err := os.Stat(o.Path); os.IsNotExist(err) {
		err := os.Mkdir(o.Path, 0755)
		if err != nil {
			return nil, err
		}
	}

	outputOptions := &output.Options{
		Path: o.Path,
		TimeToMergeLevels: &output.Levels{
			Yellow: o.TimeToMergeYellowLevel,
			Red:    o.TimeToMergeRedLevel,
		},
		MergeQueueLengthLevels: &output.Levels{
			Yellow: o.MergeQueueLengthYellowLevel,
			Red:    o.MergeQueueLengthRedLevel,
		},
		RetestsToMergeLevels: &output.Levels{
			Yellow: o.RetestsToMergeYellowLevel,
			Red:    o.RetestsToMergeRedLevel,
		},
		Source: o.Source,
	}
	metricsHandler := metrics.NewHandler()
	outputHandler := output.NewHandler(outputOptions, metricsHandler)

	err = outputHandler.Write(results)
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
