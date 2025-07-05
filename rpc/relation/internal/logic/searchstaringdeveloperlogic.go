package logic

import (
	"context"

	"github.com/wushiling50/aster/gen/relation"
	"github.com/wushiling50/aster/rpc/relation/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchStaringDeveloperLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchStaringDeveloperLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchStaringDeveloperLogic {
	return &SearchStaringDeveloperLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchStaringDeveloperLogic) SearchStaringDeveloper(in *relation.SearchStaringDeveloperReq) (*relation.SearchStaringDeveloperResp, error) {
	// todo: add your logic here and delete this line

	return &relation.SearchStaringDeveloperResp{}, nil
}
