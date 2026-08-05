package cifailures

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCIFailures(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ci-failures Suite")
}
