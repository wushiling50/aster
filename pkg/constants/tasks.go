package constants

import "time"

const Separator = "|"

// ------ APITask ------
const (
	APITaskExpireTime = time.Minute * 10
	APIMaxRetry       = 10
	APIRetryDelay     = time.Second * 10
	APIConcurrency    = 20
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

// ------ FetcherTask ------
const (
	FetchMaxRetry    = 3
	FetchRetryDelay  = time.Second * 5
	FetchConcurrency = 20
)

const (
	FetcherTaskName  = "fetch"
	FetcherTaskQueue = "fetch"
)

const (
	FetchDeveloper int = iota
	FetchCreatedRepo
	FetchStarredRepo

	FetchFollowing
	FetchFollower

	FetchIssuePROfUser
	FetchCommentOfUser
	FetchReviewOfUser

	FetchRepo
	FetchFork
)

const (
	FetchCreatedRepoCompletedDataId int = -iota - 1

	FetchStarredRepoCompletedDataId
	FetchStarringDeveloperCompletedDataId

	FetchFollowingCompletedDataId
	FetchFollowerCompletedDataId

	FetchIssuePROfUserCompletedDataId
	FetchCommentOfUserCompletedDataId
	FetchReviewOfUserCompletedDataId

	FetchForkCompletedDataId
)
