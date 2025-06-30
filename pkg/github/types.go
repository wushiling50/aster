package github

import "github.com/google/go-github/v66/github"

type CommentWithRepoId struct {
	IsIssueComment bool
	IssueComment   *github.IssueComment
	PRComment      *github.PullRequestComment
	RepoId         int64
}

type ReviewWithRepoId struct {
	Review *github.PullRequestReview
	RepoId int64
}
