package logic

import (
	"context"
	"errors"
	"math"
	"strconv"

	"github.com/wushiling50/aster/gen/analysis"
	"github.com/wushiling50/aster/gen/contribution"
	"github.com/wushiling50/aster/gen/relation"
	"github.com/wushiling50/aster/gen/repo"
	"github.com/wushiling50/aster/pkg/constants"
	"github.com/wushiling50/aster/pkg/errno"
	"github.com/wushiling50/aster/pkg/github"
	model_analysis "github.com/wushiling50/aster/pkg/model/analysis"
	"github.com/wushiling50/aster/pkg/tasks"
	"github.com/wushiling50/aster/pkg/utils"
	"github.com/wushiling50/aster/rpc/analysis/internal/pack"
	"github.com/wushiling50/aster/rpc/analysis/internal/svc"
	"golang.org/x/sync/errgroup"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateScoreLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateScoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateScoreLogic {
	return &UpdateScoreLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateScoreLogic) UpdateScore(in *analysis.UpdateAnalysisReq) (*analysis.UpdateAnalysisResp, error) {
	resp := new(analysis.UpdateAnalysisResp)

	var (
		contributionsCategorizedByRepoId = make(map[int64][]*contribution.Contribution)
		developerScore                   float64
		totalScore                       float64

		eg errgroup.Group
	)

	need, err := l.checkIfNeedUpdate(in.DeveloperId)
	if err != nil {
		logx.Errorf("service.UpdateScore: Update Score Failed: %w", err)
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	if !need {
		resp.Base = pack.BuildSuccessResp()
		return resp, nil
	}

	err = l.pushContributionTask(in.DeveloperId)
	if err != nil {
		logx.Errorf("service.UpdateScore: Failed To Enqueue Task: %v", err.Error())
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	contributions, err := l.rpcGetContributinoById(in.DeveloperId, 1000, 1)
	if err != nil {
		logx.Error(err)
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	for _, theContribution := range contributions {
		contributionsCategorizedByRepoId[theContribution.RepoId] = append(contributionsCategorizedByRepoId[theContribution.RepoId], theContribution)
	}

	for repoId, contributions := range contributionsCategorizedByRepoId {
		var (
			repoScore         float64
			contributionScore float64

			theRepo *repo.Repo

			contributionComment float64
			contributionIssue   float64
			contributionOpenPr  float64
			contributionReview  float64
			contributionMergePr float64
		)

		// repo Score
		err = l.pushRepoTask(repoId)
		if err != nil {
			logx.Errorf("service.UpdateScore: Failed To Enqueue Task: %v", err.Error())
			err = errno.InternalAsynqError.WithError(err)
			resp.Base = pack.BuildBaseResp(err)
			return resp, nil
		}

		theRepo, err := l.rpcGetRepoById(repoId)
		if err != nil {
			logx.Error(err)
			resp.Base = pack.BuildBaseResp(err)
			return resp, nil
		}

		repoScore = constants.ScoreRepoStar*float64(theRepo.StarCount) +
			constants.ScoreRepoFork*float64(theRepo.ForkCount) +
			constants.ScoreRepoCommit*float64(theRepo.CommitCount) +
			constants.ScoreRepoComment*float64(theRepo.CommentCount) +
			constants.ScoreRepoIssue*float64(theRepo.IssueCount) +
			constants.ScoreRepoOpenPR*float64(theRepo.OpenPrCount) +
			constants.ScoreRepoReview*float64(theRepo.ReviewCount) +
			constants.ScoreRepoMergePR*float64(theRepo.MergedPrCount)

		repoScore = math.Sqrt(repoScore)

		// contribution Score
		for _, theContribution := range contributions {
			switch theContribution.Category {
			case constants.CategoryComment:
				contributionComment++
			case constants.CategoryIssue:
				contributionIssue++
			case constants.CategoryOpenPullRequest:
				contributionOpenPr++
			case constants.CategoryReview:
				contributionReview++
			case constants.CategoryMergePullRequest:
				contributionMergePr++
			}
		}

		contributionScore = constants.ScoreContributionComment*contributionComment/float64(theRepo.CommentCount) +
			constants.ScoreContributionIssue*contributionIssue/float64(theRepo.IssueCount) +
			constants.ScoreContributionOpenPR*contributionOpenPr/float64(theRepo.OpenPrCount) +
			constants.ScoreContributionReview*contributionReview/float64(theRepo.ReviewCount) +
			constants.ScoreContributionMergePR*contributionMergePr/float64(theRepo.MergedPrCount)

		totalScore += repoScore * contributionScore
	}

	// developer score
	err = l.pushFollowerTask(in.DeveloperId)
	if err != nil {
		logx.Errorf("service.UpdateScore: Failed To Enqueue Task: %v", err.Error())
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	followerCount, err := l.rpcGetFollowerCountById(in.DeveloperId)
	if err != nil {
		logx.Error(err)
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	err = l.pushStarredTask(in.DeveloperId)
	if err != nil {
		logx.Errorf("service.UpdateScore: Failed To Enqueue Task: %v", err.Error())
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	starredCount, err := l.rpcGetStarredCountById(in.DeveloperId)
	if err != nil {
		logx.Error(err)
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	developerScore += constants.ScoreFollower*float64(followerCount) +
		constants.ScoreStarred*float64(starredCount)

	developerScore = math.Sqrt(developerScore * 0.1)

	totalScore += developerScore

	// update score
	eg.Go(func() error {
		return l.updateScore(&model_analysis.Score{
			DeveloperId: in.DeveloperId,
			Score:       totalScore,
		})
	})

	eg.Go(func() error {
		return l.updateRankScore(in.DeveloperId, totalScore)
	})

	if err := eg.Wait(); err != nil {
		logx.Errorf("service.UpdateScore: Update Score Failed: %w", err)
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	return resp, nil
}

func (l *UpdateScoreLogic) checkIfNeedUpdate(developerId int64) (bool, error) {
	score, err := l.svcCtx.ScoreModel.FindOneByDeveloperId(l.ctx, developerId)
	if err != nil {
		switch {
		case errors.Is(err, model_analysis.ErrNotFound):
			return true, nil
		default:
			return false, err
		}
	}

	if github.CheckIfDataExpired(score.DataUpdatedAt) {
		return true, nil
	} else {
		return false, nil
	}
}

func (l *UpdateScoreLogic) pushContributionTask(developerId int64) (err error) {
	// Comment
	commentOfUserUpdatedAt, err := l.svcCtx.CommentOfUserUpdatedAtModel.FindOneByDeveloperId(l.ctx, developerId)
	if err != nil && !errors.Is(err, model_analysis.ErrNotFound) {
		return err
	}

	if !github.CheckIfDataExpired(commentOfUserUpdatedAt.DataUpdatedAt) {
		return nil
	}

	locksCommentOfUserKey := l.svcCtx.Locks.GetNewLocksKey(constants.LockCommentOfUser, developerId)

	err = l.svcCtx.Locks.DelOldLocksKey(l.ctx, locksCommentOfUserKey)
	if err != nil {
		return err
	}

	err = tasks.FetcherTaskPusher(l.svcCtx.AsynqClient, constants.FetchCommentOfUser, developerId, "", 0)
	if err != nil {
		return err
	}

	err = l.svcCtx.Locks.Block(l.ctx, locksCommentOfUserKey)
	if err != nil {
		return err
	}

	// Issue-PR
	issuePROfUserUpdatedAt, err := l.svcCtx.IssuePROfUserUpdatedAtModel.FindOneByDeveloperId(l.ctx, developerId)
	if err != nil && !errors.Is(err, model_analysis.ErrNotFound) {
		return err
	}

	if !github.CheckIfDataExpired(issuePROfUserUpdatedAt.DataUpdatedAt) {
		return nil
	}

	locksIssuePROfUserKey := l.svcCtx.Locks.GetNewLocksKey(constants.LockIssuePROfUser, developerId)

	err = l.svcCtx.Locks.DelOldLocksKey(l.ctx, locksIssuePROfUserKey)
	if err != nil {
		return err
	}

	err = tasks.FetcherTaskPusher(l.svcCtx.AsynqClient, constants.FetchIssuePROfUser, developerId, "", 0)
	if err != nil {
		return err
	}

	err = l.svcCtx.Locks.Block(l.ctx, locksIssuePROfUserKey)
	if err != nil {
		return err
	}

	// Review
	reviewOfUserUpdatedAt, err := l.svcCtx.ReviewOfUserUpdatedAtModel.FindOneByDeveloperId(l.ctx, developerId)
	if err != nil && !errors.Is(err, model_analysis.ErrNotFound) {
		return err
	}

	if !github.CheckIfDataExpired(reviewOfUserUpdatedAt.DataUpdatedAt) {
		return nil
	}

	locksReviewOfUserKey := l.svcCtx.Locks.GetNewLocksKey(constants.LockReviewOfUser, developerId)

	err = l.svcCtx.Locks.DelOldLocksKey(l.ctx, locksReviewOfUserKey)
	if err != nil {
		return err
	}

	err = tasks.FetcherTaskPusher(l.svcCtx.AsynqClient, constants.FetchReviewOfUser, developerId, "", 0)
	if err != nil {
		return err
	}

	err = l.svcCtx.Locks.Block(l.ctx, locksReviewOfUserKey)
	if err != nil {
		return err
	}

	return nil
}

func (l *UpdateScoreLogic) rpcGetContributinoById(developerId, limit, page int64) ([]*contribution.Contribution, error) {
	var resp *contribution.SearchByDeveloperIdResp

	resp, err := l.svcCtx.ContributionRpcClient.SearchByDeveloperId(l.ctx, &contribution.SearchByDeveloperIdReq{
		DeveloperId: developerId,
		Limit:       limit,
		Page:        page,
	})
	if err != nil {
		logx.Errorf("SearchByDeveloperId: RPC Called Failed: %v", err.Error())
		err = errno.InternalServiceError.WithError(err)
		return nil, err
	}

	if !utils.IsSuccess(resp.Base) {
		err = errno.BizError.WithMessage(resp.Base.Message)
		return nil, err
	}

	return resp.Contributions, nil
}

func (l *UpdateScoreLogic) pushRepoTask(repoId int64) (err error) {
	repo, err := l.svcCtx.ReviewOfUserUpdatedAtModel.FindOneByDeveloperId(l.ctx, repoId)
	if err != nil && !errors.Is(err, model_analysis.ErrNotFound) {
		return err
	}

	if !github.CheckIfDataExpired(repo.DataUpdatedAt) {
		return nil
	}

	locksKey := l.svcCtx.Locks.GetNewLocksKey(constants.LockRepo, repoId)

	err = l.svcCtx.Locks.DelOldLocksKey(l.ctx, locksKey)
	if err != nil {
		return err
	}

	err = tasks.FetcherTaskPusher(l.svcCtx.AsynqClient, constants.FetchRepo, repoId, "", 0)
	if err != nil {
		return err
	}

	err = l.svcCtx.Locks.Block(l.ctx, locksKey)
	if err != nil {
		return err
	}

	return nil
}

func (l *UpdateScoreLogic) rpcGetRepoById(repoId int64) (*repo.Repo, error) {
	var resp *repo.GetRepoByIdResp

	resp, err := l.svcCtx.RepoRpcClient.GetRepoById(l.ctx, &repo.GetRepoByIdReq{
		Id: repoId,
	})
	if err != nil {
		logx.Errorf("GetRepoByIdRPC: RPC Called Failed: %v", err.Error())
		err = errno.InternalServiceError.WithError(err)
		return nil, err
	}

	if !utils.IsSuccess(resp.Base) {
		err = errno.BizError.WithMessage(resp.Base.Message)
		return nil, err
	}

	if resp.Repo == nil {
		err = errno.BizNotFoundError.WithMessage("Repo Not Found")
		return nil, err
	}

	return resp.Repo, nil
}

func (l *UpdateScoreLogic) pushFollowerTask(developerId int64) (err error) {
	followerUpdatedAt, err := l.svcCtx.FollowerUpdatedAtModel.FindOneByDeveloperId(l.ctx, developerId)
	if err != nil && !errors.Is(err, model_analysis.ErrNotFound) {
		return err
	}

	if !github.CheckIfDataExpired(followerUpdatedAt.DataUpdatedAt) {
		return nil
	}

	locksKey := l.svcCtx.Locks.GetNewLocksKey(constants.LockFollower, developerId)

	err = l.svcCtx.Locks.DelOldLocksKey(l.ctx, locksKey)
	if err != nil {
		return err
	}

	err = tasks.FetcherTaskPusher(l.svcCtx.AsynqClient, constants.FetchFollower, developerId, "", 0)
	if err != nil {
		return err
	}

	err = l.svcCtx.Locks.Block(l.ctx, locksKey)
	if err != nil {
		return err
	}

	return nil
}

func (l *UpdateScoreLogic) rpcGetFollowerCountById(developerId int64) (int, error) {
	var resp *relation.SearchFollowerByDeveloperIdResp

	resp, err := l.svcCtx.RelationRpcClient.SearchFollowerByDeveloperId(l.ctx, &relation.SearchFollowerByDeveloperIdReq{
		DeveloperId: developerId,
	})
	if err != nil {
		logx.Errorf("SearchFollowerByDeveloperIdRPC: RPC Called Failed: %v", err.Error())
		err = errno.InternalServiceError.WithError(err)
		return 0, err
	}

	if !utils.IsSuccess(resp.Base) {
		err = errno.BizError.WithMessage(resp.Base.Message)
		return 0, err
	}

	return len(resp.FollowerIds), nil
}

func (l *UpdateScoreLogic) pushStarredTask(developerId int64) (err error) {
	starredRepoUpdatedAt, err := l.svcCtx.StarredRepoUpdatedAtModel.FindOneByDeveloperId(l.ctx, developerId)
	if err != nil && !errors.Is(err, model_analysis.ErrNotFound) {
		return err
	}

	if !github.CheckIfDataExpired(starredRepoUpdatedAt.DataUpdatedAt) {
		return nil
	}

	locksKey := l.svcCtx.Locks.GetNewLocksKey(constants.LockStarredRepo, developerId)

	err = l.svcCtx.Locks.DelOldLocksKey(l.ctx, locksKey)
	if err != nil {
		return err
	}

	err = tasks.FetcherTaskPusher(l.svcCtx.AsynqClient, constants.FetchStarredRepo, developerId, "", 0)
	if err != nil {
		return err
	}

	err = l.svcCtx.Locks.Block(l.ctx, locksKey)
	if err != nil {
		return err
	}

	return nil
}

func (l *UpdateScoreLogic) rpcGetStarredCountById(developerId int64) (int, error) {
	var resp *relation.SearchStarredRepoResp

	resp, err := l.svcCtx.RelationRpcClient.SearchStarredRepo(l.ctx, &relation.SearchStarredRepoReq{
		DeveloperId: developerId,
	})
	if err != nil {
		logx.Errorf("SearchStarredRepo: RPC Called Failed: %v", err.Error())
		err = errno.InternalServiceError.WithError(err)
		return 0, err
	}

	if !utils.IsSuccess(resp.Base) {
		err = errno.BizError.WithMessage(resp.Base.Message)
		return 0, err
	}

	return len(resp.RepoIds), nil
}

func (l *UpdateScoreLogic) updateScore(model *model_analysis.Score) error {
	score, err := l.svcCtx.ScoreModel.FindOneByDeveloperId(l.ctx, model.DeveloperId)
	if err != nil {
		switch {
		case errors.Is(err, model_analysis.ErrNotFound):
			logx.Info("No Found This Data!")

			dataId, err := l.svcCtx.ScoreModel.CreateDataId()
			if err != nil {
				err = errno.InternalDatabaseError.WithError(err)
				return err
			}

			model.DataId = dataId
			_, err = l.svcCtx.ScoreModel.Insert(l.ctx, model)
			if err != nil {
				err = errno.InternalDatabaseError.WithError(err)
				return err
			}

			return nil
		default:
			return err
		}
	}

	model.DataId = score.DataId
	err = l.svcCtx.ScoreModel.Update(l.ctx, model)
	if err != nil {
		err = errno.InternalDatabaseError.WithError(err)
		return err
	}

	return nil
}

func (l *UpdateScoreLogic) updateRankScore(developerId int64, score float64) error {
	_, err := l.svcCtx.RedisClient.ZaddFloatCtx(l.ctx, constants.ScoreKey, score, strconv.FormatInt(developerId, 10))
	if err != nil {
		err = errno.InternalRedisError.WithError(err)
		return err
	}

	return nil
}
