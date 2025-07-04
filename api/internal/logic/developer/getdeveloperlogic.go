package developer

import (
	"context"
	"fmt"

	"github.com/hibiken/asynq"
	"github.com/wushiling50/aster/api/internal/pack"
	"github.com/wushiling50/aster/api/internal/svc"
	"github.com/wushiling50/aster/api/internal/types"
	"github.com/wushiling50/aster/pkg/constants"
	"github.com/wushiling50/aster/pkg/errno"
	"github.com/wushiling50/aster/pkg/github"
	"github.com/wushiling50/aster/pkg/tasks"
	"github.com/wushiling50/aster/pkg/utils"
	developer "github.com/wushiling50/aster/rpc/developer/developerclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDeveloperLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDeveloperLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDeveloperLogic {
	return &GetDeveloperLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDeveloperLogic) GetDeveloper(req *types.GetDeveloperReq) (resp *types.GetDeveloperResp, err error) {
	resp = new(types.GetDeveloperResp)

	developerId, err := github.GetIdByLogin(l.ctx, req.Login)
	if err != nil {
		logx.Errorf("applet.GetDeveloper: Failed To Get Id By Login %v", err.Error())
		err = errno.InternalLanguagesError.WithError(err)
		return
	}

	taskId := tasks.GetNewFetcherTaskKey(constants.FetchDeveloper, developerId)
	taskInfo, err := l.svcCtx.AsynqInspector.GetTaskInfo(constants.FetcherTaskQueue, taskId)
	if err != nil {
		logx.Errorf("applet.GetDeveloper: Failed To Get Task Info %v", err.Error())
		err = errno.InternalAsynqError.WithError(err)
		return
	}

	switch taskInfo.State {
	case asynq.TaskStatePending, asynq.TaskStateActive:
		resp.TaskState = types.TaskState{
			State: taskInfo.State.String(),
		}
	case asynq.TaskStateRetry:
		resp.TaskState = types.TaskState{
			State:  taskInfo.State.String(),
			Reason: taskInfo.LastErr,
		}
	case asynq.TaskStateArchived:
		resp.TaskState = types.TaskState{
			State:  "fail",
			Reason: taskInfo.LastErr,
		}
	case asynq.TaskStateCompleted:
		var rpcResp *types.Developer

		rpcResp, err = l.rpcGetDeveloperById(developerId)
		if err != nil {
			logx.Error(err)
			return
		}

		resp.Developer = rpcResp

	default:
		err = errno.InternalServiceError.WithMessage(fmt.Sprintf("Unexpected Task State: %v", taskInfo.State.String()))
	}

	logx.Info("Successfully Get Developer")
	return
}

func (l *GetDeveloperLogic) rpcGetDeveloperById(id int64) (typeDeveloper *types.Developer, err error) {
	var resp *developer.GetDeveloperByIdResp

	resp, err = l.svcCtx.DeveloperRpcClient.GetDeveloperById(l.ctx, &developer.GetDeveloperByIdReq{
		Id: id,
	})
	if err != nil {
		logx.Errorf("GetDeveloperByIdRPC: RPC Called Failed: %v", err.Error())
		err = errno.InternalServiceError.WithError(err)
		return
	}

	if !utils.IsSuccess(resp.Base) {
		err = errno.BizError.WithMessage(resp.Base.Message)
		return
	}

	typeDeveloper = pack.BuildTypeDeveloper(resp.Developer)

	return
}
