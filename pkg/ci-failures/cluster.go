package cifailures

import (
	"github.com/hyperjumptech/beda"
	"sort"
)

// ClusterErrors groups similar errors together.
// It takes a slice of errors and a similarity threshold (0.0 to 1.0).
func ClusterErrors(errors []*JobBuildError, threshold float64) []Cluster {
	var clusters []Cluster

	for _, jobBuildError := range errors {
		foundCluster := false
		for i, cluster := range clusters {
			if len(cluster.Representative.BuildLogErrorSnippets) == 0 {
				continue
			}
			repErrTxt := cluster.Representative.BuildLogErrorSnippets[0].ErrorText
			if len(jobBuildError.BuildLogErrorSnippets) == 0 {
				continue
			}
			errorText := jobBuildError.BuildLogErrorSnippets[0].ErrorText
			// Using Jaro-Winkler as it's good for short text like error messages.
			similarity := float64(beda.JaroWinklerDistance(repErrTxt, errorText, 0.1))

			if similarity >= threshold {
				clusters[i].Errors = append(clusters[i].Errors, jobBuildError)
				foundCluster = true
				break
			}
		}

		if !foundCluster {
			clusters = append(clusters, Cluster{
				Representative: jobBuildError,
				Errors:         []*JobBuildError{jobBuildError},
			})
		}
	}

	sort.Slice(clusters, func(i, j int) bool {
		return len(clusters[i].Errors) > len(clusters[j].Errors)
	})

	return clusters
}
