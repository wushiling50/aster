package logic

import (
	"context"

	"github.com/wushiling50/aster/gen/relation"
	"github.com/wushiling50/aster/rpc/relation/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFollowingUpdatedAtLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFollowingUpdatedAtLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFollowingUpdatedAtLogic {
	return &GetFollowingUpdatedAtLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFollowingUpdatedAtLogic) GetFollowingUpdatedAt(in *relation_relation.GetFollowingUpdatedAtReq) (*relation_relation.GetFollowingUpdatedAtResp, error) {
	// todo: add your logic here and delete this line

	return &relation_relation.GetFollowingUpdatedAtResp{}, nil
}
