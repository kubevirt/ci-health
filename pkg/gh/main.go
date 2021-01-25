package gh

import (
	"context"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/shurcooL/githubv4"
	log "github.com/sirupsen/logrus"
	"golang.org/x/oauth2"

	"github.com/fgimenez/ci-health/pkg/constants"
	"github.com/fgimenez/ci-health/pkg/mergequeue"
	"github.com/fgimenez/ci-health/pkg/types"
)

var (
	zeroDate = time.Time{}
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

// MergeQueueSizeByDate returns the merge queue size for a given date.
func (c *Client) MergeQueueSizeByDate(date time.Time) (int, error) {
	mergedQueryString := fmt.Sprintf("repo:%s created:<%[2]s type:pr merged>=%[2]s", c.source, date.Format(constants.DateFormat))
	log.Debugf("merged query: %q", mergedQueryString)
	mergedQueryResult, err := c.prQuery(mergedQueryString)
	if err != nil {
		return 0, err
	}

	notMergedQueryString := fmt.Sprintf("repo:%s created:<=%s type:pr is:open", c.source, date.Format(constants.DateFormat))
	log.Debugf("not merged query: %q", notMergedQueryString)
	notMergedQueryResult, err := c.prQuery(notMergedQueryString)
	if err != nil {
		return 0, err
	}

	log.Debugf("Merge query result length: %d, not merged query result length: %d", len(mergedQueryResult), len(notMergedQueryResult))

	result := 0
	for _, pr := range append(mergedQueryResult, notMergedQueryResult...) {
		if mergequeue.DatePREntered(&pr.PullRequestFragment, date) != zeroDate {
			result++
		}
	}

	return result, nil
}

func (c *Client) prQuery(query string) ([]struct {
	types.PullRequestFragment `graphql:"... on PullRequest"`
}, error) {
	ctx := context.Background()

	variables := map[string]interface{}{
		"querystring": githubv4.String(query),
	}

	var mergedQuery struct {
		Search struct {
			Nodes []struct {
				types.PullRequestFragment `graphql:"... on PullRequest"`
			}
		} `graphql:"search(query: $querystring, type: ISSUE, first:100)"`
	}

	err := c.inner.Query(ctx, &mergedQuery, variables)
	return mergedQuery.Search.Nodes, err
}
