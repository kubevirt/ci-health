package gh

import (
	"context"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/shurcooL/githubv4"
	log "github.com/sirupsen/logrus"
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
	mergedQueryString := fmt.Sprintf("repo:%s created:<%[2]s type:pr merged>=%[2]s", c.source, date.Format("2006-01-02"))
	log.Debugf("merged query: %q", mergedQueryString)
	mergedQueryResult, err := c.prQuery(mergedQueryString)
	if err != nil {
		return 0, err
	}

	notMergedQueryString := fmt.Sprintf("repo:%s created:<=%s type:pr is:open", c.source, date.Format("2006-01-02"))
	log.Debugf("not merged query: %q", notMergedQueryString)
	notMergedQueryResult, err := c.prQuery(notMergedQueryString)
	if err != nil {
		return 0, err
	}

	log.Debugf("Merge query result length: %d, non merged query result length: %d", len(mergedQueryResult), len(notMergedQueryResult))

	return len(mergedQueryResult) + len(notMergedQueryResult), nil
}

func (c *Client) prQuery(query string) ([]struct {
	PullRequestFragment `graphql:"... on PullRequest"`
}, error) {
	ctx := context.Background()

	variables := map[string]interface{}{
		"querystring": githubv4.String(query),
	}

	var mergedQuery struct {
		Search struct {
			Nodes []struct {
				PullRequestFragment `graphql:"... on PullRequest"`
			}
		} `graphql:"search(query: $querystring, type: ISSUE, first:100)"`
	}

	err := c.inner.Query(ctx, &mergedQuery, variables)
	return mergedQuery.Search.Nodes, err
}

func DatePREnteredMergeQueue(pr *PullRequestFragment, date time.Time) *time.Time {

	return nil
}
