package logic

import (
	"context"

	"github.com/wushiling50/aster/gen/contribution"
	"github.com/wushiling50/aster/rpc/contribution/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelAllContributionInCategoryByDeveloperIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelAllContributionInCategoryByDeveloperIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelAllContributionInCategoryByDeveloperIdLogic {
	return &DelAllContributionInCategoryByDeveloperIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelAllContributionInCategoryByDeveloperIdLogic) DelAllContributionInCategoryByDeveloperId(in *contribution.DelAllContributionInCategoryByDeveloperIdReq) (*contribution.DelAllContributionInCategoryByDeveloperIdResp, error) {
	// todo: add your logic here and delete this line

	return &contribution.DelAllContributionInCategoryByDeveloperIdResp{}, nil
}
