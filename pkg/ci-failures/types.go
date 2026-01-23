package cifailures

import (
	"regexp"
	"time"
)

// Cluster holds a group of similar error messages.
type Cluster struct {
	Representative *JobBuildError
	Errors         []*JobBuildError
}

type JobFailureData struct {
	URL string
	Sig string
	ID  string
}

type BuildLogErrorSnippet struct {
	ErrorText     string `yaml:"error_text"`
	LinkToLogLine string `yaml:"link_to_log_line"`
	StartLine     int    `yaml:"start_line"`
	ErrorLine     int    `yaml:"error_line"`
	EndLine       int    `yaml:"end_line"`
	Context       string `yaml:"context"`
}

type JobBuildError struct {
	JobURL                string                  `yaml:"job_url"`
	BuildID               int                     `yaml:"build_id"`
	Started               time.Time               `yaml:"started"`
	Finished              time.Time               `yaml:"finished"`
	BuildLogErrorSnippets []*BuildLogErrorSnippet `yaml:"build_log_error_snippets"`
}

type JobBuildErrors struct {
	JobName     string           `yaml:"job_name"`
	BuildErrors []*JobBuildError `yaml:"build_errors"`
}

var k8sVersionRegex = regexp.MustCompile(`.*-(1\.[0-9]+)-.*`)

const UnknownK8SVersion = "unknownK8SVersion"

func (j *JobBuildErrors) K8SVersion() string {
	if k8sVersionRegex.MatchString(j.JobName) {
		return k8sVersionRegex.FindString(j.JobName)
	}
	return UnknownK8SVersion
}

func (j *JobBuildErrors) SIG() string {
	return SIGForGroup(j.JobName)
}

var branchRegex = regexp.MustCompile(`.*(-[01]\.[0-9]+)$`)

const MainBranch = "main"

func (j *JobBuildErrors) Branch() string {
	if branchRegex.MatchString(j.JobName) {
		submatch := branchRegex.FindStringSubmatch(j.JobName)
		return "release" + submatch[1]
	}
	return MainBranch
}
