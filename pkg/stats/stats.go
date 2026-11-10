package stats

import (
	"cmp"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/kubevirt/ci-health/pkg/chatops"
	cifailures "github.com/kubevirt/ci-health/pkg/ci-failures"
	"github.com/kubevirt/ci-health/pkg/constants"
	"github.com/kubevirt/ci-health/pkg/mergequeue"
	"github.com/kubevirt/ci-health/pkg/quarantine"
	"github.com/kubevirt/ci-health/pkg/sigretests"
	"github.com/kubevirt/ci-health/pkg/types"
)

// analyzeBuildFn is the AnalyzeBuild implementation used to classify sig-ci
// failures. Tests replace it to avoid network calls.
var analyzeBuildFn = cifailures.AnalyzeBuild

type statsProcessor func(*types.Results) (*types.Results, error)

type HandlerOptions struct {
	Mq                *mergequeue.Handler
	Co                *chatops.Handler
	Source            string
	EndDate           time.Time
	DataDays          int
	SupportedBranches []string

	TargetMetrics []types.Metric
}

type Handler struct {
	mq                *mergequeue.Handler
	co                *chatops.Handler
	source            string
	endDate           time.Time
	dataDays          int
	supportedBranches []string
	targetMetrics     []types.Metric
}

func NewHandler(ho *HandlerOptions) *Handler {
	return &Handler{
		mq:                ho.Mq,
		co:                ho.Co,
		source:            ho.Source,
		endDate:           ho.EndDate,
		dataDays:          ho.DataDays,
		supportedBranches: ho.SupportedBranches,
		targetMetrics:     ho.TargetMetrics,
	}
}

func (h *Handler) Run() (*types.Results, error) {
	results := &types.Results{
		EndDate:  h.endDate.Format(constants.DateFormat),
		DataDays: h.dataDays,
		Source:   h.source,
		Data:     map[string]types.RunningAverageDataItem{},
	}
	var err error

	for _, targetMetric := range h.targetMetrics {
		var processor statsProcessor

		switch targetMetric {
		case types.MergeQueueLengthMetric:
			processor = h.mergeQueueProcessor
		case types.TimeToMergeMetric:
			processor = h.timeToMergeProcessor
		case types.RetestsToMergeMetric:
			processor = h.retestsToMergeProcessor
		case types.MergedPRsMetric:
			processor = h.mergedPRsProcessor
		case types.MergedPRsNoRetestMetric:
			processor = h.mergedPRsNoRetestProcessor
		case types.SIGRetestMetric:
			processor = h.sigRetestsProcessor
		case types.QuarantineMetric:
			processor = h.quarantineProcessor
		default:
			return nil, fmt.Errorf("unknown metric: %q", targetMetric)
		}

		results, err = processor(results)
		if err != nil {
			return nil, err
		}
	}
	return results, nil
}

func (h *Handler) mergeQueueProcessor(results *types.Results) (*types.Results, error) {
	currentTime, err := time.Parse(constants.DateFormat, results.EndDate)
	if err != nil {
		return results, err
	}

	dataItem := types.RunningAverageDataItem{
		DataPoints: []types.DataPoint{},
	}

	values := []float64{}
	for i := 0; i < results.DataDays; i++ {
		queryDate := currentTime.AddDate(0, 0, -1*i)
		queueLength, prs, err := h.mq.LengthAt(queryDate)
		if err != nil {
			return nil, err
		}
		values = append(values, float64(queueLength))
		dataItem.DataPoints = append(dataItem.DataPoints,
			types.DataPoint{
				Value: float64(queueLength),
				PRs:   prs,
				Date:  &queryDate,
			})
	}

	dataItem.Avg = Average(values)
	dataItem.Std = Std(values)

	results.Data[constants.MergeQueueLengthName] = dataItem

	return results, nil
}

func (h *Handler) timeToMergeProcessor(results *types.Results) (*types.Results, error) {
	currentTime, err := time.Parse(constants.DateFormat, results.EndDate)
	if err != nil {
		return results, err
	}

	dataItem := types.RunningAverageDataItem{
		DataPoints: []types.DataPoint{},
	}

	timesToMerge, err := h.mq.TimesToMerge(currentTime.AddDate(0, 0, -1*results.DataDays), currentTime)
	if err != nil {
		return nil, err
	}

	values := []float64{}

	for pr, timeToMerge := range timesToMerge {
		days := timeToMerge.Hours() / 24
		value := round(days)

		values = append(values, value)

		dataItem.DataPoints = append(dataItem.DataPoints,
			types.DataPoint{
				Value: value,
				PRs:   []types.PR{pr},
			})
	}

	dataItem.Avg = Average(values)
	dataItem.Std = Std(values)

	results.Data[constants.TimeToMergeName] = dataItem

	return results, nil
}

