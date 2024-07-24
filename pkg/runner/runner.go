package runner

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/kubevirt/ci-health/pkg/chatops"
	"github.com/kubevirt/ci-health/pkg/gh"
	"github.com/kubevirt/ci-health/pkg/mergequeue"
	"github.com/kubevirt/ci-health/pkg/metrics"
	"github.com/kubevirt/ci-health/pkg/output"
	"github.com/kubevirt/ci-health/pkg/stats"
	"github.com/kubevirt/ci-health/pkg/types"
)

func Run(o *types.Options) (*types.Results, error) {
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

	switch o.Action {
	case types.StatsAction:
		return statsRun(o, mqHandler, coHandler)
	case types.BatchAction:
		return batchRun(o, mqHandler, coHandler)
	default:
		return nil, fmt.Errorf("Unknown action: %q", o.Action)
	}
}

func statsRun(o *types.Options, mqHandler *mergequeue.Handler, coHandler *chatops.Handler) (*types.Results, error) {
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
			types.MergedPRsMetric,
			types.MergedPRsNoRetestMetric,
			types.SIGRetestMetric,
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
		MergedPRsLevels: &output.Levels{
			Yellow: o.MergedPRsYellowLevel,
			Red:    o.MergedPRsRedLevel,
		},
		MergedPRsNoRetestsLevels: &output.Levels{
			Yellow: o.MergedPRsNoRetestYellowLevel,
			Red:    o.MergedPRsNoRetestRedLevel,
		},
		SIGRetestsLevels: &output.Levels{
			Yellow: o.SIGRetestYellowLevel,
			Red:    o.SIGRetestRedLevel,
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

func batchRun(o *types.Options, mqHandler *mergequeue.Handler, coHandler *chatops.Handler) (*types.Results, error) {
	switch o.Mode {
	case types.FetchMode:
		return batchFetchRun(o, mqHandler, coHandler)
	case types.PlotMode:
		return batchPlotRun(o)
	default:
		return nil, fmt.Errorf("Unknown batch mode: %q", o.Mode)
	}
}

func setLogLevel(logLevel string) {
	if logLevel == "debug" {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}
}
