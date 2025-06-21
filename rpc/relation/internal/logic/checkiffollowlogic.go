package logic

import (
	"context"

	"github.com/wushiling50/aster/gen/relation"
	"github.com/wushiling50/aster/rpc/relation/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckIfFollowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckIfFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckIfFollowLogic {
	return &CheckIfFollowLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckIfFollowLogic) CheckIfFollow(in *relation.CheckIfFollowReq) (*relation.CheckFollowResp, error) {
	// todo: add your logic here and delete this line

	return &relation.CheckFollowResp{}, nil
}
