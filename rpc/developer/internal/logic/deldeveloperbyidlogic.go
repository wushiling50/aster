package logic

import (
	"context"

	"github.com/wushiling50/aster/gen/developer"
	"github.com/wushiling50/aster/rpc/developer/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelDeveloperByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelDeveloperByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelDeveloperByIdLogic {
	return &DelDeveloperByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelDeveloperByIdLogic) DelDeveloperById(in *developer.DelDeveloperByIdReq) (*developer.DelDeveloperByIdResp, error) {
	// todo: add your logic here and delete this line

	return &developer.DelDeveloperByIdResp{}, nil
}
