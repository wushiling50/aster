package github

import (
	"context"

	"github.com/google/go-github/v66/github"
	"github.com/wushiling50/aster/pkg/constants"
	"github.com/wushiling50/aster/pkg/errno"
	"github.com/zeromicro/go-zero/core/logx"
)

func GetAllIssuePRByLogin(ctx context.Context, login string, role string, createAfter string, searchLimit int64) (issues []*github.Issue, err error) {
	var githubClient *github.Client = githubClientInit()
	var issueSearchResult *github.IssuesSearchResult

	opts := &github.SearchOptions{
		Sort:        constants.SearchSort,
		Order:       constants.SearchOrder,
		ListOptions: github.ListOptions{PerPage: int(searchLimit), Page: 1},
	}

	issueSearchResult, _, err = githubClient.Search.Issues(ctx, role+":"+login+" created:>="+createAfter, opts)
	if err != nil {
		logx.Errorf("github.GetAllIssuePRByLogin: Fail To Fetching Issue: %v", err.Error())
		err = errno.InternalGithubError.WithError(err)
		return
	}

	issues = issueSearchResult.Issues

	return
}

func CheckIfMerged(ctx context.Context, issuePR *github.Issue, repo *github.Repository) (merged bool, err error) {
	var githubClient *github.Client = githubClientInit()
	if issuePR.IsPullRequest() {
		pr, _, err := githubClient.PullRequests.Get(ctx, repo.GetOwner().GetLogin(), repo.GetName(), issuePR.GetNumber())
		if err != nil {
			logx.Errorf("github.CheckIfMerged: Fail To Fetching PR: %v", err.Error())
			err = errno.InternalGithubError.WithError(err)
			return false, err
		}

		merged = pr.GetMerged()
	}
	return
}
