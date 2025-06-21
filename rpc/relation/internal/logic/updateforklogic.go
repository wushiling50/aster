package logic

import (
	"context"

	"github.com/wushiling50/aster/gen/relation"
	"github.com/wushiling50/aster/rpc/relation/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateForkLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateForkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateForkLogic {
	return &UpdateForkLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateForkLogic) UpdateFork(in *relation.UpdateForkReq) (*relation.UpdateForkResp, error) {
	// todo: add your logic here and delete this line

	return &relation.UpdateForkResp{}, nil
}
