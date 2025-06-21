package logic

import (
	"context"

	"github.com/wushiling50/aster/gen/contribution"
	"github.com/wushiling50/aster/rpc/contribution/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateContributionOfUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateContributionOfUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateContributionOfUserLogic {
	return &UpdateContributionOfUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateContributionOfUserLogic) UpdateContributionOfUser(in *contribution.UpdateContributionOfUserReq) (*contribution.UpdateContributionOfUserResp, error) {
	// todo: add your logic here and delete this line

	return &contribution.UpdateContributionOfUserResp{}, nil
}
