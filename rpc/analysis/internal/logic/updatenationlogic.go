package logic

import (
	"context"

	"github.com/wushiling50/aster/gen/analysis"
	"github.com/wushiling50/aster/rpc/analysis/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateNationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateNationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateNationLogic {
	return &UpdateNationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateNationLogic) UpdateNation(in *analysis.UpdateAnalysisReq) (*analysis.UpdateAnalysisResp, error) {
	// todo: add your logic here and delete this line

	return &analysis.UpdateAnalysisResp{}, nil
}
