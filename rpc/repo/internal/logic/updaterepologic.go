package logic

import (
	"context"

	"github.com/wushiling50/aster/gen/repo"
	"github.com/wushiling50/aster/rpc/repo/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateRepoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateRepoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRepoLogic {
	return &UpdateRepoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateRepoLogic) UpdateRepo(in *repo_repo.UpdateRepoReq) (*repo_repo.UpdateRepoResp, error) {
	// todo: add your logic here and delete this line

	return &repo_repo.UpdateRepoResp{}, nil
}
