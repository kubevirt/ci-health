package e2e

import (
	"os"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/fgimenez/ci-health/pkg/constants"
	"github.com/fgimenez/ci-health/pkg/runner"
	"github.com/fgimenez/ci-health/pkg/types"
)

func TestCiHealth(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ci-health Suite")
}

var (
	tokenPath string
)

const (
	source   = "kubevirt/kubevirt"
	dataDays = 1
)

var _ = BeforeSuite(func() {
	tokenPath = os.Getenv("GITHUB_TOKEN_PATH")
	if tokenPath == "" {
		Fail("Please specify an OAuth2 token in the env var GITHUB_TOKEN_PATH")
	}
})

var _ = Describe("ci-health stats", func() {
	It("Retrieves data from github", func() {
		opt := &types.Options{
			TokenPath: tokenPath,
			Source:    source,
			DataDays:  dataDays,
			LogLevel:  "debug",
		}

		results, err := runner.Run(opt)
		Expect(err).ToNot(HaveOccurred())

		Expect(results.DataDays).To(Equal(dataDays))
		Expect(results.Source).To(Equal(source))

		Expect(results.Data).To(HaveLen(2))

		names := []string{constants.MergeQueueLengthName, constants.TimeToMergeName}

		for i, name := range names {
			metricResults := results.Data[i]
			Expect(metricResults.Name).To(Equal(name))
			Expect(metricResults.Value).To(BeNumerically(">", 0))
			for _, dataPoint := range metricResults.DataPoints {
				Expect(dataPoint.Value).To(BeNumerically(">=", 0))
			}
		}
	})
})
