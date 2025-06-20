package logic

import (
	"context"

	"github.com/wushiling50/aster/gen/analysis"
	"github.com/wushiling50/aster/rpc/analysis/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelScoreLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelScoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelScoreLogic {
	return &DelScoreLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelScoreLogic) DelScore(in *analysis_analysis.DelAnalysisReq) (*analysis_analysis.DelAnalysisResp, error) {
	// todo: add your logic here and delete this line

	return &analysis_analysis.DelAnalysisResp{}, nil
}
