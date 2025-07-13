package constants

const (
	ScoreRepoStar    float64 = 0.01
	ScoreRepoFork    float64 = 0.03
	ScoreRepoCommit  float64 = 0.1
	ScoreRepoComment float64 = 1
	ScoreRepoIssue   float64 = 2
	ScoreRepoOpenPR  float64 = 3
	ScoreRepoReview  float64 = 4
	ScoreRepoMergePR float64 = 5
)

const (
	ScoreContributionComment float64 = 1
	ScoreContributionIssue   float64 = 2
	ScoreContributionOpenPR  float64 = 3
	ScoreContributionReview  float64 = 4
	ScoreContributionMergePR float64 = 5
)

const (
	ScoreStarred  float64 = 0.01
	ScoreFollower float64 = 0.1
)
