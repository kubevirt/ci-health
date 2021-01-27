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
	YellowLimit = 1.0
	RedLimit    = 2.0
)

func TestOutput(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "output Suite")
}

var badgeLevels output.Levels

var _ = BeforeSuite(func() {
	badgeLevels = output.Levels{
		Yellow: YellowLimit,
		Red:    RedLimit,
	}
})

var _ = Describe("Badge", func() {

	table.DescribeTable("should return the right color for the given values and badgeLevels",
		func(value float64, expected badge.Color) {
			actual := output.BadgeColor(value, &badgeLevels)
			Expect(actual).To(Equal(expected))
		},
		table.Entry("zero value, expected green", 0.0, badge.ColorGreen),
		table.Entry("less than yellow limit, expected green", YellowLimit-0.1, badge.ColorGreen),
		table.Entry("yellow limit, expected yellow", YellowLimit, badge.ColorYellow),
		table.Entry("more than yellow and less than red limit, expected yellow", YellowLimit+0.1, badge.ColorYellow),
		table.Entry("red limit, expected red", RedLimit, badge.ColorRed),
		table.Entry("greater than red limit, expected red", RedLimit+0.1, badge.ColorRed),
	)
})
