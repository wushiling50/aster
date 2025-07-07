package logic

import (
	"context"

	"github.com/wushiling50/aster/gen/relation"
	"github.com/wushiling50/aster/rpc/relation/internal/pack"
	"github.com/wushiling50/aster/rpc/relation/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchStarredRepoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchStarredRepoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchStarredRepoLogic {
	return &SearchStarredRepoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchStarredRepoLogic) SearchStarredRepo(in *relation.SearchStarredRepoReq) (*relation.SearchStarredRepoResp, error) {
	resp := new(relation.SearchStarredRepoResp)

	starredRepoIds, err := l.searchStarredRepo(in.DeveloperId, in.Page, in.Limit)
	if err != nil {
		logx.Errorf("service.SearchStarredRepo: Search Starred Repo Failed: %w", err)
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	if len(starredRepoIds) == 0 {
		logx.Info("service.SearchStarredRepo: No Found RepoId")
	}

	resp.Base = pack.BuildSuccessResp()
	resp.RepoIds = starredRepoIds

	return resp, nil
}

func (l *SearchStarredRepoLogic) searchStarredRepo(developerId, page, limit int64) ([]int64, error) {
	stars, err := l.svcCtx.StarModel.SearchStarredRepo(l.ctx, developerId, page, limit)
	if err != nil {
		return nil, err
	}

	starredRepoIds := make([]int64, len(stars))
	for i, star := range stars {
		starredRepoIds[i] = star.DeveloperId
	}

	return starredRepoIds, nil
}
