package logic

import (
	"context"

	"github.com/wushiling50/aster/gen/contribution"
	"github.com/wushiling50/aster/rpc/contribution/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchByDeveloperIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchByDeveloperIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchByDeveloperIdLogic {
	return &SearchByDeveloperIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchByDeveloperIdLogic) SearchByDeveloperId(in *contribution.SearchByDeveloperIdReq) (*contribution.SearchByDeveloperIdResp, error) {
	// todo: add your logic here and delete this line

	return &contribution.SearchByDeveloperIdResp{}, nil
}
