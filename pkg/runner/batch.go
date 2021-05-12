package runner

import (
	"encoding/json"
	"fmt"
	"image/color"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/fgimenez/ci-health/pkg/chatops"
	"github.com/fgimenez/ci-health/pkg/constants"
	"github.com/fgimenez/ci-health/pkg/mergequeue"
	"github.com/fgimenez/ci-health/pkg/output"
	"github.com/fgimenez/ci-health/pkg/plot"
	"github.com/fgimenez/ci-health/pkg/stats"
	"github.com/fgimenez/ci-health/pkg/timeutils"
	"github.com/fgimenez/ci-health/pkg/types"
)

func batchFetchRun(o *types.Options, mqHandler *mergequeue.Handler, coHandler *chatops.Handler) (*stats.Results, error) {
	// find closest monday to o.StartDate
	currentEndDate, err := timeutils.ClosestMonday(o.StartDate)
	if err != nil {
		return nil, err
	}

	// for each monday since closest monday to o.StartDate to current date
	now := time.Now()

	statsOptions := &stats.HandlerOptions{
		Mq:       mqHandler,
		Co:       coHandler,
		Source:   o.Source,
		DataDays: constants.DefaultDataDays,

		TargetMetrics: []types.Metric{
			types.Metric(o.TargetMetric),
		},
	}

	outputOptions := &output.Options{}

	for {
		if now.Before(currentEndDate) {
			break
		}

		// write results file
		baseDataPath := batchDataPath(o.Path, o.Source, string(o.TargetMetric))
		outputOptions.Path = path.Join(baseDataPath, currentEndDate.Format(constants.BatchDataDateFormat))
		if _, e := os.Stat(outputOptions.Path); os.IsNotExist(e) {
			// calculate results for the week
			statsOptions.EndDate = currentEndDate
			statsHandler := stats.NewHandler(statsOptions)

			results, err := statsHandler.Run()
			if err != nil {
				return nil, err
			}
			outputHandler := output.NewHandler(outputOptions, nil)
			err = outputHandler.WriteJSON(results)
			if err != nil {
				return nil, err
			}
		}

		// bump date to next monday
		currentEndDate = currentEndDate.AddDate(0, 0, 7)
	}
	return nil, nil
}

func batchPlotRun(o *types.Options) (*stats.Results, error) {
	// read batch fetch results in batch data dir
	dataBase := batchDataPath(o.Path, o.Source, string(o.TargetMetric))

	curves, err := gatherPlotData(dataBase, types.Metric(o.TargetMetric))
	if err != nil {
		return nil, err
	}
	log.Debugf("current curves: %v", curves)

	plotBase := batchPlotPath(o.Path, o.Source, string(o.TargetMetric))
	plotData := &types.PlotData{}
	plotData.Title = fmt.Sprintf("%s for %s", string(o.TargetMetric), o.Source)
	plotData.XAxisLabel = "Time"
	plotData.YAxisLabel = string(o.TargetMetric)
	plotData.Curves = curves

	return nil, plot.Draw(plotBase, plotData)
}

func batchBasePath(base, source, metric string) string {
	return path.Join(
		base,
		source,
		constants.DefaultBatchBaseOutputPath,
		metric,
	)
}

func batchDataPath(base, source, metric string) string {
	return path.Join(
		batchBasePath(base, source, metric),
		constants.DefaultBatchDataOutputPath,
	)
}

func batchPlotPath(base, source, metric string) string {
	return path.Join(
		batchBasePath(base, source, metric),
		constants.DefaultBatchPlotOutputPath,
		constants.PlotFileName,
	)
}

func gatherPlotData(basePath string, metric types.Metric) ([]types.Curve, error) {
	totalCurves := 1
	curves := make([]types.Curve, totalCurves)
	for i := 0; i < totalCurves; i++ {
		curves[i].X = []string{}
		curves[i].Y = []float64{}
		curves[i].Color = color.RGBA{0, 255, 0, 255}
	}

	err := filepath.Walk(basePath, func(entryPath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			return nil
		}
		if entryPath == basePath {
			return nil
		}

		dataFile := path.Join(entryPath, constants.JSONResultsFileName)

		jsonFile, err := os.Open(dataFile)
		if err != nil {
			return err
		}
		defer jsonFile.Close()
		byteValue, err := ioutil.ReadAll(jsonFile)
		if err != nil {
			return err
		}

		var results stats.Results
		err = json.Unmarshal(byteValue, &results)
		if err != nil {
			return err
		}

		metricName := metric.ResultsName()
		curves[0].X = append(curves[0].X, info.Name())
		curves[0].Y = append(curves[0].Y, results.Data[metricName].Avg)

		/*
			curves[1].X = append(curves[1].X, info.Name())
			curves[1].Y = append(curves[1].Y, results.Data[metricName].Avg+results.Data[metricName].Std)

			curves[2].X = append(curves[2].X, info.Name())
			if results.Data[metricName].Avg-results.Data[metricName].Std < 0 {
				curves[2].Y = append(curves[2].Y, 0)
			} else {
				curves[2].Y = append(curves[2].Y, results.Data[metricName].Avg-results.Data[metricName].Std)
			}
		*/
		return nil
	})
	if err != nil {
		return nil, err
	}

	return curves, nil
}
