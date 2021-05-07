package runner

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/fgimenez/ci-health/pkg/chatops"
	"github.com/fgimenez/ci-health/pkg/gh"
	"github.com/fgimenez/ci-health/pkg/mergequeue"
	"github.com/fgimenez/ci-health/pkg/metrics"
	"github.com/fgimenez/ci-health/pkg/output"
	"github.com/fgimenez/ci-health/pkg/stats"
	"github.com/fgimenez/ci-health/pkg/timeutils"
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

	switch o.Action {
	case types.StatsAction:
		return statsRun(o, ghClient)
	case types.BatchAction:
		return batchRun(o, ghClient)
	default:
		return nil, fmt.Errorf("Unknown action: %q", o.Action)
	}
}

func statsRun(o *types.Options, ghClient *gh.Client) (*stats.Results, error) {
	mqHandler := mergequeue.NewHandler(ghClient)
	coHandler := chatops.NewHandler(ghClient)

	options := &stats.HandlerOptions{
		Mq:       mqHandler,
		Co:       coHandler,
		Source:   o.Source,
		DataDays: o.DataDays,
		EndDate:  time.Now(),

		TargetMetrics: []types.Metric{
			types.MergeQueueLengthMetric,
			types.TimeToMergeMetric,
			types.RetestsToMergeMetric,
		},
	}

	statsHandler := stats.NewHandler(options)

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

func batchRun(o *types.Options, ghClient *gh.Client) (*stats.Results, error) {
	switch o.Mode {
	case types.FetchMode:
		return batchFetchRun(o, ghClient)
	case types.PlotMode:
		return batchPlotRun(o, ghClient)
	default:
		return nil, fmt.Errorf("Unknown batch mode: %q", o.Mode)
	}
}

func batchFetchRun(o *types.Options, ghClient *gh.Client) (*stats.Results, error) {
	// find closest monday to o.StartDate
	currentStartingDate, err := timeutils.ClosestMonday(o.StartDate)
	if err != nil {
		return nil, err
	}

	// for each monday since closest monday to o.StartDate to current date
	now := time.Now()
	for {
		if now.Before(currentStartingDate) {
			break
		}
		// calculate results for the week

		// write results file

		currentStartingDate = currentStartingDate.AddDate(0, 0, o.DataDays)
	}
	return nil, nil
}

func batchPlotRun(o *types.Options, ghClient *gh.Client) (*stats.Results, error) {
	// read batch fetch results since o.StartDate

	// generate gnuplot files using results

	return nil, nil
}

func setLogLevel(logLevel string) {
	if logLevel == "debug" {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}
}
