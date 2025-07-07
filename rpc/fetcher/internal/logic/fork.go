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
)

type FetchForkLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFetchForkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FetchForkLogic {
	return &FetchForkLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FetchForkLogic) FetchFork(repoId int64) (err error) {
	var (
		originalRepo *github.Repository
		allForks     []*github.Repository
	)

	originalRepo, _, err = githubFunc.GetRepo(l.ctx, repoId)
	if err != nil {
		logx.Error(err)
		return
	}

	if allForks, err = githubFunc.GetAllForksByRepo(l.ctx, originalRepo.GetOwner().GetLogin(), originalRepo.GetName()); err != nil {
		logx.Error(err)
		return
	}

	if err = l.rpcDelAllFork(repoId); err != nil {
		logx.Error(err)
		return
	}

	for _, forkRepo := range allForks {
		fork := pack.BuildFork(repoId, forkRepo.GetID())

		var jsonStr string

		if jsonStr, err = jsonx.MarshalToString(fork); err != nil {
			err = errno.InternalJSONError.WithError(err)
			logx.Error(err)
			continue
		}

		if err = l.svcCtx.KqForkPusher.Push(l.ctx, jsonStr); err != nil {
			err = errno.InternalKafkaError.WithError(err)
			logx.Error(err)
			continue
		}

		if err = doFetchRepo(l.ctx, l.svcCtx, forkRepo); err != nil {
			err = errno.BizError.WithError(err)
			logx.Error(err)
			continue
		}
	}

	completedFork := pack.BuildCompletedFork(constants.FetchForkCompletedDataId, repoId)

	var completedStr string
	if completedStr, err = jsonx.MarshalToString(completedFork); err != nil {
		logx.Error(err)
		err = errno.InternalJSONError.WithError(err)
		return
	}

	if err = l.svcCtx.KqForkPusher.Push(l.ctx, completedStr); err != nil {
		logx.Error(err)
		err = errno.InternalKafkaError.WithError(err)
		return
	}

	return
}

func (l *FetchForkLogic) rpcDelAllFork(repoId int64) (err error) {
	var resp *relation.DelAllForkResp

	resp, err = l.svcCtx.RelationRpcClient.DelAllFork(l.ctx, &relation.DelAllForkReq{
		OriginalRepoId: repoId,
	})

	if err != nil {
		logx.Errorf("DelAllFork: RPC called failed: %v", err.Error())
		err = errno.InternalServiceError.WithError(err)
		return
	}

	if !utils.IsSuccess(resp.Base) {
		err = errno.BizError.WithMessage(resp.Base.Message)
		return
	}

	return
}
