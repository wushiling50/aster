package constants

import "time"

const (
	GithubAPIToken string = "GITHUB_API_TOKEN"

	DefaultSearchLimit int64         = 150
	DataExpiredTime    time.Duration = 24 * time.Hour
)

const (
	RoleAuthor    = "author"
	RoleCommenter = "commenter"
	RoleReviewer  = "reviewed-by"
)

const (
	SearchSort  = "updated"
	SearchOrder = "desc"
)
