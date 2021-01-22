package gh_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGH(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "gh Suite")
}

var _ = Describe("DatePREnteredMergeQueue", func() {
	It("Determines the date a PR entered the merge queue", func() {
	})
})
