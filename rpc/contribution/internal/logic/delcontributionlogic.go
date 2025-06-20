package logic

import (
	"context"

	"github.com/wushiling50/aster/gen/contribution"
	"github.com/wushiling50/aster/rpc/contribution/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelContributionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelContributionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelContributionLogic {
	return &DelContributionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelContributionLogic) DelContribution(in *contribution_contribution.DelContributionReq) (*contribution_contribution.DelContributionResp, error) {
	// todo: add your logic here and delete this line

	return &contribution_contribution.DelContributionResp{}, nil
}
