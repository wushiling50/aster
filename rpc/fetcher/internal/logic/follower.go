package logic

import (
	"context"

	"github.com/google/go-github/v66/github"
	"github.com/wushiling50/aster/pkg/constants"
	"github.com/wushiling50/aster/pkg/errno"
	githubFunc "github.com/wushiling50/aster/pkg/github"
	"github.com/wushiling50/aster/pkg/utils"
	"github.com/wushiling50/aster/rpc/fetcher/internal/pack"
	"github.com/wushiling50/aster/rpc/fetcher/internal/svc"
	relation "github.com/wushiling50/aster/rpc/relation/relationclient"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
)

type FetchFollowerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFetchFollowerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FetchFollowerLogic {
	return &FetchFollowerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FetchFollowerLogic) FetchFollower(userId int64) (err error) {
	var (
		githubUser   *github.User
		allFollowers []*github.User
	)

	githubUser, _, err = githubFunc.GetUserById(l.ctx, userId)
	if err != nil {
		logx.Error(err)
		return
	}

	if allFollowers, err = githubFunc.GetAllFollowersByLogin(l.ctx, githubUser.GetLogin()); err != nil {
		logx.Error(err)
		return
	}

	if err = l.rpcDelAllFollower(userId); err != nil {
		logx.Error(err)
		return
	}

	for _, follower := range allFollowers {
		modelFollow := pack.BuildFollow(follower.GetID(), userId)

		var jsonStr string

		if jsonStr, err = jsonx.MarshalToString(modelFollow); err != nil {
			err = errno.InternalJSONError.WithError(err)
			logx.Error(err)
			continue
		}

		if err = l.svcCtx.KqFollowPusher.Push(l.ctx, jsonStr); err != nil {
			err = errno.InternalKafkaError.WithError(err)
			logx.Error(err)
			continue
		}

		if err = doFetchDeveloper(l.ctx, l.svcCtx, follower); err != nil {
			err = errno.BizError.WithError(err)
			logx.Error(err)
			continue
		}
	}

	completedFollow := pack.BuildCompletedFollow(constants.FetchFollowerCompletedDataId, userId)

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

func (l *FetchFollowerLogic) rpcDelAllFollower(developerId int64) (err error) {
	var resp *relation.DelAllFollowerResp

	resp, err = l.svcCtx.RelationRpcClient.DelAllFollower(l.ctx, &relation.DelAllFollowerReq{
		DeveloperId: developerId,
	})

	if err != nil {
		logx.Errorf("DelAllFollower: RPC called failed: %v", err.Error())
		err = errno.InternalServiceError.WithError(err)
		return
	}

	if !utils.IsSuccess(resp.Base) {
		err = errno.BizError.WithMessage(resp.Base.Message)
		return
	}

	return
}
