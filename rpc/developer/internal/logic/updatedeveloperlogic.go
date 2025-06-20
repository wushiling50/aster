package logic

import (
	"context"

	"github.com/wushiling50/aster/gen/developer"
	"github.com/wushiling50/aster/rpc/developer/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateDeveloperLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateDeveloperLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateDeveloperLogic {
	return &UpdateDeveloperLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateDeveloperLogic) UpdateDeveloper(in *developer_developer.UpdateDeveloperReq) (*developer_developer.UpdateDeveloperResp, error) {
	// todo: add your logic here and delete this line

	return &developer_developer.UpdateDeveloperResp{}, nil
}
