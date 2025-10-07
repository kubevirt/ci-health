package gh

import (
	"context"
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
	"time"

	"github.com/shurcooL/githubv4"
	log "github.com/sirupsen/logrus"
	"golang.org/x/oauth2"

	"github.com/kubevirt/ci-health/pkg/constants"
	"github.com/kubevirt/ci-health/pkg/types"
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
		&oauth2.Token{AccessToken: strings.TrimSpace(string(token))},
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
func (c *Client) OpenPRsAt(date time.Time) (types.MergeQueuePRList, error) {
	mergedQueryString := fmt.Sprintf("repo:%s created:<%[2]s type:pr merged:>=%[2]s", c.source, date.Format(constants.DateFormat))
	log.Debugf("merged open query: %q", mergedQueryString)
	mergedQueryResult, err := c.mergeQueuePRQuery(mergedQueryString)
	if err != nil {
		return nil, err
	}

	notMergedQueryString := fmt.Sprintf("repo:%s created:<=%s type:pr is:open", c.source, date.Format(constants.DateFormat))
	log.Debugf("not merged open query: %q", notMergedQueryString)
	notMergedQueryResult, err := c.mergeQueuePRQuery(notMergedQueryString)
	if err != nil {
		return nil, err
	}

	log.Debugf("merge query result length: %d, not merged query result length: %d", len(mergedQueryResult), len(notMergedQueryResult))

	return append(mergedQueryResult, notMergedQueryResult...), nil
}

// MergedPRsBetween returns a slice of PRs that were merged in the time frame
// defined by the start and end dates given as parameters.
func (c *Client) MergedPRsBetween(startDate, endDate time.Time) (types.MergeQueuePRList, error) {
	mergedQueryString := fmt.Sprintf("repo:%s type:pr merged:%s..%s",
		c.source,
		startDate.Format(constants.DateFormat),
		endDate.Format(constants.DateFormat),
	)
	log.Debugf("merged between query: %q", mergedQueryString)

	mergedQueryResult, err := c.mergeQueuePRQuery(mergedQueryString)
	if err != nil {
		return nil, err
	}

	log.Debugf("merged between query result length: %d", len(mergedQueryResult))

	return mergedQueryResult, nil
}

func (c *Client) mergeQueuePRQuery(query string) (types.MergeQueuePRList, error) {

	variables := map[string]interface{}{
		"querystring": githubv4.String(query),
	}

	var mergedQuery struct {
		Search struct {
			Nodes types.MergeQueuePRList
		} `graphql:"search(query: $querystring, type: ISSUE, first:100)"`
	}

	err := c.inner.Query(context.Background(), &mergedQuery, variables)
	return mergedQuery.Search.Nodes, err
}

// ChatopsMergedPRsBetween returns a slice of PRs that were merged in the time
// frame defined by the start and end dates given as parameters with data
// required by chatops tools.
func (c *Client) ChatopsMergedPRsBetween(startDate, endDate time.Time) (types.ChatopsPRList, error) {
	mergedQueryString := fmt.Sprintf("repo:%s type:pr merged:%s..%s",
		c.source,
		startDate.Format(constants.DateFormat),
		endDate.Format(constants.DateFormat),
	)
	log.Debugf("merged between query: %q", mergedQueryString)

	mergedQueryResult, err := c.chatopsPRQuery(mergedQueryString)
	if err != nil {
		return nil, err
	}

	log.Debugf("merged between query result length: %d", len(mergedQueryResult))

	return mergedQueryResult, nil
}

func (c *Client) chatopsPRQuery(query string) (types.ChatopsPRList, error) {

	variables := map[string]interface{}{
		"querystring": githubv4.String(query),
	}

	var mergedQuery struct {
		Search struct {
			Nodes types.ChatopsPRList
		} `graphql:"search(query: $querystring, type: ISSUE, first:100)"`
	}

	err := c.inner.Query(context.Background(), &mergedQuery, variables)
	return mergedQuery.Search.Nodes, err
}

func (c *Client) GetSupportedBranches(ctx context.Context) ([]string, error) {
	var query struct {
		Repository struct {
			Refs struct {
				Nodes []struct {
					Name string
				}
			} `graphql:"refs(refPrefix: \"refs/heads/\", first: 100)"`
		} `graphql:"repository(owner: \"kubevirt\", name: \"kubevirt\")"`
	}

	err := c.inner.Query(ctx, &query, nil)
	if err != nil {
		return nil, err
	}

	var releaseBranches []string
	for _, node := range query.Repository.Refs.Nodes {
		if strings.HasPrefix(node.Name, "release-") {
			releaseBranches = append(releaseBranches, node.Name)
		}
	}

	sort.Slice(releaseBranches, func(i, j int) bool {
		// Simple semver sort, assumes "release-X.Y" format
		vI := strings.TrimPrefix(releaseBranches[i], "release-")
		vJ := strings.TrimPrefix(releaseBranches[j], "release-")
		return vI > vJ
	})

	supportedBranches := []string{"main"}
	for i := 0; i < 3 && i < len(releaseBranches); i++ {
		supportedBranches = append(supportedBranches, releaseBranches[i])
	}

	return supportedBranches, nil
}
