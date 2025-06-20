package logic

import (
	"context"

	"github.com/wushiling50/aster/gen/relation"
	"github.com/wushiling50/aster/rpc/relation/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOriginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOriginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOriginLogic {
	return &GetOriginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOriginLogic) GetOrigin(in *relation_relation.GetOriginReq) (*relation_relation.GetOriginResp, error) {
	// todo: add your logic here and delete this line

	return &relation_relation.GetOriginResp{}, nil
}
