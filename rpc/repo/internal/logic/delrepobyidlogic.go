package logic

import (
	"context"

	"github.com/wushiling50/aster/gen/repo"
	"github.com/wushiling50/aster/rpc/repo/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelRepoByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelRepoByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelRepoByIdLogic {
	return &DelRepoByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelRepoByIdLogic) DelRepoById(in *repo.DelRepoByIdReq) (*repo.DelRepoByIdResp, error) {
	// todo: add your logic here and delete this line

	return &repo.DelRepoByIdResp{}, nil
}
