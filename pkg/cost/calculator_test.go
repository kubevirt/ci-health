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

var _ = Describe("CalculateJobUsage", func() {
	var cluster cost.ClusterCapacity

	BeforeEach(func() {
		cluster = cost.ClusterCapacity{
			CPUCores:    100,
			MemoryBytes: 100 * 1024 * 1024 * 1024, // 100 GiB
			NodeCount:   4,
		}
	})

	table.DescribeTable("should calculate correct CPU percentage",
		func(cpuSec, windowSec float64, expectedCPUPercent float64) {
			raw := cost.RawJobMetrics{
				PR: "1", Job: "test-job", Repo: "kubevirt", Org: "kubevirt",
				BuildID: "100", CPUSec: cpuSec, MemBytes: 0,
			}
			usage := cost.CalculateJobUsage(raw, cluster, windowSec)
			Expect(usage.CPUPercent).To(BeNumerically("~", expectedCPUPercent, 0.001))
		},
		table.Entry("one core-hour on 100 cores over 1 hour", 3600.0, 3600.0, 1.0),
		table.Entry("zero usage", 0.0, 3600.0, 0.0),
		table.Entry("half a core-hour on 100 cores over 1 hour", 1800.0, 3600.0, 0.5),
		table.Entry("full cluster for 1 hour", 360000.0, 3600.0, 100.0),
		table.Entry("one core-day on 100 cores over 1 day", 86400.0, 86400.0, 1.0),
	)

	table.DescribeTable("should calculate correct memory percentage",
		func(memBytes, expectedMemPercent float64) {
			raw := cost.RawJobMetrics{
				PR: "1", Job: "test-job", Repo: "kubevirt", Org: "kubevirt",
				BuildID: "100", CPUSec: 0, MemBytes: memBytes,
			}
			usage := cost.CalculateJobUsage(raw, cluster, 3600)
			Expect(usage.MemPercent).To(BeNumerically("~", expectedMemPercent, 0.001))
		},
		table.Entry("1 GiB on 100 GiB cluster", 1024*1024*1024.0, 1.0),
		table.Entry("zero usage", 0.0, 0.0),
		table.Entry("50 GiB on 100 GiB cluster", 50*1024*1024*1024.0, 50.0),
		table.Entry("full cluster memory", 100*1024*1024*1024.0, 100.0),
	)

	It("should handle zero cluster capacity without panic", func() {
		emptyCluster := cost.ClusterCapacity{CPUCores: 0, MemoryBytes: 0, NodeCount: 0}
		raw := cost.RawJobMetrics{
			PR: "1", Job: "test-job", Repo: "kubevirt", Org: "kubevirt",
			BuildID: "100", CPUSec: 3600, MemBytes: 1024,
		}
		usage := cost.CalculateJobUsage(raw, emptyCluster, 3600)
		Expect(usage.CPUPercent).To(Equal(0.0))
		Expect(usage.MemPercent).To(Equal(0.0))
	})

	It("should preserve raw metrics in the output", func() {
		raw := cost.RawJobMetrics{
			PR: "42", Job: "pull-kubevirt-e2e", Repo: "kubevirt", Org: "kubevirt",
			BuildID: "999", CPUSec: 7200, MemBytes: 4 * 1024 * 1024 * 1024,
		}
		usage := cost.CalculateJobUsage(raw, cluster, 3600)
		Expect(usage.PR).To(Equal("42"))
		Expect(usage.Job).To(Equal("pull-kubevirt-e2e"))
		Expect(usage.BuildID).To(Equal("999"))
		Expect(usage.CPUSec).To(Equal(7200.0))
		Expect(usage.MemBytes).To(Equal(4 * 1024 * 1024 * 1024.0))
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
		table.Entry("unmatched job", "pull-kubevirt-unit-test", ""),
		table.Entry("goveralls job", "pull-kubevirt-goveralls", ""),
	)
})

var _ = Describe("AggregateSIGUsages", func() {
	mapJobToSIG := func(jobName string) string {
		if sig := sigretests.MapJobNameToSIG(jobName); sig != "" {
			return sig
		}
		return "other"
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
		Expect(sigMap["other"].CPUPercent).To(Equal(0.5))
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
			return "other"
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
		defer os.RemoveAll(tmpDir)

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
		defer os.RemoveAll(tmpDir)

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
