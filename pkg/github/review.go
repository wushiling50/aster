package github

import (
	"context"
	"net/http"

	"github.com/google/go-github/v66/github"
	"github.com/wushiling50/aster/pkg/constants"
	"github.com/wushiling50/aster/pkg/errno"
	"github.com/zeromicro/go-zero/core/logx"
)

func GetAllReviewByLogin(ctx context.Context, login string, createAfter string, searchLimit int64) (allReviewWithRepoId []*ReviewWithRepoId, repos map[int64]*github.Repository, err error) {
	var allIssue []*github.Issue
	var githubClient *github.Client = githubClientInit()
	allIssue, err = GetAllIssuePRByLogin(ctx, login, constants.RoleReviewer, createAfter, searchLimit)

	opts := &github.ListOptions{
		PerPage: 100,
	}

	var prResp *github.Response
	repos = make(map[int64]*github.Repository)

	for _, issue := range allIssue {
		if len(allReviewWithRepoId) >= int(searchLimit) {
			break
		}

		if issue.IsPullRequest() {
			var prReviews []*github.PullRequestReview

			var repo *github.Repository
			if repo, err = GetRepoByUrl(ctx, issue.GetRepositoryURL()); err != nil {
				logx.Error(err)
				return
			}
			repos[repo.GetID()] = repo

			prReviews, prResp, err = githubClient.PullRequests.ListReviews(ctx, repo.GetOwner().GetLogin(), repo.GetName(), issue.GetNumber(), opts)
			if err != nil && (prResp == nil || prResp.StatusCode != http.StatusNotFound) {
				logx.Errorf("github.GetAllReviewByLogin: Fail To Fetching PRs: %v", err.Error())
				err = errno.InternalGithubError.WithError(err)
				return
			}

			for _, review := range prReviews {
				if review.GetUser().GetLogin() == login {
					allReviewWithRepoId = append(allReviewWithRepoId,
						&ReviewWithRepoId{
							Review: review,
							RepoId: repo.GetID(),
						})
				}
			}
		}
	}

	if len(allReviewWithRepoId) > int(searchLimit) {
		allReviewWithRepoId = allReviewWithRepoId[:int(searchLimit)]
	}

	return
}
