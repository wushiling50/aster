package logic

import (
	"context"

	"github.com/wushiling50/aster/gen/analysis"
	"github.com/wushiling50/aster/rpc/analysis/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelNationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelNationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelNationLogic {
	return &DelNationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelNationLogic) DelNation(in *analysis.DelAnalysisReq) (*analysis.DelAnalysisResp, error) {
	// todo: add your logic here and delete this line

	return &analysis.DelAnalysisResp{}, nil
}
