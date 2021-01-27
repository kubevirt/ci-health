package output_test

import (
	"testing"

	"github.com/narqo/go-badge"
	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	"github.com/fgimenez/ci-health/pkg/output"
)

const (
	YellowLimit = 2.0
	GreenLimit  = 1.0
)

func TestOutput(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "output Suite")
}

var badgeLevels output.Levels

var _ = BeforeSuite(func() {
	badgeLevels = output.Levels{
		Yellow: YellowLimit,
		Green:  GreenLimit,
	}
})

var _ = Describe("Badge", func() {

	table.DescribeTable("should return the right color for the given values and badgeLevels",
		func(value float64, expected badge.Color) {
			actual := output.BadgeColor(value, &badgeLevels)
			Expect(actual).To(Equal(expected))
		},
		table.Entry("zero value, expected green", 0.0, badge.ColorGreen),
		table.Entry("less than green limit, expected green", GreenLimit-0.1, badge.ColorGreen),
		table.Entry("green limit, expected green", GreenLimit, badge.ColorGreen),
		table.Entry("more than green and less than yellow limit, expected yellow", GreenLimit+0.1, badge.ColorYellow),
		table.Entry("yellow limit, expected yellow", YellowLimit, badge.ColorYellow),
		table.Entry("greater than yellow limit, expected red", YellowLimit+0.1, badge.ColorRed),
	)
})
