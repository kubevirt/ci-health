package output

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/narqo/go-badge"
	log "github.com/sirupsen/logrus"

	"github.com/fgimenez/ci-health/pkg/constants"
	"github.com/fgimenez/ci-health/pkg/metrics"
	"github.com/fgimenez/ci-health/pkg/stats"
)

type Levels struct {
	Yellow float64
	Red    float64
}

type Options struct {
	Path                   string
	Source                 string
	TimeToMergeLevels      *Levels
	MergeQueueLengthLevels *Levels
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

func (b *Handler) Write(results *stats.Results) error {
	err := b.writeJSON(results)
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

func (b *Handler) writeJSON(results *stats.Results) error {
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

func (b *Handler) writeMetrics(results *stats.Results) error {
	b.metricsHandler.SetAvgMergeQueueLength(results.Source, results.Data[constants.MergeQueueLengthName].Avg)
	b.metricsHandler.SetAvgTimeToMerge(results.Source, results.Data[constants.TimeToMergeName].Avg)

	m := b.metricsHandler.String()
	log.Debugf("Metrics: %s", m)

	metricsPath := filepath.Join(b.options.Path, constants.MetricsFileName)
	err := ioutil.WriteFile(metricsPath, []byte(m), 0644)

	return err
}

func (b *Handler) writeBadges(results *stats.Results) error {
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

	return err
}

func (b *Handler) writeBadge(name, filePath string, data stats.RunningAverageDataItem, levels *Levels) error {
	color := BadgeColor(data.Avg, levels)

	f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer f.Close()

	log.Debugf("Writing color %s", color)
	err = badge.Render(name, data.String(), color, f)

	return err
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
