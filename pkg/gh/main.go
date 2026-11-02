package gh

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"
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

var errTransientHTTP = errors.New("transient GitHub HTTP error")

// githubTransport exposes transient statuses to the full-query retry loop.
type githubTransport struct {
	inner http.RoundTripper
}

func (rt *githubTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	resp, err := rt.inner.RoundTrip(req)
	if err != nil {
		return resp, err
	}
	switch resp.StatusCode {
	case http.StatusBadGateway, http.StatusServiceUnavailable, http.StatusGatewayTimeout:
		if err := resp.Body.Close(); err != nil {
			log.Warnf("Closing transient GitHub HTTP response: %v", err)
		}
		return nil, fmt.Errorf("%w: %s", errTransientHTTP, resp.Status)
	default:
		return resp, nil
	}
}

// queryWithRetry retries the entire read-only operation, including response decoding.
func queryWithRetry[T any](ctx context.Context, client *githubv4.Client, result *T, variables map[string]interface{}, options ...retry.Option) error {
	settings := []retry.Option{
		retry.Attempts(5),
		retry.Delay(2 * time.Second),
		retry.MaxDelay(30 * time.Second),
		retry.LastErrorOnly(true),
		retry.Context(ctx),
		retry.RetryIf(func(err error) bool {
			if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
				return false
			}
			if errors.Is(err, errTransientHTTP) || errors.Is(err, io.EOF) || errors.Is(err, io.ErrUnexpectedEOF) {
				return true
			}
			for cause := err; cause != nil; cause = errors.Unwrap(cause) {
				if networkError, ok := cause.(net.Error); ok && networkError.Timeout() {
					return true
				}
			}
			return false
		}),
		retry.OnRetry(func(attempt uint, err error) {
			log.Warnf("GitHub GraphQL query attempt %d failed: %v", attempt+1, err)
		}),
	}
	return retry.Do(func() error {
		if err := ctx.Err(); err != nil {
			return err
		}
		// Decoding can partially populate the destination before returning an error.
		var fresh T
		if err := client.Query(ctx, &fresh, variables); err != nil {
			return err
		}
		*result = fresh
		return nil
	}, append(settings, options...)...)
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
	httpClient.Transport = &githubTransport{inner: httpClient.Transport}

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

	err := queryWithRetry(context.Background(), c.inner, &mergedQuery, variables)
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

	err := queryWithRetry(context.Background(), c.inner, &query, variables)
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

	err := queryWithRetry(context.Background(), c.inner, &mergedQuery, variables)
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

	err := queryWithRetry(ctx, c.inner, &query, nil)
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
