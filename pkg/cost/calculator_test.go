package cost_test

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	"github.com/kubevirt/ci-health/pkg/cost"
	"github.com/kubevirt/ci-health/pkg/sigretests"
)

func TestCost(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "cost Suite")
}

var _ = Describe("CalculateJobUsages", func() {
	It("should calculate CPU percentage as share of total CI usage", func() {
		jobs := []cost.RawJobMetrics{
			{PR: "1", Job: "job-a", Repo: "kubevirt", Org: "kubevirt", BuildID: "100", CPUSec: 3600, MemBytes: 1024},
			{PR: "2", Job: "job-b", Repo: "kubevirt", Org: "kubevirt", BuildID: "101", CPUSec: 7200, MemBytes: 2048},
			{PR: "3", Job: "job-c", Repo: "kubevirt", Org: "kubevirt", BuildID: "102", CPUSec: 3600, MemBytes: 1024},
		}
		// total CPU = 14400, so job-a = 25%, job-b = 50%, job-c = 25%
		usages := cost.CalculateJobUsages(jobs)
		Expect(usages).To(HaveLen(3))
		Expect(usages[0].CPUPercent).To(BeNumerically("~", 25.0, 0.001))
		Expect(usages[1].CPUPercent).To(BeNumerically("~", 50.0, 0.001))
		Expect(usages[2].CPUPercent).To(BeNumerically("~", 25.0, 0.001))
	})

	It("should calculate memory percentage as share of total CI usage", func() {
		jobs := []cost.RawJobMetrics{
			{PR: "1", Job: "job-a", Repo: "kubevirt", Org: "kubevirt", BuildID: "100", CPUSec: 100, MemBytes: 4 * 1024 * 1024 * 1024},
			{PR: "2", Job: "job-b", Repo: "kubevirt", Org: "kubevirt", BuildID: "101", CPUSec: 100, MemBytes: 6 * 1024 * 1024 * 1024},
		}
		// total mem = 10 GiB, so job-a = 40%, job-b = 60%
		usages := cost.CalculateJobUsages(jobs)
		Expect(usages[0].MemPercent).To(BeNumerically("~", 40.0, 0.001))
		Expect(usages[1].MemPercent).To(BeNumerically("~", 60.0, 0.001))
	})

	It("should handle single job as 100%", func() {
		jobs := []cost.RawJobMetrics{
			{PR: "1", Job: "job-a", Repo: "kubevirt", Org: "kubevirt", BuildID: "100", CPUSec: 5000, MemBytes: 8192},
		}
		usages := cost.CalculateJobUsages(jobs)
		Expect(usages[0].CPUPercent).To(Equal(100.0))
		Expect(usages[0].MemPercent).To(Equal(100.0))
	})

	It("should handle empty input", func() {
		usages := cost.CalculateJobUsages([]cost.RawJobMetrics{})
		Expect(usages).To(HaveLen(0))
	})

	It("should handle all-zero usage without panic", func() {
		jobs := []cost.RawJobMetrics{
			{PR: "1", Job: "job-a", Repo: "kubevirt", Org: "kubevirt", BuildID: "100", CPUSec: 0, MemBytes: 0},
		}
		usages := cost.CalculateJobUsages(jobs)
		Expect(usages[0].CPUPercent).To(Equal(0.0))
		Expect(usages[0].MemPercent).To(Equal(0.0))
	})

	It("should preserve raw metrics in the output", func() {
		jobs := []cost.RawJobMetrics{
			{PR: "42", Job: "pull-kubevirt-e2e", Repo: "kubevirt", Org: "kubevirt", BuildID: "999", CPUSec: 7200, MemBytes: 4 * 1024 * 1024 * 1024},
		}
		usages := cost.CalculateJobUsages(jobs)
		Expect(usages[0].PR).To(Equal("42"))
		Expect(usages[0].Job).To(Equal("pull-kubevirt-e2e"))
		Expect(usages[0].BuildID).To(Equal("999"))
		Expect(usages[0].CPUSec).To(Equal(7200.0))
		Expect(usages[0].MemBytes).To(Equal(4 * 1024 * 1024 * 1024.0))
	})

	It("should sum to 100% across all jobs", func() {
		jobs := []cost.RawJobMetrics{
			{PR: "1", Job: "a", Repo: "r", Org: "o", BuildID: "1", CPUSec: 1000, MemBytes: 500},
			{PR: "2", Job: "b", Repo: "r", Org: "o", BuildID: "2", CPUSec: 2000, MemBytes: 300},
			{PR: "3", Job: "c", Repo: "r", Org: "o", BuildID: "3", CPUSec: 3000, MemBytes: 200},
		}
		usages := cost.CalculateJobUsages(jobs)
		var totalCPU, totalMem float64
		for _, u := range usages {
			totalCPU += u.CPUPercent
			totalMem += u.MemPercent
		}
		Expect(totalCPU).To(BeNumerically("~", 100.0, 0.001))
		Expect(totalMem).To(BeNumerically("~", 100.0, 0.001))
	})
})

