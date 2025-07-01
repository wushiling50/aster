package pack

import (
	"github.com/google/go-github/v66/github"

	"github.com/wushiling50/aster/pkg/model/repo"
)

func BuildRepo(githubRepo *github.Repository, issueCount int64, prCount int64, commitCount int64,
	openPrCount int64, mergedPrCount int64, commentCount int64, reviewCount int64,
	languages string) *repo.Repo {
	return &repo.Repo{
		Id:            githubRepo.GetID(),
		Name:          githubRepo.GetName(),
		StarCount:     int64(githubRepo.GetStargazersCount()),
		ForkCount:     int64(githubRepo.GetForksCount()),
		IssueCount:    issueCount,
		PrCount:       prCount,
		CommitCount:   commitCount,
		Language:      languages,
		Description:   githubRepo.GetDescription(),
		MergedPrCount: mergedPrCount,
		OpenPrCount:   openPrCount,
		CommentCount:  commentCount,
		ReviewCount:   reviewCount,
	}
}