func (h *Handler) retestsToMergeProcessor(results *types.Results) (*types.Results, error) {
	currentTime, err := time.Parse(constants.DateFormat, results.EndDate)
	if err != nil {
		return results, err
	}

	dataItem := types.RunningAverageDataItem{
		DataPoints: []types.DataPoint{},
	}

	retestsToMerge, err := h.co.RetestsToMerge(currentTime.AddDate(0, 0, -1*results.DataDays), currentTime)
	if err != nil {
		return nil, err
	}

	values := []float64{}

	for pr, retestsToMerge := range retestsToMerge {
		value := float64(retestsToMerge)

		values = append(values, value)

		dataItem.DataPoints = append(dataItem.DataPoints,
			types.DataPoint{
				Value: value,
				PRs:   []types.PR{pr},
			})
	}

	dataItem.Avg = Average(values)
	dataItem.Std = Std(values)

	results.Data[constants.RetestsToMergeName] = dataItem

	return results, nil
}

func (h *Handler) mergedPRsNoRetestProcessor(results *types.Results) (*types.Results, error) {
	currentTime, err := time.Parse(constants.DateFormat, results.EndDate)
	if err != nil {
		return results, err
	}

	dataItem := types.RunningAverageDataItem{
		DataPoints: []types.DataPoint{},
	}

	retestsToMerge, err := h.co.RetestsToMerge(currentTime.AddDate(0, 0, -1*results.DataDays), currentTime)
	if err != nil {
		return nil, err
	}
	for pr, retestsToMerge := range retestsToMerge {
		if retestsToMerge == 0 {
			dataItem.DataPoints = append(dataItem.DataPoints,
				types.DataPoint{
					Value: float64(retestsToMerge),
					PRs:   []types.PR{pr},
				})
		}

	}
	dataItem.NoRetest = float64(len(dataItem.DataPoints))
	dataItem.Number = float64(len(retestsToMerge))
	results.Data[constants.MergedPRsNoRetest] = dataItem

	return results, nil
}
func (h *Handler) mergedPRsProcessor(results *types.Results) (*types.Results, error) {
	currentTime, err := time.Parse(constants.DateFormat, results.EndDate)
	if err != nil {
		return results, err
	}

	dataItem := types.RunningAverageDataItem{
		DataPoints: []types.DataPoint{},
	}

	mergedPRs, err := h.mq.MergedPRsBetween(currentTime.AddDate(0, 0, -1*results.DataDays), currentTime)
	if err != nil {
		return nil, err
	}

	for _, mergedPR := range mergedPRs {
		dataItem.DataPoints = append(dataItem.DataPoints,
			types.DataPoint{
				Value: 1,
				PRs:   []types.PR{mergedPR},
			})
	}

	dataItem.Avg = float64(len(mergedPRs)) / float64(results.DataDays)
	dataItem.Std = 0

	results.Data[constants.MergedPRsName] = dataItem

	return results, nil
}

