package logic

import (
	"context"

	"github.com/wushiling50/aster/gen/relation"
	"github.com/wushiling50/aster/rpc/relation/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelCreateRepoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelCreateRepoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelCreateRepoLogic {
	return &DelCreateRepoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelCreateRepoLogic) DelCreateRepo(in *relation_relation.DelCreateRepoReq) (*relation_relation.DelCreateRepoResp, error) {
	// todo: add your logic here and delete this line

	return &relation_relation.DelCreateRepoResp{}, nil
}
