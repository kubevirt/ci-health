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
	. "github.com/onsi/ginkgo/extensions/table"
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

	Context("DoHTTPWithRetry", func() {
		var server *httptest.Server

		AfterEach(func() {
			if server != nil {
				server.Close()
			}
		})

		It("returns nil error for 200 OK", func() {
			server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			}))
			resp, err := DoHTTPWithRetry(server.URL, http.Get)
			Expect(err).ToNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(http.StatusOK))
		})

		It("returns nil error for 404 Not Found", func() {
			server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusNotFound)
			}))
			resp, err := DoHTTPWithRetry(server.URL, http.Head)
			Expect(err).ToNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(http.StatusNotFound))
		})

		It("returns error for 403 Forbidden", func() {
			server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusForbidden)
			}))
			_, err := DoHTTPWithRetry(server.URL, http.Get)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("status 403"))
		})

		DescribeTable("retries on transient 5xx and succeeds",
			func(statusCode int, failCount int) {
				attempt := 0
				server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					attempt++
					if attempt <= failCount {
						w.WriteHeader(statusCode)
						return
					}
					w.WriteHeader(http.StatusOK)
				}))
				resp, err := DoHTTPWithRetry(server.URL, http.Get)
				Expect(err).ToNot(HaveOccurred())
				Expect(resp.StatusCode).To(Equal(http.StatusOK))
				Expect(attempt).To(Equal(failCount + 1))
			},
			Entry("500 Internal Server Error", http.StatusInternalServerError, 1),
			Entry("502 Bad Gateway", http.StatusBadGateway, 2),
			Entry("503 Service Unavailable", http.StatusServiceUnavailable, 1),
			Entry("504 Gateway Timeout", http.StatusGatewayTimeout, 1),
		)

		It("does not retry on 501 Not Implemented", func() {
			attempt := 0
			server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				attempt++
				w.WriteHeader(http.StatusNotImplemented)
			}))
			_, err := DoHTTPWithRetry(server.URL, http.Get)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("status 501"))
			Expect(attempt).To(Equal(1))
		})

		It("returns error after exhausting retries on persistent 502", func() {
			server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusBadGateway)
			}))
			_, err := DoHTTPWithRetry(server.URL, http.Get)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("transient HTTP 502"))
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
					fmt.Fprintf(w, `{"timestamp": %d, "revision": "%s", "result": "FAILURE"}`, oldTimestamp, testCommit) //nolint:errcheck
				case strings.HasSuffix(r.URL.Path, "/job-recent/200/finished.json"):
					fmt.Fprintf(w, `{"timestamp": %d, "revision": "%s", "result": "FAILURE"}`, recentTimestamp, testCommit) //nolint:errcheck
				case strings.HasSuffix(r.URL.Path, "/job-no-timestamp/300/finished.json"):
					fmt.Fprintf(w, `{"revision": "%s", "result": "SUCCESS"}`, testCommit) //nolint:errcheck
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
