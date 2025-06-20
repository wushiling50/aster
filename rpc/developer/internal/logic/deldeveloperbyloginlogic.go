package logic

import (
	"context"

	"github.com/wushiling50/aster/gen/developer"
	"github.com/wushiling50/aster/rpc/developer/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelDeveloperByLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelDeveloperByLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelDeveloperByLoginLogic {
	return &DelDeveloperByLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelDeveloperByLoginLogic) DelDeveloperByLogin(in *developer_developer.DelDeveloperByLoginReq) (*developer_developer.DelDeveloperByLoginResp, error) {
	// todo: add your logic here and delete this line

	return &developer_developer.DelDeveloperByLoginResp{}, nil
}
