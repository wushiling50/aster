package logic

import (
	"context"

	"github.com/wushiling50/aster/gen/relation"
	"github.com/wushiling50/aster/rpc/relation/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelAllForkLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelAllForkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelAllForkLogic {
	return &DelAllForkLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelAllForkLogic) DelAllFork(in *relation_relation.DelAllForkReq) (*relation_relation.DelAllForkResp, error) {
	// todo: add your logic here and delete this line

	return &relation_relation.DelAllForkResp{}, nil
}
