package logic

import (
	"context"

	"github.com/wushiling50/aster/gen/relation"
	"github.com/wushiling50/aster/rpc/relation/internal/pack"
	"github.com/wushiling50/aster/rpc/relation/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchCreatedRepoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchCreatedRepoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchCreatedRepoLogic {
	return &SearchCreatedRepoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchCreatedRepoLogic) SearchCreatedRepo(in *relation.SearchCreatedRepoReq) (*relation.SearchCreatedRepoResp, error) {
	resp := new(relation.SearchCreatedRepoResp)

	repoIds, err := l.searchCreatedRepo(in.DeveloperId, in.Page, in.Limit)
	if err != nil {
		logx.Errorf("service.SearchCreatedRepo: Search Created Repo Failed: %w", err)
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	if len(repoIds) == 0 {
		logx.Info("service.SearchCreatedRepo: No Found Created Repo")
	}

	resp.RepoIds = repoIds
	resp.Base = pack.BuildSuccessResp()

	return resp, nil
}

func (l *SearchCreatedRepoLogic) searchCreatedRepo(developerId, page, limit int64) ([]int64, error) {
	createRepos, err := l.svcCtx.CreateRepoModel.SearchCreatedRepo(l.ctx, developerId, page, limit)
	if err != nil {
		return nil, err
	}

	repoIds := make([]int64, len(createRepos))
	for i, repo := range createRepos {
		repoIds[i] = repo.RepoId
	}

	return repoIds, nil
}
