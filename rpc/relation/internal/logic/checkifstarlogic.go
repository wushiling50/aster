package logic

import (
	"context"

	"github.com/wushiling50/aster/gen/relation"
	"github.com/wushiling50/aster/rpc/relation/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckIfStarLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckIfStarLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckIfStarLogic {
	return &CheckIfStarLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckIfStarLogic) CheckIfStar(in *relation_relation.CheckIfStarReq) (*relation_relation.CheckIfStarResp, error) {
	// todo: add your logic here and delete this line

	return &relation_relation.CheckIfStarResp{}, nil
}
