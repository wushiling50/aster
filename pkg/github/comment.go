package github

import (
	"context"
	"net/http"
	"time"

	"github.com/google/go-github/v66/github"
	"github.com/wushiling50/aster/pkg/constants"
	"github.com/wushiling50/aster/pkg/errno"
	"github.com/zeromicro/go-zero/core/logx"
)

func GetAllCommentByLogin(ctx context.Context, login string, createAfter string, searchLimit int64) (allCommentWithRepoId []*CommentWithRepoId, repos map[int64]*github.Repository, err error) {
	var allIssue []*github.Issue
	var githubClient *github.Client = githubClientInit()
	allIssue, err = GetAllIssuePRByLogin(ctx, login, constants.RoleCommenter, createAfter, searchLimit)

	updated := constants.SearchSort
	desc := constants.SearchOrder
	createAfterTime, _ := time.Parse("2006-01-02", createAfter)

	issueOpts := &github.IssueListCommentsOptions{
		Sort:        &updated,
		Direction:   &desc,
		Since:       &createAfterTime,
		ListOptions: github.ListOptions{PerPage: int(searchLimit)},
	}

	prOpts := &github.PullRequestListCommentsOptions{
		Sort:        updated,
		Direction:   desc,
		Since:       createAfterTime,
		ListOptions: github.ListOptions{PerPage: int(searchLimit)},
	}

	var issueResp *github.Response
	var prResp *github.Response
	repos = make(map[int64]*github.Repository)

	// 处理 PR 评论
	for _, issue := range allIssue {
		if len(allCommentWithRepoId) >= int(searchLimit) {
			break
		}

		var repo *github.Repository
		if repo, err = GetRepoByUrl(ctx, issue.GetRepositoryURL()); err != nil {
			logx.Error(err)
			return
		}
		repos[repo.GetID()] = repo

		if issue.IsPullRequest() {
			var prComments []*github.PullRequestComment
			prComments, prResp, err = githubClient.PullRequests.ListComments(ctx, repo.GetOwner().GetLogin(), repo.GetName(), 0, prOpts)
			if err != nil && (prResp == nil || prResp.StatusCode != http.StatusNotFound) {
				logx.Errorf("github.GetAllCommentByLogin: Fail To Fetching PR: %v", err.Error())
				err = errno.InternalGithubError.WithError(err)
				return
			}

			for _, comment := range prComments {
				if comment.User.GetLogin() == login {
					allCommentWithRepoId = append(allCommentWithRepoId,
						&CommentWithRepoId{
							IsIssueComment: false,
							PRComment:      comment,
							RepoId:         repo.GetID(),
						})
				}
			}
		}
	}

	// 处理 Issue 评论
	for _, issue := range allIssue {
		if len(allCommentWithRepoId) >= int(searchLimit) {
			break
		}

		var repo *github.Repository
		if repo, err = GetRepoByUrl(ctx, issue.GetRepositoryURL()); err != nil {
			logx.Error(err)
			return
		}
		repos[repo.GetID()] = repo

		var issueComments []*github.IssueComment
		issueComments, issueResp, err = githubClient.Issues.ListComments(ctx, repo.GetOwner().GetLogin(), repo.GetName(), issue.GetNumber(), issueOpts)
		if err != nil && (issueResp == nil || issueResp.StatusCode != http.StatusNotFound) {
			logx.Errorf("github.GetAllCommentByLogin: Fail To Fetching Issue: %v", err.Error())
			err = errno.InternalGithubError.WithError(err)
			return
		}

		for _, comment := range issueComments {
			if comment.User.GetLogin() == login {
				allCommentWithRepoId = append(allCommentWithRepoId,
					&CommentWithRepoId{
						IsIssueComment: true,
						IssueComment:   comment,
						RepoId:         repo.GetID(),
					})
			}
		}
	}

	if len(allCommentWithRepoId) > int(searchLimit) {
		allCommentWithRepoId = allCommentWithRepoId[:int(searchLimit)]
	}

	return
}
