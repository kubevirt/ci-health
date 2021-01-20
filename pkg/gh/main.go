package gh

import (
	"context"
	"io/ioutil"

	"github.com/google/go-github/v33/github"
	"golang.org/x/oauth2"
)

func GetClient(tokenPath string) (*github.Client, error) {
	ctx := context.Background()
	token, err := ioutil.ReadFile(tokenPath)
	if err != nil {
		return nil, err
	}
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: string(token)},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	return client, nil
}
