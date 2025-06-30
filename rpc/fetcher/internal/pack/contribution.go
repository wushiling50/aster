package pack

import (
	"time"

	"github.com/google/go-github/v66/github"
	"github.com/wushiling50/aster/pkg/constants"
	githubFunc "github.com/wushiling50/aster/pkg/github"
	"github.com/wushiling50/aster/pkg/model/contribution"
)

// Issue - PR
func BuildIssuePR(issuePR *github.Issue, developerId int64, category string, merged bool, repo *github.Repository) *contribution.Contribution {
	return &contribution.Contribution{
		DataCreatedAt:  time.Now(),
		DataUpdatedAt:  time.Now(),
		DeveloperId:    developerId,
		RepoId:         repo.GetID(),
		Category:       category,
		Content:        issuePR.GetTitle() + " " + issuePR.GetBody(),
		CreatedAt:      issuePR.GetCreatedAt().Time,
		UpdatedAt:      issuePR.GetUpdatedAt().Time,
		ContributionId: issuePR.GetID(),
	}
}

// Comment
func BuildComment(githubCommentWithRepoId *githubFunc.CommentWithRepoId, developerId int64) *contribution.Contribution {
	if githubCommentWithRepoId.IsIssueComment {
		return &contribution.Contribution{
			DataCreatedAt:  time.Now(),
			DataUpdatedAt:  time.Now(),
			DeveloperId:    developerId,
			RepoId:         githubCommentWithRepoId.RepoId,
			Category:       constants.CategoryComment,
			Content:        githubCommentWithRepoId.IssueComment.GetBody(),
			CreatedAt:      githubCommentWithRepoId.IssueComment.GetCreatedAt().Time,
			UpdatedAt:      githubCommentWithRepoId.IssueComment.GetUpdatedAt().Time,
			ContributionId: githubCommentWithRepoId.IssueComment.GetID(),
		}
	}
	return &contribution.Contribution{
		DataCreatedAt:  time.Now(),
		DataUpdatedAt:  time.Now(),
		DeveloperId:    developerId,
		RepoId:         githubCommentWithRepoId.RepoId,
		Category:       constants.CategoryComment,
		Content:        githubCommentWithRepoId.PRComment.GetBody(),
		CreatedAt:      githubCommentWithRepoId.PRComment.GetCreatedAt().Time,
		UpdatedAt:      githubCommentWithRepoId.PRComment.GetUpdatedAt().Time,
		ContributionId: githubCommentWithRepoId.PRComment.GetID(),
	}
}

// Review
func BuildReview(githubReviewWithRepoId *githubFunc.ReviewWithRepoId, developerId int64) *contribution.Contribution {
	return &contribution.Contribution{
		DataCreatedAt:  time.Now(),
		DataUpdatedAt:  time.Now(),
		DeveloperId:    developerId,
		RepoId:         githubReviewWithRepoId.RepoId,
		Category:       constants.CategoryReview,
		Content:        githubReviewWithRepoId.Review.GetBody(),
		CreatedAt:      githubReviewWithRepoId.Review.GetSubmittedAt().Time,
		UpdatedAt:      githubReviewWithRepoId.Review.GetSubmittedAt().Time,
		ContributionId: githubReviewWithRepoId.Review.GetID(),
	}
}

func BuildCompletedContribution(dataId int, developerId int64) *contribution.Contribution {
	return &contribution.Contribution{
		DataId:      int64(dataId),
		DeveloperId: developerId,
	}
}
