package logic

import (
	"context"
	"encoding/json"

	"github.com/wushiling50/aster/pkg/errno"
	"github.com/wushiling50/aster/pkg/utils"
	analysis "github.com/wushiling50/aster/rpc/analysis/analysisclient"
	"github.com/wushiling50/aster/rpc/api_processor/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetLanguageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetLanguageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLanguageLogic {
	return &GetLanguageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetLanguageLogic) GetLanguage(developerId int64) ([]byte, error) {
	if err := l.rpcUpdateLanguage(developerId); err != nil {
		logx.Error(err)
		return nil, err
	}

	resp, err := l.rpcGetLanguage(developerId)
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	data, err := json.Marshal(resp.Languages)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (l *GetLanguageLogic) rpcUpdateLanguage(id int64) (err error) {
	var resp *analysis.UpdateAnalysisResp

	resp, err = l.svcCtx.AnalysisRpcClient.UpdateLanguage(l.ctx, &analysis.UpdateAnalysisReq{
		DeveloperId: id,
	})
	if err != nil {
		logx.Errorf("UpdateLanguageRPC: RPC called failed: %v", err.Error())
		err = errno.InternalServiceError.WithError(err)
		return
	}

	if !utils.IsSuccess(resp.Base) {
		err = errno.BizError.WithMessage(resp.Base.Message)
		return
	}

	return
}

func (l *GetLanguageLogic) rpcGetLanguage(id int64) (resp *analysis.GetLanguagesResp, err error) {
	resp, err = l.svcCtx.AnalysisRpcClient.GetLanguages(l.ctx, &analysis.GetAnalysisReq{
		DeveloperId: id,
	})
	if err != nil {
		logx.Errorf("GetLanguageRPC: RPC called failed: %v", err.Error())
		err = errno.InternalServiceError.WithError(err)
		return
	}

	if !utils.IsSuccess(resp.Base) {
		err = errno.BizError.WithMessage(resp.Base.Message)
		return
	}

	return
}
