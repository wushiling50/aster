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

func BuildCompletedCreatedRepo(dataId int64, developerId int64) *relation.CreateRepo {
	return &relation.CreateRepo{
		DataId:      dataId,
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

func BuildCompletedStarredRepo(dataId int64, developerId int64) *relation.CreateRepo {
	return &relation.CreateRepo{
		DataId:      dataId,
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

func BuildCompletedFollow(dataId int64, developerId int64) *relation.Follow {
	return &relation.Follow{
		DataId:      dataId,
		FollowerId:  developerId,
		FollowingId: 0,
	}
}
