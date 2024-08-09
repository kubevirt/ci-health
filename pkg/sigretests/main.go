package sigretests

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

const org = "kubevirt"
const repo = "kubevirt"

type failedJob struct {
	jobName     string
	buildNumber string
}

type SigRetests struct {
	SigCompute     int
	SigNetwork     int
	SigStorage     int
	SigOperator    int
	FailedJobNames []string
}

var redJobs []failedJob

func prHistoryURL(org string, repo string, prNumber string) string {
	return fmt.Sprintf("https://prow.ci.kubevirt.io/pr-history/?org=%s&repo=%s&pr=%s", org, repo, prNumber)
}

func prStorageURL(org string, repo string, prNumber string) string {
	return fmt.Sprintf("https://gcsweb.ci.kubevirt.io/gcs/kubevirt-prow/pr-logs/pull/%s_%s/%s/", org, repo, prNumber)
}

func finishedJSONURL(org string, repo string, prNumber string, jobName string, buildNumber string) string {
	return fmt.Sprintf("%s/%s/%s/finished.json", prStorageURL(org, repo, prNumber), jobName, buildNumber)
}

func filterFailedJobs(node *html.Node) (failedJobs []failedJob) {
	var redJob failedJob
	if node.Type == html.ElementNode && node.Data == "td" {
		for _, td := range node.Attr {
			if strings.Contains(td.Val, "run-failure") {
				for _, href := range node.FirstChild.Attr {
					if strings.Contains(href.Val, "e2e") {
						buildLogUrl := strings.Split(href.Val, "/")
						redJob.jobName = buildLogUrl[len(buildLogUrl)-2]
						redJob.buildNumber = buildLogUrl[len(buildLogUrl)-1]
						redJobs = append(redJobs, redJob)
					}
				}
			}
		}
	}
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		filterFailedJobs(child)
	}
	return redJobs
}

func getLatestCommit(node *html.Node) (latestCommit string) {
	if node.Type == html.ElementNode && node.Data == "a" {
		for _, td := range node.Attr {
			if strings.Contains(td.Val, "commit") {
				commitURL := strings.Split(td.Val, "/")
				latestCommit := commitURL[len(commitURL)-1]
				return latestCommit
			}
		}
	}
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		latestCommit = getLatestCommit(child)
		if latestCommit != "" {
			return latestCommit
		}
	}
	return ""
}

func getFailedJobsForLatestCommit(org string, repo string, prNumber string) (failedJobs []failedJob, err error) {
	prHistory := prHistoryURL(org, repo, prNumber)
	resp, err := http.Get(prHistory)
	if err != nil {
		return nil, fmt.Errorf("Failed to get PR history from %s", prHistory)
	}
	defer resp.Body.Close()

	prHistoryPage, err := html.Parse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Error parsing history page: %s", err)
	}
	latestCommit := getLatestCommit(prHistoryPage)
	if latestCommit == "" {
		return nil, fmt.Errorf("Failed to get latest commit from %s", prHistory)
	}
	failedJobsAllCommits := filterFailedJobs(prHistoryPage)
	redJobs = nil

	for _, job := range failedJobsAllCommits {
		finishedJSON, err := http.Get(finishedJSONURL(org, repo, prNumber, job.jobName, job.buildNumber))
		if finishedJSON.StatusCode != http.StatusOK {
			continue
		}

		if err != nil {
			return nil, fmt.Errorf("Failed to get %s finished.json : %s", job.jobName, err)
		}

		defer finishedJSON.Body.Close()

		finishedJSONData, err := io.ReadAll(finishedJSON.Body)
		if err != nil {
			return nil, fmt.Errorf("Failed to read finished JSON for %s -- %s", job.jobName, err)
		}
		var data map[string]interface{}
		err = json.Unmarshal(finishedJSONData, &data)
		if err != nil {
			return nil, fmt.Errorf("Failed to unmarshall finished JSON for %s -- %s", job.jobName, err)
		}
		if latestCommit == data["revision"] {
			failedJobs = append(failedJobs, job)
		}

	}
	return failedJobs, nil
}

func FilterJobsPerSigs(jobs []failedJob) (prSigRetests SigRetests) {
	prSigRetests = SigRetests{}
	for _, job := range jobs {
		switch {
		case strings.Contains(job.jobName, "sig-compute") || strings.Contains(job.jobName, "vgpu"):
			prSigRetests.SigCompute += 1
			prSigRetests.FailedJobNames = append(prSigRetests.FailedJobNames, job.jobName)

		case strings.Contains(job.jobName, "sig-network") || strings.Contains(job.jobName, "sriov"):
			prSigRetests.SigNetwork += 1
			prSigRetests.FailedJobNames = append(prSigRetests.FailedJobNames, job.jobName)

		case strings.Contains(job.jobName, "sig-storage"):
			prSigRetests.SigStorage += 1
			prSigRetests.FailedJobNames = append(prSigRetests.FailedJobNames, job.jobName)

		case strings.Contains(job.jobName, "sig-operator"):
			prSigRetests.SigOperator += 1
			prSigRetests.FailedJobNames = append(prSigRetests.FailedJobNames, job.jobName)
		}
	}
	return prSigRetests
}

func GetFailuresPerSIG(prNumber string, org string, repo string) (prSigRetests SigRetests, err error) {
	failedJobs, err := getFailedJobsForLatestCommit(org, repo, prNumber)
	if err != nil {
		return SigRetests{}, fmt.Errorf("Error filtering failed jobs from the latest commit - %s", err)
	}
	return FilterJobsPerSigs(failedJobs), nil
}
