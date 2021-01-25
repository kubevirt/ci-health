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
	"github.com/fgimenez/ci-health/pkg/types"
)

var (
	zeroDate = time.Time{}
)

type Client struct {
	inner  *githubv4.Client
	source string
}

func NewClient(tokenPath string, source string) (*Client, error) {
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
		source,
	}
	return client, nil
}

// OpenPRsAt returns the list of open PRs at a given date.
func (c *Client) OpenPRsAt(date time.Time) ([]struct {
	types.PullRequestFragment `graphql:"... on PullRequest"`
}, error) {
	mergedQueryString := fmt.Sprintf("repo:%s created:<%[2]s type:pr merged>=%[2]s", c.source, date.Format(constants.DateFormat))
	log.Debugf("merged open query: %q", mergedQueryString)
	mergedQueryResult, err := c.prQuery(mergedQueryString)
	if err != nil {
		return nil, err
	}

	notMergedQueryString := fmt.Sprintf("repo:%s created:<=%s type:pr is:open", c.source, date.Format(constants.DateFormat))
	log.Debugf("not merged open query: %q", notMergedQueryString)
	notMergedQueryResult, err := c.prQuery(notMergedQueryString)
	if err != nil {
		return nil, err
	}

	log.Debugf("Merge query result length: %d, not merged query result length: %d", len(mergedQueryResult), len(notMergedQueryResult))

	return append(mergedQueryResult, notMergedQueryResult...), nil
}

// MergedPRsBetween returns a slice of PRs that were merged in the time frame
// defined by the start and end dates given as parameters.
func (c *Client) MergedPRsBetween(startDate, endDate time.Time) ([]struct {
	types.PullRequestFragment `graphql:"... on PullRequest"`
}, error) {
	mergedQueryString := fmt.Sprintf("repo:%s type:pr merged:%s..%s",
		c.source,
		startDate.Format(constants.DateFormat),
		endDate.Format(constants.DateFormat),
	)
	log.Debugf("merged between query: %q", mergedQueryString)

	return c.prQuery(mergedQueryString)
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
