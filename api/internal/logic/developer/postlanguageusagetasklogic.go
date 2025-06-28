package developer

import (
	"context"

	"github.com/hibiken/asynq"
	"github.com/wushiling50/aster/api/internal/svc"
	"github.com/wushiling50/aster/api/internal/types"
	"github.com/wushiling50/aster/pkg/constants"
	"github.com/wushiling50/aster/pkg/errno"
	"github.com/wushiling50/aster/pkg/github"
	"github.com/wushiling50/aster/pkg/tasks"
	"github.com/wushiling50/aster/pkg/utils"
	"github.com/wushiling50/aster/rpc/id_generator/idgenerator"

	"github.com/zeromicro/go-zero/core/logx"
)

type PostLanguageUsageTaskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPostLanguageUsageTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostLanguageUsageTaskLogic {
	return &PostLanguageUsageTaskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PostLanguageUsageTaskLogic) PostLanguageUsageTask(req *types.PostTaskReq) (resp *types.PostTaskResp, err error) {
	resp = new(types.PostTaskResp)

	developerId, err := github.GetIdByLogin(l.ctx, req.Login)
	if err != nil {
		logx.Errorf("applet.PostLanguageUsageTask: Failed To Get Id By Login %v", err.Error())
		err = errno.InternalLanguagesError.WithError(err)
		return
	}

	reqId, err := l.rpcGetId()
	if err != nil {
		logx.Error(err)
		return
	}

	task, taskId, err := tasks.NewAPITask(constants.APIGetLanguage, developerId, reqId)
	if err != nil {
		logx.Errorf("applet.PostLanguageUsageTask: Failed To Create Task: %v", err.Error())
		err = errno.InternalAsynqError.WithError(err)
		return
	}

	_, err = l.svcCtx.AsynqClient.Enqueue(
		task,
		asynq.TaskID(taskId),
		asynq.Retention(constants.APITaskExpireTime),
		asynq.MaxRetry(constants.APIMaxRetry),
		asynq.Queue(constants.APITaskQueue),
	)
	if err != nil {
		logx.Errorf("applet.PostLanguageUsageTask: Failed To Enqueue Task: %v", err.Error())
		err = errno.InternalAsynqError.WithError(err)
		return
	}

	resp.TaskId = types.TaskId{
		TaskId: reqId,
	}

	return
}

func (l *PostLanguageUsageTaskLogic) rpcGetId() (id string, err error) {
	var resp *idgenerator.GetIdResp

	resp, err = l.svcCtx.IdGeneratorRpcClient.GetId(l.ctx, &idgenerator.GetIdReq{})
	if err != nil {
		logx.Errorf("IdGeneratorRpc: RPC called failed: %v", err.Error())
		err = errno.InternalServiceError.WithError(err)
		return
	}

	if !utils.IsSuccess(resp.Base) {
		err = errno.BizError.WithMessage(resp.Base.Message)
		return
	}

	id = resp.Id

	return
}
