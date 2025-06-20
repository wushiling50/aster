package logic

import (
	"context"

	"github.com/wushiling50/aster/gen/analysis"
	"github.com/wushiling50/aster/rpc/analysis/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetNationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetNationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetNationLogic {
	return &GetNationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetNationLogic) GetNation(in *analysis_analysis.GetAnalysisReq) (*analysis_analysis.GetNationResp, error) {
	// todo: add your logic here and delete this line

	return &analysis_analysis.GetNationResp{}, nil
}
