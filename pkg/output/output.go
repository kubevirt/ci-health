package output

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/narqo/go-badge"
	log "github.com/sirupsen/logrus"

	"github.com/fgimenez/ci-health/pkg/constants"
	"github.com/fgimenez/ci-health/pkg/stats"
)

type Levels struct {
	Yellow float64
	Red    float64
}

type Options struct {
	Path                   string
	TimeToMergeLevels      *Levels
	MergeQueueLengthLevels *Levels
}

type Handler struct {
	options *Options
}

func NewHandler(options *Options) *Handler {
	return &Handler{
		options,
	}
}

func (b *Handler) Write(results *stats.Results) error {
	resultsJSON, err := json.MarshalIndent(results, "", " ")
	if err != nil {
		return err
	}
	log.Infof("Results: %s", string(resultsJSON))
	err = ioutil.WriteFile(filepath.Join(b.options.Path, constants.JSONResultsFileName), resultsJSON, 0644)
	if err != nil {
		return err
	}

	err = b.writeMetric(
		constants.TimeToMergeBadgeName,
		filepath.Join(b.options.Path, constants.TimeToMergeBadgeFileName),
		results.Data[constants.TimeToMergeName].Value,
		b.options.TimeToMergeLevels,
	)
	if err != nil {
		return err
	}
	err = b.writeMetric(
		constants.MergeQueueLengthBadgeName,
		filepath.Join(b.options.Path, constants.MergeQueueLengthBadgeFileName),
		results.Data[constants.MergeQueueLengthName].Value,
		b.options.MergeQueueLengthLevels,
	)

	return err
}

func (b *Handler) writeMetric(name, filePath string, value string, levels *Levels) error {
	floatValue, err := strconv.ParseFloat(strings.Fields(value)[0], 64)
	if err != nil {
		return err
	}
	color := BadgeColor(floatValue, levels)

	f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer f.Close()

	log.Debugf("Writing color %s", color)
	err = badge.Render(name, value, color, f)

	return err
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
