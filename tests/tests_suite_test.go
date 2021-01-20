package tests

import (
	"os"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/fgimenez/cihealth/pkg/stats"
)

func TestCIHealth(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "cihealth Suite")
}

var (
	tokenPath string
)

const (
	source   = "kubevirt/kubevirt"
	dataDays = 7
)

var _ = BeforeSuite(func() {
	tokenPath = os.Getenv("GITHUB_TOKEN_PATH")
	if tokenPath == "" {
		Fail("Please specify an OAuth2 token in the env var GITHUB_TOKEN_PATH")
	}
})

var _ = Describe("cihealth stats", func() {
	It("Retrieves data from github", func() {
		results, err := stats.Run(tokenPath, source, dataDays)
		Expect(err).ToNot(HaveOccurred())

		Expect(results.DataDays).To(Equal(dataDays))
		Expect(results.Source).To(Equal(source))
	})
})
