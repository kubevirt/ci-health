package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/fgimenez/ci-health/pkg/constants"
	"github.com/fgimenez/ci-health/pkg/runner"
	"github.com/fgimenez/ci-health/pkg/types"
)

func main() {
	opt := &types.Options{
		Action: types.BatchAction,

		Path:      constants.DefaultPath,
		TokenPath: constants.DefaultTokenPath,
		Source:    constants.DefaultSource,
		LogLevel:  constants.DefaultLogLevel,

		Mode:         types.FetchMode,
		TargetMetric: string(types.RetestsToMergeMetric),
		StartDate:    constants.DefaultBatchStartDate,
	}

	cmd := &cobra.Command{
		Run: func(cmd *cobra.Command, arguments []string) {
			_, err := runner.Run(opt)
			if err != nil {
				log.Fatalf("error: %v", err)
			}
			log.Println("Output written to", opt.Path)
		},
	}
	flag := cmd.Flags()

	flag.StringVar(&opt.Path, "path", opt.Path, "The directory to save results to.")
	flag.StringVar(&opt.TokenPath, "gh-token", opt.TokenPath, "OAuth2 token to interact with GitHub API.")
	flag.StringVar(&opt.Source, "source", opt.Source, "GitHub repo from where retrieve the data.")
	flag.StringVar(&opt.LogLevel, "log-level", opt.LogLevel, "Log level, valid values are debug and info.")

	flag.StringVar(&opt.TargetMetric, "target-metric", opt.TargetMetric, "Which metric to generate.")
	flag.StringVar(&opt.StartDate, "start-date", opt.StartDate, "From which do we start querying data in YYYY-MM-DD format")
	flag.StringVar(&opt.Mode, "mode", opt.Mode, "Batch mode, valid values are fetch (to download data) and plot (to generate graphs).")

	if err := cmd.Execute(); err != nil {
		log.Fatalf("error: %v", err)
	}
}
