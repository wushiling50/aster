package logic

import (
	"context"

	"github.com/wushiling50/aster/gen/analysis"
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

func (l *UpdateSummaryLogic) UpdateSummary(in *analysis_analysis.UpdateAnalysisReq) (*analysis_analysis.UpdateAnalysisResp, error) {
	// todo: add your logic here and delete this line

	return &analysis_analysis.UpdateAnalysisResp{}, nil
}
