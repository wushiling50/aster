package github

import (
	"context"

	"github.com/google/go-github/v66/github"
	"github.com/wushiling50/aster/pkg/errno"
	"github.com/zeromicro/go-zero/core/logx"
)

func GetAllStarredReposByLogin(ctx context.Context, login string) (allRepos []*github.StarredRepository, err error) {
	var githubClient *github.Client = githubClientInit()
	opts := &github.ActivityListStarredOptions{
		ListOptions: github.ListOptions{PerPage: 100},
	}
	for {
		repos, resp, err := githubClient.Activity.ListStarred(ctx, login, opts)
		if err != nil {
			logx.Errorf("github.GetAllStarredReposByLogin: Fail To Fetching Starred Repo: %v", err.Error())
			err = errno.InternalGithubError.WithError(err)
			return nil, err
		}
		allRepos = append(allRepos, repos...)
		if resp.NextPage == 0 {
			break
		}
		opts.Page = resp.NextPage
	}
	return
}
