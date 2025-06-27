package constants

import "time"

const Separator = "|"

const (
	APITaskExpireTime = time.Minute * 10
	APIMaxRetry       = 10
	APIRetryDelay     = time.Second * 10
)

const (
	APITaskName  = "api"
	APITaskQueue = "api"
)

const (
	APIGetDeveloper int = iota
	APIGetLanguage
	APIGetScore
	APIGetNation
	APIGetSummary
)
