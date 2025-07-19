package logic

import (
	"context"
	"errors"

	"github.com/wushiling50/aster/gen/analysis"
	"github.com/wushiling50/aster/gen/contribution"
	"github.com/wushiling50/aster/gen/developer"
	"github.com/wushiling50/aster/pkg/constants"
	"github.com/wushiling50/aster/pkg/errno"
	"github.com/wushiling50/aster/pkg/github"
	"github.com/wushiling50/aster/pkg/llm"
	model_analysis "github.com/wushiling50/aster/pkg/model/analysis"
	model_contribution "github.com/wushiling50/aster/pkg/model/contribution"
	model_developer "github.com/wushiling50/aster/pkg/model/developer"
	"github.com/wushiling50/aster/pkg/tasks"
	"github.com/wushiling50/aster/pkg/utils"
	"github.com/wushiling50/aster/rpc/analysis/internal/pack"
	"github.com/wushiling50/aster/rpc/analysis/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSummaryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateSummaryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSummaryLogic {
	return &UpdateSummaryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateSummaryLogic) UpdateSummary(in *analysis.UpdateAnalysisReq) (*analysis.UpdateAnalysisResp, error) {
	resp := new(analysis.UpdateAnalysisResp)

	need, err := l.checkIfNeedUpdate(in.DeveloperId)
	if err != nil {
		logx.Errorf("service.UpdateSummary: Update Summary Failed: %w", err)
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	if !need {
		resp.Base = pack.BuildSuccessResp()
		return resp, nil
	}

	summary, err := l.getSummaryByLLModel(in.DeveloperId)
	if err != nil {
		logx.Errorf("service.UpdateSummary: Get Summary By LLM Failed: %w", err)
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	err = l.updateSummary(&model_analysis.Summary{
		DeveloperId: in.DeveloperId,
		Summary:     summary,
	})

	if err != nil {
		logx.Errorf("service.UpdateSummary: Update Summary Failed: %w", err)
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	return resp, nil
}

func (l *UpdateSummaryLogic) checkIfNeedUpdate(developerId int64) (bool, error) {
	summary, err := l.svcCtx.SummaryModel.FindOneByDeveloperId(l.ctx, developerId)
	if err != nil {
		switch {
		case errors.Is(err, model_analysis.ErrNotFound):
			return true, nil
		default:
			return false, err
		}
	}

	if github.CheckIfDataExpired(summary.DataUpdatedAt) {
		return true, nil
	} else {
		return false, nil
	}
}

func (l *UpdateSummaryLogic) getSummaryByLLModel(developerId int64) (summary string, err error) {
	var (
		allText          string = ""
		developerProfile *developer.Developer
		contributions    []*contribution.Contribution
		languages        *analysis.Languages
	)

	// get developer
	err = l.pushDeveloperTask(developerId)
	if err != nil {
		return
	}

	developerProfile, err = l.rpcGetDeveloperById(developerId)
	if err != nil {
		return
	}

	allText += utils.GetTextFromProfile(developerProfile)

	// get contributions
	err = l.pushContributionTask(developerId)
	if err != nil {
		return
	}

	contributions, err = l.rpcGetContributinoById(developerId, 1000, 1)
	if err != nil {
		return
	}

	allText += utils.GetTextFromContribution(contributions)

	// get languages
	languages, err = l.getLanguages(developerId)
	if err != nil {
		return
	}

	allText += utils.GetTextFromLanguages(languages)

	// get result from llm
	req := llm.BuildAnalysisSummaryReq(l.svcCtx.Config.DeepSeek, allText)

	resp, err := l.svcCtx.DeepSeekClient.CreateChatCompletion(l.ctx, req)
	if err != nil {
		err = errno.InternalLLMError.WithError(err)
		return
	}

	if len(resp.Choices) == 0 {
		err = errno.InternalLLMError.WithMessage("Not Get Vaild Request")
		return
	}

	summary = resp.Choices[0].Message.Content

	return
}

func (l *UpdateSummaryLogic) pushDeveloperTask(developerId int64) (err error) {
	developer, err := l.svcCtx.DeveloperModel.FindOneById(l.ctx, developerId)
	if err != nil && !errors.Is(err, model_developer.ErrNotFound) {
		return err
	}

	if developer != nil && !github.CheckIfDataExpired(developer.DataUpdatedAt) {
		return nil
	}

	locksKey := l.svcCtx.Locks.GetNewLocksKey(constants.LockDeveloper, developerId)

	err = l.svcCtx.Locks.DelOldLocksKey(l.ctx, locksKey)
	if err != nil {
		return err
	}

	err = tasks.FetcherTaskPusher(l.svcCtx.AsynqClient, constants.FetchDeveloper, developerId, "", 0)
	if err != nil {
		return err
	}

	err = l.svcCtx.Locks.Block(l.ctx, locksKey)
	if err != nil {
		return err
	}

	return nil
}

func (l *UpdateSummaryLogic) rpcGetDeveloperById(developerId int64) (*developer.Developer, error) {
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

func (l *UpdateSummaryLogic) pushContributionTask(developerId int64) (err error) {
	// Comment
	commentOfUserUpdatedAt, err := l.svcCtx.CommentOfUserUpdatedAtModel.FindOneByDeveloperId(l.ctx, developerId)
	if err != nil && !errors.Is(err, model_contribution.ErrNotFound) {
		return err
	}

	if commentOfUserUpdatedAt != nil && !github.CheckIfDataExpired(commentOfUserUpdatedAt.DataUpdatedAt) {
		return nil
	}

	locksCommentOfUserKey := l.svcCtx.Locks.GetNewLocksKey(constants.LockCommentOfUser, developerId)

	err = l.svcCtx.Locks.DelOldLocksKey(l.ctx, locksCommentOfUserKey)
	if err != nil {
		return err
	}

	err = tasks.FetcherTaskPusher(l.svcCtx.AsynqClient, constants.FetchCommentOfUser, developerId, github.DefaultUpdateAfterTime(), 0)
	if err != nil {
		return err
	}

	err = l.svcCtx.Locks.Block(l.ctx, locksCommentOfUserKey)
	if err != nil {
		return err
	}

	// Issue-PR
	issuePROfUserUpdatedAt, err := l.svcCtx.IssuePROfUserUpdatedAtModel.FindOneByDeveloperId(l.ctx, developerId)
	if err != nil && !errors.Is(err, model_contribution.ErrNotFound) {
		return err
	}

	if issuePROfUserUpdatedAt != nil && !github.CheckIfDataExpired(issuePROfUserUpdatedAt.DataUpdatedAt) {
		return nil
	}

	locksIssuePROfUserKey := l.svcCtx.Locks.GetNewLocksKey(constants.LockIssuePROfUser, developerId)

	err = l.svcCtx.Locks.DelOldLocksKey(l.ctx, locksIssuePROfUserKey)
	if err != nil {
		return err
	}

	err = tasks.FetcherTaskPusher(l.svcCtx.AsynqClient, constants.FetchIssuePROfUser, developerId, github.DefaultUpdateAfterTime(), 0)
	if err != nil {
		return err
	}

	err = l.svcCtx.Locks.Block(l.ctx, locksIssuePROfUserKey)
	if err != nil {
		return err
	}

	// Review
	reviewOfUserUpdatedAt, err := l.svcCtx.ReviewOfUserUpdatedAtModel.FindOneByDeveloperId(l.ctx, developerId)
	if err != nil && !errors.Is(err, model_contribution.ErrNotFound) {
		return err
	}

	if reviewOfUserUpdatedAt != nil && !github.CheckIfDataExpired(reviewOfUserUpdatedAt.DataUpdatedAt) {
		return nil
	}

	locksReviewOfUserKey := l.svcCtx.Locks.GetNewLocksKey(constants.LockReviewOfUser, developerId)

	err = l.svcCtx.Locks.DelOldLocksKey(l.ctx, locksReviewOfUserKey)
	if err != nil {
		return err
	}

	err = tasks.FetcherTaskPusher(l.svcCtx.AsynqClient, constants.FetchReviewOfUser, developerId, github.DefaultUpdateAfterTime(), 0)
	if err != nil {
		return err
	}

	err = l.svcCtx.Locks.Block(l.ctx, locksReviewOfUserKey)
	if err != nil {
		return err
	}

	return nil
}

func (l *UpdateSummaryLogic) rpcGetContributinoById(developerId, limit, page int64) ([]*contribution.Contribution, error) {
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

func (l *UpdateSummaryLogic) getLanguages(developerId int64) (*analysis.Languages, error) {
	var (
		updateLogic = NewUpdateLanguageLogic(l.ctx, l.svcCtx)
		getLogic    = NewGetLanguagesLogic(l.ctx, l.svcCtx)

		updateResp *analysis.UpdateAnalysisResp
		getResp    *analysis.GetLanguagesResp

		languages *analysis.Languages
	)

	updateResp, err := updateLogic.UpdateLanguage(&analysis.UpdateAnalysisReq{
		DeveloperId: developerId,
	})

	if err != nil {
		logx.Errorf("UpdateLanguage: Update Language Failed: %v", err.Error())
		return nil, err
	}

	if !utils.IsSuccess(updateResp.Base) {
		err = errno.BizError.WithMessage(updateResp.Base.Message)
		return nil, err
	}

	getResp, err = getLogic.GetLanguages(&analysis.GetAnalysisReq{
		DeveloperId: developerId,
	})

	if err != nil {
		logx.Errorf("GetLanguages: Get Languages Failed: %v", err.Error())
		return nil, err
	}

	if !utils.IsSuccess(updateResp.Base) {
		err = errno.BizError.WithMessage(updateResp.Base.Message)
		return nil, err
	}

	languages = getResp.Languages

	if languages == nil {
		err = errno.BizNotFoundError.WithMessage("Languages Not Found")
		return nil, err
	}

	return languages, nil
}

func (l *UpdateSummaryLogic) updateSummary(model *model_analysis.Summary) error {
	summary, err := l.svcCtx.SummaryModel.FindOneByDeveloperId(l.ctx, model.DeveloperId)
	if err != nil {
		switch {
		case errors.Is(err, model_analysis.ErrNotFound):
			logx.Info("No Found This Data!")

			dataId, err := l.svcCtx.SummaryModel.CreateDataId()
			if err != nil {
				err = errno.InternalDatabaseError.WithError(err)
				return err
			}

			model.DataId = dataId
			_, err = l.svcCtx.SummaryModel.Insert(l.ctx, model)
			if err != nil {
				err = errno.InternalDatabaseError.WithError(err)
				return err
			}

			return nil
		default:
			return err
		}
	}

	model.DataId = summary.DataId
	err = l.svcCtx.SummaryModel.Update(l.ctx, model)
	if err != nil {
		err = errno.InternalDatabaseError.WithError(err)
		return err
	}

	return nil
}
