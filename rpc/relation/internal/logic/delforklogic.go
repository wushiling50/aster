package logic

import (
	"context"

	"github.com/wushiling50/aster/gen/relation"
	"github.com/wushiling50/aster/rpc/relation/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelForkLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelForkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelForkLogic {
	return &DelForkLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelForkLogic) DelFork(in *relation_relation.DelForkReq) (*relation_relation.DelForkResp, error) {
	// todo: add your logic here and delete this line

	return &relation_relation.DelForkResp{}, nil
}
