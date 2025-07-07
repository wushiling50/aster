package logic

import (
	"context"

	"github.com/wushiling50/aster/gen/relation"
	"github.com/wushiling50/aster/rpc/relation/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelAllStaringDevLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelAllStaringDevLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelAllStaringDevLogic {
	return &DelAllStaringDevLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelAllStaringDevLogic) DelAllStaringDev(in *relation.DelAllStaringDevReq) (*relation.DelAllStaringDevResp, error) {
	// todo: add your logic here and delete this line

	return &relation.DelAllStaringDevResp{}, nil
}
