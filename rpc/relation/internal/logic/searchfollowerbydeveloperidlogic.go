package logic

import (
	"context"

	"github.com/wushiling50/aster/gen/relation"
	"github.com/wushiling50/aster/rpc/relation/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchFollowerByDeveloperIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchFollowerByDeveloperIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchFollowerByDeveloperIdLogic {
	return &SearchFollowerByDeveloperIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchFollowerByDeveloperIdLogic) SearchFollowerByDeveloperId(in *relation_relation.SearchFollowerByDeveloperIdReq) (*relation_relation.SearchFollowerByDeveloperIdResp, error) {
	// todo: add your logic here and delete this line

	return &relation_relation.SearchFollowerByDeveloperIdResp{}, nil
}
