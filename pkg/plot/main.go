package plot

import (
	"os"
	"path"
	"time"

	log "github.com/sirupsen/logrus"
	gonumplot "gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"

	"github.com/fgimenez/ci-health/pkg/constants"
	"github.com/fgimenez/ci-health/pkg/types"
)

func Draw(filePath string, data *types.PlotData) error {
	xticks := gonumplot.TimeTicks{Format: constants.BatchDataDateFormat}

	p := gonumplot.New()
	p.Title.Text = data.Title
	p.X.Tick.Marker = xticks
	p.X.Label.Text = data.XAxisLabel
	p.Y.Label.Text = data.YAxisLabel
	p.Add(plotter.NewGrid())

	for _, curve := range data.Curves {
		data, err := transform(curve.X, curve.Y)
		if err != nil {
			return err
		}

		line, points, err := plotter.NewLinePoints(data)
		if err != nil {
			return err
		}
		line.Color = curve.Color
		points.Shape = draw.CircleGlyph{}
		points.Color = curve.Color

		p.Add(line, points)
	}
	log.Debugf("before saving image to %s", filePath)
	dir := path.Dir(filePath)
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return err
	}

	return p.Save(30*vg.Centimeter, 15*vg.Centimeter, filePath)
}

func transform(x []string, y []float64) (plotter.XYs, error) {
	pts := make(plotter.XYs, len(x))

	for cont := 0; cont < len(x); cont++ {
		parsed, err := time.Parse(constants.BatchDataDateFormat, x[cont])
		log.Debugf("parsed date: %v", parsed)
		if err != nil {
			return nil, err
		}
		date := time.Date(parsed.Year(), parsed.Month(), parsed.Day(), 0, 0, 0, 0, time.UTC).Unix()
		pts[cont].X = float64(date)
		pts[cont].Y = y[cont]

	}

	return pts, nil
}
