package logic

import (
	"context"

	"github.com/google/go-github/v66/github"
	"github.com/wushiling50/aster/pkg/constants"
	"github.com/wushiling50/aster/pkg/errno"
	githubFunc "github.com/wushiling50/aster/pkg/github"
	"github.com/wushiling50/aster/rpc/fetcher/internal/pack"
	"github.com/wushiling50/aster/rpc/fetcher/internal/svc"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
)

type FetchCreatedRepoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFetchCreatedRepoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FetchCreatedRepoLogic {
	return &FetchCreatedRepoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FetchCreatedRepoLogic) FetchCreatedRepo(userId int64) (err error) {
	var (
		githubUser *github.User
		allRepos   []*github.Repository
	)

	githubUser, _, err = githubFunc.GetUserById(l.ctx, userId)
	if err != nil {
		logx.Error(err)
		return
	}

	if allRepos, err = githubFunc.GetAllReposByLogin(l.ctx, githubUser.GetLogin()); err != nil {
		logx.Error(err)
		return
	}

	for _, githubRepo := range allRepos {
		if githubRepo.GetFork() {
			logx.Info("Skipped Fork Repo !")
			continue
		}

		createdRepo := pack.BuildCreatedRepo(githubRepo, userId)

		var jsonStr string

		if jsonStr, err = jsonx.MarshalToString(createdRepo); err != nil {
			err = errno.InternalJSONError.WithError(err)
			logx.Error(err)
			continue
		}

		if err = l.svcCtx.KqCreateRepoPusher.Push(l.ctx, jsonStr); err != nil {
			err = errno.InternalKafkaError.WithError(err)
			logx.Error(err)
			continue
		}
	}

	completedCreatedRepo := pack.BuildCompletedCreatedRepo(constants.FetchCreatedRepoCompletedDataId, userId)

	var completedStr string
	if completedStr, err = jsonx.MarshalToString(completedCreatedRepo); err != nil {
		logx.Error(err)
		err = errno.InternalJSONError.WithError(err)
		return
	}

	if err = l.svcCtx.KqCreateRepoPusher.Push(l.ctx, completedStr); err != nil {
		logx.Error(err)
		err = errno.InternalKafkaError.WithError(err)
		return
	}

	return
}
