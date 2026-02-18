package sigretests

import (
	"bytes"
	"golang.org/x/net/html"
	"os"
	"strconv"
	"testing"

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
			_, err := filterForLastCommit(org, repo, prNumber, commit, jobsAllCommits)
			Expect(err).To(Succeed())
		})

		It("filterOptionalJobs", func() {
			commit := getLatestCommit(rootHtmlNode)
			Expect(commit).ToNot(BeNil())
			jobsAllCommits := filterJobs(rootHtmlNode)
			Expect(jobsAllCommits).ToNot(BeNil())
			jobsLatestCommit, err := filterForLastCommit(org, repo, prNumber, commit, jobsAllCommits)
			Expect(err).To(Succeed())
			_, err = filterOptionalJobs(org, repo, prNumber, jobsLatestCommit)
			Expect(err).To(Succeed())
		})
	})
})
