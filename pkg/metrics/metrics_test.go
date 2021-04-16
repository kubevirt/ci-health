package metrics_test

import (
	"fmt"
	"testing"

	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	"github.com/fgimenez/ci-health/pkg/constants"
	"github.com/fgimenez/ci-health/pkg/metrics"
)

func TestMetrics(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "metrics Suite")
}

var (
	subject *metrics.Handler
)

const (
	defaultSource = "kubevirt/kubevirt"
)

var _ = Describe("Handler", func() {
	BeforeSuite(func() {
		subject = metrics.NewHandler()
	})

	table.DescribeTable("String: should return a string with the current metric values",
		func(avgmql, avgttm, avgrtm, stdmql, stdttm, stdrtm float64, expectedItems []string) {
			subject.SetAvgMergeQueueLength(defaultSource, avgmql)
			subject.SetAvgTimeToMerge(defaultSource, avgttm)
			subject.SetAvgRetestsToMerge(defaultSource, avgrtm)
			subject.SetStdMergeQueueLength(defaultSource, stdmql)
			subject.SetStdTimeToMerge(defaultSource, stdttm)
			subject.SetStdRetestsToMerge(defaultSource, stdrtm)

			actual := subject.String()
			for _, expected := range expectedItems {
				Expect(actual).To(ContainSubstring(expected))
			}
		},
		table.Entry("help string for avg mql is present", 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, []string{fmt.Sprintf("# HELP %s", constants.AvgMergeQueueLengthMetricName)}),
		table.Entry("help string for avg ttm is present", 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, []string{fmt.Sprintf("# HELP %s", constants.AvgTimeToMergeMetricName)}),
		table.Entry("help string for avg rtm is present", 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, []string{fmt.Sprintf("# HELP %s", constants.AvgRetestsToMergeMetricName)}),
		table.Entry("help string for std mql is present", 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, []string{fmt.Sprintf("# HELP %s", constants.StdMergeQueueLengthMetricName)}),
		table.Entry("help string for std ttm is present", 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, []string{fmt.Sprintf("# HELP %s", constants.StdTimeToMergeMetricName)}),
		table.Entry("help string for std rtm is present", 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, []string{fmt.Sprintf("# HELP %s", constants.StdRetestsToMergeMetricName)}),
		table.Entry("avg mql can be set", 10.0, 0.0, 0.0, 0.0, 0.0, 0.0, []string{fmt.Sprintf(`%s{source="%s"} 10`, constants.AvgMergeQueueLengthMetricName, defaultSource)}),
		table.Entry("avg ttm can be set", 0.0, 10.0, 0.0, 0.0, 0.0, 0.0, []string{fmt.Sprintf(`%s{source="%s"} 10`, constants.AvgTimeToMergeMetricName, defaultSource)}),
		table.Entry("avg rtm can be set", 0.0, 0.0, 10.0, 0.0, 0.0, 0.0, []string{fmt.Sprintf(`%s{source="%s"} 10`, constants.AvgRetestsToMergeMetricName, defaultSource)}),
		table.Entry("std mql can be set", 0.0, 0.0, 0.0, 10.0, 0.0, 0.0, []string{fmt.Sprintf(`%s{source="%s"} 10`, constants.StdMergeQueueLengthMetricName, defaultSource)}),
		table.Entry("std ttm can be set", 0.0, 0.0, 0.0, 0.0, 10.0, 0.0, []string{fmt.Sprintf(`%s{source="%s"} 10`, constants.StdTimeToMergeMetricName, defaultSource)}),
		table.Entry("std rtm can be set", 0.0, 0.0, 0.0, 0.0, 0.0, 10.0, []string{fmt.Sprintf(`%s{source="%s"} 10`, constants.StdRetestsToMergeMetricName, defaultSource)}),
		table.Entry("mql, ttm and rtm can be set", 10.0, 10.0, 10.0, 10.0, 10.0, 10.0, []string{
			fmt.Sprintf(`%s{source="%s"} 10`, constants.AvgMergeQueueLengthMetricName, defaultSource),
			fmt.Sprintf(`%s{source="%s"} 10`, constants.AvgTimeToMergeMetricName, defaultSource),
			fmt.Sprintf(`%s{source="%s"} 10`, constants.AvgRetestsToMergeMetricName, defaultSource),
			fmt.Sprintf(`%s{source="%s"} 10`, constants.StdMergeQueueLengthMetricName, defaultSource),
			fmt.Sprintf(`%s{source="%s"} 10`, constants.StdTimeToMergeMetricName, defaultSource),
			fmt.Sprintf(`%s{source="%s"} 10`, constants.StdRetestsToMergeMetricName, defaultSource),
		}),
	)

	table.DescribeTable("String: should not return generic metric values",
		func(notExpected string) {
			subject.SetAvgMergeQueueLength(defaultSource, 1.0)
			subject.SetAvgTimeToMerge(defaultSource, 1.0)

			actual := subject.String()
			Expect(actual).ToNot(ContainSubstring(notExpected))
		},
		table.Entry("help string for process_virtual_memory_max_bytes is not present", "# HELP process_virtual_memory_max_bytes"),
		table.Entry("help string for go_gc_duration_seconds is not present", "# HELP go_gc_duration_seconds"),
	)
})
