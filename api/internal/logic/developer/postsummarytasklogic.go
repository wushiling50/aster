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
	"github.com/wushiling50/aster/rpc/id_generator/idgenerator"

	"github.com/zeromicro/go-zero/core/logx"
)

type PostSummaryTaskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPostSummaryTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostSummaryTaskLogic {
	return &PostSummaryTaskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PostSummaryTaskLogic) PostSummaryTask(req *types.PostTaskReq) (resp *types.PostTaskResp, err error) {
	resp = new(types.PostTaskResp)

	developerId, err := github.GetIdByLogin(l.ctx, req.Login)
	if err != nil {
		logx.Errorf("applet.PostSummaryTask: Failed To Get Id By Login %v", err.Error())
		err = errno.InternalLanguagesError.WithError(err)
		return
	}

	getIdResp, err := l.svcCtx.IdGeneratorRpcClient.GetId(l.ctx, &idgenerator.GetIdReq{})
	if err != nil {
		logx.Errorf("IdGeneratorRpc: RPC called failed: %v", err.Error())
		err = errno.InternalServiceError.WithError(err)
		return nil, err
	}
	reqId := getIdResp.Id

	task, taskId, err := tasks.NewAPITask(constants.APIGetSummary, developerId, reqId)
	if err != nil {
		logx.Errorf("applet.PostSummaryTask: Failed To Create Task: %v", err.Error())
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
		logx.Errorf("applet.PostSummaryTask: Failed To Enqueue Task: %v", err.Error())
		err = errno.InternalAsynqError.WithError(err)
		return
	}

	resp.TaskId = types.TaskId{
		TaskId: reqId,
	}

	return
}
