package pack

import (
	"github.com/wushiling50/aster/gen/repo"
	model_repo "github.com/wushiling50/aster/pkg/model/repo"
)

func BuildModelRepo(repo *repo.Repo) *model_repo.Repo {
	return &model_repo.Repo{
		DataId:                  repo.DataId,
		Id:                      repo.Id,
		Name:                    repo.Name,
		StarCount:               repo.StarCount,
		ForkCount:               repo.ForkCount,
		IssueCount:              repo.IssueCount,
		CommitCount:             repo.CommitCount,
		PrCount:                 repo.PrCount,
		Language:                repo.Language,
		Description:             repo.Description,
		LastFetchForkAt:         repo.LastFetchForkAt,
		LastFetchContributionAt: repo.LastFetchContributionAt,
		MergedPrCount:           repo.MergedPrCount,
		OpenPrCount:             repo.OpenPrCount,
		CommentCount:            repo.CommentCount,
		ReviewCount:             repo.ReviewCount,
	}
}

func BuildGenRepo(modelRepo *model_repo.Repo) *repo.Repo {
	return &repo.Repo{
		DataId:                  modelRepo.DataId,
		Id:                      modelRepo.Id,
		Name:                    modelRepo.Name,
		StarCount:               modelRepo.StarCount,
		ForkCount:               modelRepo.ForkCount,
		IssueCount:              modelRepo.IssueCount,
		CommitCount:             modelRepo.CommitCount,
		PrCount:                 modelRepo.PrCount,
		Language:                modelRepo.Language,
		Description:             modelRepo.Description,
		LastFetchForkAt:         modelRepo.LastFetchForkAt,
		LastFetchContributionAt: modelRepo.LastFetchContributionAt,
		MergedPrCount:           modelRepo.MergedPrCount,
		OpenPrCount:             modelRepo.OpenPrCount,
		CommentCount:            modelRepo.CommentCount,
		ReviewCount:             modelRepo.ReviewCount,
	}
}
