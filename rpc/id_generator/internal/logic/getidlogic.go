package logic

import (
	"context"

	"github.com/wushiling50/aster/gen/id_generator"
	"github.com/wushiling50/aster/rpc/id_generator/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetIdLogic {
	return &GetIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetIdLogic) GetId(in *id_generator_id_generator.GetIdReq) (*id_generator_id_generator.GetIdResp, error) {
	// todo: add your logic here and delete this line

	return &id_generator_id_generator.GetIdResp{}, nil
}
