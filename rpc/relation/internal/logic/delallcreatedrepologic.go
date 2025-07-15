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

type DelAllCreatedRepoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelAllCreatedRepoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelAllCreatedRepoLogic {
	return &DelAllCreatedRepoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelAllCreatedRepoLogic) DelAllCreatedRepo(in *relation.DelAllCreatedRepoReq) (*relation.DelAllCreatedRepoResp, error) {
	resp := new(relation.DelAllCreatedRepoResp)

	err := l.delAllCreatedRepo(in.DeveloperId, 1, math.MaxInt64)
	if err != nil {
		logx.Errorf("service.DelAllCreatedRepo: Del All Created Repo Failed: %w", err)
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.Base = pack.BuildSuccessResp()

	return resp, nil
}

func (l *DelAllCreatedRepoLogic) delAllCreatedRepo(developerId, page, limit int64) error {
	createRepos, err := l.svcCtx.CreateRepoModel.SearchCreatedRepo(l.ctx, developerId, page, limit)
	if err != nil {
		return err
	}

	eg := errgroup.Group{}
	for _, createdRepo := range createRepos {
		dataId := createdRepo.DataId

		eg.Go(func() error {
			return l.svcCtx.CreateRepoModel.Delete(l.ctx, dataId)
		})
	}

	if err := eg.Wait(); err != nil {
		return err
	}

	return nil
}
