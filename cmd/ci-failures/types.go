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
	Anchor              string
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
		return result[i].Clusters[0].Representative.Started.After(result[j].Clusters[0].Representative.Started)
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

// ReasonGroup holds errors grouped by category_reason within a category.
type ReasonGroup struct {
	Reason         string
	Count          int
	Share          float64
	JobBuildErrors []*cifailures.JobBuildError
	Clusters       []cifailures.Cluster
}

// CategoryGroup holds a category with its reason sub-groups.
type CategoryGroup struct {
	Category     string
	Count        int
	Share        float64
	ReasonGroups []*ReasonGroup
}

// groupByReason groups a slice of errors by category, then by reason within each category.
func groupByReason(errors []*cifailures.JobBuildError, total int) []*CategoryGroup {
	grouped := map[string]map[string][]*cifailures.JobBuildError{}
	for _, e := range errors {
		cat := e.Category
		reason := e.CategoryReason
		if grouped[cat] == nil {
			grouped[cat] = map[string][]*cifailures.JobBuildError{}
		}
		grouped[cat][reason] = append(grouped[cat][reason], e)
	}

	var result []*CategoryGroup
	for cat, reasons := range grouped {
		catCount := 0
		var reasonGroups []*ReasonGroup
		for reason, errs := range reasons {
			catCount += len(errs)
			clusters := cifailures.ClusterErrors(errs, 0.9)
			reasonGroups = append(reasonGroups, &ReasonGroup{
				Reason:         reason,
				Count:          len(errs),
				Share:          float64(len(errs)) / float64(total) * 100,
				JobBuildErrors: errs,
				Clusters:       clusters,
			})
		}
		sort.Slice(reasonGroups, func(i, j int) bool {
			return reasonGroups[i].Count > reasonGroups[j].Count
		})
		result = append(result, &CategoryGroup{
			Category:     cat,
			Count:        catCount,
			Share:        float64(catCount) / float64(total) * 100,
			ReasonGroups: reasonGroups,
		})
	}
	sort.Slice(result, func(i, j int) bool {
		return len(result[i].ReasonGroups) > len(result[j].ReasonGroups)
	})
	return result
}

// GetGroupedByReason returns failures grouped by category, then by category_reason within each category.
func (f *Failures) GetGroupedByReason() []*CategoryGroup {
	var all []*cifailures.JobBuildError
	for _, errors := range f.CategorizedFailures {
		all = append(all, errors...)
	}
	return groupByReason(all, f.Total)
}

// GetGroupedByReason returns errors within this categorized failure grouped by category, then by reason.
func (cf *CategorizedFailure) GetGroupedByReason() []*CategoryGroup {
	return groupByReason(cf.JobBuildErrors, len(cf.JobBuildErrors))
}

type Category struct {
	Value string
}

type TemplateData struct {
	Date                string
	Org                 string
	Repo                string
	FailuresPerBranch   Failures
	FailuresPerDay      Failures
	FailuresPerSIG      Failures
	FailuresPerCategory Failures
}
