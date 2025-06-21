package logic

import (
	"context"

	"github.com/wushiling50/aster/gen/relation"
	"github.com/wushiling50/aster/rpc/relation/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelAllFollowingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelAllFollowingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelAllFollowingLogic {
	return &DelAllFollowingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelAllFollowingLogic) DelAllFollowing(in *relation.DelAllFollowingReq) (*relation.DelAllFollowingResp, error) {
	// todo: add your logic here and delete this line

	return &relation.DelAllFollowingResp{}, nil
}
