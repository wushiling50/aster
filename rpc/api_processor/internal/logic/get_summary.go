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

type GetSummaryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSummaryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSummaryLogic {
	return &GetSummaryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSummaryLogic) GetSummary(developerId int64) ([]byte, error) {
	if err := l.rpcUpdateSummary(developerId); err != nil {
		logx.Error(err)
		return nil, err
	}

	resp, err := l.rpcGetSummary(developerId)
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	data, err := json.Marshal(resp.Summary)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (l *GetSummaryLogic) rpcUpdateSummary(id int64) (err error) {
	var resp *analysis.UpdateAnalysisResp

	resp, err = l.svcCtx.AnalysisRpcClient.UpdateSummary(l.ctx, &analysis.UpdateAnalysisReq{
		DeveloperId: id,
	})
	if err != nil {
		logx.Errorf("UpdateSummaryRPC: RPC called failed: %v", err.Error())
		err = errno.InternalServiceError.WithError(err)
		return
	}

	if !utils.IsSuccess(resp.Base) {
		err = errno.BizError.WithMessage(resp.Base.Message)
		return
	}

	return
}

func (l *GetSummaryLogic) rpcGetSummary(id int64) (resp *analysis.GetSummaryResp, err error) {
	resp, err = l.svcCtx.AnalysisRpcClient.GetSummary(l.ctx, &analysis.GetAnalysisReq{
		DeveloperId: id,
	})
	if err != nil {
		logx.Errorf("GetSummaryRPC: RPC called failed: %v", err.Error())
		err = errno.InternalServiceError.WithError(err)
		return
	}

	if !utils.IsSuccess(resp.Base) {
		err = errno.BizError.WithMessage(resp.Base.Message)
		return
	}

	return
}
