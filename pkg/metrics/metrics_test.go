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

var _ = Describe("Handler", func() {
	BeforeSuite(func() {
		subject = metrics.NewHandler()
	})

	table.DescribeTable("String: should a string with the current metric values",
		func(mql, ttm float64, source string, expectedItems []string) {
			subject.SetAvgMergeQueueLength(source, mql)
			subject.SetAvgTimeToMerge(source, ttm)

			actual := subject.String()
			for _, expected := range expectedItems {
				Expect(actual).To(ContainSubstring(expected))
			}
		},
		table.Entry("help string for mql is present", 0.0, 0.0, "kubevirt/kubevirt", []string{fmt.Sprintf("# HELP %s", constants.AvgMergeQueueLengthMetricName)}),
		table.Entry("help string for ttm is present", 0.0, 0.0, "kubevirt/kubevirt", []string{fmt.Sprintf("# HELP %s", constants.AvgTimeToMergeMetricName)}),
		table.Entry("mql can be set", 10.0, 0.0, "kubevirt/kubevirt", []string{fmt.Sprintf(`%s{source="kubevirt/kubevirt"} 10`, constants.AvgMergeQueueLengthMetricName)}),
		table.Entry("ttm can be set", 0.0, 10.0, "kubevirt/kubevirt", []string{fmt.Sprintf(`%s{source="kubevirt/kubevirt"} 10`, constants.AvgTimeToMergeMetricName)}),
		table.Entry("mql and ttm can be set", 10.0, 10.0, "kubevirt/kubevirt", []string{
			fmt.Sprintf(`%s{source="kubevirt/kubevirt"} 10`, constants.AvgMergeQueueLengthMetricName),
			fmt.Sprintf(`%s{source="kubevirt/kubevirt"} 10`, constants.AvgTimeToMergeMetricName)}),
	)
})
