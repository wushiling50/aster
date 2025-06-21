package logic

import (
	"context"

	"github.com/wushiling50/aster/gen/repo"
	"github.com/wushiling50/aster/rpc/repo/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRepoByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRepoByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRepoByIdLogic {
	return &GetRepoByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetRepoByIdLogic) GetRepoById(in *repo.GetRepoByIdReq) (*repo.GetRepoByIdResp, error) {
	// todo: add your logic here and delete this line

	return &repo.GetRepoByIdResp{}, nil
}
