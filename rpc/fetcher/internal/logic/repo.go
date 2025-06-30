package logic

import (
	"context"

	"github.com/google/go-github/v66/github"
	"github.com/wushiling50/aster/pkg/errno"
	githubFunc "github.com/wushiling50/aster/pkg/github"
	"github.com/wushiling50/aster/rpc/fetcher/internal/pack"
	"github.com/wushiling50/aster/rpc/fetcher/internal/svc"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
)

type FetchRepoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFetchRepoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FetchRepoLogic {
	return &FetchRepoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FetchRepoLogic) FetchRepo(repoId int64) (err error) {
	var githubRepo *github.Repository
	if githubRepo, _, err = githubFunc.GetRepo(l.ctx, repoId); err != nil {
		logx.Error(err)
		return
	}
	return doFetchRepo(l.ctx, l.svcCtx, githubRepo)
}

func doFetchRepo(ctx context.Context, svcCtx *svc.ServiceContext, githubRepo *github.Repository) (err error) {
	var (
		issueCount    int64
		prCount       int64
		commitCount   int64
		openPrCount   int64
		mergedPrCount int64
		commentCount  int64
		reviewCount   int64
		languages     string

		jsonStr string
	)

	repoOwner := githubRepo.GetOwner().GetLogin()
	repoName := githubRepo.GetName()
	issueCount, prCount, err = githubFunc.GetIssuePrCountByRepo(ctx, repoOwner, repoName)
	if err != nil {
		logx.Error(err)
		return
	}

	commitCount, err = githubFunc.GetCommitCountByRepo(ctx, repoOwner, repoName)
	if err != nil {
		logx.Error(err)
		return
	}

	openPrCount, err = githubFunc.GetOpenPrCountByRepo(ctx, repoOwner, repoName)
	if err != nil {
		logx.Error(err)
		return
	}

	mergedPrCount, err = githubFunc.GetMergedPrCountByRepo(ctx, repoOwner, repoName)
	if err != nil {
		logx.Error(err)
		return
	}

	commentCount, err = githubFunc.GetCommentCountByRepo(ctx, repoOwner, repoName)
	if err != nil {
		logx.Error(err)
		return
	}

	reviewCount, err = githubFunc.GetReviewCountByRepo(ctx, repoOwner, repoName)
	if err != nil {
		logx.Error(err)
		return
	}

	languages, err = githubFunc.GetLanguagesByRepo(ctx, repoOwner, repoName)
	if err != nil {
		logx.Error(err)
		return
	}

	modelRepo := pack.BuildRepo(githubRepo, issueCount, prCount, commitCount, openPrCount, mergedPrCount,
		commentCount, reviewCount, languages)

	if jsonStr, err = jsonx.MarshalToString(modelRepo); err != nil {
		logx.Error(err)
		err = errno.InternalJSONError.WithError(err)
		return
	}

	if err = svcCtx.KqRepoPusher.Push(ctx, jsonStr); err != nil {
		logx.Error(err)
		err = errno.InternalKafkaError.WithError(err)
		return
	}

	logx.Info("Successfully Push Repo")
	return
}
