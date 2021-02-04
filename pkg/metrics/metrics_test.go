package metrics_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

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
			subject.SetMergeQueueLength(source, mql)
			subject.SetTimeToMerge(source, ttm)

			actual := subject.String()
			for _, expected := range expectedItems {
				Expect(actual).To(ContainSubstring(expected))
			}
		},
		table.Entry("help string for mql is present", 0.0, 0.0, "kubevirt/kubevirt", []string{`# HELP cihealth_merge_queue_lenght_average The average number of PRs in the merge queue.`}),
		table.Entry("help string for ttm is present", 0.0, 0.0, "kubevirt/kubevirt", []string{`# HELP cihealth_time_to_merge_seconds The average time it took to merge PRs.`}),
		table.Entry("mql can be set", 10.0, 0.0, "kubevirt/kubevirt", []string{`cihealth_merge_queue_lenght_average{method="kubevirt/kubevirt"} 10`}),
		table.Entry("ttm can be set", 0.0, 10.0, "kubevirt/kubevirt", []string{`cihealth_time_to_merge_seconds{method="kubevirt/kubevirt"} 10`}),
		table.Entry("mql and ttm can be set", 10.0, 10.0, "kubevirt/kubevirt", []string{
			`cihealth_merge_queue_lenght_average{method="kubevirt/kubevirt"} 10`,
			`cihealth_time_to_merge_seconds{method="kubevirt/kubevirt"} 10`}),
	)
})