func (h *Handler) sigRetestsProcessor(results *types.Results) (*types.Results, error) {
	currentTime, err := time.Parse(constants.DateFormat, results.EndDate)
	var failedJobNames []string
	var failedJobURLs []string
	var successJobNames []string
	if err != nil {
		return results, err
	}

	dataItem := types.RunningAverageDataItem{
		DataPoints: []types.DataPoint{},
	}

	startTime := currentTime.AddDate(0, 0, -1*results.DataDays)
	mergedPRs, err := h.mq.MergedPRsBetween(startTime, currentTime)
	if err != nil {
		return results, err
	}

	prFailures := make(map[int][]types.PRJobFailure, len(mergedPRs))
	allRetests := map[int]int{}
	if h.co != nil {
		if counts, err := h.co.AllRetestsBetween(startTime, currentTime); err != nil {
			log.WithError(err).Warn("failed to count all /retest comments; PR retest report will omit merged PRs")
		} else {
			allRetests = counts
		}
	}
	e2e := newE2EClassifier(results.DataDays)
	for _, mergedPR := range mergedPRs {
		org, repo := h.sourceOrgRepo()
		jobsPerSIG, err := sigretests.GetJobsPerSIG(strconv.Itoa(mergedPR.Number), org, repo, h.supportedBranches, startTime)
		if err != nil {
			return results, err
		}
		dataItem.SIGComputeRetest = dataItem.SIGComputeRetest + float64(jobsPerSIG.SigComputeFailure)
		dataItem.SIGNetworkRetest = dataItem.SIGNetworkRetest + float64(jobsPerSIG.SigNetworkFailure)
		dataItem.SIGStorageRetest = dataItem.SIGStorageRetest + float64(jobsPerSIG.SigStorageFailure)
		dataItem.SIGOperatorRetest = dataItem.SIGOperatorRetest + float64(jobsPerSIG.SigOperatorFailure)
		ciCauses := classifyCIFailures(jobsPerSIG.SigCIFailureURLs)
		dataItem.SIGCIRetest = dataItem.SIGCIRetest + float64(jobsPerSIG.SigCIFailure)
		dataItem.SIGCIExternalRetest = dataItem.SIGCIExternalRetest + float64(countExternalCauses(ciCauses))
		dataItem.SIGMonitoringRetest = dataItem.SIGMonitoringRetest + float64(jobsPerSIG.SigMonitoringFailure)
		dataItem.SIGComputeTotal = dataItem.SIGComputeTotal + float64(jobsPerSIG.SigComputeFailure) + float64(jobsPerSIG.SigComputeSuccess)
		dataItem.SIGNetworkTotal = dataItem.SIGNetworkTotal + float64(jobsPerSIG.SigNetworkFailure) + float64(jobsPerSIG.SigNetworkSuccess)
		dataItem.SIGStorageTotal = dataItem.SIGStorageTotal + float64(jobsPerSIG.SigStorageFailure) + float64(jobsPerSIG.SigStorageSuccess)
		dataItem.SIGOperatorTotal = dataItem.SIGOperatorTotal + float64(jobsPerSIG.SigOperatorFailure) + float64(jobsPerSIG.SigOperatorSuccess)
		dataItem.SIGMonitoringTotal = dataItem.SIGMonitoringTotal + float64(jobsPerSIG.SigMonitoringFailure) + float64(jobsPerSIG.SigMonitoringSuccess)
		dataItem.DataPoints = append(dataItem.DataPoints,
			types.DataPoint{
				Value: float64(len(jobsPerSIG.FailedJobNames)),
				PRs:   []types.PR{mergedPR},
			})
		failedJobNames = slices.Concat(failedJobNames, jobsPerSIG.FailedJobNames)
		successJobNames = slices.Concat(successJobNames, jobsPerSIG.SuccessJobNames)
		failedJobURLs = slices.Concat(failedJobURLs, jobsPerSIG.FailedJobURLs)
		if allRetests[mergedPR.Number] > 0 {
			prFailures[mergedPR.Number] = h.collectPRJobFailures(mergedPR.Number, startTime, e2e)
		}
	}
	dataItem.SIGCITotal = dataItem.SIGComputeTotal + dataItem.SIGStorageTotal + dataItem.SIGNetworkTotal + dataItem.SIGOperatorTotal + dataItem.SIGCIRetest + dataItem.SIGMonitoringTotal
	sortedFailedJobs := types.SortByMostFailed(countFailedJobs(failedJobNames))
	for i, job := range sortedFailedJobs {
		for _, success := range successJobNames {
			if job.JobName == success {
				sortedFailedJobs[i].SuccessCount++
			}
		}
		for _, failedJobURL := range failedJobURLs {
			jobNameInURL := strings.Split(failedJobURL, "/")[len(strings.Split(failedJobURL, "/"))-2]
			if job.JobName == jobNameInURL {
				sortedFailedJobs[i].FailureURLs = append(sortedFailedJobs[i].FailureURLs, failedJobURL)
			}
		}
		// Sort by Prow job ID descending (newest first)
		slices.SortFunc(sortedFailedJobs[i].FailureURLs, func(a, b string) int {
			return cmp.Compare(prowJobID(b), prowJobID(a))
		})
	}

	failedJobSet := make(map[string]struct{}, len(sortedFailedJobs))
	for _, job := range sortedFailedJobs {
		failedJobSet[job.JobName] = struct{}{}
	}
	successOnlyCounts := make(map[string]int)
	for _, name := range successJobNames {
		if _, hasFailed := failedJobSet[name]; !hasFailed {
			successOnlyCounts[name]++
		}
	}
	successOnlyNames := make([]string, 0, len(successOnlyCounts))
	for name := range successOnlyCounts {
		successOnlyNames = append(successOnlyNames, name)
	}
	slices.Sort(successOnlyNames)
	for _, name := range successOnlyNames {
		sortedFailedJobs = append(sortedFailedJobs, types.FailedJob{JobName: name, SuccessCount: successOnlyCounts[name]})
	}

	dataItem.FailedJobLeaderBoard = sortedFailedJobs
	results.Data[constants.SIGRetests] = dataItem
	results.MergedPRCount = len(mergedPRs)
	results.PRRetestReport = BuildPRRetestReport(mergedPRs, allRetests, prFailures)

	openSummaries, openCount, openErr := h.collectOpenPRRetestReport(startTime, e2e)
	if openErr != nil {
		log.WithError(openErr).Warn("failed to collect open PR retest report")
	} else {
		results.OpenPRRetestReport = openSummaries
		results.OpenPRCount = openCount
	}

	return results, nil
}

