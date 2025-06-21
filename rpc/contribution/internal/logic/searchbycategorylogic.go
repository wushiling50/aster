package logic

import (
	"context"

	"github.com/wushiling50/aster/gen/contribution"
	"github.com/wushiling50/aster/rpc/contribution/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchByCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchByCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchByCategoryLogic {
	return &SearchByCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchByCategoryLogic) SearchByCategory(in *contribution.SearchByCategoryReq) (*contribution.SearchByCategoryResp, error) {
	// todo: add your logic here and delete this line

	return &contribution.SearchByCategoryResp{}, nil
}
