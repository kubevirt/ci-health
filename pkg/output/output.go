package output

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"

	"github.com/narqo/go-badge"
	log "github.com/sirupsen/logrus"

	"github.com/kubevirt/ci-health/pkg/constants"
	"github.com/kubevirt/ci-health/pkg/metrics"
	"github.com/kubevirt/ci-health/pkg/types"
)

type Levels struct {
	Yellow float64
	Red    float64
}

type Options struct {
	Path                     string
	Source                   string
	TimeToMergeLevels        *Levels
	MergeQueueLengthLevels   *Levels
	RetestsToMergeLevels     *Levels
	MergedPRsLevels          *Levels
	MergedPRsNoRetestsLevels *Levels
	SIGRetestsLevels         *Levels
}

type Handler struct {
	options        *Options
	metricsHandler *metrics.Handler
}

func NewHandler(options *Options, metricsHandler *metrics.Handler) *Handler {
	return &Handler{
		options,
		metricsHandler,
	}
}

func (b *Handler) Write(results *types.Results) error {
	err := b.WriteJSON(results)
	if err != nil {
		return err
	}

	err = b.writeBadges(results)
	if err != nil {
		return err
	}

	err = b.writeMetrics(results)

	return err
}

func (b *Handler) WriteJSON(results *types.Results) error {
	basePath, err := b.initializeSourcePath()
	if err != nil {
		return err
	}

	resultsJSON, err := json.MarshalIndent(results, "", " ")
	if err != nil {
		return err
	}
	log.Infof("Results: %s", string(resultsJSON))
	resultsPath := filepath.Join(basePath, constants.JSONResultsFileName)
	err = ioutil.WriteFile(resultsPath, resultsJSON, 0644)

	return err
}

func (b *Handler) writeMetrics(results *types.Results) error {
	b.metricsHandler.SetAvgMergeQueueLength(results.Source, results.Data[constants.MergeQueueLengthName].Avg)
	b.metricsHandler.SetAvgTimeToMerge(results.Source, results.Data[constants.TimeToMergeName].Avg)
	b.metricsHandler.SetAvgRetestsToMerge(results.Source, results.Data[constants.RetestsToMergeName].Avg)
	b.metricsHandler.SetStdMergeQueueLength(results.Source, results.Data[constants.MergeQueueLengthName].Std)
	b.metricsHandler.SetStdTimeToMerge(results.Source, results.Data[constants.TimeToMergeName].Std)
	b.metricsHandler.SetStdRetestsToMerge(results.Source, results.Data[constants.RetestsToMergeName].Std)

	m := b.metricsHandler.String()
	log.Debugf("Metrics: %s", m)

	metricsPath := filepath.Join(b.options.Path, constants.MetricsFileName)
	err := ioutil.WriteFile(metricsPath, []byte(m), 0644)

	return err
}

