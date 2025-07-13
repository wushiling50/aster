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

type FetchIssuePROfUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFetchIssuePROfUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FetchIssuePROfUserLogic {
	return &FetchIssuePROfUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FetchIssuePROfUserLogic) FetchIssuePROfUser(userId int64, createAfter string, serachLimit int64) (err error) {
	var (
		githubUser *github.User

		allIssuePR []*github.Issue
		repos      map[int64]*github.Repository = make(map[int64]*github.Repository)
	)

	githubUser, _, err = githubFunc.GetUserById(l.ctx, userId)
	if err != nil {
		logx.Error(err)
		return
	}

	if allIssuePR, err = githubFunc.GetAllIssuePRByLogin(l.ctx, githubUser.GetLogin(), constants.RoleAuthor, createAfter, serachLimit); err != nil {
		logx.Error(err)
		return
	}

	if err = l.rpcDelAllContributionInCategory(userId, constants.CategoryIssue); err != nil {
		logx.Error(err)
		return
	}

	if err = l.rpcDelAllContributionInCategory(userId, constants.CategoryOpenPullRequest); err != nil {
		logx.Error(err)
		return
	}

	if err = l.rpcDelAllContributionInCategory(userId, constants.CategoryMergePullRequest); err != nil {
		logx.Error(err)
		return
	}

	for _, issuePR := range allIssuePR {
		var (
			merged   bool
			repo     *github.Repository
			category string
		)

		if repo, err = githubFunc.GetRepoByUrl(l.ctx, issuePR.GetRepositoryURL()); err != nil {
			logx.Error(err)
			continue
		}

		if merged, err = githubFunc.CheckIfMerged(l.ctx, issuePR, repo); err != nil {
			continue
		}

		if issuePR.IsPullRequest() {
			if merged {
				category = constants.CategoryMergePullRequest
			} else {
				category = constants.CategoryOpenPullRequest
			}
		} else {
			category = constants.CategoryIssue
		}

		genIssuePR := pack.BuildIssuePR(issuePR, userId, category, merged, repo)
		var jsonStr string

		if jsonStr, err = jsonx.MarshalToString(genIssuePR); err != nil {
			err = errno.InternalJSONError.WithError(err)
			logx.Error(err)
			continue
		}

		if err = l.svcCtx.KqContributionPusher.Push(l.ctx, jsonStr); err != nil {
			err = errno.InternalKafkaError.WithError(err)
			logx.Error(err)
			continue
		}

		repos[repo.GetID()] = repo
	}

	for _, repo := range repos {
		if err = doFetchRepo(l.ctx, l.svcCtx, repo); err != nil {
			continue
		}
	}

	completedContribution := pack.BuildCompletedContribution(constants.FetchIssuePROfUserCompletedDataId, userId)

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

func (l *FetchIssuePROfUserLogic) rpcDelAllContributionInCategory(developerId int64, category string) (err error) {
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
