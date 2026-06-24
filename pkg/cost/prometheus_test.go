package cost

import (
	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("resultValue", func() {
	table.DescribeTable("should parse prometheus result values",
		func(value interface{}, expected float64, shouldErr bool) {
			r := prometheusResult{Value: [2]interface{}{1234567890.0, value}}
			val, err := resultValue(r)
			if shouldErr {
				Expect(err).To(HaveOccurred())
			} else {
				Expect(err).NotTo(HaveOccurred())
				Expect(val).To(BeNumerically("~", expected, 0.001))
			}
		},
		table.Entry("integer string", "42", 42.0, false),
		table.Entry("float string", "3.14", 3.14, false),
		table.Entry("zero string", "0", 0.0, false),
		table.Entry("large number", "1234567890.123", 1234567890.123, false),
		table.Entry("non-string type errors", 42.0, 0.0, true),
		table.Entry("nil errors", nil, 0.0, true),
	)
})
