package cifailures

import (
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

// normalizeJobURL ensures the URL uses the double-slash format expected by storeBuildLogIfNotExists.
func normalizeJobURL(url string) string {
	url = strings.TrimRight(url, "/")
	const singleSlash = "https://prow.ci.kubevirt.io/view/gs/"
	const doubleSlash = "https://prow.ci.kubevirt.io//view/gs/"
	if strings.HasPrefix(url, singleSlash) && !strings.HasPrefix(url, doubleSlash) {
		url = doubleSlash + strings.TrimPrefix(url, singleSlash)
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
	}

	lines := strings.Split(log.LogContent, "\n")
	for i := len(lines) - 1; i >= 0; i-- {
		line := lines[i]
		if rgExpression.MatchString(line) {
			start := max(0, i-3)
			end := min(len(lines)-1, i+3)
			snippet := &BuildLogErrorSnippet{
				StartLine:     start + 1,
				EndLine:       end + 1,
				ErrorLine:     i + 1,
				ErrorText:     strings.TrimSpace(lines[i]),
				LinkToLogLine: fmt.Sprintf("%s#1:build-log.txt%%3A%d", prowJobURL, i+1),
			}
			for j := start; j <= end; j++ {
				snippet.Context += strings.TrimSpace(lines[j]) + "\n"
			}
			snippet.Context = strings.TrimSuffix(snippet.Context, "\n")
			buildError.BuildLogErrorSnippets = append(buildError.BuildLogErrorSnippets, snippet)
		}
	}

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
	defer outputFile.Close()

	encoder := yaml.NewEncoder(outputFile)
	if err := encoder.Encode(jobBuildErrors); err != nil {
		return fmt.Errorf("failed to encode YAML: %v", err)
	}
	return nil
}
