package logic

import (
	"context"
	"errors"

	"github.com/wushiling50/aster/gen/analysis"
	"github.com/wushiling50/aster/gen/contribution"
	"github.com/wushiling50/aster/pkg/constants"
	"github.com/wushiling50/aster/pkg/errno"
	"github.com/wushiling50/aster/pkg/github"
	model_analysis "github.com/wushiling50/aster/pkg/model/analysis"
	"github.com/wushiling50/aster/pkg/tasks"
	"github.com/wushiling50/aster/pkg/utils"
	"github.com/wushiling50/aster/rpc/analysis/internal/pack"
	"github.com/wushiling50/aster/rpc/analysis/internal/svc"

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
		err = errno.InternalAsynqError.WithError(err)
		resp.Base = pack.BuildSuccessResp()
		return resp, nil
	}

	contributions, err := l.rpcGetContributinoById(in.DeveloperId, 1000, 1)
	if err != nil {
		logx.Error(err)
		resp.Base = pack.BuildSuccessResp()
		return resp, nil
	}

	for _, theContribution := range contributions {
		contributionsCategorizedByRepoId[theContribution.RepoId] = append(contributionsCategorizedByRepoId[theContribution.RepoId], theContribution)
	}

	// TODO: analysis score

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
