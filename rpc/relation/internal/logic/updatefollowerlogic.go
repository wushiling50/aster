package logic

import (
	"context"

	"github.com/wushiling50/aster/gen/relation"
	"github.com/wushiling50/aster/rpc/relation/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateFollowerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateFollowerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateFollowerLogic {
	return &UpdateFollowerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateFollowerLogic) UpdateFollower(in *relation_relation.UpdateFollowerReq) (*relation_relation.UpdateFollowerResp, error) {
	// todo: add your logic here and delete this line

	return &relation_relation.UpdateFollowerResp{}, nil
}
