package cifailures

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/kubevirt/ci-health/pkg/prow"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"io"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

const logsDir = "output/tmp/build-logs"

var (
	jobNamePattern = regexp.MustCompile(`.*(pull-kubevirt-[^/]+).*`)

	// Regex to find errors in log files
	rgExpression = regexp.MustCompile(`^E\d{4} \d\d:\d\d:\d\d\.\d+|(Error|ERROR|error)s?:|(FAIL|Failure \[)\b|timed out|panic\b|\[FAILED\]|fatal: `)
)

type JobFailureData struct {
	URL string
	Sig string
	ID  string
}

// ShowCIFailureJobs fetches URLs for the CI failure runs from the data of the latest run,
// covering the last 7 days. Resulting URLs are sorted by sig and then by id section,
// where both are captured via regular expression matching on the job url.
// It then returns only the URLs for which the junit.functest.xml artifact does not exist.
func ShowCIFailureJobs() ([]string, error) {
	// 1. Read and parse the JSON file
	jsonFile, err := os.Open("./output/kubevirt/kubevirt/results.json")
	if err != nil {
		return nil, fmt.Errorf("failed to open results.json: %v", err)
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	var result struct {
		Data struct {
			SIGRetests struct {
				FailedJobLeaderBoard []struct {
					FailureURLs []string `json:"FailureURLs"`
				} `json:"FailedJobLeaderBoard"`
			} `json:"SIGRetests"`
		} `json:"Data"`
	}

	if err := json.Unmarshal(byteValue, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal json: %v", err)
	}

	// 2. Extract, capture groups, and flatten URLs
	var jobFailures []JobFailureData
	re := regexp.MustCompile(`.*pull-kubevirt-.*-[0-9]+.[0-9]+-(?P<id>(?P<prefix>(ipv6-)?)(?P<sig>sig-[a-z]+)?[^/]*).*`)
	names := re.SubexpNames()

	for _, board := range result.Data.SIGRetests.FailedJobLeaderBoard {
		for _, url := range board.FailureURLs {
			match := re.FindStringSubmatch(url)
			if match == nil {
				continue
			}

			jobFailure := JobFailureData{URL: url}
			for i, n := range names {
				if i > 0 && i < len(match) {
					switch n {
					case "sig":
						jobFailure.Sig = match[i]
					case "id":
						jobFailure.ID = match[i]
					}
				}
			}
			jobFailures = append(jobFailures, jobFailure)
		}
	}

	// 3. Sort the data
	sort.Slice(jobFailures, func(i, j int) bool {
		if jobFailures[i].Sig != jobFailures[j].Sig {
			return jobFailures[i].Sig < jobFailures[j].Sig
		}
		return jobFailures[i].ID < jobFailures[j].ID
	})

	// 4. Check for artifact existence and add URL to list if not found
	var ciFailureJobURLs []string
	for _, jf := range jobFailures {
		artifactURL := strings.Replace(jf.URL, "https://prow.ci.kubevirt.io//view/gs/", "https://storage.googleapis.com/", 1)
		artifactURL = fmt.Sprintf("%s/artifacts/junit.functest.xml", artifactURL)

		resp, err := http.Head(artifactURL)
		if err != nil {
			ciFailureJobURLs = append(ciFailureJobURLs, jf.URL)
			continue
		}

		if resp.StatusCode != http.StatusOK {
			ciFailureJobURLs = append(ciFailureJobURLs, jf.URL)
		}
		err = resp.Body.Close()
		if err != nil {
			return nil, fmt.Errorf("failed to close output file: %v", err)
		}
	}

	return ciFailureJobURLs, nil
}

type buildLog struct {
	ProwJobURL  string    `yaml:"prow_job_url"`
	BuildLogURL string    `yaml:"build_log_url"`
	LogContent  string    `yaml:"log_content"`
	Started     time.Time `yaml:"started"`
	Finished    time.Time `yaml:"finished"`
}

// DownloadBuildLogs downloads the log files for urls of failed jobs located in output/tmp/ci-failure-jobs.txt
// into the folder output/tmp/build-logs. Each build log file can be identified by the
// file name that holds the build number being the same as the last part of the build url from the
// output/tmp/ci-failure-jobs.txt.
func DownloadBuildLogs(ciFailureJobURLs []string) ([]string, error) {
	if err := os.MkdirAll(logsDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create build logs directory: %w", err)
	}

	var logFiles []string
	for _, url := range ciFailureJobURLs {

		logFilePath, err := storeBuildLogIfNotExists(url)
		if err != nil {
			return nil, err
		}
		logFiles = append(logFiles, logFilePath)
	}

	return logFiles, nil
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

// ExtractErrors fetches failures from the build logs of build urls given through the file.
// It writes matching lines into one file per group, prefixed with 'output/tmp/errors-'.
// It returns the file names of the files created.
func ExtractErrors(ciFailureJobURLs []string) ([]string, error) {
	groups := map[string]string{
		"sig-compute":    "sig-compute|sig-operator|sev|vgpu|windows",
		"sig-network":    "sig-network|sriov",
		"sig-storage":    "sig-storage",
		"sig-monitoring": "sig-monitoring",
	}
	var outputFiles []string

	for sigName, group := range groups {
		urlsMatchingGroup := filterURLsByGroup(ciFailureJobURLs, group)
		if len(urlsMatchingGroup) == 0 {
			continue
		}

		jobErrorsByJobName := map[string]*JobBuildErrors{}

		for _, ciFailureJobURL := range urlsMatchingGroup {
			var logFilePath string
			logFilePath, err := storeBuildLogIfNotExists(ciFailureJobURL)
			if err != nil {
				return nil, fmt.Errorf("failed to fetch build log: %v", err)
			}

			var log *buildLog
			log, err = readBuildLog(logFilePath)
			if err != nil {
				return nil, fmt.Errorf("failed to read build log: %v", err)
			}

			lines := strings.Split(log.LogContent, "\n")

			jobNameSubmatch := jobNamePattern.FindStringSubmatch(ciFailureJobURL)
			if len(jobNameSubmatch) < 2 {
				return nil, fmt.Errorf("could not extract job name from URL: %s", ciFailureJobURL)
			}
			jobName := jobNameSubmatch[1]

			jobErrorsForJob, exists := jobErrorsByJobName[jobName]
			if !exists {
				jobErrorsForJob = &JobBuildErrors{JobName: jobName}
				jobErrorsByJobName[jobName] = jobErrorsForJob
			}

			var buildId int
			buildId, err = buildIDFromJobURL(ciFailureJobURL)
			if err != nil {
				return nil, fmt.Errorf("failed to extract build id from job url: %v", err)
			}
			nextBuildError := &JobBuildError{
				JobURL:   ciFailureJobURL,
				BuildID:  buildId,
				Started:  log.Started,
				Finished: log.Finished,
			}
			jobErrorsForJob.BuildErrors = append(jobErrorsForJob.BuildErrors, nextBuildError)

			for i := len(lines) - 1; i >= 0; i-- {
				line := lines[i]
				if rgExpression.MatchString(line) {
					start := max(0, i-3)
					end := min(len(lines)-1, i+3)
					// to align with what prow build log shows related to line numbers, all values are translated to 1 based
					buildErrorLogSnippet := &BuildLogErrorSnippet{
						StartLine:     start + 1,
						EndLine:       end + 1,
						ErrorLine:     i + 1,
						ErrorText:     strings.TrimSpace(lines[i]),
						LinkToLogLine: fmt.Sprintf("%s#1:build-log.txt%%3A%d", ciFailureJobURL, i+1),
					}
					for j := start; j <= end; j++ {
						// trimming space since we want to avoid double quoted strings for readability
						buildErrorLogSnippet.Context += strings.TrimSpace(lines[j]) + "\n"
					}
					buildErrorLogSnippet.Context = strings.TrimSuffix(buildErrorLogSnippet.Context, "\n")
					nextBuildError.BuildLogErrorSnippets = append(nextBuildError.BuildLogErrorSnippets, buildErrorLogSnippet)
				}
			}
		}

		if len(jobErrorsByJobName) == 0 {
			continue
		}

		outputDirName := fmt.Sprintf("output/tmp/errors-%s", sigName)
		if err := os.MkdirAll(outputDirName, 0755); err != nil {
			return nil, fmt.Errorf("failed to create output directory: %w", err)
		}

		for jobName, jobErrs := range jobErrorsByJobName {
			outputFileName := filepath.Join(outputDirName, fmt.Sprintf("%s.yaml", jobName))
			var outputFile *os.File
			outputFile, err := os.Create(outputFileName)
			if err != nil {
				return nil, fmt.Errorf("failed to create errors file %s: %v", outputFileName, err)
			}
			encoder := yaml.NewEncoder(outputFile)
			err = encoder.Encode(&jobErrs)
			if err != nil {
				return nil, fmt.Errorf("failed to encode errors file %s: %v", outputFileName, err)
			}

			err = outputFile.Close()
			if err != nil {
				return nil, fmt.Errorf("failed to close output file: %v", err)
			}
			outputFiles = append(outputFiles, outputFileName)
		}
	}
	return outputFiles, nil
}

func readBuildLog(logFilePath string) (*buildLog, error) {
	var buildLogFile *os.File
	buildLogFile, err := os.Open(logFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open log file: %v", err)
	}

	var log *buildLog
	decoder := yaml.NewDecoder(buildLogFile)
	err = decoder.Decode(&log)
	if err != nil {
		return nil, fmt.Errorf("failed to read log file: %v", err)
	}
	err = buildLogFile.Close()
	if err != nil {
		return nil, fmt.Errorf("failed to close log file: %v", err)
	}
	return log, nil
}

func storeBuildLogIfNotExists(url string) (string, error) {
	buildID, err := buildIDFromJobURL(url)
	if err != nil {
		return "", err
	}

	logFilePath := filepath.Join(logsDir, fmt.Sprintf("%d.yaml", buildID))
	_, err = os.Stat(logFilePath)
	if errors.Is(err, fs.ErrNotExist) {
		gcsBaseURL := strings.Replace(url, "https://prow.ci.kubevirt.io//view/gs/", "https://storage.googleapis.com/", 1)

		var gcsBuildLogFileContent []byte
		gcsBuildLogFileContent, err = retrieveFileContentFromGCS(gcsBaseURL + "/build-log.txt")
		if err != nil {
			return "", fmt.Errorf("failed to read build log content body: %v", err)
		}

		var gcsStartedFileContent []byte
		gcsStartedFileContent, err = retrieveFileContentFromGCS(gcsBaseURL + "/started.json")
		if err != nil {
			return "", fmt.Errorf("failed to read build log content body: %v", err)
		}

		var started prow.Started
		err = json.Unmarshal(gcsStartedFileContent, &started)
		if err != nil {
			return "", fmt.Errorf("failed to decode started: %v", err)
		}

		var gcsFinishedFileContent []byte
		gcsFinishedFileContent, err = retrieveFileContentFromGCS(gcsBaseURL + "/finished.json")
		if err != nil {
			return "", fmt.Errorf("failed to read build log content body: %v", err)
		}

		var finished prow.Started
		err = json.Unmarshal(gcsFinishedFileContent, &finished)
		if err != nil {
			return "", fmt.Errorf("failed to decode finished: %v", err)
		}

		var logFile *os.File
		logFile, err = os.Create(logFilePath)
		if err != nil {
			return "", fmt.Errorf("failed to create log file for build %d: %v", buildID, err)
		}
		encoder := yaml.NewEncoder(logFile)

		l := buildLog{
			ProwJobURL:  url,
			BuildLogURL: gcsBaseURL + "/build-log.txt",
			LogContent:  string(gcsBuildLogFileContent),
			Started:     started.Time(),
			Finished:    finished.Time(),
		}
		err = encoder.Encode(&l)
		if err != nil {
			return "", fmt.Errorf("failed to write build log: %v", err)
		}
		err = logFile.Close()
		if err != nil {
			return "", fmt.Errorf("failed to close log file: %v", err)
		}
	}
	return logFilePath, nil
}

func closeAndLogErr(closer io.Closer) {
	err := closer.Close()
	if err != nil {
		log.WithError(err).Warn("failed to close")
	}
}

func retrieveFileContentFromGCS(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to download content from %s: %v", url, err)
	}
	defer closeAndLogErr(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to download content from %s: status code %d", url, resp.StatusCode)
	}

	var gcsFileContent []byte
	gcsFileContent, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read content body from %s: %v", url, err)
	}
	return gcsFileContent, nil
}

func buildIDFromJobURL(url string) (int, error) {
	buildID := filepath.Base(url)
	if buildID == "." || buildID == "/" {
		return 0, fmt.Errorf("could not determine build ID from URL: %s", url)
	}
	atoi, err := strconv.Atoi(buildID)
	if err != nil {
		return 0, fmt.Errorf("could not convert build ID to int: %s", url)
	}
	return atoi, nil
}

func filterURLsByGroup(ciFailureJobURLs []string, groupPattern string) []string {
	groupRegex := regexp.MustCompile(groupPattern)

	var urls []string
	for _, url := range ciFailureJobURLs {
		if groupRegex.MatchString(url) {
			urls = append(urls, url)
		}
	}
	return urls
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