func (b *Handler) writeBadges(results *types.Results) error {
	basePath, err := b.initializeSourcePath()
	if err != nil {
		return err
	}

	err = b.writeBadge(
		constants.TimeToMergeBadgeName,
		filepath.Join(basePath, constants.TimeToMergeBadgeFileName),
		results.Data[constants.TimeToMergeName],
		b.options.TimeToMergeLevels,
	)
	if err != nil {
		return err
	}

	err = b.writeBadge(
		constants.MergeQueueLengthBadgeName,
		filepath.Join(basePath, constants.MergeQueueLengthBadgeFileName),
		results.Data[constants.MergeQueueLengthName],
		b.options.MergeQueueLengthLevels,
	)

	err = b.writeBadge(
		constants.RetestsToMergeBadgeName,
		filepath.Join(basePath, constants.RetestsToMergeBadgeFileName),
		results.Data[constants.RetestsToMergeName],
		b.options.RetestsToMergeLevels,
	)

	err = b.writeBadge(
		constants.MergedPRsNoRetestBadgeName,
		filepath.Join(basePath, constants.MergedPRsNoRetestBadgeFileName),
		results.Data[constants.MergedPRsNoRetest],
		b.options.MergedPRsNoRetestsLevels,
	)

	err = b.writeSIGRetestBadge(
		constants.SIGComputeBadgeName,
		filepath.Join(basePath, constants.SIGComputeRetestBadgeFileName),
		results.Data[constants.SIGRetests],
		b.options.SIGRetestsLevels,
	)

	err = b.writeSIGRetestBadge(
		constants.SIGStorageBadgeName,
		filepath.Join(basePath, constants.SIGStorageRetestBadgeFileName),
		results.Data[constants.SIGRetests],
		b.options.SIGRetestsLevels,
	)

	err = b.writeSIGRetestBadge(
		constants.SIGNetworkBadgeName,
		filepath.Join(basePath, constants.SIGNetworkRetestBadgeFileName),
		results.Data[constants.SIGRetests],
		b.options.SIGRetestsLevels,
	)

	err = b.writeSIGRetestBadge(
		constants.SIGOperatorBadgeName,
		filepath.Join(basePath, constants.SIGOperatorRetestBadgeFileName),
		results.Data[constants.SIGRetests],
		b.options.SIGRetestsLevels,
	)

	err = b.writeSIGRetestBadge(
		constants.SIGCIBadgeName,
		filepath.Join(basePath, constants.SIGCIRetestBadgeFileName),
		results.Data[constants.SIGRetests],
		b.options.SIGRetestsLevels,
	)

	err = b.writeSIGRetestBadge(
		constants.SIGMonitoringBadgeName,
		filepath.Join(basePath, constants.SIGMonitoringRetestBadgeFileName),
		results.Data[constants.SIGRetests],
		b.options.SIGRetestsLevels,
	)

	err = b.writeQuarantineBadge(
		constants.SIGMonitoringBadgeName,
		filepath.Join(basePath, constants.QuarantineMonitoringBadgeFileName),
		results.Data[constants.QuarantineStats],
		b.options.SIGRetestsLevels,
	)

	err = b.writeQuarantineBadge(
		constants.SIGComputeBadgeName,
		filepath.Join(basePath, constants.QuarantineComputeBadgeFileName),
		results.Data[constants.QuarantineStats],
		b.options.SIGRetestsLevels,
	)

	err = b.writeQuarantineBadge(
		constants.SIGNetworkBadgeName,
		filepath.Join(basePath, constants.QuarantineNetworkBadgeFileName),
		results.Data[constants.QuarantineStats],
		b.options.SIGRetestsLevels,
	)

	err = b.writeQuarantineBadge(
		constants.SIGStorageBadgeName,
		filepath.Join(basePath, constants.QuarantineStorageBadgeFileName),
		results.Data[constants.QuarantineStats],
		b.options.SIGRetestsLevels,
	)

	err = b.writeQuarantineBadge(
		constants.SIGStorageBadgeName,
		filepath.Join(basePath, constants.QuarantineBadgeFileName),
		results.Data[constants.QuarantineStats],
		b.options.SIGRetestsLevels,
	)

	err = b.writeJobFailureBadges(
		results.Data[constants.SIGRetests],
		b.options.SIGRetestsLevels,
	)

	return err
}

func (b *Handler) writeBadge(name, filePath string, data types.RunningAverageDataItem, levels *Levels) error {
	color := BadgeColor(data.Avg, levels)
	if name == constants.MergedPRsNoRetestBadgeName {
		color = NoRetestBadgeColor(data.NoRetest, data.Number, levels)
	}

	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	log.Debugf("Writing color %s", color)
	badgeString := data.String()
	if name == constants.MergedPRsNoRetestBadgeName {
		badgeString = data.SimpleBadgeString()
	}
	if name == constants.TimeToMergeBadgeName {
		badgeString = fmt.Sprintf("%.2f ± std %.2f", (data.Avg * 24), (data.Std * 24))
	}
	err = badge.Render(name, badgeString, color, f)

	return err
}

