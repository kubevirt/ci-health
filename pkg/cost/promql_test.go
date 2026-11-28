package cost_test

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"text/template"
	"time"

	"github.com/kubevirt/ci-health/pkg/cost"
)

// The fixture contains data and expectations only. Expressions always come
// from the HTTP requests sent by the production client, preventing query drift.
func TestCostPromQL(t *testing.T) {
	for _, tc := range []struct {
		name, namespace, nodeSelector, window  string
		cpuCapacity, memoryCapacity, nodeCount int
		cpuA, cpuB, memoryA                    int
	}{
		{"workers/5m", "prow-jobs", "worker-.*", "5m", 12, 3000, 2, 450, 75, 300},
		{"single-node/10m", "ci-jobs", "worker-a", "10m", 4, 1000, 1, 900, 150, 1000},
	} {
		t.Run(tc.name, func(t *testing.T) {
			queries := captureCostQueries(t, func(client *cost.PrometheusClient) error {
				if _, err := client.GetClusterCapacity(tc.nodeSelector); err != nil {
					return err
				}
				_, err := client.GetJobMetrics(tc.namespace, tc.window)
				return err
			})
			if len(queries) != 6 {
				t.Fatalf("expected all six cost queries, got %d: %v", len(queries), queries)
			}
			fixture, err := template.ParseFiles("testdata/promql.yaml.tmpl")
			if err != nil {
				t.Fatal(err)
			}
			var data bytes.Buffer
			err = fixture.Execute(&data, map[string]interface{}{
				"Queries": queries, "Namespace": tc.namespace,
				"CPUCapacity": tc.cpuCapacity, "MemoryCapacity": tc.memoryCapacity,
				"NodeCount": tc.nodeCount, "CPUA": tc.cpuA, "CPUB": tc.cpuB, "MemoryA": tc.memoryA,
				"HistoricalLabels": tc.window == "10m",
			})
			if err != nil {
				t.Fatal(err)
			}
			if output, err := runPromQLFixture(t, data.Bytes()); err != nil {
				t.Fatalf("cost query validation failed: %v\n%s\nqueries: %v", err, output, queries)
			}
		})
	}
}

func TestCostPromQLRejectsMalformedQuery(t *testing.T) {
	queries := captureCostQueries(t, func(client *cost.PrometheusClient) error {
		_, err := client.Query("sum(")
		return err
	})
	if len(queries) != 1 {
		t.Fatalf("expected one malformed query, got %v", queries)
	}
	fixture := fmt.Sprintf("tests:\n- promql_expr_test:\n  - expr: %q\n    eval_time: 0m\n    exp_samples: []\n", queries[0])
	output, err := runPromQLFixture(t, []byte(fixture))
	if err == nil || !strings.Contains(output, "parse error") {
		t.Fatalf("expected a PromQL parse error, got %v\n%s", err, output)
	}
}

func captureCostQueries(t *testing.T, query func(*cost.PrometheusClient) error) []string {
	t.Helper()
	requests := make(chan string, 16)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet || r.URL.Path != "/api/v1/query" {
			t.Errorf("unexpected request: %s %s", r.Method, r.URL.Path)
			w.WriteHeader(http.StatusNotFound)
			return
		}
		requests <- r.URL.Query().Get("query")
		w.Header().Set("Content-Type", "application/json")
		// This response only allows the client to issue all queries; promtool
		// independently parses and evaluates each captured expression below.
		_, _ = fmt.Fprint(w, `{"status":"success","data":{"resultType":"vector","result":[{"metric":{},"value":[0,"1"]}]}}`)
	}))
	t.Cleanup(server.Close)
	if err := query(cost.NewPrometheusClient(server.URL, "", false)); err != nil {
		t.Fatal(err)
	}
	var queries []string
	for len(requests) > 0 {
		queries = append(queries, <-requests)
	}
	return queries
}

func runPromQLFixture(t *testing.T, fixture []byte) (string, error) {
	t.Helper()
	promtool, err := exec.LookPath("promtool")
	if err != nil {
		t.Fatal("promtool is required: run sh hack/install-promtool.sh <directory> from the repository root and add that directory to PATH (see README Development)")
	}
	path := filepath.Join(t.TempDir(), "promql.yaml")
	if err := os.WriteFile(path, fixture, 0600); err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	output, err := exec.CommandContext(ctx, promtool, "test", "rules", path).CombinedOutput()
	return string(output), err
}
