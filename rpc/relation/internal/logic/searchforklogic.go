package logic

import (
	"context"

	"github.com/wushiling50/aster/gen/relation"
	"github.com/wushiling50/aster/rpc/relation/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchForkLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchForkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchForkLogic {
	return &SearchForkLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchForkLogic) SearchFork(in *relation.SearchForkReq) (*relation.SearchForkResp, error) {
	// todo: add your logic here and delete this line

	return &relation.SearchForkResp{}, nil
}
