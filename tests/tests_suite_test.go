package tests

import (
	"os"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCIHealth(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "cihealth Suite")
}

var (
	token string
)

var _ = BeforeSuite(func() {
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		Fail("Please specify an OAuth2 token in the env var GITHUB_TOKEN")
	}
})

var _ = Describe("cihealth stats", func() {
	It("Retrieves data from github", func() {

	})
})
