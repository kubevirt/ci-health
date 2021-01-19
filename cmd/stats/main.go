package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/fgimenez/cihealth/pkg/ghclient"
	"github.com/spf13/cobra"
	yaml "gopkg.in/yaml.v2"
)

type options struct {
	Path      string
	TokenPath string
	Source    string
	DataDays  int
}

func main() {
	opt := &options{
		Path:      "/tmp/test",
		TokenPath: "",
		Source:    "kubevirt/kubevirt",
		DataDays:  7,
	}

	cmd := &cobra.Command{
		Run: func(cmd *cobra.Command, arguments []string) {
			if err := opt.Run(); err != nil {
				log.Fatalf("error: %v", err)
			}
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

func (o *options) Run() error {
	if o.TokenPath == "" {
		return fmt.Errorf("You need to specify the GitHub token path with --gh-token")
	}

	results, err := ghclient.Run(o.TokenPath, o.Source, o.DataDays)
	if err != nil {
		return err
	}

	d, err := yaml.Marshal(&results)
	if err != nil {
		return err
	}

	log.Printf("Writing output file %s", o.Path)
	err = ioutil.WriteFile(o.Path, d, 0644)
	if err != nil {
		return err
	}

	return nil
}
