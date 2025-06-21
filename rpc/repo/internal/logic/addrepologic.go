package logic

import (
	"context"

	"github.com/wushiling50/aster/gen/repo"
	"github.com/wushiling50/aster/rpc/repo/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddRepoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddRepoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddRepoLogic {
	return &AddRepoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddRepoLogic) AddRepo(in *repo.AddRepoReq) (*repo.AddRepoResp, error) {
	// todo: add your logic here and delete this line

	return &repo.AddRepoResp{}, nil
}
