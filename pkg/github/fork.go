package github

import (
	"context"

	"github.com/google/go-github/v66/github"
	"github.com/wushiling50/aster/pkg/errno"
	"github.com/zeromicro/go-zero/core/logx"
)

func GetAllForksByRepo(ctx context.Context, login string, repoName string) (allForks []*github.Repository, err error) {
	var githubClient *github.Client = githubClientInit()
	opts := &github.RepositoryListForksOptions{ListOptions: github.ListOptions{PerPage: 100}}
	for {
		repos, resp, err := githubClient.Repositories.ListForks(ctx, login, repoName, opts)
		if err != nil {
			logx.Errorf("github.GetAllForksByRepo: Fail To Fetching Fork: %v", err.Error())
			err = errno.InternalGithubError.WithError(err)
			return nil, err
		}
		allForks = append(allForks, repos...)
		if resp.NextPage == 0 {
			break
		}
		opts.Page = resp.NextPage
	}

	return
}
