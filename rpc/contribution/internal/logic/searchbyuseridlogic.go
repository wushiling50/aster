package logic

import (
	"context"

	"github.com/wushiling50/aster/gen/contribution"
	"github.com/wushiling50/aster/rpc/contribution/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchByUserIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchByUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchByUserIdLogic {
	return &SearchByUserIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchByUserIdLogic) SearchByUserId(in *contribution_contribution.SearchByUserIdReq) (*contribution_contribution.SearchByUserIdResp, error) {
	// todo: add your logic here and delete this line

	return &contribution_contribution.SearchByUserIdResp{}, nil
}
