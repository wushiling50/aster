package logic

import (
	"context"

	"github.com/wushiling50/aster/gen/contribution"
	"github.com/wushiling50/aster/rpc/contribution/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateIssuePROfUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateIssuePROfUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateIssuePROfUserLogic {
	return &UpdateIssuePROfUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateIssuePROfUserLogic) UpdateIssuePROfUser(in *contribution_contribution.UpdateIssuePROfUserReq) (*contribution_contribution.UpdateIssuePROfUserResp, error) {
	// todo: add your logic here and delete this line

	return &contribution_contribution.UpdateIssuePROfUserResp{}, nil
}
