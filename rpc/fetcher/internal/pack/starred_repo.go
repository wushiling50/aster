package pack

import (
	"github.com/google/go-github/v66/github"
	"github.com/wushiling50/aster/pkg/model/relation"
)

func BuildStarredRepo(githubStarredRepo *github.StarredRepository, userId int64) *relation.Star {
	return &relation.Star{
		DeveloperId: userId,
		RepoId:      githubStarredRepo.Repository.GetID(),
	}
}

func BuildCompletedStarredRepo(dataId int, userId int64) *relation.CreateRepo {
	return &relation.CreateRepo{
		DataId:      int64(dataId),
		DeveloperId: userId,
	}
}
