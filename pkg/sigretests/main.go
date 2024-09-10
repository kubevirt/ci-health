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

type job struct {
	jobName     string
	buildNumber string
	buildURL    string
	failure     bool
}

type SigRetests struct {
	SigComputeFailure  int
	SigNetworkFailure  int
	SigStorageFailure  int
	SigOperatorFailure int
	SigComputeSuccess  int
	SigNetworkSuccess  int
	SigStorageSuccess  int
	SigOperatorSuccess int
	FailedJobNames     []string
	SuccessJobNames    []string
}

var prowjobs []job

func prHistoryURL(org string, repo string, prNumber string) string {
	return fmt.Sprintf("https://prow.ci.kubevirt.io/pr-history/?org=%s&repo=%s&pr=%s", org, repo, prNumber)
}

func prStorageURL(org string, repo string, prNumber string) string {
	return fmt.Sprintf("https://gcsweb.ci.kubevirt.io/gcs/kubevirt-prow/pr-logs/pull/%s_%s/%s/", org, repo, prNumber)
}

func finishedJSONURL(org string, repo string, prNumber string, jobName string, buildNumber string) string {
	return fmt.Sprintf("%s/%s/%s/finished.json", prStorageURL(org, repo, prNumber), jobName, buildNumber)
}

func filterJobs(node *html.Node) (jobs []job) {
	var e2eJob job
	if node.Type == html.ElementNode && node.Data == "td" {
		for _, td := range node.Attr {
			if strings.Contains(td.Val, "run-failure") {
				for _, href := range node.FirstChild.Attr {
					if strings.Contains(href.Val, "e2e") {
						e2eJob.failure = true
						e2eJob.buildURL = href.Val
						buildLogUrl := strings.Split(href.Val, "/")
						e2eJob.jobName = buildLogUrl[len(buildLogUrl)-2]
						e2eJob.buildNumber = buildLogUrl[len(buildLogUrl)-1]
						prowjobs = append(prowjobs, e2eJob)
						continue
					}
				}
			}
			if strings.Contains(td.Val, "run-success") {
				for _, href := range node.FirstChild.Attr {
					if strings.Contains(href.Val, "e2e") {
						e2eJob.failure = false
						e2eJob.buildURL = href.Val
						buildLogUrl := strings.Split(href.Val, "/")
						e2eJob.jobName = buildLogUrl[len(buildLogUrl)-2]
						e2eJob.buildNumber = buildLogUrl[len(buildLogUrl)-1]
						prowjobs = append(prowjobs, e2eJob)
						continue
					}
				}
			}
		}
	}
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		filterJobs(child)
	}
	return prowjobs
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

func filterForLastCommit(org string, repo string, prNumber string, latestCommit string, jobList []job) (filteredJobList []job, err error) {
	for _, job := range jobList {
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
			filteredJobList = append(filteredJobList, job)
		}
	}
	return filteredJobList, nil
}

func getJobsForLatestCommit(org string, repo string, prNumber string) (jobsLastCommit []job, err error) {
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
	jobsAllCommits := filterJobs(prHistoryPage)
	prowjobs = nil

	jobsLatestCommit, err := filterForLastCommit(org, repo, prNumber, latestCommit, jobsAllCommits)
	if err != nil {
		return nil, err
	}
	return jobsLatestCommit, nil
}

func sortJobNamesOnResult(job job, sigRetests SigRetests) (jobCounts SigRetests) {
	if job.failure == true {
		sigRetests.FailedJobNames = append(sigRetests.FailedJobNames, job.jobName)
	} else {
		sigRetests.SuccessJobNames = append(sigRetests.SuccessJobNames, job.jobName)
	}
	return sigRetests
}

func FilterJobsPerSigs(jobs []job) (prSigRetests SigRetests) {
	prSigRetests = SigRetests{}
	for _, job := range jobs {
		switch {
		case strings.Contains(job.jobName, "sig-compute") || strings.Contains(job.jobName, "vgpu"):
			if job.failure {
				prSigRetests.SigComputeFailure += 1
			} else {
				prSigRetests.SigComputeSuccess += 1
			}
		case strings.Contains(job.jobName, "sig-network") || strings.Contains(job.jobName, "sriov"):
			if job.failure {
				prSigRetests.SigNetworkFailure += 1
			} else {
				prSigRetests.SigNetworkSuccess += 1
			}
		case strings.Contains(job.jobName, "sig-storage"):
			if job.failure {
				prSigRetests.SigStorageFailure += 1
			} else {
				prSigRetests.SigStorageSuccess += 1
			}
		case strings.Contains(job.jobName, "sig-operator"):
			if job.failure {
				prSigRetests.SigOperatorFailure += 1
			} else {
				prSigRetests.SigOperatorSuccess += 1
			}
		}
		prSigRetests = sortJobNamesOnResult(job, prSigRetests)
	}
	return prSigRetests
}

func GetJobsPerSIG(prNumber string, org string, repo string) (prSigRetests SigRetests, err error) {
	prowJobs, err := getJobsForLatestCommit(org, repo, prNumber)
	if err != nil {
		return SigRetests{}, fmt.Errorf("Error filtering failed jobs from the latest commit - %s", err)
	}
	return FilterJobsPerSigs(prowJobs), nil
}
