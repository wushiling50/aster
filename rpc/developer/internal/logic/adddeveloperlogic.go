package logic

import (
	"context"

	"github.com/wushiling50/aster/gen/developer"
	"github.com/wushiling50/aster/rpc/developer/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddDeveloperLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddDeveloperLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddDeveloperLogic {
	return &AddDeveloperLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------developer-----------------------
func (l *AddDeveloperLogic) AddDeveloper(in *developer_developer.AddDeveloperReq) (*developer_developer.AddDeveloperResp, error) {
	// todo: add your logic here and delete this line

	return &developer_developer.AddDeveloperResp{}, nil
}
