package pack

import (
	"github.com/google/go-github/v66/github"
	"github.com/wushiling50/aster/gen/relation"
)

// CreatedRepo
func BuildCreatedRepo(githubRepo *github.Repository, developerId int64) *relation.CreateRepo {
	return &relation.CreateRepo{
		DeveloperId: developerId,
		RepoId:      githubRepo.GetID(),
	}
}

func BuildCompletedCreatedRepo(dataId int, developerId int64) *relation.CreateRepo {
	return &relation.CreateRepo{
		DataId:      int64(dataId),
		DeveloperId: developerId,
	}
}

// Star
func BuildStarredRepo(githubStarredRepo *github.StarredRepository, developerId int64) *relation.Star {
	return &relation.Star{
		DeveloperId: developerId,
		RepoId:      githubStarredRepo.Repository.GetID(),
	}
}

func BuildCompletedStarredRepo(dataId int, developerId int64) *relation.CreateRepo {
	return &relation.CreateRepo{
		DataId:      int64(dataId),
		DeveloperId: developerId,
	}
}

// Follow
func BuildFollow(followerId int64, followingId int64) *relation.Follow {
	return &relation.Follow{
		FollowerId:  followerId,
		FollowingId: followingId,
	}
}

func BuildCompletedFollow(dataId int, developerId int64) *relation.Follow {
	return &relation.Follow{
		DataId:      int64(dataId),
		FollowerId:  developerId,
		FollowingId: 0,
	}
}

// Fork
func BuildFork(originalRepoId int64, forkRepoId int64) *relation.Fork {
	return &relation.Fork{
		OriginalRepoId: originalRepoId,
		ForkRepoId:     forkRepoId,
	}
}

func BuildCompletedFork(dataId int, originalRepoId int64) *relation.Fork {
	return &relation.Fork{
		DataId:         int64(dataId),
		OriginalRepoId: originalRepoId,
	}
}
