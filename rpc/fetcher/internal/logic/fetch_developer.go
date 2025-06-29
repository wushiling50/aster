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

type FetchDeveloperLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFetchDeveloperLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FetchDeveloperLogic {
	return &FetchDeveloperLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FetchDeveloperLogic) FetchDeveloper(userId int64) (err error) {
	var githubUser *github.User
	githubUser, _, err = githubFunc.GetUserById(l.ctx, userId)
	if err != nil {
		logx.Error(err)
		return
	}
	return doFetchDeveloper(l.ctx, l.svcCtx, githubUser)
}

func doFetchDeveloper(ctx context.Context, svcCtx *svc.ServiceContext, githubUser *github.User) (err error) {
	var starredRepoCount int64

	if starredRepoCount, err = githubFunc.GetStarredRepoCountByLogin(ctx, githubUser.GetLogin()); err != nil {
		logx.Error(err)
		return
	}

	modelDeveloper := pack.BuildDeveloperProfile(githubUser, starredRepoCount)

	var jsonStr string

	if jsonStr, err = jsonx.MarshalToString(modelDeveloper); err != nil {
		logx.Error(err)
		err = errno.InternalJSONError.WithError(err)
		return
	}

	if err = svcCtx.KqDeveloperPusher.Push(ctx, jsonStr); err != nil {
		logx.Error(err)
		err = errno.InternalKafkaError.WithError(err)
		return
	}

	logx.Info("Successfully Push Developer Profile")
	return
}
