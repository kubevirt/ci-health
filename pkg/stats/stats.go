package stats

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/kubevirt/ci-health/pkg/chatops"
	"github.com/kubevirt/ci-health/pkg/constants"
	"github.com/kubevirt/ci-health/pkg/mergequeue"
	"github.com/kubevirt/ci-health/pkg/sigretests"
	"github.com/kubevirt/ci-health/pkg/types"
)

type statsProcessor func(*types.Results) (*types.Results, error)

type HandlerOptions struct {
	Mq       *mergequeue.Handler
	Co       *chatops.Handler
	Source   string
	EndDate  time.Time
	DataDays int

	TargetMetrics []types.Metric
}

type Handler struct {
	mq            *mergequeue.Handler
	co            *chatops.Handler
	source        string
	endDate       time.Time
	dataDays      int
	targetMetrics []types.Metric
}

func NewHandler(ho *HandlerOptions) *Handler {
	return &Handler{
		mq:            ho.Mq,
		co:            ho.Co,
		source:        ho.Source,
		endDate:       ho.EndDate,
		dataDays:      ho.DataDays,
		targetMetrics: ho.TargetMetrics,
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
		default:
			return nil, fmt.Errorf("Unknown metric: %q", targetMetric)
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

	mergedPRs, err := h.mq.MergedPRsBetween(currentTime.AddDate(0, 0, -1*results.DataDays), currentTime)
	if err != nil {
		return results, err
	}

	for _, mergedPR := range mergedPRs {
		jobsPerSIG, err := sigretests.GetJobsPerSIG(strconv.Itoa(mergedPR.Number), "kubevirt", "kubevirt")
		if err != nil {
			return results, err
		}
		dataItem.SIGComputeRetest = dataItem.SIGComputeRetest + float64(jobsPerSIG.SigComputeFailure)
		dataItem.SIGNetworkRetest = dataItem.SIGNetworkRetest + float64(jobsPerSIG.SigNetworkFailure)
		dataItem.SIGStorageRetest = dataItem.SIGStorageRetest + float64(jobsPerSIG.SigStorageFailure)
		dataItem.SIGOperatorRetest = dataItem.SIGOperatorRetest + float64(jobsPerSIG.SigOperatorFailure)
		dataItem.SIGComputeTotal = dataItem.SIGComputeTotal + float64(jobsPerSIG.SigComputeFailure) + float64(jobsPerSIG.SigComputeSuccess)
		dataItem.SIGNetworkTotal = dataItem.SIGNetworkTotal + float64(jobsPerSIG.SigNetworkFailure) + float64(jobsPerSIG.SigNetworkSuccess)
		dataItem.SIGStorageTotal = dataItem.SIGStorageTotal + float64(jobsPerSIG.SigStorageFailure) + float64(jobsPerSIG.SigStorageSuccess)
		dataItem.SIGOperatorTotal = dataItem.SIGOperatorTotal + float64(jobsPerSIG.SigOperatorFailure) + float64(jobsPerSIG.SigOperatorSuccess)
		dataItem.DataPoints = append(dataItem.DataPoints,
			types.DataPoint{
				Value: float64(len(jobsPerSIG.FailedJobNames)),
				PRs:   []types.PR{mergedPR},
			})
		failedJobNames = slices.Concat(failedJobNames, jobsPerSIG.FailedJobNames)
		successJobNames = slices.Concat(successJobNames, jobsPerSIG.SuccessJobNames)
		failedJobURLs = slices.Concat(failedJobURLs, jobsPerSIG.FailedJobURLs)
	}
	sortedFailedJobs := types.SortByMostFailed(countFailedJobs(failedJobNames))
	for i, job := range sortedFailedJobs {
		for _, success := range successJobNames {
			if job.JobName == success {
				sortedFailedJobs[i].SuccesCount++
			}
		}
		for _, failedJobURL := range failedJobURLs {
			jobNameInURL := strings.Split(failedJobURL, "/")[len(strings.Split(failedJobURL, "/"))-2]
			if job.JobName == jobNameInURL {
				sortedFailedJobs[i].FailureURLs = append(sortedFailedJobs[i].FailureURLs, failedJobURL)
			}
		}
	}
	dataItem.FailedJobLeaderBoard = sortedFailedJobs
	results.Data[constants.SIGRetests] = dataItem

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
		total += math.Pow((v - avg), 2)
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

func round(value float64) float64 {
	return math.Round(value*100) / 100
}
