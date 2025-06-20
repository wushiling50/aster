package logic

import (
	"context"

	"github.com/wushiling50/aster/gen/relation"
	"github.com/wushiling50/aster/rpc/relation/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateStarredRepoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateStarredRepoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateStarredRepoLogic {
	return &UpdateStarredRepoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateStarredRepoLogic) UpdateStarredRepo(in *relation_relation.UpdateStarredRepoReq) (*relation_relation.UpdateStarredRepoResp, error) {
	// todo: add your logic here and delete this line

	return &relation_relation.UpdateStarredRepoResp{}, nil
}
