package logic

import (
	"context"

	"github.com/wushiling50/aster/gen/relation"
	"github.com/wushiling50/aster/rpc/relation/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCreateRepoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCreateRepoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCreateRepoLogic {
	return &UpdateCreateRepoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateCreateRepoLogic) UpdateCreateRepo(in *relation_relation.UpdateCreateRepoReq) (*relation_relation.UpdateCreateRepoResp, error) {
	// todo: add your logic here and delete this line

	return &relation_relation.UpdateCreateRepoResp{}, nil
}
