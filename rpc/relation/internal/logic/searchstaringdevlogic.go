package logic

import (
	"context"

	"github.com/wushiling50/aster/gen/relation"
	"github.com/wushiling50/aster/rpc/relation/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchStaringDevLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchStaringDevLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchStaringDevLogic {
	return &SearchStaringDevLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchStaringDevLogic) SearchStaringDev(in *relation_relation.SearchStaringDevReq) (*relation_relation.SearchStaringDevResp, error) {
	// todo: add your logic here and delete this line

	return &relation_relation.SearchStaringDevResp{}, nil
}
