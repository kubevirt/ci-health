package prow

import (
	"encoding/json"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"os"
	"testing"
)

func TestProw(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "prow pkg Suite")
}

var _ = Describe("types", func() {
	When("started", func() {
		It("can be deserialized", func() {
			file, err := os.Open("testdata/started.json")
			Expect(err).ToNot(HaveOccurred())
			decoder := json.NewDecoder(file)
			var started *Started
			err = decoder.Decode(&started)
			Expect(err).ToNot(HaveOccurred())
		})
	})
	When("finished", func() {
		It("can be deserialized", func() {
			file, err := os.Open("testdata/finished.json")
			Expect(err).ToNot(HaveOccurred())
			decoder := json.NewDecoder(file)
			var finished *Finished
			err = decoder.Decode(&finished)
			Expect(err).ToNot(HaveOccurred())
		})
	})
})
