package logic

import (
	"context"

	"github.com/wushiling50/aster/gen/relation"
	"github.com/wushiling50/aster/rpc/relation/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStarredRepoUpdatedAtLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetStarredRepoUpdatedAtLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStarredRepoUpdatedAtLogic {
	return &GetStarredRepoUpdatedAtLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetStarredRepoUpdatedAtLogic) GetStarredRepoUpdatedAt(in *relation.GetStarredRepoUpdatedAtReq) (*relation.GetStarredRepoUpdatedAtResp, error) {
	// todo: add your logic here and delete this line

	return &relation.GetStarredRepoUpdatedAtResp{}, nil
}