var _ = Describe("AggregatePRUsages", func() {
	It("should aggregate jobs by PR", func() {
		jobs := []cost.JobUsage{
			{PR: "1", Repo: "kubevirt", Org: "kubevirt", Job: "e2e-test", CPUPercent: 2.0, MemPercent: 1.5},
			{PR: "1", Repo: "kubevirt", Org: "kubevirt", Job: "unit-test", CPUPercent: 0.5, MemPercent: 0.3},
			{PR: "2", Repo: "kubevirt", Org: "kubevirt", Job: "e2e-test", CPUPercent: 3.0, MemPercent: 2.0},
		}

		prs := cost.AggregatePRUsages(jobs)
		Expect(prs).To(HaveLen(2))

		prMap := map[string]cost.PRUsage{}
		for _, pr := range prs {
			prMap[pr.PR] = pr
		}

		Expect(prMap["1"].CPUPercent).To(Equal(2.5))
		Expect(prMap["1"].MemPercent).To(Equal(1.8))
		Expect(prMap["1"].RunCount).To(Equal(2))
		Expect(prMap["1"].Jobs).To(HaveLen(2))

		Expect(prMap["2"].CPUPercent).To(Equal(3.0))
		Expect(prMap["2"].RunCount).To(Equal(1))
	})

	It("should return empty slice for no jobs", func() {
		prs := cost.AggregatePRUsages([]cost.JobUsage{})
		Expect(prs).To(HaveLen(0))
	})
})

var _ = Describe("TopNPRs", func() {
	It("should return top N PRs sorted by CPU percentage", func() {
		prs := []cost.PRUsage{
			{PR: "1", CPUPercent: 1.0},
			{PR: "2", CPUPercent: 5.0},
			{PR: "3", CPUPercent: 3.0},
			{PR: "4", CPUPercent: 10.0},
			{PR: "5", CPUPercent: 2.0},
		}

		top := cost.TopNPRs(prs, 3)
		Expect(top).To(HaveLen(3))
		Expect(top[0].PR).To(Equal("4"))
		Expect(top[1].PR).To(Equal("2"))
		Expect(top[2].PR).To(Equal("3"))
	})

	It("should return all PRs when N exceeds length", func() {
		prs := []cost.PRUsage{
			{PR: "1", CPUPercent: 1.0},
			{PR: "2", CPUPercent: 2.0},
		}
		top := cost.TopNPRs(prs, 10)
		Expect(top).To(HaveLen(2))
	})

	It("should not mutate original slice", func() {
		prs := []cost.PRUsage{
			{PR: "1", CPUPercent: 1.0},
			{PR: "2", CPUPercent: 5.0},
			{PR: "3", CPUPercent: 3.0},
		}
		cost.TopNPRs(prs, 2)
		Expect(prs[0].PR).To(Equal("1"))
	})
})

