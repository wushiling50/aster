package logic

import (
	"context"

	"github.com/wushiling50/aster/gen/relation"
	"github.com/wushiling50/aster/rpc/relation/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelFollowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelFollowLogic {
	return &DelFollowLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelFollowLogic) DelFollow(in *relation_relation.DelFollowReq) (*relation_relation.DelFollowResp, error) {
	// todo: add your logic here and delete this line

	return &relation_relation.DelFollowResp{}, nil
}
