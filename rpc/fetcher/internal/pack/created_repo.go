package pack

import (
	"github.com/google/go-github/v66/github"
	"github.com/wushiling50/aster/pkg/model/relation"
)

func BuildCreatedRepo(githubRepo *github.Repository, userId int64) *relation.CreateRepo {
	return &relation.CreateRepo{
		DeveloperId: userId,
		RepoId:      githubRepo.GetID(),
	}
}

func BuildCompletedCreatedRepo(dataId int, userId int64) *relation.CreateRepo {
	return &relation.CreateRepo{
		DataId:      int64(dataId),
		DeveloperId: userId,
	}
}
