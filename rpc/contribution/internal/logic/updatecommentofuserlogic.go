package logic

import (
	"context"

	"github.com/wushiling50/aster/gen/contribution"
	"github.com/wushiling50/aster/rpc/contribution/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCommentOfUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCommentOfUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCommentOfUserLogic {
	return &UpdateCommentOfUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateCommentOfUserLogic) UpdateCommentOfUser(in *contribution_contribution.UpdateCommentOfUserReq) (*contribution_contribution.UpdateCommentOfUserResp, error) {
	// todo: add your logic here and delete this line

	return &contribution_contribution.UpdateCommentOfUserResp{}, nil
}
