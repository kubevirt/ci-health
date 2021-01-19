package ghclient

import "github.com/fgimenez/cihealth/pkg/stats"

type GHClient struct{}

func New() *GHClient {
	return &GHClient{}
}

func (client *GHClient) Run(token string) (*stats.Results, error) {
	results := &stats.Results{}

	return results, nil
}
