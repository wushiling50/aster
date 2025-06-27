package developer

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

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

type GetSummaryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSummaryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSummaryLogic {
	return &GetSummaryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSummaryLogic) GetSummary(req *types.GetSummaryReq) (resp *types.GetSummaryResp, err error) {
	resp = new(types.GetSummaryResp)

	reqId := req.TaskId

	id, err := github.GetIdByLogin(l.ctx, req.Login)
	if err != nil {
		logx.Errorf("applet.GetSummary: Failed To Get Id By Login %v", err.Error())
		err = errno.InternalLanguagesError.WithError(err)
		return
	}

	taskId := tasks.GetNewAPITaskKey(constants.APIGetRegion, id, reqId)
	taskInfo, err := l.svcCtx.AsynqInspector.GetTaskInfo(constants.APITaskQueue, taskId)
	if err != nil {
		logx.Errorf("applet.GetSummary: Failed To Get Task Info %v", err.Error())
		err = errno.InternalAsynqError.WithError(err)
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
		var summary = analysis.Summary{}

		err = json.Unmarshal(taskInfo.Result, &summary)
		if err != nil {
			logx.Errorf("applet.GetSummary: Failed To Unmarshal Task Result %v", err.Error())
			err = errno.InternalJSONError.WithError(err)
			return
		}

		resp.Summary = types.Summary{
			Id:        id,
			Summary:   summary.Summary,
			UpdatedAt: time.Unix(summary.DataUpdatedAt, 0).Format(time.RFC3339),
		}
	default:
		err = errno.InternalServiceError.WithMessage(fmt.Sprintf("Unexpected task state: %v", taskInfo.State.String()))
	}

	return
}
