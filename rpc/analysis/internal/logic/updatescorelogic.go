package logic

import (
	"context"
	"errors"
	"math"
	"strconv"
	"sync"

	"github.com/wushiling50/aster/gen/analysis"
	"github.com/wushiling50/aster/gen/contribution"
	"github.com/wushiling50/aster/gen/developer"
	"github.com/wushiling50/aster/gen/repo"
	"github.com/wushiling50/aster/pkg/constants"
	"github.com/wushiling50/aster/pkg/errno"
	"github.com/wushiling50/aster/pkg/github"
	model_analysis "github.com/wushiling50/aster/pkg/model/analysis"
	model_contribution "github.com/wushiling50/aster/pkg/model/contribution"
	model_developer "github.com/wushiling50/aster/pkg/model/developer"
	model_repo "github.com/wushiling50/aster/pkg/model/repo"
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
		mu sync.Mutex
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
		theRepoId := repoId
		theContributions := contributions
		eg.Go(func() error {
			var (
				err error

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
			err = l.pushRepoTask(theRepoId)
			if err != nil {
				err = errno.InternalAsynqError.WithError(err)
				return err
			}

			theRepo, err = l.rpcGetRepoById(theRepoId)
			if err != nil {
				return err
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
			for _, theContribution := range theContributions {
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

			if theRepo.CommentCount != 0 {
				radio := contributionComment / float64(theRepo.CommentCount)
				contributionScore += constants.ScoreContributionComment * radio
			}

			if theRepo.IssueCount != 0 {
				radio := contributionIssue / float64(theRepo.IssueCount)
				contributionScore += constants.ScoreContributionIssue * radio
			}

			if theRepo.OpenPrCount != 0 {
				radio := contributionOpenPr / float64(theRepo.OpenPrCount)
				contributionScore += constants.ScoreContributionOpenPR * radio
			}

			if theRepo.ReviewCount != 0 {
				radio := contributionReview / float64(theRepo.ReviewCount)
				contributionScore += constants.ScoreContributionReview * radio
			}

			if theRepo.MergedPrCount != 0 {
				radio := contributionMergePr / float64(theRepo.MergedPrCount)
				contributionScore += constants.ScoreContributionMergePR * radio
			}

			mu.Lock()
			defer mu.Unlock()

			totalScore += repoScore * contributionScore

			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		switch {
		case errors.Is(err, errno.InternalAsynqError):
			logx.Errorf("service.UpdateScore: Failed To Enqueue Task: %v", err.Error())

			resp.Base = pack.BuildBaseResp(err)
			return resp, nil
		default:
			logx.Error(err)
			resp.Base = pack.BuildBaseResp(err)
			return resp, nil
		}
	}

	// developer score
	err = l.pushDeveloperTask(in.DeveloperId)
	if err != nil {
		logx.Errorf("service.UpdateScore: Failed To Enqueue Task: %v", err.Error())
		resp.Base = pack.BuildBaseResp(err)
		return resp, err
	}

	theDeveloper, err := l.rpcGetDeveloperById(in.DeveloperId)
	if err != nil {
		logx.Error(err)
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	developerScore += constants.ScoreFollower*float64(theDeveloper.Followers) +
		constants.ScoreStarred*float64(theDeveloper.Stars)

	developerScore = math.Sqrt(developerScore * 0.1)

	totalScore += developerScore

	// update score
	err = l.updateScore(&model_analysis.Score{
		DeveloperId: in.DeveloperId,
		Score:       totalScore,
	})
	if err != nil {
		logx.Errorf("service.UpdateScore: Update Score Failed: %w", err)
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	err = l.updateRankScore(in.DeveloperId, totalScore)
	if err != nil {
		logx.Errorf("service.UpdateScore: Update Rank Score Failed: %w", err)
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.Base = pack.BuildSuccessResp()

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

	if utils.CheckIfDataExpired(score.UpdatedAt) {
		return true, nil
	} else {
		return false, nil
	}
}

func (l *UpdateScoreLogic) pushContributionTask(developerId int64) error {
	var eg errgroup.Group

	// Comment
	eg.Go(func() error {
		var err error

		locksKey := l.svcCtx.Locks.GetNewLocksKey(constants.LockCommentOfUser, developerId)
		getLock, err := l.svcCtx.Locks.TryLock(l.ctx, locksKey)
		if err != nil {
			return err
		}

		defer l.svcCtx.Locks.TryUnLock(l.ctx, locksKey)

		if !getLock {
			return nil
		}

		commentOfUserUpdatedAt, err := l.svcCtx.CommentOfUserUpdatedAtModel.FindOneByDeveloperId(l.ctx, developerId)
		if err != nil && !errors.Is(err, model_contribution.ErrNotFound) {
			return err
		}

		if commentOfUserUpdatedAt != nil && !utils.CheckIfDataExpired(commentOfUserUpdatedAt.UpdatedAt) {
			return nil
		}

		blocksCommentOfUserKey := l.svcCtx.Locks.GetNewLocksKey(constants.BlockCommentOfUser, developerId)

		err = l.svcCtx.Locks.DelOldLocksKey(l.ctx, blocksCommentOfUserKey)
		if err != nil {
			return err
		}

		err = tasks.FetcherTaskPusher(l.svcCtx.AsynqClient, constants.FetchCommentOfUser, developerId, github.DefaultUpdateAfterTime(), 0)
		if err != nil {
			return err
		}

		err = l.svcCtx.Locks.Block(l.ctx, blocksCommentOfUserKey)
		if err != nil {
			return err
		}

		return nil
	})

	// Issue-PR
	eg.Go(func() error {
		var err error

		locksKey := l.svcCtx.Locks.GetNewLocksKey(constants.LockIssuePROfUser, developerId)
		getLock, err := l.svcCtx.Locks.TryLock(l.ctx, locksKey)
		if err != nil {
			return err
		}

		defer l.svcCtx.Locks.TryUnLock(l.ctx, locksKey)

		if !getLock {
			return nil
		}

		issuePROfUserUpdatedAt, err := l.svcCtx.IssuePROfUserUpdatedAtModel.FindOneByDeveloperId(l.ctx, developerId)
		if err != nil && !errors.Is(err, model_contribution.ErrNotFound) {
			return err
		}

		if issuePROfUserUpdatedAt != nil && !utils.CheckIfDataExpired(issuePROfUserUpdatedAt.UpdatedAt) {
			return nil
		}

		blocksIssuePROfUserKey := l.svcCtx.Locks.GetNewLocksKey(constants.BlockIssuePROfUser, developerId)

		err = l.svcCtx.Locks.DelOldLocksKey(l.ctx, blocksIssuePROfUserKey)
		if err != nil {
			return err
		}

		err = tasks.FetcherTaskPusher(l.svcCtx.AsynqClient, constants.FetchIssuePROfUser, developerId, github.DefaultUpdateAfterTime(), 0)
		if err != nil {
			return err
		}

		err = l.svcCtx.Locks.Block(l.ctx, blocksIssuePROfUserKey)
		if err != nil {
			return err
		}

		return nil
	})

	// Review
	eg.Go(func() error {
		var err error

		locksKey := l.svcCtx.Locks.GetNewLocksKey(constants.LockReviewOfUser, developerId)
		getLock, err := l.svcCtx.Locks.TryLock(l.ctx, locksKey)
		if err != nil {
			return err
		}

		defer l.svcCtx.Locks.TryUnLock(l.ctx, locksKey)

		if !getLock {
			return nil
		}

		reviewOfUserUpdatedAt, err := l.svcCtx.ReviewOfUserUpdatedAtModel.FindOneByDeveloperId(l.ctx, developerId)
		if err != nil && !errors.Is(err, model_contribution.ErrNotFound) {
			return err
		}

		if reviewOfUserUpdatedAt != nil && !utils.CheckIfDataExpired(reviewOfUserUpdatedAt.UpdatedAt) {
			return nil
		}

		blocksReviewOfUserKey := l.svcCtx.Locks.GetNewLocksKey(constants.BlockReviewOfUser, developerId)

		err = l.svcCtx.Locks.DelOldLocksKey(l.ctx, blocksReviewOfUserKey)
		if err != nil {
			return err
		}

		err = tasks.FetcherTaskPusher(l.svcCtx.AsynqClient, constants.FetchReviewOfUser, developerId, github.DefaultUpdateAfterTime(), 0)
		if err != nil {
			return err
		}

		err = l.svcCtx.Locks.Block(l.ctx, blocksReviewOfUserKey)
		if err != nil {
			return err
		}

		return nil
	})

	if err := eg.Wait(); err != nil {
		logx.Error(err)
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
	locksKey := l.svcCtx.Locks.GetNewLocksKey(constants.LockRepo, repoId)
	getLock, err := l.svcCtx.Locks.TryLock(l.ctx, locksKey)
	if err != nil {
		return err
	}

	defer l.svcCtx.Locks.TryUnLock(l.ctx, locksKey)

	if !getLock {
		return nil
	}

	repo, err := l.svcCtx.RepoModel.FindOneById(l.ctx, repoId)
	if err != nil && !errors.Is(err, model_repo.ErrNotFound) {
		return err
	}

	if repo != nil && !utils.CheckIfDataExpired(repo.UpdatedAt) {
		return nil
	}

	blocksKey := l.svcCtx.Locks.GetNewLocksKey(constants.BlockRepo, repoId)

	err = l.svcCtx.Locks.DelOldLocksKey(l.ctx, blocksKey)
	if err != nil {
		return err
	}

	err = tasks.FetcherTaskPusher(l.svcCtx.AsynqClient, constants.FetchRepo, repoId, "", 0)
	if err != nil {
		return err
	}

	err = l.svcCtx.Locks.Block(l.ctx, blocksKey)
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

func (l *UpdateScoreLogic) pushDeveloperTask(developerId int64) (err error) {
	locksKey := l.svcCtx.Locks.GetNewLocksKey(constants.LockDeveloper, developerId)
	getLock, err := l.svcCtx.Locks.TryLock(l.ctx, locksKey)
	if err != nil {
		return err
	}

	defer l.svcCtx.Locks.TryUnLock(l.ctx, locksKey)

	if !getLock {
		return nil
	}

	developer, err := l.svcCtx.DeveloperModel.FindOneById(l.ctx, developerId)
	if err != nil && !errors.Is(err, model_developer.ErrNotFound) {
		return err
	}

	if developer != nil && !utils.CheckIfDataExpired(developer.UpdatedAt) {
		return nil
	}

	blocksKey := l.svcCtx.Locks.GetNewLocksKey(constants.BlockDeveloper, developerId)

	err = l.svcCtx.Locks.DelOldLocksKey(l.ctx, blocksKey)
	if err != nil {
		return err
	}

	err = tasks.FetcherTaskPusher(l.svcCtx.AsynqClient, constants.FetchDeveloper, developerId, "", 0)
	if err != nil {
		return err
	}

	err = l.svcCtx.Locks.Block(l.ctx, blocksKey)
	if err != nil {
		return err
	}

	return nil
}

func (l *UpdateScoreLogic) rpcGetDeveloperById(developerId int64) (*developer.Developer, error) {
	var resp *developer.GetDeveloperByIdResp

	resp, err := l.svcCtx.DeveloperRpcClient.GetDeveloperById(l.ctx, &developer.GetDeveloperByIdReq{
		Id: developerId,
	})
	if err != nil {
		logx.Errorf("GetDeveloperByIdRPC: RPC Called Failed: %v", err.Error())
		err = errno.InternalServiceError.WithError(err)
		return nil, err
	}

	if !utils.IsSuccess(resp.Base) {
		err = errno.BizError.WithMessage(resp.Base.Message)
		return nil, err
	}

	if resp.Developer == nil {
		err = errno.BizNotFoundError.WithMessage("Developer Not Found")
		return nil, err
	}

	return resp.Developer, nil
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
