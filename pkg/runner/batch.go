package runner

import (
	"encoding/json"
	"fmt"
	"image/color"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/kubevirt/ci-health/pkg/chatops"
	"github.com/kubevirt/ci-health/pkg/constants"
	"github.com/kubevirt/ci-health/pkg/mergequeue"
	"github.com/kubevirt/ci-health/pkg/output"
	"github.com/kubevirt/ci-health/pkg/plot"
	"github.com/kubevirt/ci-health/pkg/stats"
	"github.com/kubevirt/ci-health/pkg/timeutils"
	"github.com/kubevirt/ci-health/pkg/types"
)

func batchFetchRun(o *types.Options, mqHandler *mergequeue.Handler, coHandler *chatops.Handler) (*types.Results, error) {
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
		outputOptions.Path = path.Join(baseDataPath, currentEndDate.Format(constants.DateFormat))
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

func batchPlotRun(o *types.Options) (*types.Results, error) {
	// read batch fetch results in batch data dir
	dataBase := batchDataPath(o.Path, o.Source, string(o.TargetMetric))

	curves, err := gatherPlotData(dataBase, types.Metric(o.TargetMetric), o.StartDate)
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

func gatherPlotData(basePath string, metric types.Metric, startDate string) ([]types.Curve, error) {
	totalCurves := 2
	if metric == types.QuarantineMetric {
		totalCurves = 5 // Total + 4 SIGs
	}
	curves := make([]types.Curve, totalCurves)
	for i := 0; i < totalCurves; i++ {
		curves[i].X = []string{}
		curves[i].Y = []float64{}
		if metric == types.QuarantineMetric {
			switch i {
			case 0:
				curves[i].Color = color.RGBA{255, 0, 0, 255}
				curves[i].Title = "Total Quarantined"
			case 1:
				curves[i].Color = color.RGBA{0, 0, 255, 255}
				curves[i].Title = "SIG Compute"
			case 2:
				curves[i].Color = color.RGBA{0, 255, 0, 255}
				curves[i].Title = "SIG Storage"
			case 3:
				curves[i].Color = color.RGBA{255, 165, 0, 255}
				curves[i].Title = "SIG Network"
			case 4:
				curves[i].Color = color.RGBA{128, 0, 128, 255}
				curves[i].Title = "SIG Monitoring"
			}
		} else {
			if i == 0 {
				curves[i].Color = color.RGBA{0, 255, 0, 255}
			} else {
				curves[i].Color = color.RGBA{0, 0, 255, 255}
			}
		}
	}

	metricName := metric.ResultsName()

	dateExtractRegex := regexp.MustCompile(`([0-9]{4}-[0-9]{2}-[0-9]{2})`)
	startDateToCheck, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		return nil, err
	}
	shouldFilterByStartDate := func(entryPath string) (bool, error) {
		if startDate == "" {
			return false, nil
		}
		if !dateExtractRegex.MatchString(entryPath) {
			return false, nil
		}
		parsedDate, err := time.Parse("2006-01-02", dateExtractRegex.FindString(entryPath))
		if err != nil {
			return false, err
		}
		if !parsedDate.Before(startDateToCheck) {
			return false, nil
		}
		return true, nil
	}

	err = filepath.Walk(basePath, func(entryPath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			return nil
		}
		if entryPath == basePath {
			return nil
		}
		filter, err := shouldFilterByStartDate(entryPath)
		if err != nil {
			return err
		}
		if filter {
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

		var results types.Results
		err = json.Unmarshal(byteValue, &results)
		if err != nil {
			return err
		}

		switch metricName {
		case constants.MergeQueueLengthName:
			addPRRangeResults(results, curves, metricName)
		case constants.RetestsToMergeName, constants.TimeToMergeName:
			addPRUnitResults(results, curves, metricName)
		case constants.QuarantineStats:
			addQuarantineResults(results, curves, metricName)
		}

		if metricName != constants.QuarantineStats {
			// Only add average line for non-quarantine metrics
			curves[1].X = append(curves[1].X, info.Name())
			curves[1].Y = append(curves[1].Y, results.Data[metricName].Avg)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return curves, nil
}

func addPRRangeResults(results types.Results, curves []types.Curve, metricName string) {
	for _, dataPoint := range results.Data[metricName].DataPoints {
		if dataPoint.Date != nil {
			curves[0].X = append(curves[0].X, dataPoint.Date.Format(constants.DateFormat))
			curves[0].Y = append(curves[0].Y, dataPoint.Value)
		}
	}
}

func addPRUnitResults(results types.Results, curves []types.Curve, metricName string) {
	for _, dataPoint := range results.Data[metricName].DataPoints {
		for _, pr := range dataPoint.PRs {
			curves[0].X = append(curves[0].X, pr.MergedAt)
			curves[0].Y = append(curves[0].Y, dataPoint.Value)
		}
	}
}

func addQuarantineResults(results types.Results, curves []types.Curve, metricName string) {
	data := results.Data[metricName]
	for _, dataPoint := range data.DataPoints {
		if dataPoint.Date != nil {
			dateStr := dataPoint.Date.Format(constants.DateFormat)

			curves[0].X = append(curves[0].X, dateStr)
			curves[0].Y = append(curves[0].Y, data.QuarantineTotal)

			curves[1].X = append(curves[1].X, dateStr)
			curves[1].Y = append(curves[1].Y, data.QuarantineSigCompute)

			curves[2].X = append(curves[2].X, dateStr)
			curves[2].Y = append(curves[2].Y, data.QuarantineSigStorage)

			curves[3].X = append(curves[3].X, dateStr)
			curves[3].Y = append(curves[3].Y, data.QuarantineSigNetwork)

			curves[4].X = append(curves[4].X, dateStr)
			curves[4].Y = append(curves[4].Y, data.QuarantineSigMonitoring)
		}
	}
}
