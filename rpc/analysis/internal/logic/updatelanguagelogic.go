package logic

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/wushiling50/aster/gen/analysis"
	"github.com/wushiling50/aster/gen/relation"
	"github.com/wushiling50/aster/gen/repo"
	"github.com/wushiling50/aster/pkg/constants"
	"github.com/wushiling50/aster/pkg/errno"
	model_analysis "github.com/wushiling50/aster/pkg/model/analysis"
	model_relation "github.com/wushiling50/aster/pkg/model/relation"
	"github.com/wushiling50/aster/pkg/tasks"
	"github.com/wushiling50/aster/pkg/utils"
	"github.com/wushiling50/aster/rpc/analysis/internal/pack"
	"github.com/wushiling50/aster/rpc/analysis/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLanguageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateLanguageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLanguageLogic {
	return &UpdateLanguageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateLanguageLogic) UpdateLanguage(in *analysis.UpdateAnalysisReq) (*analysis.UpdateAnalysisResp, error) {
	resp := new(analysis.UpdateAnalysisResp)

	var (
		allLanguageBytes = make(map[string]int64)   // 使用该 language 的 byte 数
		allMetrics       = make(map[string]float64) // 某个 language 的占比
		totalBytes       int64
	)

	need, err := l.checkIfNeedUpdate(in.DeveloperId)
	if err != nil {
		logx.Errorf("service.UpdateLanguage: Update Language Failed: %w", err)
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	if !need {
		resp.Base = pack.BuildSuccessResp()
		return resp, nil
	}

	locksKey := l.svcCtx.Locks.GetNewLocksKey(constants.LockLanguage, in.DeveloperId)
	getLock, err := l.svcCtx.Locks.TryLock(l.ctx, locksKey)
	if err != nil {
		logx.Errorf("service.UpdateLanguage: Get Lock Failed: %w", err)
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	if !getLock {
		_, err = l.svcCtx.LanguagesModel.FindOneByDeveloperId(l.ctx, in.DeveloperId)
		if err != nil {
			switch {
			case errors.Is(err, model_analysis.ErrNotFound):
				l.svcCtx.Locks.Check(l.ctx, locksKey)
			default:
				resp.Base = pack.BuildBaseResp(err)
				return resp, nil
			}
		}
		resp.Base = pack.BuildSuccessResp()
		return resp, nil
	}

	defer l.svcCtx.Locks.TryUnLock(l.ctx, locksKey)

	err = l.pushCreatedRepoTask(in.DeveloperId)
	if err != nil {
		logx.Errorf("service.UpdateLanguage: Failed To Enqueue Task: %v", err.Error())
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	createRepoIds, err := l.rpcSearchCreatedRepo(in.DeveloperId, 1000, 1)
	if err != nil {
		logx.Error(err)
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	for _, repoId := range createRepoIds {
		var languageBytes map[string]int64

		repoLanguages, err := l.rpcGetRepoById(repoId)
		if err != nil {
			logx.Error(err)
			continue
		}

		if repoLanguages == "" {
			continue
		}

		err = json.Unmarshal([]byte(repoLanguages), &languageBytes)
		if err != nil {
			logx.Error(err)
			continue
		}

		for language, bytes := range languageBytes {
			allLanguageBytes[language] += bytes
			totalBytes += bytes
		}
	}

	total := float64(totalBytes)
	for language, bytes := range allLanguageBytes {
		allMetrics[language] = (float64(bytes) / total) * 100
	}

	jsonBytes, err := json.Marshal(allMetrics)
	if err != nil {
		logx.Error(err)

		err = errno.InternalJSONError.WithError(err)
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	err = l.updateLanguage(&model_analysis.Languages{
		DeveloperId: in.DeveloperId,
		Language:    string(jsonBytes),
	})

	if err != nil {
		logx.Errorf("service.UpdateLanguage: Update Language Failed: %w", err)
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.Base = pack.BuildSuccessResp()

	return resp, nil
}

func (l *UpdateLanguageLogic) checkIfNeedUpdate(developerId int64) (bool, error) {
	langauge, err := l.svcCtx.LanguagesModel.FindOneByDeveloperId(l.ctx, developerId)
	if err != nil {
		switch {
		case errors.Is(err, model_analysis.ErrNotFound):
			return true, nil
		default:
			return false, err
		}
	}

	if utils.CustomizeCheckIfDataExpired(langauge.UpdatedAt, 3*constants.ONE_DAY) {
		return true, nil
	} else {
		return false, nil
	}
}

func (l *UpdateLanguageLogic) pushCreatedRepoTask(developerId int64) error {
	locksKey := l.svcCtx.Locks.GetNewLocksKey(constants.LockCreatedRepo, developerId)
	getLock, err := l.svcCtx.Locks.TryLock(l.ctx, locksKey)
	if err != nil {
		return err
	}

	if !getLock {
		l.svcCtx.Locks.Check(l.ctx, locksKey)
		return nil
	}

	defer l.svcCtx.Locks.TryUnLock(l.ctx, locksKey)

	createdRepoUpdatedAt, err := l.svcCtx.CreatedRepoUpdatedAtModel.FindOneByDeveloperId(l.ctx, developerId)
	if err != nil && !errors.Is(err, model_relation.ErrNotFound) {
		err = errno.InternalDatabaseError.WithError(err)
		return err
	}

	if createdRepoUpdatedAt != nil && !utils.CheckIfDataExpired(createdRepoUpdatedAt.UpdatedAt) {
		return nil
	}

	blocksKey := l.svcCtx.Locks.GetNewLocksKey(constants.BlockCreatedRepo, developerId)

	err = l.svcCtx.Locks.DelOldLocksKey(l.ctx, blocksKey)
	if err != nil {
		return err
	}

	err = tasks.FetcherTaskPusher(l.svcCtx.AsynqClient, constants.FetchCreatedRepo, developerId, "", 0)
	if err != nil {
		return err
	}

	err = l.svcCtx.Locks.Block(l.ctx, blocksKey)
	if err != nil {
		return err
	}

	return nil
}

func (l *UpdateLanguageLogic) rpcSearchCreatedRepo(developerId, limit, page int64) (createRepoIds []int64, err error) {
	var resp *relation.SearchCreatedRepoResp

	resp, err = l.svcCtx.RelationRpcClient.SearchCreatedRepo(l.ctx, &relation.SearchCreatedRepoReq{
		DeveloperId: developerId,
		Limit:       limit,
		Page:        page,
	})
	if err != nil {
		logx.Errorf("SearchCreatedRepoRPC: RPC Called Failed: %v", err.Error())
		err = errno.InternalServiceError.WithError(err)
		return
	}

	if !utils.IsSuccess(resp.Base) {
		err = errno.BizError.WithMessage(resp.Base.Message)
		return
	}

	createRepoIds = resp.RepoIds

	return
}

func (l *UpdateLanguageLogic) rpcGetRepoById(repodId int64) (languages string, err error) {
	var resp *repo.GetRepoByIdResp

	resp, err = l.svcCtx.RepoRpcClient.GetRepoById(l.ctx, &repo.GetRepoByIdReq{
		Id: repodId,
	})
	if err != nil {
		logx.Errorf("GetRepoByIdRPC: RPC Called Failed: %v", err.Error())
		err = errno.InternalServiceError.WithError(err)
		return
	}

	if !utils.IsSuccess(resp.Base) {
		err = errno.BizError.WithMessage(resp.Base.Message)
		return
	}

	if resp.Repo == nil {
		logx.Info("Repo Is Empty")
		return "", err
	}

	languages = resp.GetRepo().GetLanguage()

	return
}

func (l *UpdateLanguageLogic) updateLanguage(model *model_analysis.Languages) error {
	language, err := l.svcCtx.LanguagesModel.FindOneByDeveloperId(l.ctx, model.DeveloperId)
	if err != nil {
		switch {
		case errors.Is(err, model_analysis.ErrNotFound):
			logx.Info("No Found This Data!")

			dataId, err := l.svcCtx.LanguagesModel.CreateDataId()
			if err != nil {
				err = errno.InternalDatabaseError.WithError(err)
				return err
			}

			model.DataId = dataId
			_, err = l.svcCtx.LanguagesModel.Insert(l.ctx, model)
			if err != nil {
				err = errno.InternalDatabaseError.WithError(err)
				return err
			}

			return nil
		default:
			return err
		}
	}

	model.DataId = language.DataId
	err = l.svcCtx.LanguagesModel.Update(l.ctx, model)
	if err != nil {
		err = errno.InternalDatabaseError.WithError(err)
		return err
	}

	return nil
}
