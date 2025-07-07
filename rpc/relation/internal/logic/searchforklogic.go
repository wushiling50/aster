package logic

import (
	"context"

	"github.com/wushiling50/aster/gen/relation"
	"github.com/wushiling50/aster/rpc/relation/internal/pack"
	"github.com/wushiling50/aster/rpc/relation/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchForkLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchForkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchForkLogic {
	return &SearchForkLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchForkLogic) SearchFork(in *relation.SearchForkReq) (*relation.SearchForkResp, error) {
	resp := new(relation.SearchForkResp)

	forkRepoIds, err := l.searchFork(in.OriginalRepoId, in.Page, in.Limit)
	if err != nil {
		logx.Errorf("service.SearchFork: Search Fork Failed: %w", err)
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	if len(forkRepoIds) == 0 {
		logx.Info("service.SearchFork: No Found Fork")
	}

	resp.Base = pack.BuildSuccessResp()
	resp.ForkRepoIds = forkRepoIds

	return resp, nil
}

func (l *SearchForkLogic) searchFork(originalRepoId, page, limit int64) ([]int64, error) {
	forks, err := l.svcCtx.ForkModel.SearchFork(l.ctx, originalRepoId, page, limit)
	if err != nil {
		return nil, err
	}

	forkRepoIds := make([]int64, len(forks))
	for i, fork := range forks {
		forkRepoIds[i] = fork.ForkRepoId
	}

	return forkRepoIds, nil
}
