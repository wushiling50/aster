package pack

import (
	"github.com/google/go-github/v66/github"
	"github.com/wushiling50/aster/gen/contribution"
	"github.com/wushiling50/aster/pkg/constants"
	githubFunc "github.com/wushiling50/aster/pkg/github"
)

// Issue - PR
func BuildIssuePR(issuePR *github.Issue, developerId int64, category string, merged bool, repo *github.Repository) *contribution.Contribution {
	return &contribution.Contribution{
		DeveloperId:           developerId,
		RepoId:                repo.GetID(),
		Category:              category,
		Content:               issuePR.GetTitle() + " " + issuePR.GetBody(),
		ContributionCreatedAt: issuePR.GetCreatedAt().Unix(),
		ContributionUpdatedAt: issuePR.GetUpdatedAt().Unix(),
		ContributionId:        issuePR.GetID(),
	}
}

// Comment
func BuildComment(githubCommentWithRepoId *githubFunc.CommentWithRepoId, developerId int64) *contribution.Contribution {
	if githubCommentWithRepoId.IsIssueComment {
		return &contribution.Contribution{
			DeveloperId:           developerId,
			RepoId:                githubCommentWithRepoId.RepoId,
			Category:              constants.CategoryComment,
			Content:               githubCommentWithRepoId.IssueComment.GetBody(),
			ContributionCreatedAt: githubCommentWithRepoId.IssueComment.GetCreatedAt().Unix(),
			ContributionUpdatedAt: githubCommentWithRepoId.IssueComment.GetUpdatedAt().Unix(),
			ContributionId:        githubCommentWithRepoId.IssueComment.GetID(),
		}
	}
	return &contribution.Contribution{
		DeveloperId:           developerId,
		RepoId:                githubCommentWithRepoId.RepoId,
		Category:              constants.CategoryComment,
		Content:               githubCommentWithRepoId.PRComment.GetBody(),
		ContributionCreatedAt: githubCommentWithRepoId.PRComment.GetCreatedAt().Unix(),
		ContributionUpdatedAt: githubCommentWithRepoId.PRComment.GetUpdatedAt().Unix(),
		ContributionId:        githubCommentWithRepoId.PRComment.GetID(),
	}
}

// Review
func BuildReview(githubReviewWithRepoId *githubFunc.ReviewWithRepoId, developerId int64) *contribution.Contribution {
	return &contribution.Contribution{
		DeveloperId:           developerId,
		RepoId:                githubReviewWithRepoId.RepoId,
		Category:              constants.CategoryReview,
		Content:               githubReviewWithRepoId.Review.GetBody(),
		ContributionCreatedAt: githubReviewWithRepoId.Review.GetSubmittedAt().Unix(),
		ContributionUpdatedAt: githubReviewWithRepoId.Review.GetSubmittedAt().Unix(),
		ContributionId:        githubReviewWithRepoId.Review.GetID(),
	}
}

func BuildCompletedContribution(dataId int64, developerId int64) *contribution.Contribution {
	return &contribution.Contribution{
		DataId:      dataId,
		DeveloperId: developerId,
	}
}
