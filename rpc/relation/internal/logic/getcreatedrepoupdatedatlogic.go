package logic

import (
	"context"

	"github.com/wushiling50/aster/gen/relation"
	"github.com/wushiling50/aster/rpc/relation/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCreatedRepoUpdatedAtLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCreatedRepoUpdatedAtLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCreatedRepoUpdatedAtLogic {
	return &GetCreatedRepoUpdatedAtLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCreatedRepoUpdatedAtLogic) GetCreatedRepoUpdatedAt(in *relation.GetCreatedRepoUpdatedAtReq) (*relation.GetCreatedRepoUpdatedAtResp, error) {
	// todo: add your logic here and delete this line

	return &relation.GetCreatedRepoUpdatedAtResp{}, nil
}