var _ = Describe("MapJobNameToSIG", func() {
	table.DescribeTable("should map job names to correct SIG",
		func(jobName, expectedSIG string) {
			Expect(sigretests.MapJobNameToSIG(jobName)).To(Equal(expectedSIG))
		},
		table.Entry("sig-compute job", "pull-kubevirt-e2e-sig-compute-k8s-1.35", "compute"),
		table.Entry("vgpu job maps to compute", "pull-kubevirt-e2e-vgpu", "compute"),
		table.Entry("sig-network job", "pull-kubevirt-e2e-sig-network-k8s-1.35", "network"),
		table.Entry("sriov job maps to network", "pull-kubevirt-e2e-sriov", "network"),
		table.Entry("sig-storage job", "pull-kubevirt-e2e-sig-storage", "storage"),
		table.Entry("sig-operator job", "pull-kubevirt-e2e-sig-operator", "operator"),
		table.Entry("sig-monitoring job", "pull-kubevirt-e2e-sig-monitoring", "monitoring"),
		table.Entry("sig-performance job", "pull-kubevirt-e2e-sig-performance", "performance"),
		table.Entry("performance periodic", "periodic-kubevirt-performance-cluster-scale", "performance"),
		table.Entry("unmatched job", "pull-kubevirt-unit-test", ""),
		table.Entry("goveralls job", "pull-kubevirt-goveralls", ""),
	)
})

var _ = Describe("AggregateSIGUsages", func() {
	mapJobToSIG := func(jobName string) string {
		if sig := sigretests.MapJobNameToSIG(jobName); sig != "" {
			return sig
		}
		return "non-sig"
	}

	It("should group jobs by SIG", func() {
		jobs := []cost.JobUsage{
			{Job: "pull-kubevirt-e2e-sig-compute", CPUPercent: 2.0, MemPercent: 1.0},
			{Job: "pull-kubevirt-e2e-sig-compute-vgpu", CPUPercent: 1.0, MemPercent: 0.5},
			{Job: "pull-kubevirt-e2e-sig-network", CPUPercent: 3.0, MemPercent: 2.0},
			{Job: "pull-kubevirt-unit-test", CPUPercent: 0.5, MemPercent: 0.2},
		}

		sigs := cost.AggregateSIGUsages(jobs, mapJobToSIG)
		sigMap := map[string]cost.SIGUsage{}
		for _, s := range sigs {
			sigMap[s.Name] = s
		}

		Expect(sigMap["compute"].CPUPercent).To(Equal(3.0))
		Expect(sigMap["compute"].RunCount).To(Equal(2))
		Expect(sigMap["network"].CPUPercent).To(Equal(3.0))
		Expect(sigMap["network"].RunCount).To(Equal(1))
		Expect(sigMap["non-sig"].CPUPercent).To(Equal(0.5))
	})

	It("should sort by CPU percentage descending", func() {
		jobs := []cost.JobUsage{
			{Job: "pull-kubevirt-unit-test", CPUPercent: 0.5},
			{Job: "pull-kubevirt-e2e-sig-compute", CPUPercent: 5.0},
			{Job: "pull-kubevirt-e2e-sig-network", CPUPercent: 3.0},
		}
		sigs := cost.AggregateSIGUsages(jobs, mapJobToSIG)
		Expect(sigs[0].Name).To(Equal("compute"))
		Expect(sigs[1].Name).To(Equal("network"))
	})
})

