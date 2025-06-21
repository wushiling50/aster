package logic

import (
	"context"

	"github.com/wushiling50/aster/gen/contribution"
	"github.com/wushiling50/aster/rpc/contribution/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddContributionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddContributionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddContributionLogic {
	return &AddContributionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddContributionLogic) AddContribution(in *contribution.AddContributionReq) (*contribution.AddContributionResp, error) {
	// todo: add your logic here and delete this line

	return &contribution.AddContributionResp{}, nil
}
