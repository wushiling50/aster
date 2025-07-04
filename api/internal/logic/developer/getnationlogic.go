package developer

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hibiken/asynq"
	"github.com/wushiling50/aster/api/internal/svc"
	"github.com/wushiling50/aster/api/internal/types"
	"github.com/wushiling50/aster/pkg/constants"
	"github.com/wushiling50/aster/pkg/errno"
	"github.com/wushiling50/aster/pkg/github"
	"github.com/wushiling50/aster/pkg/tasks"
	analysis "github.com/wushiling50/aster/rpc/analysis/analysisclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetNationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetNationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetNationLogic {
	return &GetNationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetNationLogic) GetNation(req *types.GetNationReq) (resp *types.GetNationResp, err error) {
	resp = new(types.GetNationResp)

	reqId := req.TaskId

	developerId, err := github.GetIdByLogin(l.ctx, req.Login)
	if err != nil {
		logx.Errorf("applet.GetNation: Failed To Get Id By Login %v", err.Error())
		err = errno.InternalLanguagesError.WithError(err)
		return
	}

	taskId := tasks.GetNewAPITaskKey(constants.APIGetNation, developerId, reqId)
	taskInfo, err := l.svcCtx.AsynqInspector.GetTaskInfo(constants.APITaskQueue, taskId)
	if err != nil {
		logx.Errorf("applet.GetNation: Failed To Get Task Info %v", err.Error())
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
		var nation = analysis.Nation{}

		err = json.Unmarshal(taskInfo.Result, &nation)
		if err != nil {
			logx.Errorf("applet.GetNation: Failed To Unmarshal Task Result %v", err.Error())
			err = errno.InternalJSONError.WithError(err)
			return
		}

		resp.Nation = &types.Nation{
			Id:         developerId,
			Nation:     nation.Nation,
			Confidence: nation.Confidence,
		}
	default:
		err = errno.InternalServiceError.WithMessage(fmt.Sprintf("Unexpected Task State: %v", taskInfo.State.String()))
	}

	logx.Info("Successfully Get Nation")
	return
}
