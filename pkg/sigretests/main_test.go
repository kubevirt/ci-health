package sigretests

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	"golang.org/x/net/html"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSIGRetests(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "sigretests suite")
}

var _ = Describe("main", func() {
	Context("parse html", func() {

		var rootHtmlNode *html.Node
		var prNumber string

		BeforeEach(func() {
			prNumber = strconv.Itoa(16831)
			content, err := os.ReadFile("testdata/pr-history-16831.html")
			Expect(err).To(Succeed())
			rootHtmlNode, err = html.Parse(bytes.NewReader(content))
			Expect(err).To(Succeed())
		})

		It("getLatestCommit", func() {
			commit := getLatestCommit(rootHtmlNode)
			Expect(commit).ToNot(BeNil())
		})

		It("filterJobs", func() {
			jobsAllCommits := filterJobs(rootHtmlNode)
			Expect(jobsAllCommits).ToNot(BeNil())
		})

		It("filterForLastCommit", func() {
			commit := getLatestCommit(rootHtmlNode)
			Expect(commit).ToNot(BeNil())
			jobsAllCommits := filterJobs(rootHtmlNode)
			Expect(jobsAllCommits).ToNot(BeNil())
			_, err := filterForLastCommit(defaultStorageBaseURL, org, repo, prNumber, commit, jobsAllCommits, time.Time{})
			Expect(err).To(Succeed())
		})

		It("filterOptionalJobs", func() {
			commit := getLatestCommit(rootHtmlNode)
			Expect(commit).ToNot(BeNil())
			jobsAllCommits := filterJobs(rootHtmlNode)
			Expect(jobsAllCommits).ToNot(BeNil())
			jobsLatestCommit, err := filterForLastCommit(defaultStorageBaseURL, org, repo, prNumber, commit, jobsAllCommits, time.Time{})
			Expect(err).To(Succeed())
			_, err = filterOptionalJobs(org, repo, prNumber, jobsLatestCommit)
			Expect(err).To(Succeed())
		})
	})

	Context("filterForLastCommit time filtering", func() {
		var server *httptest.Server
		var testCommit string
		var testJobs []job

		BeforeEach(func() {
			testCommit = "abc123"

			testJobs = []job{
				{jobName: "job-old", buildNumber: "100", failure: true},
				{jobName: "job-recent", buildNumber: "200", failure: true},
				{jobName: "job-no-timestamp", buildNumber: "300", failure: false},
			}

			oldTimestamp := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC).Unix()
			recentTimestamp := time.Date(2025, 3, 15, 0, 0, 0, 0, time.UTC).Unix()

			server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				switch {
				case strings.HasSuffix(r.URL.Path, "/job-old/100/finished.json"):
					fmt.Fprintf(w, `{"timestamp": %d, "revision": "%s", "result": "FAILURE"}`, oldTimestamp, testCommit)
				case strings.HasSuffix(r.URL.Path, "/job-recent/200/finished.json"):
					fmt.Fprintf(w, `{"timestamp": %d, "revision": "%s", "result": "FAILURE"}`, recentTimestamp, testCommit)
				case strings.HasSuffix(r.URL.Path, "/job-no-timestamp/300/finished.json"):
					fmt.Fprintf(w, `{"revision": "%s", "result": "SUCCESS"}`, testCommit)
				default:
					http.NotFound(w, r)
				}
			}))
		})

		AfterEach(func() {
			server.Close()
		})

		It("returns all jobs when notBefore is zero", func() {
			filtered, err := filterForLastCommit(server.URL, org, repo, "42", testCommit, testJobs, time.Time{})
			Expect(err).To(Succeed())
			Expect(filtered).To(HaveLen(3))
		})

		It("filters out jobs that finished before notBefore", func() {
			cutoff := time.Date(2025, 3, 1, 0, 0, 0, 0, time.UTC)
			filtered, err := filterForLastCommit(server.URL, org, repo, "42", testCommit, testJobs, cutoff)
			Expect(err).To(Succeed())
			Expect(filtered).To(HaveLen(2))
			Expect(filtered[0].jobName).To(Equal("job-recent"))
			Expect(filtered[1].jobName).To(Equal("job-no-timestamp"))
		})

		It("keeps jobs with no timestamp field", func() {
			cutoff := time.Date(2025, 6, 1, 0, 0, 0, 0, time.UTC)
			filtered, err := filterForLastCommit(server.URL, org, repo, "42", testCommit, testJobs, cutoff)
			Expect(err).To(Succeed())
			Expect(filtered).To(HaveLen(1))
			Expect(filtered[0].jobName).To(Equal("job-no-timestamp"))
		})
	})
})
