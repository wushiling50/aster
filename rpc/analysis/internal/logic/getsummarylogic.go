package logic

import (
	"context"

	"github.com/wushiling50/aster/gen/analysis"
	"github.com/wushiling50/aster/rpc/analysis/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSummaryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSummaryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSummaryLogic {
	return &GetSummaryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetSummaryLogic) GetSummary(in *analysis.GetAnalysisReq) (*analysis.GetSummaryResp, error) {
	// todo: add your logic here and delete this line

	return &analysis.GetSummaryResp{}, nil
}
