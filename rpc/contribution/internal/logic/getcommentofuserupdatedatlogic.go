package logic

import (
	"context"

	"github.com/wushiling50/aster/gen/contribution"
	"github.com/wushiling50/aster/rpc/contribution/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentOfUserUpdatedAtLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCommentOfUserUpdatedAtLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentOfUserUpdatedAtLogic {
	return &GetCommentOfUserUpdatedAtLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCommentOfUserUpdatedAtLogic) GetCommentOfUserUpdatedAt(in *contribution_contribution.GetCommentOfUserUpdatedAtReq) (*contribution_contribution.GetCommentOfUserUpdatedAtResp, error) {
	// todo: add your logic here and delete this line

	return &contribution_contribution.GetCommentOfUserUpdatedAtResp{}, nil
}
