package gh

import "time"

type PageInfo struct {
	HasNextPage bool
	EndCursor   string
}

type Comment struct {
	CreatedAt time.Time
	BodyText  string
}

type PullRequestFragment struct {
	Number   int
	Comments struct {
		Nodes []Comment
	} `graphql:"comments(first: 100, orderBy: {field: UPDATED_AT, direction: DESC})"`
}
