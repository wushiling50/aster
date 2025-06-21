package logic

import (
	"context"

	"github.com/wushiling50/aster/gen/relation"
	"github.com/wushiling50/aster/rpc/relation/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchStarredRepoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchStarredRepoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchStarredRepoLogic {
	return &SearchStarredRepoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchStarredRepoLogic) SearchStarredRepo(in *relation.SearchStarredRepoReq) (*relation.SearchStarredRepoResp, error) {
	// todo: add your logic here and delete this line

	return &relation.SearchStarredRepoResp{}, nil
}
