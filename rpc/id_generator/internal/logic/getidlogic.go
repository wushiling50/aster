package logic

import (
	"context"

	gonanoid "github.com/matoous/go-nanoid"
	"github.com/wushiling50/aster/gen/id_generator"
	"github.com/wushiling50/aster/rpc/id_generator/internal/pack"
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

func (l *GetIdLogic) GetId(in *id_generator.GetIdReq) (*id_generator.GetIdResp, error) {
	resp := new(id_generator.GetIdResp)

	id, err := gonanoid.Nanoid()
	if err != nil {
		logx.Errorf("service.GetId: %w", err)
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.Id = id
	resp.Base = pack.BuildSuccessResp()

	return resp, nil
}
