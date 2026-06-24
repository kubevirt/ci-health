package main

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/kubevirt/ci-health/pkg/cost"
)

const (
	defaultThanosURL    = "https://thanos-querier.openshift-monitoring.svc:9091"
	defaultNamespace    = "kubevirt-prow-jobs"
	defaultNodeSelector = "kubevirt-worker-.*"
	defaultDataDays     = 7
	defaultSource       = "kubevirt/kubevirt"
	defaultPath         = "./output"
	defaultLogLevel     = "info"
)

func main() {
	opts := &cost.HandlerOptions{
		ThanosURL:    defaultThanosURL,
		Namespace:    defaultNamespace,
		NodeSelector: defaultNodeSelector,
		DataDays:     defaultDataDays,
		Source:       defaultSource,
		Path:         defaultPath,
	}

	var bearerTokenPath string
	var logLevel string

	cmd := &cobra.Command{
		Use:   "cost",
		Short: "Generate CI resource usage report per PR",
		Long:  "Queries Prometheus/Thanos for Prow job pod resource usage and generates a report showing what percentage of cluster capacity each PR consumes.",
		Run: func(cmd *cobra.Command, args []string) {
			if logLevel == "debug" {
				log.SetLevel(log.DebugLevel)
			}

			if bearerTokenPath != "" {
				token, err := os.ReadFile(bearerTokenPath)
				if err != nil {
					log.Fatalf("reading bearer token: %v", err)
				}
				opts.BearerToken = string(token)
			}

			handler := cost.NewHandler(opts)
			report, err := handler.Run()
			if err != nil {
				log.Fatalf("error: %v", err)
			}

			log.Infof("report complete: %d PRs, %d runs, cluster: %d nodes (%.0f cores)",
				report.PRCount, report.RunCount, report.Cluster.NodeCount, report.Cluster.CPUCores)
		},
	}

	flags := cmd.Flags()
	flags.StringVar(&opts.ThanosURL, "thanos-url", opts.ThanosURL, "Thanos querier URL")
	flags.StringVar(&bearerTokenPath, "bearer-token-path", "", "Path to a file containing a bearer token for Thanos auth")
	flags.StringVar(&opts.Namespace, "namespace", opts.Namespace, "Namespace where Prow job pods run")
	flags.StringVar(&opts.NodeSelector, "node-selector", opts.NodeSelector, "Regex to match worker node names for capacity calculation")
	flags.IntVar(&opts.DataDays, "data-days", opts.DataDays, "Number of days to look back")
	flags.StringVar(&opts.Source, "source", opts.Source, "GitHub org/repo to report on")
	flags.StringVar(&opts.Path, "path", opts.Path, "Output directory for results")
	flags.Float64Var(&opts.MonthlyCost, "monthly-cost", 0, "Optional total monthly infrastructure cost in USD for dollar conversion")
	flags.StringVar(&logLevel, "log-level", defaultLogLevel, "Log level (debug or info)")

	if err := cmd.Execute(); err != nil {
		log.Fatalf("error: %v", err)
	}
}
