package main

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/fgimenez/cihealth/pkg/runner"
)

func main() {
	opt := &runner.Options{
		Path:      "/tmp/test",
		TokenPath: "",
		Source:    "kubevirt/kubevirt",
		DataDays:  7,
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

	if err := cmd.Execute(); err != nil {
		log.Fatalf("error: %v", err)
	}
}