var _ = Describe("ApplyCostRates", func() {
	It("should apply monthly cost pro-rated to data window", func() {
		report := &cost.UsageReport{
			DataDays: 30,
			PRUsages: []cost.PRUsage{
				{PR: "1", CPUPercent: 10.0},
				{PR: "2", CPUPercent: 5.0},
			},
			SIGUsages: []cost.SIGUsage{
				{Name: "compute", CPUPercent: 15.0},
			},
			TopPRs: []cost.PRUsage{
				{PR: "1", CPUPercent: 10.0},
			},
		}

		cost.ApplyCostRates(report, 10000.0)

		Expect(report.TotalCost).NotTo(BeNil())
		Expect(*report.TotalCost).To(Equal(1500.0))
		Expect(*report.PRUsages[0].TotalCost).To(Equal(1000.0))
		Expect(*report.PRUsages[1].TotalCost).To(Equal(500.0))
		Expect(*report.SIGUsages[0].TotalCost).To(Equal(1500.0))
	})

	It("should pro-rate cost for a 7-day window", func() {
		report := &cost.UsageReport{
			DataDays: 7,
			PRUsages: []cost.PRUsage{
				{PR: "1", CPUPercent: 10.0},
			},
			SIGUsages: []cost.SIGUsage{},
			TopPRs:    []cost.PRUsage{},
		}

		cost.ApplyCostRates(report, 30000.0)

		// 30000 × 7/30 = 7000 window cost; 10% of 7000 = 700
		Expect(*report.PRUsages[0].TotalCost).To(Equal(700.0))
	})

	It("should recompute TopPRs with cost fields populated", func() {
		report := &cost.UsageReport{
			DataDays: 30,
			PRUsages: []cost.PRUsage{
				{PR: "1", CPUPercent: 10.0},
				{PR: "2", CPUPercent: 5.0},
			},
			TopPRs: []cost.PRUsage{},
		}

		cost.ApplyCostRates(report, 10000.0)

		Expect(report.TopPRs).To(HaveLen(2))
		Expect(report.TopPRs[0].PR).To(Equal("1"))
		Expect(report.TopPRs[0].TotalCost).NotTo(BeNil())
		Expect(*report.TopPRs[0].TotalCost).To(Equal(1000.0))
	})

	It("should not set cost fields when monthly cost is zero", func() {
		report := &cost.UsageReport{
			PRUsages: []cost.PRUsage{
				{PR: "1", CPUPercent: 10.0},
			},
		}

		cost.ApplyCostRates(report, 0)

		Expect(report.TotalCost).To(BeNil())
		Expect(report.PRUsages[0].TotalCost).To(BeNil())
	})
})

var _ = Describe("BuildReport", func() {
	It("should build a complete report from raw metrics", func() {
		cluster := cost.ClusterCapacity{
			CPUCores:    100,
			MemoryBytes: 100 * 1024 * 1024 * 1024,
			NodeCount:   4,
		}
		jobs := []cost.RawJobMetrics{
			{PR: "1", Job: "pull-kubevirt-e2e-sig-compute", Repo: "kubevirt", Org: "kubevirt", BuildID: "100", CPUSec: 3600, MemBytes: 1024 * 1024 * 1024},
			{PR: "1", Job: "pull-kubevirt-unit-test", Repo: "kubevirt", Org: "kubevirt", BuildID: "101", CPUSec: 1800, MemBytes: 512 * 1024 * 1024},
			{PR: "2", Job: "pull-kubevirt-e2e-sig-network", Repo: "kubevirt", Org: "kubevirt", BuildID: "102", CPUSec: 7200, MemBytes: 2 * 1024 * 1024 * 1024},
		}

		mapSIG := func(jobName string) string {
			if sig := sigretests.MapJobNameToSIG(jobName); sig != "" {
				return sig
			}
			return "non-sig"
		}
		endTime := time.Date(2026, 6, 24, 12, 0, 0, 0, time.UTC)
		report := cost.BuildReport(jobs, cluster, 1, "kubevirt/kubevirt", endTime, mapSIG)

		Expect(report.PRCount).To(Equal(2))
		Expect(report.RunCount).To(Equal(3))
		Expect(report.DataDays).To(Equal(1))
		Expect(report.Source).To(Equal("kubevirt/kubevirt"))
		Expect(report.Cluster.CPUCores).To(Equal(100.0))
		Expect(report.PRUsages).To(HaveLen(2))
		Expect(report.SIGUsages).NotTo(BeEmpty())
		Expect(report.TopPRs).NotTo(BeEmpty())
		Expect(report.JobTypeUsage).NotTo(BeEmpty())
		Expect(report.TotalCPUPercent).To(BeNumerically(">", 0))
		Expect(report.AvgCPUPerPR).To(BeNumerically(">", 0))
	})
})

