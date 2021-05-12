package timeutils_test

import (
	"fmt"
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	"github.com/fgimenez/ci-health/pkg/timeutils"
)

func TestTimeutils(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "timeutils Suite")
}

var _ = Describe("ClosestMonday", func() {
	table.DescribeTable("should return the closest monday to the given date",
		func(date string, expectedDate time.Time, expectedErr error) {
			actualDate, actualErr := timeutils.ClosestMonday(date)
			Expect(actualDate).To(Equal(expectedDate))
			if expectedErr == nil {
				Expect(actualErr).To(BeNil())
			} else if actualErr == nil {
				Fail(fmt.Sprintf("Actual error is nil, expected to be %v", expectedErr))
			} else {
				Expect(actualErr.Error()).To(Equal(expectedErr.Error()))
			}
		},
		table.Entry("invalid date",
			"not a valid date",
			time.Time{},
			fmt.Errorf("Not a valid date: %q", "not a valid date"),
		),
		table.Entry("valid date, monday",
			"2021-05-03",
			time.Date(2021, time.Month(5), 3, 0, 0, 0, 0, time.UTC),
			nil,
		),
		table.Entry("valid date, sunday",
			"2021-05-02",
			time.Date(2021, time.Month(5), 3, 0, 0, 0, 0, time.UTC),
			nil,
		),
		table.Entry("valid date, different month",
			"2021-04-28",
			time.Date(2021, time.Month(5), 3, 0, 0, 0, 0, time.UTC),
			nil,
		),
		table.Entry("valid date, different year",
			"2020-12-30",
			time.Date(2021, time.Month(1), 4, 0, 0, 0, 0, time.UTC),
			nil,
		),
		table.Entry("valid date, february in leap year",
			"2020-02-28",
			time.Date(2020, time.Month(3), 2, 0, 0, 0, 0, time.UTC),
			nil,
		),
	)
})
