package metrics

import (
	"bufio"
	"bytes"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/common/expfmt"
	log "github.com/sirupsen/logrus"

	"github.com/fgimenez/ci-health/pkg/constants"
)

type Handler struct {
	avgMergeQueueLength *prometheus.GaugeVec
	stdMergeQueueLength *prometheus.GaugeVec
	avgTimeToMerge      *prometheus.GaugeVec
	stdTimeToMerge      *prometheus.GaugeVec
	avgRetestsToMerge   *prometheus.GaugeVec
	stdRetestsToMerge   *prometheus.GaugeVec
}

func NewHandler() *Handler {
	avgMergeQueueLenght := promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: constants.AvgMergeQueueLengthMetricName,
		Help: "The average number of PRs in the merge queue",
	},
		[]string{"source"},
	)
	stdMergeQueueLenght := promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: constants.StdMergeQueueLengthMetricName,
		Help: "The standard deviation of the number of PRs in the merge queue",
	},
		[]string{"source"},
	)

	avgTimeToMerge := promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: constants.AvgTimeToMergeMetricName,
		Help: "The average days it took to merge PRs",
	},
		[]string{"source"},
	)
	stdTimeToMerge := promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: constants.StdTimeToMergeMetricName,
		Help: "The standard deviation of the days it took to merge PRs",
	},
		[]string{"source"},
	)

	avgRetestsToMerge := promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: constants.AvgRetestsToMergeMetricName,
		Help: "The average retests it took to merge PRs",
	},
		[]string{"source"},
	)
	stdRetestsToMerge := promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: constants.StdRetestsToMergeMetricName,
		Help: "The standard deviation of the retests it took to merge PRs",
	},
		[]string{"source"},
	)

	return &Handler{
		avgMergeQueueLenght,
		stdMergeQueueLenght,
		avgTimeToMerge,
		stdTimeToMerge,
		avgRetestsToMerge,
		stdRetestsToMerge,
	}
}

func (h *Handler) SetAvgMergeQueueLength(source string, value float64) {
	h.avgMergeQueueLength.With(prometheus.Labels{"source": source}).Set(value)
}

func (h *Handler) SetAvgTimeToMerge(source string, value float64) {
	h.avgTimeToMerge.With(prometheus.Labels{"source": source}).Set(value)
}

func (h *Handler) SetAvgRetestsToMerge(source string, value float64) {
	h.avgRetestsToMerge.With(prometheus.Labels{"source": source}).Set(value)
}

func (h *Handler) SetStdMergeQueueLength(source string, value float64) {
	h.stdMergeQueueLength.With(prometheus.Labels{"source": source}).Set(value)
}

func (h *Handler) SetStdTimeToMerge(source string, value float64) {
	h.stdTimeToMerge.With(prometheus.Labels{"source": source}).Set(value)
}

func (h *Handler) SetStdRetestsToMerge(source string, value float64) {
	h.stdRetestsToMerge.With(prometheus.Labels{"source": source}).Set(value)
}

func (h *Handler) String() string {
	var buffer bytes.Buffer
	w := bufio.NewWriter(&buffer)

	contentType := expfmt.FmtText
	enc := expfmt.NewEncoder(w, contentType)

	reg := prometheus.DefaultGatherer
	mfs, err := reg.Gather()
	if err != nil {
		log.Errorf("Could not gather metrics: %s", err)
		return ""
	}
	for _, mf := range mfs {
		if !strings.HasPrefix(*mf.Name, "cihealth") {
			continue
		}
		err = enc.Encode(mf)
		if err != nil {
			log.Errorf("Could not encode metric %s: %v", mf, err)
		}
	}
	err = w.Flush()
	if err != nil {
		log.Errorf("Could not flush writer %v", err)
	}

	return buffer.String()
}
