package github

import (
	"context"
	"fmt"

	"github.com/google/go-github/v66/github"

	"github.com/wushiling50/aster/pkg/errno"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
)

func GetRepo(ctx context.Context, repoId int64) (githubRepo *github.Repository, githubResp *github.Response, err error) {
	var githubClient *github.Client = githubClientInit()
	if githubRepo, githubResp, err = githubClient.Repositories.GetByID(ctx, repoId); err != nil {
		logx.Errorf("github.GetRepo: Fail To Fetching Repo: %v", err.Error())
		err = errno.InternalGithubError.WithError(err)
		return
	}

	logx.Infof("Successfully Get Repo")
	return
}

func GetIssuePrCountByRepo(ctx context.Context, owner string, name string) (issueCount int64, prCount int64, err error) {
	var (
		githubClient    *github.Client = githubClientInit()
		githubPrResp    *github.Response
		githubIssueResp *github.Response
	)

	prOpts := &github.PullRequestListOptions{
		ListOptions: github.ListOptions{PerPage: 1},
	}

	issueOpts := &github.IssueListByRepoOptions{
		ListOptions: github.ListOptions{PerPage: 1},
	}

	if _, githubPrResp, err = githubClient.PullRequests.List(ctx, owner, name, prOpts); err != nil {
		logx.Errorf("github.GetIssuePrCountByRepo: Fail To Fetching PRs: %v", err.Error())
		err = errno.InternalGithubError.WithError(err)
		return
	}

	if _, githubIssueResp, err = githubClient.Issues.ListByRepo(ctx, owner, name, issueOpts); err != nil {
		logx.Errorf("github.GetIssuePrCountByRepo: Fail To Fetching Issues: %v", err.Error())
		err = errno.InternalGithubError.WithError(err)
		return
	}

	prCount = int64(githubPrResp.LastPage)
	issueCount = int64(githubIssueResp.LastPage) - prCount

	return
}

func GetCommitCountByRepo(ctx context.Context, owner string, name string) (commitCount int64, err error) {
	var githubCommitResp *github.Response
	var githubClient *github.Client = githubClientInit()
	opts := &github.CommitsListOptions{
		ListOptions: github.ListOptions{PerPage: 1},
	}
	if _, githubCommitResp, err = githubClient.Repositories.ListCommits(ctx, owner, name, opts); err != nil {
		logx.Errorf("github.GetCommitCountByRepo: Fail To Fetching Commit: %v", err.Error())
		err = errno.InternalGithubError.WithError(err)
		return
	}

	commitCount = int64(githubCommitResp.LastPage)

	return
}

func GetOpenPrCountByRepo(ctx context.Context, owner string, name string) (openPrCount int64, err error) {
	var githubPrResp *github.Response
	var githubClient *github.Client = githubClientInit()
	opts := &github.PullRequestListOptions{
		State:       "open",
		ListOptions: github.ListOptions{PerPage: 1},
	}
	if _, githubPrResp, err = githubClient.PullRequests.List(ctx, owner, name, opts); err != nil {
		logx.Errorf("github.GetOpenPrCountByRepo: Fail To Fetching OpenPrs: %v", err.Error())
		err = errno.InternalGithubError.WithError(err)
		return
	}

	openPrCount = int64(githubPrResp.LastPage)

	return
}

func GetMergedPrCountByRepo(ctx context.Context, owner string, name string) (mergedPrCount int64, err error) {
	var githubPrResp *github.Response
	var githubClient *github.Client = githubClientInit()
	opts := &github.SearchOptions{
		ListOptions: github.ListOptions{PerPage: 1},
	}
	query := fmt.Sprintf("repo:%s/%s is:pr is:merged", owner, name)
	if _, githubPrResp, err = githubClient.Search.Issues(ctx, query, opts); err != nil {
		logx.Errorf("github.GetMergedPrCountByRepo: Fail To Fetching MergedPrs: %v", err.Error())
		err = errno.InternalGithubError.WithError(err)
		return
	}

	mergedPrCount = int64(githubPrResp.LastPage)

	return
}

func GetCommentCountByRepo(ctx context.Context, owner string, name string) (commentCount int64, err error) {
	var githubIssueResp *github.Response
	var githubClient *github.Client = githubClientInit()
	opts := &github.IssueListCommentsOptions{
		ListOptions: github.ListOptions{PerPage: 1},
	}
	if _, githubIssueResp, err = githubClient.Issues.ListComments(ctx, owner, name, 0, opts); err != nil {
		logx.Errorf("github.GetCommentCountByRepo: Fail To Fetching Comment: %v", err.Error())
		err = errno.InternalGithubError.WithError(err)
		return
	}

	commentCount = int64(githubIssueResp.LastPage)

	return
}

func GetReviewCountByRepo(ctx context.Context, owner string, name string) (reviewCount int64, err error) {
	var githubPrResp *github.Response
	var githubClient *github.Client = githubClientInit()
	opts := &github.PullRequestListCommentsOptions{
		ListOptions: github.ListOptions{PerPage: 1},
	}
	if _, githubPrResp, err = githubClient.PullRequests.ListComments(ctx, owner, name, 0, opts); err != nil {
		logx.Errorf("github.GetReviewCountByRepo: Fail To Fetching Review: %v", err.Error())
		err = errno.InternalGithubError.WithError(err)
		return
	}

	reviewCount = int64(githubPrResp.LastPage)

	return
}

func GetLanguagesByRepo(ctx context.Context, owner string, name string) (languages string, err error) {
	var githubLanguages map[string]int
	var githubClient *github.Client = githubClientInit()

	if githubLanguages, _, err = githubClient.Repositories.ListLanguages(ctx, owner, name); err != nil {
		logx.Errorf("github.GetLanguagesByRepo: Fail To Fetching Languages: %v", err.Error())
		err = errno.InternalGithubError.WithError(err)
		return
	}

	if languages, err = jsonx.MarshalToString(githubLanguages); err != nil {
		logx.Errorf("Unexpected error when marshalling languages: %v", err.Error())
		err = errno.InternalJSONError.WithError(err)
		return
	}

	return
}