func (b *Handler) writeSIGRetestBadge(name, filePath string, data types.RunningAverageDataItem, levels *Levels) error {
	var value float64
	var total float64

	switch name {
	case constants.SIGComputeBadgeName:
		value = data.SIGComputeRetest
		total = data.SIGComputeTotal
	case constants.SIGNetworkBadgeName:
		value = data.SIGNetworkRetest
		total = data.SIGNetworkTotal
	case constants.SIGStorageBadgeName:
		value = data.SIGStorageRetest
		total = data.SIGStorageTotal
	case constants.SIGOperatorBadgeName:
		value = data.SIGOperatorRetest
		total = data.SIGOperatorTotal
	case constants.SIGCIBadgeName:
		value = data.SIGCIRetest
		total = data.SIGCITotal
	case constants.SIGMonitoringBadgeName:
		value = data.SIGMonitoringRetest
		total = data.SIGMonitoringTotal
	}

	color := BadgeColor(((value / total) * 100), levels)

	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	badgeString := fmt.Sprintf("%.0f / %.0f    |    %.2f%s", value, total, (value/total)*100, "%")

	return badge.Render(name, badgeString, color, f)
}

func (b *Handler) writeJobFailureBadges(data types.RunningAverageDataItem, levels *Levels) error {
	failedJobLeaders := data.FailedJobLeaderBoard
	basePath, err := b.initializeSourcePath()
	if err != nil {
		return err
	}
	for i := 1; i <= 10; i++ {
		filePath := filepath.Join(basePath, fmt.Sprintf("failedjob%s.svg", strconv.Itoa(i)))
		f, err := os.Create(filePath)
		if err != nil {
			return err
		}
		defer f.Close()

		if i < len(failedJobLeaders) {
			job := failedJobLeaders[i-1]
			color := BadgeColor(float64(job.FailureCount), levels)
			badgeString := fmt.Sprintf("%.0f / %.0f    |    %.0f%s",
				float64(job.FailureCount),
				float64(job.FailureCount+job.SuccessCount),
				(float64(job.FailureCount)/float64(job.FailureCount+job.SuccessCount))*100, "%")

			err = badge.Render(job.JobName, badgeString, color, f)
		} else {
			color := badge.ColorGreen
			err = badge.Render("-", "-", color, f)
		}
	}
	return err
}

func (b *Handler) writeQuarantineBadge(name, filePath string, data types.RunningAverageDataItem, levels *Levels) error {
	var value float64

	switch name {
	case constants.SIGComputeBadgeName:
		value = data.QuarantineSigCompute
	case constants.SIGNetworkBadgeName:
		value = data.QuarantineSigNetwork
	case constants.SIGStorageBadgeName:
		value = data.QuarantineSigStorage
	case constants.SIGMonitoringBadgeName:
		value = data.QuarantineSigMonitoring
	case constants.QuarantineBadgeName:
		value = data.QuarantineTotal
	}

	color := BadgeColor(value, levels)

	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	badgeString := fmt.Sprintf("%.0f", value)

	return badge.Render(name, badgeString, color, f)
}

func (b *Handler) initializeSourcePath() (string, error) {
	basePath := filepath.Join(b.options.Path, b.options.Source)
	err := os.MkdirAll(basePath, 0755)
	return basePath, err
}
func BadgeColor(value float64, levels *Levels) badge.Color {
	var color badge.Color

	if value < levels.Yellow {
		color = badge.ColorGreen
	} else if value >= levels.Yellow && value < levels.Red {
		color = badge.ColorYellow
	} else {
		color = badge.ColorRed
	}

	return color
}
func NoRetestBadgeColor(value float64, mergedPRs float64, levels *Levels) badge.Color {
	var color badge.Color
	var percentVal float64
	if value == 0 {
		percentVal = 0
	} else if mergedPRs == 0 {
		percentVal = 0
	} else {
		percentVal = value / mergedPRs
	}
	if percentVal > levels.Yellow {
		color = badge.ColorGreen
	} else if percentVal <= levels.Yellow && percentVal > levels.Red {
		color = badge.ColorYellow
	} else {
		color = badge.ColorRed
	}

	return color
}