func (h *Handler) sourceOrgRepo() (string, string) {
	org, repo, ok := strings.Cut(h.source, "/")
	if !ok || org == "" || repo == "" {
		return "kubevirt", "kubevirt"
	}
	return org, repo
}

// collectOpenPRRetestReport lists currently open PRs updated in the stats window
// and classifies required e2e failures for those with a /retest.
// Failures here are not added to SIG badge counters.
func (h *Handler) collectOpenPRRetestReport(startTime time.Time, e2e *e2eClassifier) ([]types.PRRetestSummary, int, error) {
	if h.co == nil {
		return nil, 0, nil
	}
	openRetests, considered, err := h.co.OpenPRRetests(startTime, h.endDate)
	if err != nil {
		return nil, 0, err
	}
	prs := make([]types.PR, 0, len(openRetests))
	counts := make(map[int]int, len(openRetests))
	failures := make(map[int][]types.PRJobFailure)
	for pr, count := range openRetests {
		counts[pr.Number] = count
		if count == 0 {
			continue
		}
		prs = append(prs, pr)
		failures[pr.Number] = h.collectPRJobFailures(pr.Number, startTime, e2e)
	}
	return BuildPRRetestReport(prs, counts, failures), considered, nil
}

func (h *Handler) collectPRJobFailures(prNumber int, notBefore time.Time, e2e *e2eClassifier) []types.PRJobFailure {
	org, repo := h.sourceOrgRepo()
	details, err := sigretests.ListRequiredE2EFailures(org, repo, strconv.Itoa(prNumber), h.supportedBranches, notBefore)
	if err != nil {
		log.WithError(err).Warnf("PR %d: failed to list required e2e failures", prNumber)
		return nil
	}
	var ciURLs []string
	for _, d := range details {
		if d.SIG == "ci" {
			ciURLs = append(ciURLs, d.URL)
		}
	}
	return jobFailuresForPR(prNumber, details, classifyCIFailures(ciURLs), e2e)
}

func prowJobID(url string) int {
	parts := strings.Split(url, "/")
	id, _ := strconv.Atoi(parts[len(parts)-1])
	return id
}

func (h *Handler) quarantineProcessor(results *types.Results) (*types.Results, error) {
	currentTime, err := time.Parse(constants.DateFormat, results.EndDate)
	if err != nil {
		return results, err
	}

	dataItem := types.RunningAverageDataItem{
		DataPoints: []types.DataPoint{},
	}
	qStats, err := quarantine.GetQuarantineStats()
	if err != nil {
		return results, err
	}

	dataItem.QuarantineTotal = qStats.TotalQuarantineCount
	dataItem.QuarantineSigCompute = qStats.SigComputeQuarantine
	dataItem.QuarantineSigStorage = qStats.SigStorageQuarantine
	dataItem.QuarantineSigNetwork = qStats.SigNetworkQuarantine
	dataItem.QuarantineSigMonitoring = qStats.SigMonitoringQuarantine

	dataItem.DataPoints = append(dataItem.DataPoints,
		types.DataPoint{
			Value: qStats.TotalQuarantineCount,
			Date:  &currentTime,
		})

	dataItem.Avg = qStats.TotalQuarantineCount

	results.Data[constants.QuarantineStats] = dataItem

	return results, nil
}

// Average returns the average of the given floats.
func Average(xs []float64) float64 {
	if len(xs) == 0 {
		return 0
	}
	total := 0.0
	for _, v := range xs {
		total += v
	}
	result := total / float64(len(xs))
	return round(result)
}

// Std returns the standard deviation of the given floats.
func Std(xs []float64) float64 {
	if len(xs) == 0 {
		return 0
	}
	avg := Average(xs)
	total := 0.0
	for _, v := range xs {
		total += (v - avg) * (v - avg)
	}
	variance := total / float64(len(xs))
	result := math.Sqrt(variance)
	return round(result)
}

func countFailedJobs(jobNames []string) map[string]int {
	countFailedJobs := make(map[string]int)
	for _, name := range jobNames {
		countFailedJobs[name]++
	}
	return countFailedJobs
}

