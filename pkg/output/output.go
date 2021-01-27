package output

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/narqo/go-badge"
	log "github.com/sirupsen/logrus"

	"github.com/fgimenez/ci-health/pkg/constants"
	"github.com/fgimenez/ci-health/pkg/stats"
)

type Levels struct {
	Yellow float64
	Red    float64
}

type BadgeOptions struct {
	Path                   string
	TimeToMergeLevels      *Levels
	MergeQueueLengthLevels *Levels
}

type BadgeHandler struct {
	options *BadgeOptions
}

func NewBadgeHandler(options *BadgeOptions) *BadgeHandler {
	return &BadgeHandler{
		options,
	}
}

func (b *BadgeHandler) Write(results *stats.Results) error {

	err := b.writeMetric(
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

func (b *BadgeHandler) writeMetric(name, filePath string, value float64, levels *Levels) error {
	color := BadgeColor(value, levels)

	f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer f.Close()

	log.Debugf("Writing color %s", color)
	return badge.Render(name, fmt.Sprintf("%.2f", value), color, f)
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
