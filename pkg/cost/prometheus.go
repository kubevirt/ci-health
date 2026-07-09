package cost

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/avast/retry-go"
	log "github.com/sirupsen/logrus"
)

// PrometheusClient queries a Prometheus-compatible API (including Thanos Querier).
type PrometheusClient struct {
	baseURL    string
	httpClient *http.Client
}

type prometheusResponse struct {
	Status string         `json:"status"`
	Data   prometheusData `json:"data"`
	Error  string         `json:"error,omitempty"`
}

type prometheusData struct {
	ResultType string             `json:"resultType"`
	Result     []PrometheusResult `json:"result"`
}

type PrometheusResult struct {
	Metric map[string]string `json:"metric"`
	Value  [2]interface{}    `json:"value"`
}

// NewPrometheusClient creates a client for querying Prometheus/Thanos.
// If bearerToken is non-empty, it is sent as an Authorization header.
// TLS verification is skipped for in-cluster OpenShift Thanos endpoints.
func NewPrometheusClient(baseURL, bearerToken string, insecureSkipVerify bool) *PrometheusClient {
	if insecureSkipVerify && bearerToken == "" {
		log.Warn("TLS verification disabled with no bearer token — connections are unauthenticated and unverified")
	}

	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: insecureSkipVerify},
	}

	client := &http.Client{
		Transport: &authRoundTripper{
			token: bearerToken,
			inner: transport,
		},
		Timeout: 30 * time.Second,
	}

	return &PrometheusClient{
		baseURL:    baseURL,
		httpClient: client,
	}
}

type authRoundTripper struct {
	token string
	inner http.RoundTripper
}

func (a *authRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	if a.token != "" {
		req.Header.Set("Authorization", "Bearer "+a.token)
	}
	return a.inner.RoundTrip(req)
}

// Query executes an instant PromQL query and returns the results.
func (c *PrometheusClient) Query(query string) ([]PrometheusResult, error) {
	u, err := url.Parse(c.baseURL + "/api/v1/query")
	if err != nil {
		return nil, fmt.Errorf("parsing URL: %w", err)
	}
	q := u.Query()
	q.Set("query", query)
	u.RawQuery = q.Encode()

	queryLog := log.WithField("query", query)

	var results []PrometheusResult
	err = retry.Do(
		func() error {
			resp, err := c.httpClient.Get(u.String())
			if err != nil {
				queryLog.Warnf("request failed (will retry): %v", err)
				return err
			}
			defer resp.Body.Close()

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				return retry.Unrecoverable(fmt.Errorf("reading response: %w", err))
			}

			if resp.StatusCode == http.StatusServiceUnavailable || resp.StatusCode == http.StatusGatewayTimeout {
				queryLog.Infof("got %d, will retry", resp.StatusCode)
				return fmt.Errorf("prometheus returned %d", resp.StatusCode)
			}
			if resp.StatusCode != http.StatusOK {
				return retry.Unrecoverable(fmt.Errorf("prometheus returned %d: %s", resp.StatusCode, string(body)))
			}

			var promResp prometheusResponse
			if err := json.Unmarshal(body, &promResp); err != nil {
				return retry.Unrecoverable(fmt.Errorf("decoding response: %w", err))
			}
			if promResp.Status != "success" {
				return retry.Unrecoverable(fmt.Errorf("query failed: %s", promResp.Error))
			}

			results = promResp.Data.Result
			return nil
		},
		retry.Attempts(3),
		retry.Delay(2*time.Second),
	)

	if err != nil {
		return nil, fmt.Errorf("querying prometheus: %w", err)
	}

	queryLog.Debugf("returned %d results", len(results))
	return results, nil
}

// ResultValue extracts the float64 value from a Prometheus instant query result.
func ResultValue(r PrometheusResult) (float64, error) {
	valStr, ok := r.Value[1].(string)
	if !ok {
		return 0, fmt.Errorf("unexpected value type: %T", r.Value[1])
	}
	return strconv.ParseFloat(valStr, 64)
}

