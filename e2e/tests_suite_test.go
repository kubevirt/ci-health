package e2e

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/fgimenez/ci-health/pkg/constants"
	"github.com/fgimenez/ci-health/pkg/runner"
	"github.com/fgimenez/ci-health/pkg/stats"
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
	It("Retrieves data from github and writes badges", func() {
		artifactsDir, err := ioutil.TempDir("", "e2e-ci-health")
		Expect(err).ToNot(HaveOccurred())

		opt := &types.Options{
			Path:      artifactsDir,
			TokenPath: tokenPath,
			Source:    source,
			DataDays:  dataDays,
			LogLevel:  "debug",
		}

		checkResults := func(results *stats.Results) {
			Expect(results.DataDays).To(Equal(dataDays))
			Expect(results.Source).To(Equal(source))

			Expect(results.Data).To(HaveLen(2))

			names := []string{constants.MergeQueueLengthName, constants.TimeToMergeName}

			for _, name := range names {
				metricResults := results.Data[name]
				avg := metricResults.Avg
				Expect(avg).To(BeNumerically(">=", 0))
				std := metricResults.Std
				Expect(std).To(BeNumerically(">=", 0))
				for _, dataPoint := range metricResults.DataPoints {
					Expect(dataPoint.Value).To(BeNumerically(">=", 0))
				}
			}
		}

		results, err := runner.Run(opt)
		Expect(err).ToNot(HaveOccurred())

		By("Checking returned data")
		checkResults(results)

		By("Checking badge files")
		badgeFileNames := []string{
			filepath.Join(artifactsDir, constants.MergeQueueLengthBadgeFileName),
			filepath.Join(artifactsDir, constants.TimeToMergeBadgeFileName),
		}
		for _, badgeFileName := range badgeFileNames {
			_, err := os.Stat(badgeFileName)
			Expect(err).ToNot(HaveOccurred())
		}

		By("Checking JSON file")
		jsonFileName := filepath.Join(artifactsDir, constants.JSONResultsFileName)
		jsonData, err := ioutil.ReadFile(jsonFileName)
		Expect(err).ToNot(HaveOccurred())

		var jsonResults stats.Results
		err = json.Unmarshal(jsonData, &jsonResults)
		Expect(err).ToNot(HaveOccurred())

		checkResults(&jsonResults)

		By("Checking metrics file")
		metricsFileName := filepath.Join(artifactsDir, constants.MetricsFileName)
		metricsDataBytes, err := ioutil.ReadFile(metricsFileName)
		Expect(err).ToNot(HaveOccurred())
		metricsData := string(metricsDataBytes)

		expectedMetricsStrings := []string{
			fmt.Sprintf("# HELP %s", constants.AvgMergeQueueLengthMetricName),
			fmt.Sprintf("# HELP %s", constants.AvgTimeToMergeMetricName),
			fmt.Sprintf(`%s{source="kubevirt/kubevirt"}`, constants.AvgMergeQueueLengthMetricName),
			fmt.Sprintf(`%s{source="kubevirt/kubevirt"}`, constants.AvgTimeToMergeMetricName),
		}
		for _, expected := range expectedMetricsStrings {
			Expect(metricsData).To(ContainSubstring(expected))
		}
	})
})
