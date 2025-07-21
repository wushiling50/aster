package logic

import (
	"context"

	"github.com/google/go-github/v66/github"
	"github.com/wushiling50/aster/gen/relation"
	"github.com/wushiling50/aster/pkg/constants"
	"github.com/wushiling50/aster/pkg/errno"
	githubFunc "github.com/wushiling50/aster/pkg/github"
	"github.com/wushiling50/aster/pkg/utils"
	"github.com/wushiling50/aster/rpc/fetcher/internal/pack"
	"github.com/wushiling50/aster/rpc/fetcher/internal/svc"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/sync/errgroup"
)

type FetchStarredRepoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFetchStarredRepoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FetchStarredRepoLogic {
	return &FetchStarredRepoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FetchStarredRepoLogic) FetchStarredRepo(userId int64) (err error) {
	var (
		githubUser *github.User
		allRepos   []*github.StarredRepository

		eg errgroup.Group
	)

	githubUser, _, err = githubFunc.GetUserById(l.ctx, userId)
	if err != nil {
		logx.Error(err)
		return
	}

	if allRepos, err = githubFunc.GetAllStarredReposByLogin(l.ctx, githubUser.GetLogin()); err != nil {
		logx.Error(err)
		return
	}

	if err = l.rpcDelAllStarredRepo(userId); err != nil {
		logx.Error(err)
		return
	}

	for _, githubRepo := range allRepos {
		starredRepo := pack.BuildStarredRepo(githubRepo, userId)

		var jsonStr string

		if jsonStr, err = jsonx.MarshalToString(starredRepo); err != nil {
			err = errno.InternalJSONError.WithError(err)
			logx.Error(err)
			continue
		}

		if err = l.svcCtx.KqStarPusher.Push(l.ctx, jsonStr); err != nil {
			err = errno.InternalKafkaError.WithError(err)
			logx.Error(err)
			continue
		}

		eg.Go(func() error {
			return doFetchRepo(l.ctx, l.svcCtx, githubRepo.Repository)
		})
	}

	if err := eg.Wait(); err != nil {
		err = errno.BizError.WithError(err)
		logx.Error(err)
		return err
	}

	completedStarredRepo := pack.BuildCompletedStarredRepo(constants.FetchStarredRepoCompletedDataId, userId)

	var completedStr string
	if completedStr, err = jsonx.MarshalToString(completedStarredRepo); err != nil {
		err = errno.InternalJSONError.WithError(err)
		logx.Error(err)
		return
	}

	if err = l.svcCtx.KqStarPusher.Push(l.ctx, completedStr); err != nil {
		err = errno.InternalKafkaError.WithError(err)
		logx.Error(err)
		return
	}

	return
}

func (l *FetchStarredRepoLogic) rpcDelAllStarredRepo(developerId int64) (err error) {
	var resp *relation.DelAllStarredRepoResp

	resp, err = l.svcCtx.RelationRpcClient.DelAllStarredRepo(l.ctx, &relation.DelAllStarredRepoReq{
		DeveloperId: developerId,
	})

	if err != nil {
		logx.Errorf("DelDelAllStarredRepo: RPC called failed: %v", err.Error())
		err = errno.InternalServiceError.WithError(err)
		return
	}

	if !utils.IsSuccess(resp.Base) {
		err = errno.BizError.WithMessage(resp.Base.Message)
		return
	}

	return
}
