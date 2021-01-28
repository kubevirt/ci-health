package stats_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	"github.com/fgimenez/ci-health/pkg/stats"
)

func TestStats(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "stats Suite")
}

var _ = Describe("Average", func() {
	table.DescribeTable("should return the average of the given floats",
		func(input []float64, expected float64) {
			actual := stats.Average(input)
			Expect(actual).To(Equal(expected))
		},
		table.Entry("no rounding, no entries", []float64{}, 0.0),
		table.Entry("no rounding, single entry", []float64{1}, 1.0),
		table.Entry("no rounding, multiple entries", []float64{2, 4, 6}, 4.0),
		table.Entry("rounding to the nearest, 2 decimal places", []float64{8, 9, 9, 15, 10, 7}, 9.67),
		table.Entry("rounding to the nearest, 2 decimal places, zero entries", []float64{0, 9, 9, 15, 10, 7}, 8.33),
		table.Entry("rounding to the nearest, 2 decimal places, negative entries", []float64{-5, 9, 9, 15, 10, 7}, 7.5),
	)
})
var _ = Describe("Std", func() {
	table.DescribeTable("should return the std of the given floats",
		func(input []float64, expected float64) {
			actual := stats.Std(input)
			Expect(actual).To(Equal(expected))
		},
		table.Entry("no entries", []float64{}, 0.0),
		table.Entry("single entry", []float64{1}, 0.0),
		table.Entry("multiple entries", []float64{2, 4, 6}, 1.63),
		table.Entry("multiple entries, 0 average", []float64{2, 4, 6, -2, -4, -6}, 4.32),
		table.Entry("multiple entries, outliers", []float64{0, 9, 9, 15, 10, 100}, 34.35),
		table.Entry("multiple entries, negative", []float64{0, -9, 9, 15, 10, -100}, 39.89),
	)
})
