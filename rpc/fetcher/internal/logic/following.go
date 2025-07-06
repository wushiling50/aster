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

type FetchFollowingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFetchFollowingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FetchFollowingLogic {
	return &FetchFollowingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FetchFollowingLogic) FetchFollowing(userId int64) (err error) {
	var (
		githubUser   *github.User
		allFollowing []*github.User
	)

	githubUser, _, err = githubFunc.GetUserById(l.ctx, userId)
	if err != nil {
		logx.Error(err)
		return
	}

	if allFollowing, err = githubFunc.GetAllFollowingByLogin(l.ctx, githubUser.GetLogin()); err != nil {
		logx.Error(err)
		return
	}

	for _, following := range allFollowing {
		follow := pack.BuildFollow(userId, following.GetID())

		var jsonStr string

		if jsonStr, err = jsonx.MarshalToString(follow); err != nil {
			err = errno.InternalJSONError.WithError(err)
			logx.Error(err)
			continue
		}

		if err = l.svcCtx.KqFollowPusher.Push(l.ctx, jsonStr); err != nil {
			err = errno.InternalKafkaError.WithError(err)
			logx.Error(err)
			continue
		}

		if err = doFetchDeveloper(l.ctx, l.svcCtx, following); err != nil {
			err = errno.BizError.WithError(err)
			logx.Error(err)
			continue
		}
	}

	completedFollow := pack.BuildCompletedFollow(constants.FetchFollowingCompletedDataId, userId)

	var completedStr string
	if completedStr, err = jsonx.MarshalToString(completedFollow); err != nil {
		logx.Error(err)
		err = errno.InternalJSONError.WithError(err)
		return
	}

	if err = l.svcCtx.KqFollowPusher.Push(l.ctx, completedStr); err != nil {
		logx.Error(err)
		err = errno.InternalKafkaError.WithError(err)
		return
	}

	return
}
