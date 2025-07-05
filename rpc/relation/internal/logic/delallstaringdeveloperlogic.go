package logic

import (
	"context"

	"github.com/wushiling50/aster/gen/relation"
	"github.com/wushiling50/aster/rpc/relation/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelAllStaringDeveloperLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelAllStaringDeveloperLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelAllStaringDeveloperLogic {
	return &DelAllStaringDeveloperLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelAllStaringDeveloperLogic) DelAllStaringDeveloper(in *relation.DelAllStaringDeveloperReq) (*relation.DelAllStaringDeveloperResp, error) {
	// todo: add your logic here and delete this line

	return &relation.DelAllStaringDeveloperResp{}, nil
}
