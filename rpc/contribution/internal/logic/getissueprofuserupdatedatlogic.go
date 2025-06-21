package logic

import (
	"context"

	"github.com/wushiling50/aster/gen/contribution"
	"github.com/wushiling50/aster/rpc/contribution/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetIssuePROfUserUpdatedAtLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetIssuePROfUserUpdatedAtLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetIssuePROfUserUpdatedAtLogic {
	return &GetIssuePROfUserUpdatedAtLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetIssuePROfUserUpdatedAtLogic) GetIssuePROfUserUpdatedAt(in *contribution.GetIssuePROfUserUpdatedAtReq) (*contribution.GetIssuePROfUserUpdatedAtResp, error) {
	// todo: add your logic here and delete this line

	return &contribution.GetIssuePROfUserUpdatedAtResp{}, nil
}
