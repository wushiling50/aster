package developer

import (
	"context"

	"github.com/wushiling50/aster/api/internal/svc"
	"github.com/wushiling50/aster/api/internal/types"
	"github.com/wushiling50/aster/pkg/constants"
	"github.com/wushiling50/aster/pkg/errno"
	"github.com/wushiling50/aster/pkg/github"
	"github.com/wushiling50/aster/pkg/tasks"

	"github.com/zeromicro/go-zero/core/logx"
)

type PostDeveloperLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPostDeveloperLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostDeveloperLogic {
	return &PostDeveloperLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PostDeveloperLogic) PostDeveloper(req *types.PostTaskReq) (err error) {
	developerId, err := github.GetIdByLogin(l.ctx, req.Login)
	if err != nil {
		logx.Errorf("applet.PostLanguageUsageTask: Failed To Get Id By Login %v", err.Error())
		err = errno.InternalLanguagesError.WithError(err)
		return
	}

	err = tasks.FetcherTaskPusher(l.svcCtx.AsynqClient, constants.FetchDeveloper, developerId, "", 0)

	if err != nil {
		logx.Errorf("applet.PostDeveloper: Failed To Enqueue Task: %v", err.Error())
		err = errno.InternalAsynqError.WithError(err)
		return
	}

	return
}
