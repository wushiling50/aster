package logic

import (
	"context"

	"github.com/wushiling50/aster/gen/contribution"
	"github.com/wushiling50/aster/rpc/contribution/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelAllContributionInCategoryByUserIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelAllContributionInCategoryByUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelAllContributionInCategoryByUserIdLogic {
	return &DelAllContributionInCategoryByUserIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelAllContributionInCategoryByUserIdLogic) DelAllContributionInCategoryByUserId(in *contribution_contribution.DelAllContributionInCategoryByUserIdReq) (*contribution_contribution.DelAllContributionInCategoryByUserIdResp, error) {
	// todo: add your logic here and delete this line

	return &contribution_contribution.DelAllContributionInCategoryByUserIdResp{}, nil
}
