package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/fgimenez/cihealth/pkg/runner"
	"github.com/fgimenez/cihealth/pkg/types"
)

func main() {
	opt := &types.Options{
		Path:      "/tmp/test",
		TokenPath: "",
		Source:    "kubevirt/kubevirt",
		DataDays:  7,
		LogLevel:  "info",
	}

	cmd := &cobra.Command{
		Run: func(cmd *cobra.Command, arguments []string) {
			path, err := runner.Run(opt)
			if err != nil {
				log.Fatalf("error: %v", err)
			}
			log.Println("Output written to", path)
		},
	}
	flag := cmd.Flags()

	flag.StringVar(&opt.Path, "path", opt.Path, "The directory to save results to.")
	flag.StringVar(&opt.TokenPath, "gh-token", opt.TokenPath, "OAuth2 token to interact with GitHub API.")
	flag.StringVar(&opt.Source, "source", opt.Source, "GitHub repo from where retrieve the data.")
	flag.IntVar(&opt.DataDays, "data-days", opt.DataDays, "Number of days to retrieve data from.")
	flag.StringVar(&opt.LogLevel, "log-level", opt.LogLevel, "Log level, valid values are debug and info.")

	if err := cmd.Execute(); err != nil {
		log.Fatalf("error: %v", err)
	}
}
