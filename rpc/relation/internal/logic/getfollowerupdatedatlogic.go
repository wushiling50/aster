package logic

import (
	"context"

	"github.com/wushiling50/aster/gen/relation"
	"github.com/wushiling50/aster/rpc/relation/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFollowerUpdatedAtLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFollowerUpdatedAtLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFollowerUpdatedAtLogic {
	return &GetFollowerUpdatedAtLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFollowerUpdatedAtLogic) GetFollowerUpdatedAt(in *relation_relation.GetFollowerUpdatedAtReq) (*relation_relation.GetFollowerUpdatedAtResp, error) {
	// todo: add your logic here and delete this line

	return &relation_relation.GetFollowerUpdatedAtResp{}, nil
}
