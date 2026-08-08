package cost

import (
	"sort"
	"time"
)

// CalculateJobUsages converts raw metrics into JobUsages with percentages of total CI usage.
// Percentages represent each job's share of all CI resource consumption, not cluster capacity.
func CalculateJobUsages(jobs []RawJobMetrics) []JobUsage {
	var totalCPU, totalMem float64
	for _, raw := range jobs {
		totalCPU += raw.CPUSec
		totalMem += raw.MemBytes
	}

	usages := make([]JobUsage, 0, len(jobs))
	for _, raw := range jobs {
		var cpuPercent, memPercent float64
		if totalCPU > 0 {
			cpuPercent = raw.CPUSec / totalCPU * 100
		}
		if totalMem > 0 {
			memPercent = raw.MemBytes / totalMem * 100
		}
		usages = append(usages, JobUsage{
			PR:         raw.PR,
			Job:        raw.Job,
			Repo:       raw.Repo,
			Org:        raw.Org,
			BuildID:    raw.BuildID,
			CPUSec:     raw.CPUSec,
			MemBytes:   raw.MemBytes,
			CPUPercent: cpuPercent,
			MemPercent: memPercent,
		})
	}
	return usages
}

// AggregatePRUsages groups job usages by PR and computes totals.
func AggregatePRUsages(jobs []JobUsage) []PRUsage {
	prMap := map[string]*PRUsage{}
	for _, j := range jobs {
		key := j.Org + "/" + j.Repo + "#" + j.PR
		pr, ok := prMap[key]
		if !ok {
			pr = &PRUsage{
				PR:   j.PR,
				Repo: j.Repo,
				Org:  j.Org,
			}
			prMap[key] = pr
		}
		pr.Jobs = append(pr.Jobs, j)
		pr.CPUPercent += j.CPUPercent
		pr.MemPercent += j.MemPercent
		pr.RunCount++
	}

	result := make([]PRUsage, 0, len(prMap))
	for _, pr := range prMap {
		result = append(result, *pr)
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].CPUPercent > result[j].CPUPercent
	})
	return result
}

// AggregateSIGUsages groups job usages by SIG using a job-to-SIG mapping function.
func AggregateSIGUsages(jobs []JobUsage, mapJobToSIG func(string) string) []SIGUsage {
	sigMap := map[string]*SIGUsage{}
	for _, j := range jobs {
		sig := mapJobToSIG(j.Job)
		if sig == "" {
			sig = "other"
		}
		s, ok := sigMap[sig]
		if !ok {
			s = &SIGUsage{Name: sig}
			sigMap[sig] = s
		}
		s.CPUPercent += j.CPUPercent
		s.MemPercent += j.MemPercent
		s.RunCount++
	}

	result := make([]SIGUsage, 0, len(sigMap))
	for _, s := range sigMap {
		result = append(result, *s)
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].CPUPercent > result[j].CPUPercent
	})
	return result
}

// AggregateJobTypeUsages groups usage percentages by job name.
func AggregateJobTypeUsages(jobs []JobUsage) map[string]float64 {
	result := map[string]float64{}
	for _, j := range jobs {
		result[j.Job] += j.CPUPercent
	}
	return result
}

// TopNPRs returns the N most resource-intensive PRs sorted by CPU percentage descending.
func TopNPRs(prs []PRUsage, n int) []PRUsage {
	sorted := make([]PRUsage, len(prs))
	copy(sorted, prs)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].CPUPercent > sorted[j].CPUPercent
	})
	if n > len(sorted) {
		n = len(sorted)
	}
	return sorted[:n]
}

// ApplyCostRates applies optional dollar cost to all usage structs.
// monthlyCost is pro-rated to the report's DataDays window:
// windowCost = monthlyCost × dataDays / 30, then each PR's cost = CPU% × windowCost / 100.
// TopPRs is recomputed from PRUsages so the cost fields are always consistent.
func ApplyCostRates(report *UsageReport, monthlyCost float64) {
	if monthlyCost <= 0 {
		return
	}

	windowCost := monthlyCost * float64(report.DataDays) / 30.0

	totalCost := 0.0
	for i := range report.PRUsages {
		cost := report.PRUsages[i].CPUPercent * windowCost / 100
		report.PRUsages[i].TotalCost = &cost
		totalCost += cost
	}

	for i := range report.SIGUsages {
		cost := report.SIGUsages[i].CPUPercent * windowCost / 100
		report.SIGUsages[i].TotalCost = &cost
	}

	report.TopPRs = TopNPRs(report.PRUsages, 10)
	report.TotalCost = &totalCost
}

// BuildReport constructs a full UsageReport from raw job metrics and cluster capacity.
// endTime is the report end timestamp; the start is computed as endTime minus dataDays.
// Percentages represent each PR's share of total CI resource consumption.
func BuildReport(jobs []RawJobMetrics, cluster ClusterCapacity, dataDays int, source string, endTime time.Time, mapJobToSIG func(string) string) *UsageReport {
	jobUsages := CalculateJobUsages(jobs)

	prUsages := AggregatePRUsages(jobUsages)
	sigUsages := AggregateSIGUsages(jobUsages, mapJobToSIG)
	jobTypeUsage := AggregateJobTypeUsages(jobUsages)
	topPRs := TopNPRs(prUsages, 10)

	var totalCPU, totalMem float64
	for _, pr := range prUsages {
		totalCPU += pr.CPUPercent
		totalMem += pr.MemPercent
	}

	var avgCPU, avgMem float64
	if len(prUsages) > 0 {
		avgCPU = totalCPU / float64(len(prUsages))
		avgMem = totalMem / float64(len(prUsages))
	}

	return &UsageReport{
		StartDate:       endTime.AddDate(0, 0, -dataDays),
		EndDate:         endTime,
		DataDays:        dataDays,
		Source:          source,
		Cluster:         cluster,
		TotalCPUPercent: totalCPU,
		TotalMemPercent: totalMem,
		AvgCPUPerPR:     avgCPU,
		AvgMemPerPR:     avgMem,
		PRCount:         len(prUsages),
		RunCount:        len(jobUsages),
		PRUsages:        prUsages,
		SIGUsages:       sigUsages,
		TopPRs:          topPRs,
		JobTypeUsage:    jobTypeUsage,
	}
}
