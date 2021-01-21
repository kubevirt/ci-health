package gh

import (
	"context"
	"io/ioutil"

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

func (c *Client) GetMergeQueueLength() int {
	_ = `query {
  search(query: "repo:kubevirt/kubevirt type:pr state:open label:lgtm label:approved -label:do-not-merge/hold ", type: ISSUE, last:100){
    issueCount
    pageInfo {
      hasNextPage
      startCursor
      endCursor
    }
    edges {
      node {
        ... on PullRequest {
          number
					baseRef {
            target {
              commitUrl
            }
          }
          comments(first:100, orderBy:{field:UPDATED_AT, direction:DESC}) {
            edges {
             	node {
                createdAt
                bodyText
              }
            }
          }
        }
      }
    }
  }
}
`
	return 0
}
