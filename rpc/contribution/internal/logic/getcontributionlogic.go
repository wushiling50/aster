package logic

import (
	"context"

	"github.com/wushiling50/aster/gen/contribution"
	"github.com/wushiling50/aster/rpc/contribution/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetContributionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetContributionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetContributionLogic {
	return &GetContributionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetContributionLogic) GetContribution(in *contribution_contribution.GetContributionReq) (*contribution_contribution.GetContributionResp, error) {
	// todo: add your logic here and delete this line

	return &contribution_contribution.GetContributionResp{}, nil
}
