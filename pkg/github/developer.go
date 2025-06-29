package github

import (
	"context"

	"github.com/google/go-github/v66/github"
	"github.com/wushiling50/aster/pkg/errno"
	"github.com/zeromicro/go-zero/core/logx"
)

func GetStarredRepoCountByLogin(ctx context.Context, login string) (starredRepoCount int64, err error) {
	var githubClient *github.Client = githubClientInit()
	var githubStarredRepoResp *github.Response
	opts := &github.ActivityListStarredOptions{
		ListOptions: github.ListOptions{PerPage: 1},
	}

	if _, githubStarredRepoResp, err = githubClient.Activity.ListStarred(ctx, login, opts); err != nil {
		logx.Errorf("github.GetStarredRepoCountByLogin: Fail To Fetching Starred Repo: %v", err.Error())
		err = errno.InternalGithubError.WithError(err)
		return
	}

	starredRepoCount = int64(githubStarredRepoResp.LastPage)

	return
}
