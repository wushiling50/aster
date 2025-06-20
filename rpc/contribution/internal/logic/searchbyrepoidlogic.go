package logic

import (
	"context"

	"github.com/wushiling50/aster/gen/contribution"
	"github.com/wushiling50/aster/rpc/contribution/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchByRepoIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchByRepoIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchByRepoIdLogic {
	return &SearchByRepoIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchByRepoIdLogic) SearchByRepoId(in *contribution_contribution.SearchByRepoIdReq) (*contribution_contribution.SearchByRepoIdResp, error) {
	// todo: add your logic here and delete this line

	return &contribution_contribution.SearchByRepoIdResp{}, nil
}
