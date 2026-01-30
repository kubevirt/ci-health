package main

import (
	cifailures "github.com/kubevirt/ci-health/pkg/ci-failures"
	"sort"
)

type ShareCategory struct {
	CSSClassName       string
	MinPercentageValue float64
}

var shareCategories = []ShareCategory{
	{
		"darkred",
		50.0,
	},
	{
		"red",
		25.0,
	},
	{
		"orangered",
		10.0,
	},
	{
		"orange",
		5.0,
	},
	{
		"yellow",
		2.5,
	},
	{
		"lightyellow",
		0.25,
	},
}

type CategorizedFailure struct {
	CategoryValue  string
	Share          float64
	JobBuildErrors []*cifailures.JobBuildError
	Clusters       []cifailures.Cluster
	ShareCategory
}

type Failures struct {
	CategoryName        string
	Total               int
	CategorizedFailures map[string][]*cifailures.JobBuildError
}

func (f *Failures) GetSortedFailures() []*CategorizedFailure {
	var result []*CategorizedFailure
	for categoryValue, jobBuildErrors := range f.CategorizedFailures {
		share := float64(len(jobBuildErrors)) / float64(f.Total) * float64(100)
		clusters := cifailures.ClusterErrors(jobBuildErrors, 0.9)
		categorizedFailure := &CategorizedFailure{
			CategoryValue:  categoryValue,
			JobBuildErrors: jobBuildErrors,
			Clusters:       clusters,
			Share:          share,
		}
		for _, shareCategory := range shareCategories {
			if shareCategory.MinPercentageValue <= share {
				categorizedFailure.ShareCategory = shareCategory
				break
			}
		}
		result = append(result, categorizedFailure)
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].Share > result[j].Share
	})
	return result
}

func (f *Failures) AddAll(category string, errors *cifailures.JobBuildErrors) {
	f.Total += len(errors.BuildErrors)
	f.initCategorizedFailuresForCategory(category)
	f.CategorizedFailures[category] = append(f.CategorizedFailures[category], errors.BuildErrors...)
}

func (f *Failures) AddError(category string, error *cifailures.JobBuildError) {
	f.Total += 1
	f.initCategorizedFailuresForCategory(category)
	f.CategorizedFailures[category] = append(f.CategorizedFailures[category], error)
}

func (f *Failures) initCategorizedFailuresForCategory(category string) {
	if f.CategorizedFailures == nil {
		f.CategorizedFailures = map[string][]*cifailures.JobBuildError{}
	}
	if _, exists := f.CategorizedFailures[category]; !exists {
		f.CategorizedFailures[category] = []*cifailures.JobBuildError{}
	}
}

type Category struct {
	Value string
}

type TemplateData struct {
	Date              string
	Org               string
	Repo              string
	FailuresPerBranch Failures
	FailuresPerDay    Failures
	FailuresPerSIG    Failures
}