type ciFailureCause struct {
	cause  string
	reason string
}

func classifyCIFailures(urls []string) map[string]ciFailureCause {
	classified := make(map[string]ciFailureCause, len(urls))
	for _, url := range urls {
		result, err := analyzeBuildFn(url)
		if err != nil {
			log.WithError(err).Warnf("failed to analyze build %s; treating as ci", url)
		}
		cause, reason := CauseFromCIAnalysis(result, err)
		classified[url] = ciFailureCause{cause: cause, reason: reason}
	}
	return classified
}

func countExternalCauses(classified map[string]ciFailureCause) int {
	external := 0
	for _, c := range classified {
		if c.cause == types.FailureCauseExternal {
			external++
		}
	}
	return external
}

// CauseFromCIAnalysis maps an AnalyzeBuild result onto a retest cause.
// Registry / GitHub / cache failures stay external. kubeadm livez noise and
// cluster-down teardown that AnalyzeBuild tagged external are treated as ci
// so a later snippet does not hide the real cluster-up/sync failure.
// Compile/pr-build is left unlabeled. Analysis errors are ci.
func CauseFromCIAnalysis(result *cifailures.JobBuildErrors, err error) (cause, reason string) {
	if err != nil {
		return types.FailureCauseCI, "failed to analyze build; treating as ci"
	}
	if result == nil || len(result.BuildErrors) == 0 {
		return types.FailureCauseCI, "no build errors; treating as ci"
	}
	be := result.BuildErrors[0]
	reason = be.CategoryReason
	if reason == "" {
		reason = be.Category
	}
	switch be.Category {
	case string(cifailures.CategoryExternal):
		if externalNoiseIsCI(reason) {
			return types.FailureCauseCI, reason
		}
		return types.FailureCauseExternal, reason
	case string(cifailures.CategoryPRBuild):
		return "", reason
	default:
		return types.FailureCauseCI, reason
	}
}

func externalNoiseIsCI(reason string) bool {
	r := strings.ToLower(reason)
	if strings.Contains(r, "kube-apiserver body decode") {
		return true
	}
	if strings.Contains(r, "podman container removal timeout") {
		return true
	}
	return false
}

func jobFailuresForPR(prNumber int, details []sigretests.FailedJobDetail, ciCauses map[string]ciFailureCause, clf *e2eClassifier) []types.PRJobFailure {
	if len(details) == 0 {
		return nil
	}
	failures := make([]types.PRJobFailure, 0, len(details))
	for _, d := range details {
		failure := types.PRJobFailure{
			JobName: d.JobName,
			URL:     d.URL,
			SIG:     d.SIG,
		}
		if d.SIG == "ci" {
			if c, ok := ciCauses[d.URL]; ok {
				failure.Cause = c.cause
				failure.Reason = c.reason
			} else {
				failure.Cause = types.FailureCauseCI
				failure.Reason = "unclassified sig-ci failure"
			}
		} else {
			failure.Reason = "e2e test failure"
			if clf != nil {
				failure.Cause, failure.Reason = clf.classify(d.URL)
				cause := failure.Cause
				if cause == "" {
					cause = "other"
				}
				log.Infof("PR %d %s: %s (%s)", prNumber, d.JobName, cause, failure.Reason)
			}
		}
		failures = append(failures, failure)
	}
	slices.SortFunc(failures, func(a, b types.PRJobFailure) int {
		if n := cmp.Compare(a.JobName, b.JobName); n != 0 {
			return n
		}
		return cmp.Compare(a.URL, b.URL)
	})
	return failures
}

// BuildPRRetestReport joins PRs, window-scoped /retest counts, and per-job causes.
// PRs with a zero /retest count are omitted.
func BuildPRRetestReport(prs []types.PR, retestCounts map[int]int, failuresByPR map[int][]types.PRJobFailure) []types.PRRetestSummary {
	if retestCounts == nil {
		retestCounts = map[int]int{}
	}
	report := make([]types.PRRetestSummary, 0, len(prs))
	for _, pr := range prs {
		count := retestCounts[pr.Number]
		if count <= 0 {
			continue
		}
		report = append(report, types.PRRetestSummary{
			Number:      pr.Number,
			MergedAt:    pr.MergedAt,
			RetestCount: count,
			Failures:    failuresByPR[pr.Number],
		})
	}
	slices.SortFunc(report, func(a, b types.PRRetestSummary) int {
		if n := cmp.Compare(b.RetestCount, a.RetestCount); n != 0 {
			return n
		}
		return cmp.Compare(b.Number, a.Number)
	})
	return report
}

func round(value float64) float64 {
	return math.Round(value*100) / 100
}
