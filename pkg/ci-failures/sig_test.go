package cifailures

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("sig", func() {
	Context("SIGForGroup", func() {
		DescribeTable("can extract SIG",
			func(laneName, expectedSIG string) {
				sig, err := SIGForGroup(laneName)
				Expect(sig).To(BeEquivalentTo(expectedSIG))
				Expect(err).ToNot(HaveOccurred())
			},
			Entry("sig-compute-arm64", "pull-kubevirt-e2e-kind-1.36-sig-compute-arm64", "sig-compute"),
			Entry("vgpu", "pull-kubevirt-e2e-kind-1.35-vgpu", "sig-compute"),
			Entry("sev", "pull-kubevirt-e2e-kind-1.36-sev", "sig-compute"),
			Entry("sig-operator", "pull-kubevirt-e2e-k8s-sig-operator", "sig-compute"),
			Entry("sig-network", "pull-kubevirt-e2e-k8s-1.34-sig-network", "sig-network"),
			Entry("sig-network", "pull-kubevirt-e2e-k8s-1.36-sig-network-smoke", "sig-network"),
			Entry("sriov", "pull-kubevirt-e2e-kind-sriov", "sig-network"),
			Entry("sig-storage", "pull-kubevirt-e2e-k8s-1.35-sig-storage", "sig-storage"),
			Entry("sig-performance", "pull-kubevirt-e2e-k8s-1.36-sig-performance-kube-burner", "sig-performance"),
			Entry("sig-monitoring", "pull-kubevirt-e2e-k8s-1.34-sig-monitoring", "sig-monitoring"),
		)

		DescribeTable("can not extract SIG",
			func(laneName string) {
				sig, err := SIGForGroup(laneName)
				Expect(err).To(HaveOccurred())
				Expect(sig).To(BeEquivalentTo(UnknownSIG))
			},
			Entry("lint", "pull-kubevirt-metrics-lint"),
			Entry("prom-rules", "pull-kubevirt-prom-rules-verify"),
		)
	})
})
