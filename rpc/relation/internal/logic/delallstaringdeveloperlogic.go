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
	resp := new(relation.DelAllStaringDeveloperResp)

	err := l.delAllStaringDeveloper(in.RepoId, 1, math.MaxInt64)
	if err != nil {
		logx.Errorf("service.DelAllStaringDeveloper: Del All Staring Developer Failed: %w", err)
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.Base = pack.BuildSuccessResp()

	return resp, nil
}

func (l *DelAllStaringDeveloperLogic) delAllStaringDeveloper(repoId, page, limit int64) error {
	stars, err := l.svcCtx.StarModel.SearchStaringDeveloper(l.ctx, repoId, page, limit)
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
