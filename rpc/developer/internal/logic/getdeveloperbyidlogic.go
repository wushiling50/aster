package logic

import (
	"context"

	"github.com/wushiling50/aster/gen/developer"
	"github.com/wushiling50/aster/rpc/developer/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDeveloperByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetDeveloperByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDeveloperByIdLogic {
	return &GetDeveloperByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetDeveloperByIdLogic) GetDeveloperById(in *developer_developer.GetDeveloperByIdReq) (*developer_developer.GetDeveloperByIdResp, error) {
	// todo: add your logic here and delete this line

	return &developer_developer.GetDeveloperByIdResp{}, nil
}
