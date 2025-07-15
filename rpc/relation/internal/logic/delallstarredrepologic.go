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

type DelAllStarredRepoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelAllStarredRepoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelAllStarredRepoLogic {
	return &DelAllStarredRepoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelAllStarredRepoLogic) DelAllStarredRepo(in *relation.DelAllStarredRepoReq) (*relation.DelAllStarredRepoResp, error) {
	resp := new(relation.DelAllStarredRepoResp)

	err := l.delAllStarredRepo(in.DeveloperId, 1, math.MaxInt64)
	if err != nil {
		logx.Errorf("service.DelAllStarredRepo: Del All Starred Repo Failed: %w", err)
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.Base = pack.BuildSuccessResp()

	return resp, nil
}

func (l *DelAllStarredRepoLogic) delAllStarredRepo(developerId, page, limit int64) error {
	stars, err := l.svcCtx.StarModel.SearchStarredRepo(l.ctx, developerId, page, limit)
	if err != nil {
		return err
	}

	eg := errgroup.Group{}
	for _, star := range stars {
		dataId := star.DataId

		eg.Go(func() error {
			return l.svcCtx.StarModel.Delete(l.ctx, dataId)
		})
	}

	if err := eg.Wait(); err != nil {
		return err
	}

	return nil
}
