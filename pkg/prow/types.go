package prow

import "time"

type Started struct {
	Timestamp int64  `json:"timestamp"`
	Pull      string `json:"pull"`
	Repos     struct {
		repo string `json:"kubevirt/kubevirt"`
	} `json:"repos"`
	RepoCommit  string `json:"repo-commit"`
	RepoVersion string `json:"repo-version"`
}

func (s Started) Time() time.Time {
	return time.Unix(s.Timestamp, 0)
}

type Finished struct {
	Timestamp int64  `json:"timestamp"`
	Passed    bool   `json:"passed"`
	Result    string `json:"result"`
	Revision  string `json:"revision"`
}

func (f Finished) Time() time.Time {
	return time.Unix(f.Timestamp, 0)
}
