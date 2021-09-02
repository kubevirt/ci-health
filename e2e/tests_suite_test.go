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
	"github.com/fgimenez/ci-health/pkg/types"
)

func TestCiHealth(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ci-health Suite")
}

var (
	tokenPath string
	names     = []string{
		constants.MergeQueueLengthName,
		constants.TimeToMergeName,
		constants.RetestsToMergeName,
		constants.MergedPRsName,
	}
	badgeFileNames = []string{
		constants.MergeQueueLengthBadgeFileName,
		constants.TimeToMergeBadgeFileName,
		constants.RetestsToMergeBadgeFileName,
		constants.MergedPRsBadgeFileName,
	}
	avgNames = []string{
		constants.AvgMergeQueueLengthMetricName,
		constants.AvgTimeToMergeMetricName,
		constants.AvgRetestsToMergeMetricName,
	}
	stdNames = []string{
		constants.StdMergeQueueLengthMetricName,
		constants.StdTimeToMergeMetricName,
		constants.StdRetestsToMergeMetricName,
	}
	statNames = append(avgNames, stdNames...)
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
			Action: types.StatsAction,

			Path:      artifactsDir,
			TokenPath: tokenPath,
			Source:    source,
			DataDays:  dataDays,
			LogLevel:  "debug",
		}

		checkResults := func(results *types.Results) {
			Expect(results.DataDays).To(Equal(dataDays))
			Expect(results.Source).To(Equal(source))

			Expect(results.Data).To(HaveLen(len(names)))

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

		sourceArtifactsDir := filepath.Join(artifactsDir, source)

		By("Checking badge files")
		for _, badgeFileName := range badgeFileNames {
			if badgeFileName == constants.MergedPRsBadgeFileName {
				continue
			}
			fullBadgeFileName := filepath.Join(sourceArtifactsDir, badgeFileName)
			_, err := os.Stat(fullBadgeFileName)
			Expect(err).ToNot(HaveOccurred())
		}

		By("Checking JSON file")
		jsonFileName := filepath.Join(sourceArtifactsDir, constants.JSONResultsFileName)
		jsonData, err := ioutil.ReadFile(jsonFileName)
		Expect(err).ToNot(HaveOccurred())

		var jsonResults types.Results
		err = json.Unmarshal(jsonData, &jsonResults)
		Expect(err).ToNot(HaveOccurred())

		checkResults(&jsonResults)

		By("Checking metrics file")
		metricsFileName := filepath.Join(artifactsDir, constants.MetricsFileName)
		metricsDataBytes, err := ioutil.ReadFile(metricsFileName)
		Expect(err).ToNot(HaveOccurred())
		metricsData := string(metricsDataBytes)

		expectedMetricsStrings := []string{}

		for _, statName := range statNames {
			expectedMetricsStrings = append(expectedMetricsStrings, fmt.Sprintf("# HELP %s", statName))
			expectedMetricsStrings = append(expectedMetricsStrings, fmt.Sprintf(`%s{source="kubevirt/kubevirt"}`, statName))
		}
		for _, expected := range expectedMetricsStrings {
			Expect(metricsData).To(ContainSubstring(expected))
		}
		notExpectedMetricsStrings := []string{
			"# HELP process_virtual_memory_max_bytes",
			"# HELP go_gc_duration_seconds",
		}
		for _, notExpected := range notExpectedMetricsStrings {
			Expect(metricsData).ToNot(ContainSubstring(notExpected))
		}
	})
})
