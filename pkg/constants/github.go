package constants

import "time"

const (
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
