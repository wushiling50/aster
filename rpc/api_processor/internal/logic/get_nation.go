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

type GetNationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetNationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetNationLogic {
	return &GetNationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetNationLogic) GetNation(developerId int64) ([]byte, error) {
	if err := l.rpcUpdateNation(developerId); err != nil {
		logx.Error(err)
		return nil, err
	}

	resp, err := l.rpcGetNation(developerId)
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	data, err := json.Marshal(resp.Nation)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (l *GetNationLogic) rpcUpdateNation(id int64) (err error) {
	var resp *analysis.UpdateAnalysisResp

	resp, err = l.svcCtx.AnalysisRpcClient.UpdateNation(l.ctx, &analysis.UpdateAnalysisReq{
		DeveloperId: id,
	})
	if err != nil {
		logx.Errorf("UpdateNationRPC: RPC called failed: %v", err.Error())
		err = errno.InternalServiceError.WithError(err)
		return
	}

	if !utils.IsSuccess(resp.Base) {
		err = errno.BizError.WithMessage(resp.Base.Message)
		return
	}

	return
}

func (l *GetNationLogic) rpcGetNation(id int64) (resp *analysis.GetNationResp, err error) {
	resp, err = l.svcCtx.AnalysisRpcClient.GetNation(l.ctx, &analysis.GetAnalysisReq{
		DeveloperId: id,
	})
	if err != nil {
		logx.Errorf("GetNationRPC: RPC called failed: %v", err.Error())
		err = errno.InternalServiceError.WithError(err)
		return
	}

	if !utils.IsSuccess(resp.Base) {
		err = errno.BizError.WithMessage(resp.Base.Message)
		return
	}

	return
}
