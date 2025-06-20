package logic

import (
	"context"

	"github.com/wushiling50/aster/gen/relation"
	"github.com/wushiling50/aster/rpc/relation/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelStarLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelStarLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelStarLogic {
	return &DelStarLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelStarLogic) DelStar(in *relation_relation.DelStarReq) (*relation_relation.DelStarResp, error) {
	// todo: add your logic here and delete this line

	return &relation_relation.DelStarResp{}, nil
}
