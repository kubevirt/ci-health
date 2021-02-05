package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

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

	return ""
}
