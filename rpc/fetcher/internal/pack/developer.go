package pack

import (
	"github.com/google/go-github/v66/github"
	"github.com/wushiling50/aster/pkg/model/developer"
)

func BuildDeveloperProfile(githubUser *github.User, starredRepoCount int64) *developer.Developer {
	return &developer.Developer{
		Id:              githubUser.GetID(),
		Name:            githubUser.GetName(),
		Login:           githubUser.GetLogin(),
		AvatarUrl:       githubUser.GetAvatarURL(),
		Company:         githubUser.GetCompany(),
		Location:        githubUser.GetLocation(),
		Bio:             githubUser.GetBio(),
		Blog:            githubUser.GetBlog(),
		Email:           githubUser.GetEmail(),
		CreatedAt:       githubUser.GetCreatedAt().Time,
		UpdatedAt:       githubUser.GetUpdatedAt().Time,
		TwitterUsername: githubUser.GetTwitterUsername(),
		Repos:           int64(githubUser.GetPublicRepos()),
		Following:       int64(githubUser.GetFollowing()),
		Followers:       int64(githubUser.GetFollowers()),
		Gists:           int64(githubUser.GetPublicGists()),
		Stars:           starredRepoCount,
	}
}