// GetClusterCapacity queries total allocatable CPU and memory from worker nodes.
// nodeSelector is a PromQL regex pattern (e.g. "kubevirt-worker-.*").
func (c *PrometheusClient) GetClusterCapacity(nodeSelector string) (*ClusterCapacity, error) {
	cpuQuery := fmt.Sprintf(`sum(kube_node_status_allocatable{resource="cpu", node=~"%s"})`, nodeSelector)
	memQuery := fmt.Sprintf(`sum(kube_node_status_allocatable{resource="memory", node=~"%s"})`, nodeSelector)
	countQuery := fmt.Sprintf(`count(kube_node_status_allocatable{resource="cpu", node=~"%s"})`, nodeSelector)

	cpuResults, err := c.Query(cpuQuery)
	if err != nil {
		return nil, fmt.Errorf("querying CPU capacity: %w", err)
	}
	memResults, err := c.Query(memQuery)
	if err != nil {
		return nil, fmt.Errorf("querying memory capacity: %w", err)
	}
	countResults, err := c.Query(countQuery)
	if err != nil {
		return nil, fmt.Errorf("querying node count: %w", err)
	}

	if len(cpuResults) == 0 || len(memResults) == 0 || len(countResults) == 0 {
		return nil, fmt.Errorf("no capacity data returned for nodes matching %q", nodeSelector)
	}

	cpuCores, err := ResultValue(cpuResults[0])
	if err != nil {
		return nil, fmt.Errorf("parsing CPU capacity: %w", err)
	}
	memBytes, err := ResultValue(memResults[0])
	if err != nil {
		return nil, fmt.Errorf("parsing memory capacity: %w", err)
	}
	nodeCount, err := ResultValue(countResults[0])
	if err != nil {
		return nil, fmt.Errorf("parsing node count: %w", err)
	}

	return &ClusterCapacity{
		CPUCores:    cpuCores,
		MemoryBytes: memBytes,
		NodeCount:   int(nodeCount),
	}, nil
}

// RawJobMetrics holds the raw Prometheus data for a single job pod.
type RawJobMetrics struct {
	Pod      string
	PR       string
	Job      string
	Repo     string
	Org      string
	BuildID  string
	CPUSec   float64
	MemBytes float64
}

// GetJobMetrics queries per-pod CPU and memory usage for completed Prow jobs
// in the given namespace over the specified window.
func (c *PrometheusClient) GetJobMetrics(namespace, window string) ([]RawJobMetrics, error) {
	cpuQuery := fmt.Sprintf(
		`sum(increase(container_cpu_usage_seconds_total{namespace="%s", container!="POD", container!=""}[%s])) by (pod)`,
		namespace, window,
	)
	memQuery := fmt.Sprintf(
		`avg_over_time(container_memory_working_set_bytes{namespace="%s", container!="POD", container!=""}[%s]) by (pod)`,
		namespace, window,
	)
	labelsQuery := fmt.Sprintf(
		`kube_pod_labels{namespace="%s", label_prow_k8s_io_refs_pull!=""}`,
		namespace,
	)

	cpuResults, err := c.Query(cpuQuery)
	if err != nil {
		return nil, fmt.Errorf("querying CPU usage: %w", err)
	}
	memResults, err := c.Query(memQuery)
	if err != nil {
		return nil, fmt.Errorf("querying memory usage: %w", err)
	}
	labelResults, err := c.Query(labelsQuery)
	if err != nil {
		return nil, fmt.Errorf("querying pod labels: %w", err)
	}

	cpuByPod := make(map[string]float64)
	for _, r := range cpuResults {
		pod := r.Metric["pod"]
		val, err := ResultValue(r)
		if err != nil {
			log.Warnf("skipping CPU result for pod %s: %v", pod, err)
			continue
		}
		cpuByPod[pod] = val
	}

	memByPod := make(map[string]float64)
	for _, r := range memResults {
		pod := r.Metric["pod"]
		val, err := ResultValue(r)
		if err != nil {
			log.Warnf("skipping memory result for pod %s: %v", pod, err)
			continue
		}
		memByPod[pod] = val
	}

	labelsByPod := make(map[string]map[string]string)
	for _, r := range labelResults {
		pod := r.Metric["pod"]
		labelsByPod[pod] = r.Metric
	}

	var jobs []RawJobMetrics
	for pod, labels := range labelsByPod {
		cpu, hasCPU := cpuByPod[pod]
		mem, hasMem := memByPod[pod]
		if !hasCPU && !hasMem {
			continue
		}

		jobs = append(jobs, RawJobMetrics{
			Pod:     pod,
			PR:      labels["label_prow_k8s_io_refs_pull"],
			Job:     labels["label_prow_k8s_io_job"],
			Repo:    labels["label_prow_k8s_io_refs_repo"],
			Org:     labels["label_prow_k8s_io_refs_org"],
			BuildID: labels["label_prow_k8s_io_build_id"],
			CPUSec:  cpu,
			MemBytes: mem,
		})
	}

	log.Infof("found %d job pods with metrics in namespace %s", len(jobs), namespace)
	return jobs, nil
}
