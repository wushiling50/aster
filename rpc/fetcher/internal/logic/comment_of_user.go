package logic

import (
	"context"

	"github.com/google/go-github/v66/github"
	"github.com/wushiling50/aster/pkg/constants"
	"github.com/wushiling50/aster/pkg/errno"
	githubFunc "github.com/wushiling50/aster/pkg/github"
	"github.com/wushiling50/aster/pkg/utils"
	contribution "github.com/wushiling50/aster/rpc/contribution/contributionclient"
	"github.com/wushiling50/aster/rpc/fetcher/internal/pack"
	"github.com/wushiling50/aster/rpc/fetcher/internal/svc"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
)

type FetchCommentOfUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFetchCommentOfUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FetchCommentOfUserLogic {
	return &FetchCommentOfUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FetchCommentOfUserLogic) FetchCommentOfUser(userId int64, createAfter string, searchLimit int64) (err error) {
	var (
		githubUser *github.User
		allComment []*githubFunc.CommentWithRepoId
		allRepo    map[int64]*github.Repository
	)

	githubUser, _, err = githubFunc.GetUserById(l.ctx, userId)
	if err != nil {
		logx.Error(err)
		return
	}

	if allComment, allRepo, err = githubFunc.GetAllCommentByLogin(l.ctx, githubUser.GetLogin(), createAfter, searchLimit); err != nil {
		logx.Error(err)
		return
	}

	if err = l.rpcDelAllContributionInCategory(userId, constants.CategoryComment); err != nil {
		logx.Error(err)
		return
	}

	for _, comment := range allComment {
		comment := pack.BuildComment(comment, userId)

		var jsonStr string

		if jsonStr, err = jsonx.MarshalToString(comment); err != nil {
			err = errno.InternalJSONError.WithError(err)
			logx.Error(err)
			continue
		}

		if err = l.svcCtx.KqContributionPusher.Push(l.ctx, jsonStr); err != nil {
			err = errno.InternalKafkaError.WithError(err)
			logx.Error(err)
			continue
		}
	}

	for _, repo := range allRepo {
		if err = doFetchRepo(l.ctx, l.svcCtx, repo); err != nil {
			continue
		}
	}

	completedContribution := pack.BuildCompletedContribution(constants.FetchCommentOfUserCompletedDataId, userId)

	var completedStr string
	if completedStr, err = jsonx.MarshalToString(completedContribution); err != nil {
		logx.Error(err)
		err = errno.InternalJSONError.WithError(err)
		return
	}

	if err = l.svcCtx.KqContributionPusher.Push(l.ctx, completedStr); err != nil {
		logx.Error(err)
		err = errno.InternalKafkaError.WithError(err)
		return
	}

	return
}

func (l *FetchCommentOfUserLogic) rpcDelAllContributionInCategory(developerId int64, category string) (err error) {
	var resp *contribution.DelAllContributionInCategoryByDeveloperIdResp

	resp, err = l.svcCtx.ContributionRpcClient.DelAllContributionInCategoryByDeveloperId(l.ctx, &contribution.DelAllContributionInCategoryByDeveloperIdReq{
		Category:    category,
		DeveloperId: developerId,
	})

	if err != nil {
		logx.Errorf("DelAllContributionInCategory: RPC called failed: %v", err.Error())
		err = errno.InternalServiceError.WithError(err)
		return
	}

	if !utils.IsSuccess(resp.Base) {
		err = errno.BizError.WithMessage(resp.Base.Message)
		return
	}

	return
}
