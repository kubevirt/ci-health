package cifailures

import (
	"fmt"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

// normalizeJobURL ensures the URL uses the canonical single-slash format.
func normalizeJobURL(url string) string {
	url = strings.TrimRight(url, "/")
	const singleSlash = "https://prow.ci.kubevirt.io/view/gs/"
	const doubleSlash = "https://prow.ci.kubevirt.io//view/gs/"
	if after, found := strings.CutPrefix(url, doubleSlash); found {
		url = singleSlash + after
	}
	return url
}

// AnalyzeBuild downloads the build log for a single prow job URL, extracts errors,
// categorizes them, and returns the result as a JobBuildErrors struct.
func AnalyzeBuild(prowJobURL string) (*JobBuildErrors, error) {
	prowJobURL = normalizeJobURL(prowJobURL)

	if err := os.MkdirAll(logsDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create build logs directory: %w", err)
	}

	logFilePath, err := storeBuildLogIfNotExists(prowJobURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch build log: %v", err)
	}

	log, err := readBuildLog(logFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read build log: %v", err)
	}

	jobName, err := jobNameFromURL(prowJobURL)
	if err != nil {
		return nil, err
	}

	buildId, err := buildIDFromJobURL(prowJobURL)
	if err != nil {
		return nil, fmt.Errorf("failed to extract build id from job url: %v", err)
	}

	buildError := &JobBuildError{
		JobURL:   prowJobURL,
		BuildID:  buildId,
		Started:  log.Started,
		Finished: log.Finished,
		Result:   log.Result,
		Passed:   log.Passed,
	}

	buildError.BuildLogErrorSnippets = extractErrorSnippets(log.LogContent, prowJobURL)

	buildError.Category = string(CategorizeJobBuildError(buildError))

	return &JobBuildErrors{
		JobName:     jobName,
		BuildErrors: []*JobBuildError{buildError},
	}, nil
}

// WriteJobBuildErrorsYAML writes a JobBuildErrors struct to a file in YAML format.
func WriteJobBuildErrorsYAML(outputPath string, jobBuildErrors *JobBuildErrors) error {
	outputFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file %s: %v", outputPath, err)
	}
	defer func() {
		if err := outputFile.Close(); err != nil {
			log.WithError(err).Warn("failed closing output file")
		}
	}()

	encoder := yaml.NewEncoder(outputFile)
	if err := encoder.Encode(jobBuildErrors); err != nil {
		return fmt.Errorf("failed to encode YAML: %v", err)
	}
	return nil
}
