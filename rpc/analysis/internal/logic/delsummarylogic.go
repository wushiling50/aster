package logic

import (
	"context"

	"github.com/wushiling50/aster/gen/analysis"
	"github.com/wushiling50/aster/rpc/analysis/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelSummaryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelSummaryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelSummaryLogic {
	return &DelSummaryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelSummaryLogic) DelSummary(in *analysis.DelAnalysisReq) (*analysis.DelAnalysisResp, error) {
	// todo: add your logic here and delete this line

	return &analysis.DelAnalysisResp{}, nil
}
