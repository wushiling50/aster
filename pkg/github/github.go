package github

import (
	"context"
	"os"
	"time"

	"github.com/google/go-github/v66/github"
	"github.com/wushiling50/aster/pkg/constants"
	"github.com/wushiling50/aster/pkg/errno"
	"github.com/zeromicro/go-zero/core/logx"
)

func GetIdByLogin(ctx context.Context, login string) (id int64, err error) {
	var githubClient *github.Client = github.NewClient(nil).WithAuthToken(os.Getenv(constants.GithubAPIToken))
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
	var githubClient *github.Client = github.NewClient(nil).WithAuthToken(os.Getenv(constants.GithubAPIToken))
	var githubUser *github.User

	if githubUser, _, err = githubClient.Users.GetByID(ctx, id); err != nil {
		logx.Errorf("github.GetIdByLogin: Fail To Fetching user: %v", err.Error())
		err = errno.InternalGithubError.WithError(err)
		return
	}

	login = githubUser.GetLogin()

	logx.Infof("Successfully Get Login %s Of Id %v", login, id)
	return
}

func DefaultUpdateAfterTime() string {
	return time.Unix(time.Now().Unix()-int64(constants.ONE_WEEK.Seconds()), 0).Format("2006-01-02")
}

func CheckIfDataExpired(lastUpdate time.Time) bool {
	return time.Since(lastUpdate) > constants.DataExpiredTime
}
