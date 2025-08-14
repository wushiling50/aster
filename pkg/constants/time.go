package constants

import "time"

const (
	ONE_SECOND = time.Second
	ONE_MINUTE = time.Minute
	ONE_HOUR   = time.Hour
	ONE_DAY    = 24 * ONE_HOUR
	ONE_WEEK   = 7 * ONE_DAY

	DataExpiredTime time.Duration = 24 * time.Hour
)
