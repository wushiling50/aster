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

type GetScoreLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetScoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetScoreLogic {
	return &GetScoreLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetScoreLogic) GetScore(developerId int64) ([]byte, error) {
	if err := l.rpcUpdateScore(developerId); err != nil {
		logx.Error(err)
		return nil, err
	}

	resp, err := l.rpcGetScore(developerId)
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	data, err := json.Marshal(resp.Score)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (l *GetScoreLogic) rpcUpdateScore(id int64) (err error) {
	var resp *analysis.UpdateAnalysisResp

	resp, err = l.svcCtx.AnalysisRpcClient.UpdateScore(l.ctx, &analysis.UpdateAnalysisReq{
		DeveloperId: id,
	})
	if err != nil {
		logx.Errorf("UpdateScoreRPC: RPC called failed: %v", err.Error())
		err = errno.InternalServiceError.WithError(err)
		return
	}

	if !utils.IsSuccess(resp.Base) {
		err = errno.BizError.WithMessage(resp.Base.Message)
		return
	}

	return
}

func (l *GetScoreLogic) rpcGetScore(id int64) (resp *analysis.GetScoreResp, err error) {
	resp, err = l.svcCtx.AnalysisRpcClient.GetScore(l.ctx, &analysis.GetAnalysisReq{
		DeveloperId: id,
	})
	if err != nil {
		logx.Errorf("GetScoreRPC: RPC called failed: %v", err.Error())
		err = errno.InternalServiceError.WithError(err)
		return
	}

	if !utils.IsSuccess(resp.Base) {
		err = errno.BizError.WithMessage(resp.Base.Message)
		return
	}

	return
}
