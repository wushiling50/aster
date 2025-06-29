package github

import (
	"context"

	"github.com/google/go-github/v66/github"
	"github.com/wushiling50/aster/pkg/errno"
	"github.com/zeromicro/go-zero/core/logx"
)

func GetAllFollowersByLogin(ctx context.Context, login string) (allFollowers []*github.User, err error) {
	var githubClient *github.Client = githubClientInit()
	opts := &github.ListOptions{PerPage: 100}
	for {
		followers, resp, err := githubClient.Users.ListFollowers(ctx, login, opts)
		if err != nil {
			logx.Errorf("Unexpected Error When Fetching Follower: %v" + err.Error())
			err = errno.InternalGithubError.WithError(err)
			return nil, err
		}
		allFollowers = append(allFollowers, followers...)
		if resp.NextPage == 0 {
			break
		}
		opts.Page = resp.NextPage
	}

	return
}

func GetAllFollowingByLogin(ctx context.Context, login string) (allFollowing []*github.User, err error) {
	var githubClient *github.Client = githubClientInit()
	opts := &github.ListOptions{PerPage: 100}
	for {
		following, resp, err := githubClient.Users.ListFollowing(ctx, login, opts)
		if err != nil {
			logx.Errorf("Unexpected Error When Fetching Following: %v" + err.Error())
			err = errno.InternalGithubError.WithError(err)
			return nil, err
		}
		allFollowing = append(allFollowing, following...)
		if resp.NextPage == 0 {
			break
		}
		opts.Page = resp.NextPage
	}

	return
}
