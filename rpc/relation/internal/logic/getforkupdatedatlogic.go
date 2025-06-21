package logic

import (
	"context"

	"github.com/wushiling50/aster/gen/relation"
	"github.com/wushiling50/aster/rpc/relation/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetForkUpdatedAtLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetForkUpdatedAtLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetForkUpdatedAtLogic {
	return &GetForkUpdatedAtLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetForkUpdatedAtLogic) GetForkUpdatedAt(in *relation.GetForkUpdatedAtReq) (*relation.GetForkUpdatedAtResp, error) {
	// todo: add your logic here and delete this line

	return &relation.GetForkUpdatedAtResp{}, nil
}
