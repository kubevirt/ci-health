package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/kubevirt/ci-health/pkg/constants"
	"github.com/kubevirt/ci-health/pkg/htmlreport"
	"github.com/kubevirt/ci-health/pkg/types"
)

func main() {
	opt := &types.Options{
		Path:        constants.DefaultPath,
		Sig:         "compute",
		ResultsPath: constants.DefaultResultsPath,
	}
	kind := "sig"

	cmd := &cobra.Command{
		Run: func(cmd *cobra.Command, arguments []string) {
			switch kind {
			case "pr-retests":
				if err := htmlreport.GeneratePRRetest(opt); err != nil {
					log.Fatalf("error generating HTML report: %v", err)
				}
				log.Printf("HTML report written to %s/%s", opt.Path, constants.PRRetestReportFileName)
			case "sig":
				if err := htmlreport.Generate(opt); err != nil {
					log.Fatalf("error generating HTML report: %v", err)
				}
				reportPath := fmt.Sprintf("%s/sig-%s-failure-report.html", opt.Path, opt.Sig)
				log.Printf("HTML report written to %s", reportPath)
			default:
				log.Fatalf("unknown --kind %q (want sig or pr-retests)", kind)
			}
		},
	}
	flag := cmd.Flags()

	flag.StringVar(&opt.Path, "path", opt.Path, "Path to output HTML reports to")
	flag.StringVar(&opt.Sig, "sig", opt.Sig, "Name of SIG to generate failure report for (kind=sig)")
	flag.StringVar(&kind, "kind", kind, "Report kind: sig or pr-retests")
	flag.StringVar(&opt.ResultsPath, "results-path", opt.ResultsPath, "Path to the results.json file created by ci-health")

	if err := cmd.Execute(); err != nil {
		log.Fatalf("error: %v", err)
	}
}
