package gh

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"sort"
	"strings"
	"time"

	retry "github.com/avast/retry-go"
	"github.com/shurcooL/githubv4"
	log "github.com/sirupsen/logrus"
	"golang.org/x/oauth2"

	"github.com/kubevirt/ci-health/pkg/constants"
	"github.com/kubevirt/ci-health/pkg/types"
)

// retryRoundTripper wraps an http.RoundTripper and retries on transient 5xx errors.
type retryRoundTripper struct {
	inner http.RoundTripper
}

func (rt *retryRoundTripper) RoundTrip(req *http.Request) (resp *http.Response, err error) {
	retryErr := retry.Do(
		func() error {
			// Reset the request body for retries (POST bodies are consumed after the first attempt).
			if req.GetBody != nil {
				body, bodyErr := req.GetBody()
				if bodyErr != nil {
					return retry.Unrecoverable(bodyErr)
				}
				req.Body = body
			}
			resp, err = rt.inner.RoundTrip(req)
			if err != nil {
				return retry.Unrecoverable(err)
			}
			switch resp.StatusCode {
			case http.StatusBadGateway, http.StatusServiceUnavailable, http.StatusGatewayTimeout:
				log.Warnf("GitHub API returned %d, will retry", resp.StatusCode)
				return fmt.Errorf("transient HTTP %d from GitHub API", resp.StatusCode)
			default:
				return nil
			}
		},
		retry.Attempts(5),
		retry.Delay(2*time.Second),
		retry.MaxDelay(30*time.Second),
		retry.LastErrorOnly(true),
	)
	if retryErr != nil && err == nil {
		err = retryErr
	}
	return
}

type Client struct {
	inner  *githubv4.Client
	source string
}

func NewClient(tokenPath string, source string) (*Client, error) {
	token, err := os.ReadFile(tokenPath)
	if err != nil {
		return nil, err
	}
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: strings.TrimSpace(string(token))},
	)
	httpClient := oauth2.NewClient(context.Background(), src)
	httpClient.Transport = &retryRoundTripper{inner: httpClient.Transport}

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

const maxOpenPRSearchResults = 500

// OpenPRNumbers returns currently open PR numbers, optionally limited to those
// updated at or after updatedSince. Results are capped so a large backlog of
// stale PRs cannot explode follow-up GraphQL and Prow queries.
func (c *Client) OpenPRNumbers(updatedSince time.Time) ([]int, error) {
	queryString := fmt.Sprintf("repo:%s type:pr is:open", c.source)
	if !updatedSince.IsZero() {
		queryString += fmt.Sprintf(" updated:>=%s", updatedSince.Format(constants.DateFormat))
	}
	log.Debugf("open PRs query: %q", queryString)

	var numbers []int
	var cursor *githubv4.String
	for {
		variables := map[string]interface{}{
			"querystring": githubv4.String(queryString),
			"cursor":      cursor,
		}
		var searchQuery struct {
			Search struct {
				PageInfo struct {
					HasNextPage bool
					EndCursor   githubv4.String
				}
				Nodes types.BarePRList
			} `graphql:"search(query: $querystring, type: ISSUE, first: 100, after: $cursor)"`
		}
		if err := c.inner.Query(context.Background(), &searchQuery, variables); err != nil {
			return nil, err
		}
		for _, node := range searchQuery.Search.Nodes {
			numbers = append(numbers, node.Number)
			if len(numbers) >= maxOpenPRSearchResults {
				log.Warnf("open PR search hit cap of %d results; later PRs are omitted", maxOpenPRSearchResults)
				return numbers, nil
			}
		}
		if !searchQuery.Search.PageInfo.HasNextPage {
			break
		}
		cursor = githubv4.NewString(searchQuery.Search.PageInfo.EndCursor)
	}
	log.Debugf("open PRs query result length: %d", len(numbers))
	return numbers, nil
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

// FetchPRTimelineItems fetches the labeled/unlabeled timeline items for a
// single PR via the repository query. This avoids the GitHub GraphQL
// complexity budget that silently truncates nested timeline data when many
// PRs are returned through the search API.
func (c *Client) FetchPRTimelineItems(number int) (*types.MergeQueuePullRequestFragment, error) {
	parts := strings.SplitN(c.source, "/", 2)
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid source %q, expected owner/repo", c.source)
	}

	variables := map[string]interface{}{
		"owner":  githubv4.String(parts[0]),
		"repo":   githubv4.String(parts[1]),
		"number": githubv4.Int(number),
	}

	var query struct {
		Repository struct {
			PullRequest types.MergeQueuePullRequestFragment `graphql:"pullRequest(number: $number)"`
		} `graphql:"repository(owner: $owner, name: $repo)"`
	}

	err := c.inner.Query(context.Background(), &query, variables)
	if err != nil {
		return nil, err
	}
	return &query.Repository.PullRequest, nil
}

// FetchChatopsPR fetches commit, force-push, and comment timeline items for a
// single PR. Per-PR queries avoid the search API truncating nested timelines.
func (c *Client) FetchChatopsPR(number int) (*types.ChatopsPullRequestFragment, error) {
	parts := strings.SplitN(c.source, "/", 2)
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid source %q, expected owner/repo", c.source)
	}

	variables := map[string]interface{}{
		"owner":  githubv4.String(parts[0]),
		"repo":   githubv4.String(parts[1]),
		"number": githubv4.Int(number),
	}

	var query struct {
		Repository struct {
			PullRequest types.ChatopsPullRequestFragment `graphql:"pullRequest(number: $number)"`
		} `graphql:"repository(owner: $owner, name: $repo)"`
	}

	err := c.inner.Query(context.Background(), &query, variables)
	if err != nil {
		return nil, err
	}
	return &query.Repository.PullRequest, nil
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
