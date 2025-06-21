package logic

import (
	"context"

	"github.com/wushiling50/aster/gen/relation"
	"github.com/wushiling50/aster/rpc/relation/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCreatorIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCreatorIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCreatorIdLogic {
	return &GetCreatorIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCreatorIdLogic) GetCreatorId(in *relation.GetCreatorIdReq) (*relation.GetCreatorIdResp, error) {
	// todo: add your logic here and delete this line

	return &relation.GetCreatorIdResp{}, nil
}
