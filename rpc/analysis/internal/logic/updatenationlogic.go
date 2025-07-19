package logic

import (
	"context"
	"encoding/json"
	"errors"
	"os/exec"
	"strings"

	"github.com/biter777/countries"
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

type UpdateNationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateNationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateNationLogic {
	return &UpdateNationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateNationLogic) UpdateNation(in *analysis.UpdateAnalysisReq) (*analysis.UpdateAnalysisResp, error) {
	resp := new(analysis.UpdateAnalysisResp)

	var (
		login            string
		nationConfidence = make(map[string]float64)

		mostPossibleNation     string
		mostPossibleConfidence float64 = 0
	)

	need, err := l.checkIfNeedUpdate(in.DeveloperId)
	if err != nil {
		logx.Errorf("service.UpdateNation: Update Nation Failed: %w", err)
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	if !need {
		resp.Base = pack.BuildSuccessResp()
		return resp, nil
	}

	login, err = github.GetLoginById(l.ctx, in.DeveloperId)
	if err != nil {
		logx.Errorf("service.UpdateNation: Failed To Get Login By Id :%v", err.Error())
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	// get nation with confidence by script and llm
	nationConfidence, err = l.getNationWithConfidenceByScript(login)
	if err != nil {
		logx.Info("service.UpdateNation: Get Nation And Confidence By Script Failed")
		err = nil
	}

	if nationConfidence == nil {
		nationConfidence = make(map[string]float64)
	}

	nation, confidence, err := l.getNationWithConfidenceByLLModel(in.DeveloperId)
	if err != nil {
		logx.Errorf("service.UpdateNation: Get Nation And Confidence By LLM Failed: %w", err)
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}
	nationConfidence[nation] += confidence * constants.LLModelConfidenceWeight

	// update
	for nation, confidence := range nationConfidence {
		if confidence > mostPossibleConfidence {
			mostPossibleNation = nation
			mostPossibleConfidence = confidence
		}
	}

	err = l.updateNation(&model_analysis.Nation{
		DeveloperId: in.DeveloperId,
		Nation:      mostPossibleNation,
		Confidence:  mostPossibleConfidence,
	})

	if err != nil {
		logx.Errorf("service.UpdateNation: Update Nation Failed: %w", err)
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	return resp, nil
}

func (l *UpdateNationLogic) checkIfNeedUpdate(developerId int64) (bool, error) {
	nation, err := l.svcCtx.NationModel.FindOneByDeveloperId(l.ctx, developerId)
	if err != nil {
		switch {
		case errors.Is(err, model_analysis.ErrNotFound):
			return true, nil
		default:
			return false, err
		}
	}

	if github.CheckIfDataExpired(nation.DataUpdatedAt) {
		return true, nil
	} else {
		return false, nil
	}
}

func (l *UpdateNationLogic) getNationWithConfidenceByScript(login string) (nationConfidence map[string]float64, err error) {
	var (
		cmd *exec.Cmd
		out []byte
	)

	cmd = exec.Command("venv/bin/python", "script/nation/main.py", login)

	out, err = cmd.CombinedOutput()
	if err != nil {
		err = errno.InternalScriptError.WithError(err)
		return
	}

	err = json.Unmarshal(out, &nationConfidence)
	if err != nil {
		err = errno.InternalJSONError.WithError(err)
		return
	}

	return
}

func (l *UpdateNationLogic) getNationWithConfidenceByLLModel(developerId int64) (nation string, confidence float64, err error) {
	var (
		allText          string = ""
		nationConfidence *llm.NationConfidence
		developerProfile *developer.Developer
		contributions    []*contribution.Contribution
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

	// get result from llm
	req := llm.BuildAnalysisNationReq(l.svcCtx.Config.DeepSeek, allText)

	resp, err := l.svcCtx.DeepSeekClient.CreateChatCompletion(l.ctx, req)
	if err != nil {
		err = errno.InternalLLMError.WithError(err)
		return
	}

	if len(resp.Choices) == 0 {
		err = errno.InternalLLMError.WithMessage("Not Get Vaild Request")
		return
	}

	modelOutput := resp.Choices[0].Message.Content

	if err = json.Unmarshal([]byte(modelOutput), &nationConfidence); err != nil {
		err = errno.InternalJSONError.WithError(err)
		return
	}

	nation = strings.ToLower(countries.ByName(nationConfidence.Nation).Alpha2())
	if nationConfidence.Confidence < 0 {
		confidence = 0
	} else if nationConfidence.Confidence > 1 {
		confidence = 1
	} else {
		confidence = nationConfidence.Confidence
	}

	return
}

func (l *UpdateNationLogic) pushDeveloperTask(developerId int64) (err error) {
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

func (l *UpdateNationLogic) rpcGetDeveloperById(developerId int64) (*developer.Developer, error) {
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

func (l *UpdateNationLogic) pushContributionTask(developerId int64) (err error) {
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

func (l *UpdateNationLogic) rpcGetContributinoById(developerId, limit, page int64) ([]*contribution.Contribution, error) {
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

func (l *UpdateNationLogic) updateNation(model *model_analysis.Nation) error {
	nation, err := l.svcCtx.NationModel.FindOneByDeveloperId(l.ctx, model.DeveloperId)
	if err != nil {
		switch {
		case errors.Is(err, model_analysis.ErrNotFound):
			logx.Info("No Found This Data!")

			dataId, err := l.svcCtx.NationModel.CreateDataId()
			if err != nil {
				err = errno.InternalDatabaseError.WithError(err)
				return err
			}

			model.DataId = dataId
			_, err = l.svcCtx.NationModel.Insert(l.ctx, model)
			if err != nil {
				err = errno.InternalDatabaseError.WithError(err)
				return err
			}

			return nil
		default:
			return err
		}
	}

	model.DataId = nation.DataId
	err = l.svcCtx.NationModel.Update(l.ctx, model)
	if err != nil {
		err = errno.InternalDatabaseError.WithError(err)
		return err
	}

	return nil
}
