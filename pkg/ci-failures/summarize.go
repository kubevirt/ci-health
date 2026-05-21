package cifailures

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

type SummaryErrorSnippet struct {
	ErrorText string `yaml:"error_text"`
	Link      string `yaml:"link"`
}

type SummaryJob struct {
	JobName        string                `yaml:"job_name"`
	JobURL         string                `yaml:"job_url"`
	BuildID        int                   `yaml:"build_id"`
	Category       string                `yaml:"category"`
	CategoryReason string                `yaml:"category_reason"`
	ErrorSnippets  []SummaryErrorSnippet `yaml:"error_snippets"`
}

type SummaryK8sFinding struct {
	Detector  string `yaml:"detector"`
	Severity  string `yaml:"severity"`
	Kind      string `yaml:"kind"`
	Namespace string `yaml:"namespace,omitempty"`
	Name      string `yaml:"name"`
	Reason    string `yaml:"reason"`
	Message   string `yaml:"message"`
}

type SummaryK8s struct {
	BuildID           int                 `yaml:"build_id"`
	JobName           string              `yaml:"job_name"`
	TotalFindings     int                 `yaml:"total_findings"`
	Findings          []SummaryK8sFinding `yaml:"findings,omitempty"`
	ContainerLogFiles []ContainerLogFile  `yaml:"container_log_files,omitempty"`
}

type SessionSummary struct {
	Jobs        []SummaryJob `yaml:"jobs"`
	K8sAnalyses []SummaryK8s `yaml:"k8s_analyses,omitempty"`
}

var excludePrefixes = []string{
	"k8s-analysis-",
	"test-rate-",
	"lane-rate-",
	"flake-overview",
	"change-relevance-",
}

func SummarizeSession(sessionDir string) (*SessionSummary, error) {
	allFiles, err := filepath.Glob(filepath.Join(sessionDir, "*.yaml"))
	if err != nil {
		return nil, fmt.Errorf("failed to glob session directory: %w", err)
	}

	summary := &SessionSummary{}

	for _, f := range allFiles {
		base := filepath.Base(f)
		if strings.HasPrefix(base, "k8s-analysis-") {
			k8s, err := readK8sFile(f)
			if err != nil {
				return nil, fmt.Errorf("failed to read %s: %w", base, err)
			}
			summary.K8sAnalyses = append(summary.K8sAnalyses, k8s)
		} else if !isExcluded(base) {
			job, err := readJobFile(f)
			if err != nil {
				return nil, fmt.Errorf("failed to read %s: %w", base, err)
			}
			summary.Jobs = append(summary.Jobs, job)
		}
	}

	return summary, nil
}

func isExcluded(basename string) bool {
	for _, prefix := range excludePrefixes {
		if strings.HasPrefix(basename, prefix) {
			return true
		}
	}
	return false
}

func readJobFile(path string) (SummaryJob, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return SummaryJob{}, err
	}

	var jbe JobBuildErrors
	if err := yaml.Unmarshal(data, &jbe); err != nil {
		return SummaryJob{}, err
	}

	job := SummaryJob{
		JobName: jbe.JobName,
	}

	if len(jbe.BuildErrors) > 0 {
		be := jbe.BuildErrors[0]
		job.JobURL = be.JobURL
		job.BuildID = be.BuildID
		job.Category = be.Category
		job.CategoryReason = be.CategoryReason

		job.ErrorSnippets = make([]SummaryErrorSnippet, 0, len(be.BuildLogErrorSnippets))
		for _, s := range be.BuildLogErrorSnippets {
			job.ErrorSnippets = append(job.ErrorSnippets, SummaryErrorSnippet{
				ErrorText: s.ErrorText,
				Link:      s.LinkToLogLine,
			})
		}
	}

	return job, nil
}

func readK8sFile(path string) (SummaryK8s, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return SummaryK8s{}, err
	}

	var result K8sAnalysisResult
	if err := yaml.Unmarshal(data, &result); err != nil {
		return SummaryK8s{}, err
	}

	k8s := SummaryK8s{
		BuildID:           result.BuildID,
		JobName:           result.JobName,
		TotalFindings:     result.Summary.TotalFindings,
		ContainerLogFiles: result.ContainerLogFiles,
	}

	k8s.Findings = make([]SummaryK8sFinding, 0, len(result.Findings))
	for _, f := range result.Findings {
		k8s.Findings = append(k8s.Findings, SummaryK8sFinding{
			Detector:  f.Detector,
			Severity:  f.Severity,
			Kind:      f.Kind,
			Namespace: f.Namespace,
			Name:      f.Name,
			Reason:    f.Reason,
			Message:   f.Message,
		})
	}

	return k8s, nil
}
