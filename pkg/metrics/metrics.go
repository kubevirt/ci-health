package metrics

import (
	"bufio"
	"bytes"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/common/expfmt"
	log "github.com/sirupsen/logrus"

	"github.com/fgimenez/ci-health/pkg/constants"
)

type Handler struct {
	avgMergeQueueLength *prometheus.GaugeVec
	avgTimeToMerge      *prometheus.GaugeVec
}

func NewHandler() *Handler {
	avgMergeQueueLenght := promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: constants.AvgMergeQueueLengthMetricName,
		Help: "The average number of PRs in the merge queue",
	},
		[]string{"source"},
	)

	avgTimeToMerge := promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: constants.AvgTimeToMergeMetricName,
		Help: "The average time it took to merge PRs",
	},
		[]string{"source"},
	)

	return &Handler{
		avgMergeQueueLenght,
		avgTimeToMerge,
	}
}

func (h *Handler) SetAvgMergeQueueLength(source string, value float64) {
	h.avgMergeQueueLength.With(prometheus.Labels{"source": source}).Set(value)
}

func (h *Handler) SetAvgTimeToMerge(source string, value float64) {
	h.avgTimeToMerge.With(prometheus.Labels{"source": source}).Set(value)
}

func (h *Handler) String() string {
	reg := prometheus.DefaultGatherer
	mfs, err := reg.Gather()
	if err != nil {
		log.Errorf("Could not gather metrics: %s", err)
		return ""
	}

	var buffer bytes.Buffer
	w := bufio.NewWriter(&buffer)

	contentType := expfmt.FmtText
	enc := expfmt.NewEncoder(w, contentType)

	for _, mf := range mfs {
		err = enc.Encode(mf)
		if err != nil {
			log.Errorf("Could not encode metric %s: %v", mf, err)
		}
	}

	return buffer.String()
}
