package logic

import (
	"context"

	"github.com/wushiling50/aster/gen/relation"
	"github.com/wushiling50/aster/rpc/relation/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelAllFollowerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelAllFollowerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelAllFollowerLogic {
	return &DelAllFollowerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelAllFollowerLogic) DelAllFollower(in *relation.DelAllFollowerReq) (*relation.DelAllFollowerResp, error) {
	// todo: add your logic here and delete this line

	return &relation.DelAllFollowerResp{}, nil
}
