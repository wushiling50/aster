package logic

import (
	"context"

	"github.com/wushiling50/aster/gen/analysis"
	"github.com/wushiling50/aster/rpc/analysis/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetScoreLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetScoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetScoreLogic {
	return &GetScoreLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetScoreLogic) GetScore(in *analysis_analysis.GetAnalysisReq) (*analysis_analysis.GetScoreResp, error) {
	// todo: add your logic here and delete this line

	return &analysis_analysis.GetScoreResp{}, nil
}
