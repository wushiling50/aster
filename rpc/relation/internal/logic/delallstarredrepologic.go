package logic

import (
	"context"

	"github.com/wushiling50/aster/gen/relation"
	"github.com/wushiling50/aster/rpc/relation/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelAllStarredRepoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelAllStarredRepoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelAllStarredRepoLogic {
	return &DelAllStarredRepoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelAllStarredRepoLogic) DelAllStarredRepo(in *relation.DelAllStarredRepoReq) (*relation.DelAllStarredRepoResp, error) {
	// todo: add your logic here and delete this line

	return &relation.DelAllStarredRepoResp{}, nil
}
