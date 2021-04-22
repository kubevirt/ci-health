package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/fgimenez/ci-health/pkg/constants"
	"github.com/fgimenez/ci-health/pkg/runner"
	"github.com/fgimenez/ci-health/pkg/types"
)

func main() {
	opt := &types.StatsOptions{
		BaseOptions: types.BaseOptions{
			Path:      constants.DefaultPath,
			TokenPath: constants.DefaultTokenPath,
			Source:    constants.DefaultSource,
			DataDays:  constants.DefaultDataDays,
			LogLevel:  constants.DefaultLogLevel,
		},
		TimeToMergeYellowLevel:      constants.DefaultTimeToMergeYellowLevel,
		TimeToMergeRedLevel:         constants.DefaultTimeToMergeRedLevel,
		MergeQueueLengthYellowLevel: constants.DefaultMergeQueueLengthYellowLevel,
		MergeQueueLengthRedLevel:    constants.DefaultMergeQueueLengthRedLevel,
		RetestsToMergeYellowLevel:   constants.DefaultRetestsToMergeYellowLevel,
		RetestsToMergeRedLevel:      constants.DefaultRetestsToMergeRedLevel,
	}

	cmd := &cobra.Command{
		Run: func(cmd *cobra.Command, arguments []string) {
			_, err := runner.RunStats(opt)
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
	flag.IntVar(&opt.DataDays, "data-days", opt.DataDays, "Number of days to retrieve data from.")
	flag.StringVar(&opt.LogLevel, "log-level", opt.LogLevel, "Log level, valid values are debug and info.")
	flag.Float64Var(&opt.TimeToMergeYellowLevel, "time-to-merge-yellow", opt.TimeToMergeYellowLevel, "Lower limit of the TimeToMerge metric above which the label will be yellow")
	flag.Float64Var(&opt.TimeToMergeRedLevel, "time-to-merge-red", opt.TimeToMergeRedLevel, "Lower limit of the TimeToMerge metric above which the label will be red")
	flag.Float64Var(&opt.MergeQueueLengthYellowLevel, "merge-queue-length-yellow", opt.MergeQueueLengthYellowLevel, "Lower limit of the MergeQueueLength metric above which the label will be yellow")
	flag.Float64Var(&opt.MergeQueueLengthRedLevel, "merge-queue-length-red", opt.MergeQueueLengthRedLevel, "Lower limit of the MergeQueueLength metric above which the label will be red")
	flag.Float64Var(&opt.RetestsToMergeYellowLevel, "retests-to-merge-yellow", opt.RetestsToMergeYellowLevel, "Lower limit of the RetestsToMerge metric above which the label will be yellow")
	flag.Float64Var(&opt.RetestsToMergeRedLevel, "retests-to-merge-red", opt.RetestsToMergeRedLevel, "Lower limit of the RetestsToMerge metric above which the label will be red")

	if err := cmd.Execute(); err != nil {
		log.Fatalf("error: %v", err)
	}
}