var _ = Describe("GenerateHTMLReport", func() {
	It("should render an HTML report without error", func() {
		report := &cost.UsageReport{
			StartDate:       time.Date(2026, 6, 17, 0, 0, 0, 0, time.UTC),
			EndDate:         time.Date(2026, 6, 24, 0, 0, 0, 0, time.UTC),
			DataDays:        7,
			Source:          "kubevirt/kubevirt",
			Cluster:         cost.ClusterCapacity{CPUCores: 100, MemoryBytes: 100 * 1024 * 1024 * 1024, NodeCount: 4},
			TotalCPUPercent: 42.5,
			TotalMemPercent: 30.0,
			AvgCPUPerPR:     2.5,
			AvgMemPerPR:     1.5,
			PRCount:         17,
			RunCount:        85,
			PRUsages: []cost.PRUsage{
				{PR: "100", Repo: "kubevirt", Org: "kubevirt", CPUPercent: 5.0, MemPercent: 3.0, RunCount: 4},
			},
			SIGUsages: []cost.SIGUsage{
				{Name: "compute", CPUPercent: 20.0, MemPercent: 15.0, RunCount: 40},
			},
			TopPRs: []cost.PRUsage{
				{PR: "100", Repo: "kubevirt", Org: "kubevirt", CPUPercent: 5.0, MemPercent: 3.0, RunCount: 4},
			},
		}

		tmpDir, err := os.MkdirTemp("", "cost-report-test-*")
		Expect(err).NotTo(HaveOccurred())
		defer func() { Expect(os.RemoveAll(tmpDir)).To(Succeed()) }()

		err = cost.GenerateHTMLReport(report, tmpDir)
		Expect(err).NotTo(HaveOccurred())

		reportPath := filepath.Join(tmpDir, "cost-report.html")
		info, err := os.Stat(reportPath)
		Expect(err).NotTo(HaveOccurred())
		Expect(info.Size()).To(BeNumerically(">", 0))
	})

	It("should render cost columns when TotalCost is set", func() {
		prCost := 500.0
		sigCost := 1500.0
		totalCost := 1500.0
		report := &cost.UsageReport{
			StartDate:       time.Date(2026, 6, 17, 0, 0, 0, 0, time.UTC),
			EndDate:         time.Date(2026, 6, 24, 0, 0, 0, 0, time.UTC),
			DataDays:        7,
			Source:          "kubevirt/kubevirt",
			Cluster:         cost.ClusterCapacity{CPUCores: 100, MemoryBytes: 100 * 1024 * 1024 * 1024, NodeCount: 4},
			TotalCPUPercent: 42.5,
			TotalMemPercent: 30.0,
			AvgCPUPerPR:     2.5,
			AvgMemPerPR:     1.5,
			PRCount:         1,
			RunCount:        4,
			TotalCost:       &totalCost,
			PRUsages: []cost.PRUsage{
				{PR: "100", Repo: "kubevirt", Org: "kubevirt", CPUPercent: 5.0, MemPercent: 3.0, RunCount: 4, TotalCost: &prCost},
			},
			SIGUsages: []cost.SIGUsage{
				{Name: "compute", CPUPercent: 20.0, MemPercent: 15.0, RunCount: 40, TotalCost: &sigCost},
			},
			TopPRs: []cost.PRUsage{
				{PR: "100", Repo: "kubevirt", Org: "kubevirt", CPUPercent: 5.0, MemPercent: 3.0, RunCount: 4, TotalCost: &prCost},
			},
		}

		tmpDir, err := os.MkdirTemp("", "cost-report-cost-test-*")
		Expect(err).NotTo(HaveOccurred())
		defer func() { Expect(os.RemoveAll(tmpDir)).To(Succeed()) }()

		err = cost.GenerateHTMLReport(report, tmpDir)
		Expect(err).NotTo(HaveOccurred())

		reportPath := filepath.Join(tmpDir, "cost-report.html")
		data, err := os.ReadFile(reportPath)
		Expect(err).NotTo(HaveOccurred())
		html := string(data)

		Expect(html).To(ContainSubstring("$1500.00"))
		Expect(html).To(ContainSubstring("$500.00"))
		Expect(html).To(ContainSubstring("Est. Cost"))
	})
})
