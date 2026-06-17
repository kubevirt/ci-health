package cost

import "time"

// ClusterCapacity represents the total allocatable resources across worker nodes.
type ClusterCapacity struct {
	CPUCores    float64 `json:"cpuCores"`
	MemoryBytes float64 `json:"memoryBytes"`
	NodeCount   int     `json:"nodeCount"`
}

// JobUsage represents resource usage for a single Prow job run.
type JobUsage struct {
	PR         string        `json:"pr"`
	Job        string        `json:"job"`
	Repo       string        `json:"repo"`
	Org        string        `json:"org"`
	BuildID    string        `json:"buildId"`
	CPUSec     float64       `json:"cpuSeconds"`
	MemBytes   float64       `json:"memBytes"`
	CPUPercent float64       `json:"cpuPercent"`
	MemPercent float64       `json:"memPercent"`
	Duration   time.Duration `json:"duration"`
	Start      time.Time     `json:"start"`
	End        time.Time     `json:"end"`
	CPUCost    *float64      `json:"cpuCost,omitempty"`
	MemCost    *float64      `json:"memCost,omitempty"`
	TotalCost  *float64      `json:"totalCost,omitempty"`
}

// PRUsage aggregates resource usage across all job runs for a single PR.
type PRUsage struct {
	PR         string      `json:"pr"`
	Repo       string      `json:"repo"`
	Org        string      `json:"org"`
	Jobs       []JobUsage  `json:"jobs"`
	CPUPercent float64     `json:"cpuPercent"`
	MemPercent float64     `json:"memPercent"`
	RunCount   int         `json:"runCount"`
	Duration   time.Duration `json:"totalDuration"`
	TotalCost  *float64    `json:"totalCost,omitempty"`
}

// SIGUsage aggregates resource usage by SIG.
type SIGUsage struct {
	Name       string  `json:"name"`
	CPUPercent float64 `json:"cpuPercent"`
	MemPercent float64 `json:"memPercent"`
	RunCount   int     `json:"runCount"`
	TotalCost  *float64 `json:"totalCost,omitempty"`
}

// UsageReport is the top-level result for a resource usage report run.
type UsageReport struct {
	StartDate     time.Time          `json:"startDate"`
	EndDate       time.Time          `json:"endDate"`
	DataDays      int                `json:"dataDays"`
	Source        string             `json:"source"`
	Cluster       ClusterCapacity    `json:"cluster"`
	TotalCPUPercent float64          `json:"totalCpuPercent"`
	TotalMemPercent float64          `json:"totalMemPercent"`
	AvgCPUPerPR   float64            `json:"avgCpuPercentPerPR"`
	AvgMemPerPR   float64            `json:"avgMemPercentPerPR"`
	PRCount       int                `json:"prCount"`
	RunCount      int                `json:"runCount"`
	PRUsages      []PRUsage          `json:"prUsages"`
	SIGUsages     []SIGUsage         `json:"sigUsages"`
	TopPRs        []PRUsage          `json:"topPRs"`
	JobTypeUsage  map[string]float64 `json:"jobTypeUsage"`
	TotalCost     *float64           `json:"totalCost,omitempty"`
}

// Rates holds optional cost rates for converting usage percentages to dollar amounts.
type Rates struct {
	CPUPerHour    float64
	MemPerGiBHour float64
	MonthlyCost   float64
}
