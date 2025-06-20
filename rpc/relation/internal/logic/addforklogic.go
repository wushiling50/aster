package logic

import (
	"context"

	"github.com/wushiling50/aster/gen/relation"
	"github.com/wushiling50/aster/rpc/relation/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddForkLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddForkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddForkLogic {
	return &AddForkLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------Fork-----------------------
func (l *AddForkLogic) AddFork(in *relation_relation.AddForkReq) (*relation_relation.AddForkResp, error) {
	// todo: add your logic here and delete this line

	return &relation_relation.AddForkResp{}, nil
}
