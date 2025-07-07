package logic

import (
	"context"
	"math"

	"github.com/wushiling50/aster/gen/relation"
	"github.com/wushiling50/aster/rpc/relation/internal/pack"
	"github.com/wushiling50/aster/rpc/relation/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelAllFollowingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelAllFollowingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelAllFollowingLogic {
	return &DelAllFollowingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelAllFollowingLogic) DelAllFollowing(in *relation.DelAllFollowingReq) (*relation.DelAllFollowingResp, error) {
	resp := new(relation.DelAllFollowingResp)

	err := l.delAllFollowing(in.DeveloperId, 1, math.MaxInt64)
	if err != nil {
		logx.Errorf("service.DelAllFollowing: Del All Following Failed: %w", err)
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.Base = pack.BuildSuccessResp()

	return resp, nil
}

func (l *DelAllFollowingLogic) delAllFollowing(developerId, page, limit int64) error {
	followings, err := l.svcCtx.FollowModel.SearchFollowingByDeveloperId(l.ctx, developerId, page, limit)
	if err != nil {
		return err
	}

	for _, following := range followings {
		err := l.svcCtx.FollowModel.Delete(l.ctx, following.DataId)
		if err != nil {
			return err
		}
	}

	return nil
}
