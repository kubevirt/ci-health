package gh

import (
	"context"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

type Client struct {
	inner    *githubv4.Client
	dataDays int
	source   string
}

func NewClient(tokenPath, source string, dataDays int) (*Client, error) {
	token, err := ioutil.ReadFile(tokenPath)
	if err != nil {
		return nil, err
	}
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: string(token)},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	inner := githubv4.NewClient(httpClient)

	client := &Client{
		inner,
		dataDays,
		source,
	}
	return client, nil
}

func (c *Client) DataDays() int {
	return c.dataDays
}

func (c *Client) Source() string {
	return c.source
}

func (c *Client) GetOpenApprovedPRsByDate(date time.Time) (int, error) {
	ctx := context.Background()

	mergedQueryString := fmt.Sprintf("repo:%s created:<=%[1]s type:pr status:CLOSED merged>%[1]s", c.source, date.Format("2006-01-02"))
	variables := map[string]interface{}{
		"querystring": githubv4.String(mergedQueryString),
	}

	var mergedQuery struct {
		Search struct {
			IssueCount int
			Nodes      []struct {
				PullRequestFragment struct {
					Number   int
					Comments struct {
						Nodes []Comment
					} `graphql:"comments(first: 100, orderBy: {field: UPDATED_AT, direction: DESC})"`
				} `graphql:"... on PullRequest"`
			}
		} `graphql:"search(query: $querystring, type: ISSUE, first:100)"`
	}

	err := c.inner.Query(ctx, &mergedQuery, variables)
	if err != nil {
		return 0, err
	}

	notMergedQueryString := fmt.Sprintf("repo:%s created:<=%s type:pr is:open", c.source, date.Format("2006-01-02"))
	variables = map[string]interface{}{
		"querystring": githubv4.String(notMergedQueryString),
	}
	var notMergedQuery struct {
		Search struct {
			IssueCount int
			Nodes      []struct {
				PullRequestFragment struct {
					Number   int
					Comments struct {
						Nodes []Comment
					} `graphql:"comments(first: 100, orderBy: {field: UPDATED_AT, direction: DESC})"`
				} `graphql:"... on PullRequest"`
			}
		} `graphql:"search(query: $querystring, type: ISSUE, first:100)"`
	}
	err = c.inner.Query(ctx, &notMergedQuery, variables)
	if err != nil {
		return 0, err
	}

	return notMergedQuery.Search.IssueCount, nil
}

func IsPRInMergeQueueAtDate(pr *PullRequestFragment, date time.Time) bool {
	var hasLGTMLabel, hasApprobedLabel, doesNotHaveDoNotMergeLabel, doesNotHaveNeedsRebaseLabel bool

	return hasLGTMLabel && hasApprobedLabel && doesNotHaveDoNotMergeLabel && doesNotHaveNeedsRebaseLabel
}
