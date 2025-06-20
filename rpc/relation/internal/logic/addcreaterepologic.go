package logic

import (
	"context"

	"github.com/wushiling50/aster/gen/relation"
	"github.com/wushiling50/aster/rpc/relation/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCreateRepoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddCreateRepoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCreateRepoLogic {
	return &AddCreateRepoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------CreateRepo-----------------------
func (l *AddCreateRepoLogic) AddCreateRepo(in *relation_relation.AddCreateRepoReq) (*relation_relation.AddCreateRepoResp, error) {
	// todo: add your logic here and delete this line

	return &relation_relation.AddCreateRepoResp{}, nil
}
