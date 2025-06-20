package logic

import (
	"context"

	"github.com/wushiling50/aster/gen/developer"
	"github.com/wushiling50/aster/rpc/developer/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDeveloperByLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetDeveloperByLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDeveloperByLoginLogic {
	return &GetDeveloperByLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetDeveloperByLoginLogic) GetDeveloperByLogin(in *developer_developer.GetDeveloperByLoginReq) (*developer_developer.GetDeveloperByLoginResp, error) {
	// todo: add your logic here and delete this line

	return &developer_developer.GetDeveloperByLoginResp{}, nil
}
