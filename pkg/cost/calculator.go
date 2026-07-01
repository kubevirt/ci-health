package cost

import (
	"sort"
	"time"
)

// CalculateJobUsage converts raw metrics into a JobUsage with percentage of cluster capacity.
func CalculateJobUsage(raw RawJobMetrics, cluster ClusterCapacity, windowSeconds float64) JobUsage {
	var cpuPercent float64
	if cluster.CPUCores > 0 && windowSeconds > 0 {
		cpuPercent = (raw.CPUSec / windowSeconds) / cluster.CPUCores * 100
	}

	var memPercent float64
	if cluster.MemoryBytes > 0 {
		memPercent = raw.MemBytes / cluster.MemoryBytes * 100
	}

	return JobUsage{
		PR:         raw.PR,
		Job:        raw.Job,
		Repo:       raw.Repo,
		Org:        raw.Org,
		BuildID:    raw.BuildID,
		CPUSec:     raw.CPUSec,
		MemBytes:   raw.MemBytes,
		CPUPercent: cpuPercent,
		MemPercent: memPercent,
	}
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
// If monthlyCost > 0, each PR's dollar cost = its CPU% × monthlyCost / 100.
func ApplyCostRates(report *UsageReport, monthlyCost float64) {
	if monthlyCost <= 0 {
		return
	}

	totalCost := 0.0
	for i := range report.PRUsages {
		cost := report.PRUsages[i].CPUPercent * monthlyCost / 100
		report.PRUsages[i].TotalCost = &cost
		totalCost += cost
	}

	for i := range report.SIGUsages {
		cost := report.SIGUsages[i].CPUPercent * monthlyCost / 100
		report.SIGUsages[i].TotalCost = &cost
	}

	for i := range report.TopPRs {
		cost := report.TopPRs[i].CPUPercent * monthlyCost / 100
		report.TopPRs[i].TotalCost = &cost
	}

	report.TotalCost = &totalCost
}

// BuildReport constructs a full UsageReport from raw job metrics and cluster capacity.
// endTime is the report end timestamp; the start is computed as endTime minus dataDays.
func BuildReport(jobs []RawJobMetrics, cluster ClusterCapacity, dataDays int, source string, endTime time.Time, mapJobToSIG func(string) string) *UsageReport {
	windowSeconds := float64(dataDays) * 24 * 3600

	jobUsages := make([]JobUsage, 0, len(jobs))
	for _, raw := range jobs {
		jobUsages = append(jobUsages, CalculateJobUsage(raw, cluster, windowSeconds))
	}

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
