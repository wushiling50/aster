package logic

import (
	"context"

	"github.com/wushiling50/aster/gen/relation"
	"github.com/wushiling50/aster/rpc/relation/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateFollowingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateFollowingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateFollowingLogic {
	return &UpdateFollowingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateFollowingLogic) UpdateFollowing(in *relation_relation.UpdateFollowingReq) (*relation_relation.UpdateFollowingResp, error) {
	// todo: add your logic here and delete this line

	return &relation_relation.UpdateFollowingResp{}, nil
}
