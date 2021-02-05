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
	mergeQueueLength *prometheus.GaugeVec
	timeToMerge      *prometheus.GaugeVec
}

func NewHandler() *Handler {
	mergeQueueLenght := promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: constants.MergeQueueLengthMetricName,
		Help: "The average number of PRs in the merge queue",
	},
		[]string{"source"},
	)

	timeToMerge := promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: constants.TimeToMergeMetricName,
		Help: "The average time it took to merge PRs",
	},
		[]string{"source"},
	)

	return &Handler{
		mergeQueueLenght,
		timeToMerge,
	}
}

func (h *Handler) SetMergeQueueLength(source string, value float64) {
	h.mergeQueueLength.With(prometheus.Labels{"source": source}).Set(value)
}

func (h *Handler) SetTimeToMerge(source string, value float64) {
	h.timeToMerge.With(prometheus.Labels{"source": source}).Set(value)
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
