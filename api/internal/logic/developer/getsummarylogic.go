package developer

import (
	"context"

	"github.com/wushiling50/aster/api/internal/svc"
	"github.com/wushiling50/aster/api/internal/types"

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

func (l *GetSummaryLogic) GetSummary(req *types.GetSummaryReq) (resp *types.GetSummaryResp, err error) {
	// todo: add your logic here and delete this line

	return
}
