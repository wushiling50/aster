package logic

import (
	"context"

	"github.com/wushiling50/aster/gen/contribution"
	"github.com/wushiling50/aster/rpc/contribution/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateReviewOfUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateReviewOfUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateReviewOfUserLogic {
	return &UpdateReviewOfUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateReviewOfUserLogic) UpdateReviewOfUser(in *contribution_contribution.UpdateReviewOfUserReq) (*contribution_contribution.UpdateReviewOfUserResp, error) {
	// todo: add your logic here and delete this line

	return &contribution_contribution.UpdateReviewOfUserResp{}, nil
}
