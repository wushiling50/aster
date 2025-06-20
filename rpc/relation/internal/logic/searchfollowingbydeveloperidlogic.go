package logic

import (
	"context"

	"github.com/wushiling50/aster/gen/relation"
	"github.com/wushiling50/aster/rpc/relation/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchFollowingByDeveloperIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchFollowingByDeveloperIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchFollowingByDeveloperIdLogic {
	return &SearchFollowingByDeveloperIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchFollowingByDeveloperIdLogic) SearchFollowingByDeveloperId(in *relation_relation.SearchFollowingByDeveloperIdReq) (*relation_relation.SearchFollowingByDeveloperIdResp, error) {
	// todo: add your logic here and delete this line

	return &relation_relation.SearchFollowingByDeveloperIdResp{}, nil
}
