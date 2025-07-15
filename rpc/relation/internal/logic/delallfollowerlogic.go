package logic

import (
	"context"
	"math"

	"github.com/wushiling50/aster/gen/relation"
	"github.com/wushiling50/aster/rpc/relation/internal/pack"
	"github.com/wushiling50/aster/rpc/relation/internal/svc"
	"golang.org/x/sync/errgroup"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelAllFollowerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelAllFollowerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelAllFollowerLogic {
	return &DelAllFollowerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelAllFollowerLogic) DelAllFollower(in *relation.DelAllFollowerReq) (*relation.DelAllFollowerResp, error) {
	resp := new(relation.DelAllFollowerResp)

	err := l.delAllFollower(in.DeveloperId, 1, math.MaxInt64)
	if err != nil {
		logx.Errorf("service.DelAllFollower: Del All Follower Failed: %w", err)
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.Base = pack.BuildSuccessResp()

	return resp, nil
}

func (l *DelAllFollowerLogic) delAllFollower(developerId, page, limit int64) error {
	followers, err := l.svcCtx.FollowModel.SearchFollowerByDeveloperId(l.ctx, developerId, page, limit)
	if err != nil {
		return err
	}

	eg := errgroup.Group{}
	for _, follower := range followers {
		dataId := follower.DataId

		eg.Go(func() error {
			return l.svcCtx.FollowModel.Delete(l.ctx, dataId)
		})
	}

	if err := eg.Wait(); err != nil {
		return err
	}

	return nil
}
