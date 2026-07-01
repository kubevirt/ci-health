package cost

import (
	"fmt"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

// HandlerOptions configures the cost report handler.
type HandlerOptions struct {
	ThanosURL    string
	BearerToken  string
	Namespace    string
	NodeSelector string
	DataDays     int
	Source       string
	Path         string
	MonthlyCost        float64
	InsecureSkipVerify bool
}

// Handler orchestrates querying Prometheus and building the usage report.
type Handler struct {
	client  *PrometheusClient
	options *HandlerOptions
}

// NewHandler creates a cost report handler.
func NewHandler(opts *HandlerOptions) *Handler {
	client := NewPrometheusClient(opts.ThanosURL, opts.BearerToken, opts.InsecureSkipVerify)
	return &Handler{
		client:  client,
		options: opts,
	}
}

// Run queries Prometheus, calculates usage, and writes the report.
func (h *Handler) Run() (*UsageReport, error) {
	log.Infof("querying cluster capacity for nodes matching %q", h.options.NodeSelector)
	cluster, err := h.client.GetClusterCapacity(h.options.NodeSelector)
	if err != nil {
		return nil, fmt.Errorf("getting cluster capacity: %w", err)
	}
	log.Infof("cluster capacity: %d nodes, %.0f CPU cores, %.1f GiB memory",
		cluster.NodeCount, cluster.CPUCores, cluster.MemoryBytes/(1024*1024*1024))

	window := fmt.Sprintf("%dd", h.options.DataDays)
	log.Infof("querying job metrics for namespace %q over %s window", h.options.Namespace, window)
	rawJobs, err := h.client.GetJobMetrics(h.options.Namespace, window)
	if err != nil {
		return nil, fmt.Errorf("getting job metrics: %w", err)
	}

	if len(rawJobs) == 0 {
		log.Warn("no job metrics found")
	}

	report := BuildReport(rawJobs, *cluster, h.options.DataDays, h.options.Source, time.Now(), MapJobToSIG)

	if h.options.MonthlyCost > 0 {
		ApplyCostRates(report, h.options.MonthlyCost)
		log.Infof("applied monthly cost of $%.2f", h.options.MonthlyCost)
	}

	if err := GenerateReport(report, h.options.Path, h.options.Source); err != nil {
		return nil, fmt.Errorf("writing output: %w", err)
	}

	return report, nil
}

// MapJobToSIG maps a Prow job name to a SIG, matching the logic in pkg/sigretests.
func MapJobToSIG(jobName string) string {
	switch {
	case strings.Contains(jobName, "sig-compute") || strings.Contains(jobName, "vgpu"):
		return "compute"
	case strings.Contains(jobName, "sig-network") || strings.Contains(jobName, "sriov"):
		return "network"
	case strings.Contains(jobName, "sig-storage"):
		return "storage"
	case strings.Contains(jobName, "sig-operator"):
		return "operator"
	case strings.Contains(jobName, "sig-monitoring"):
		return "monitoring"
	default:
		return "other"
	}
}
