package logic

import (
	"context"

	"github.com/wushiling50/aster/gen/contribution"
	"github.com/wushiling50/aster/rpc/contribution/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetReviewOfUserUpdatedAtLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetReviewOfUserUpdatedAtLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetReviewOfUserUpdatedAtLogic {
	return &GetReviewOfUserUpdatedAtLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetReviewOfUserUpdatedAtLogic) GetReviewOfUserUpdatedAt(in *contribution.GetReviewOfUserUpdatedAtReq) (*contribution.GetReviewOfUserUpdatedAtResp, error) {
	// todo: add your logic here and delete this line

	return &contribution.GetReviewOfUserUpdatedAtResp{}, nil
}
