package logic

import (
	"context"

	"github.com/wushiling50/aster/gen/relation"
	"github.com/wushiling50/aster/rpc/relation/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelAllCreatedRepoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelAllCreatedRepoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelAllCreatedRepoLogic {
	return &DelAllCreatedRepoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelAllCreatedRepoLogic) DelAllCreatedRepo(in *relation.DelAllCreatedRepoReq) (*relation.DelAllCreatedRepoResp, error) {
	// todo: add your logic here and delete this line

	return &relation.DelAllCreatedRepoResp{}, nil
}
