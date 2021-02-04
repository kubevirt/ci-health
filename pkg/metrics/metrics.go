package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type Handler struct {
	mergeQueueLength *prometheus.GaugeVec
	timeToMerge      *prometheus.GaugeVec
}

func NewHandler() *Handler {
	mergeQueueLenght := promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "cihealth_merge_queue_lenght_total",
		Help: "The average number of PRs in the merge queue",
	},
		[]string{"source"},
	)

	timeToMerge := promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "cihealth_time_to_merge_seconds",
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
