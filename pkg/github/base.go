package github

import (
	"context"
	"strings"
	"time"

	"github.com/google/go-github/v66/github"
	"github.com/wushiling50/aster/pkg/constants"
	"github.com/wushiling50/aster/pkg/errno"
	"github.com/zeromicro/go-zero/core/logx"
)

func githubClientInit() *github.Client {
	return github.NewClient(nil).WithAuthToken("*********")
}

func GetIdByLogin(ctx context.Context, login string) (id int64, err error) {
	var githubClient *github.Client = githubClientInit()
	var githubUser *github.User

	if githubUser, _, err = githubClient.Users.Get(ctx, login); err != nil {
		logx.Errorf("github.GetIdByLogin: Fail To Fetching user: %v", err.Error())
		err = errno.InternalGithubError.WithError(err)
		return
	}

	id = githubUser.GetID()

	logx.Infof("Successfully Get Id %v Of Login %s", id, login)
	return
}

func GetLoginById(ctx context.Context, id int64) (login string, err error) {
	var githubClient *github.Client = githubClientInit()
	var githubUser *github.User

	if githubUser, _, err = githubClient.Users.GetByID(ctx, id); err != nil {
		logx.Errorf("github.GetIdByLogin: Fail To Fetching User: %v", err.Error())
		err = errno.InternalGithubError.WithError(err)
		return
	}

	login = githubUser.GetLogin()

	logx.Infof("Successfully Get Login %s Of Id %v", login, id)
	return
}

func GetUserById(ctx context.Context, id int64) (githubUser *github.User, githubResp *github.Response, err error) {
	var githubClient *github.Client = githubClientInit()
	if githubUser, githubResp, err = githubClient.Users.GetByID(ctx, id); err != nil {
		logx.Errorf("github.GetUserById: Fail To Fetching User: %v", err.Error())
		err = errno.InternalGithubError.WithError(err)
		return
	}

	logx.Infof("Successfully Get User")
	return
}

func GetRepo(ctx context.Context, repoId int64) (githubRepo *github.Repository, githubResp *github.Response, err error) {
	var githubClient *github.Client = githubClientInit()
	if githubRepo, githubResp, err = githubClient.Repositories.GetByID(ctx, repoId); err != nil {
		logx.Errorf("github.GetRepo: Fail To Fetching Repo: %v", err.Error())
		err = errno.InternalGithubError.WithError(err)
		return
	}

	logx.Infof("Successfully Get Repo")
	return
}

func GetRepoByUrl(ctx context.Context, repoUrl string) (repo *github.Repository, err error) {
	var (
		githubClient *github.Client = githubClientInit()
		split        []string
		owner        string
		repoName     string
	)
	split = strings.Split(repoUrl, "/")
	owner = split[len(split)-2]
	repoName = split[len(split)-1]

	if repo, _, err = githubClient.Repositories.Get(ctx, owner, repoName); err != nil {
		logx.Errorf("github.GetRepoByUrl: Fail To Fetching Repo: %v", err.Error())
		err = errno.InternalGithubError.WithError(err)
		return
	}

	return
}

func DefaultUpdateAfterTime() string {
	return time.Unix(time.Now().Unix()-int64(constants.ONE_WEEK.Seconds()), 0).Format("2006-01-02")
}

func CheckIfDataExpired(lastUpdate time.Time) bool {
	return time.Since(lastUpdate) > constants.DataExpiredTime
}
