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

type GetScoreLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetScoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetScoreLogic {
	return &GetScoreLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetScoreLogic) GetScore(req *types.GetScoreReq) (resp *types.GetScoreResp, err error) {
	resp = new(types.GetScoreResp)

	reqId := req.TaskId

	developerId, err := github.GetIdByLogin(l.ctx, req.Login)
	if err != nil {
		logx.Errorf("applet.GetScore: Failed To Get Id By Login %v", err.Error())
		err = errno.InternalLanguagesError.WithError(err)
		return
	}

	taskId := tasks.GetNewAPITaskKey(constants.APIGetScore, developerId, reqId)
	taskInfo, err := l.svcCtx.AsynqInspector.GetTaskInfo(constants.APITaskQueue, taskId)
	if err != nil {
		logx.Errorf("applet.GetScore: Failed To Get Task Info %v", err.Error())
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
		var score = analysis.Score{}

		err = json.Unmarshal(taskInfo.Result, &score)
		if err != nil {
			logx.Errorf("applet.GetScore: Failed To Unmarshal Task Result %v", err.Error())
			err = errno.InternalJSONError.WithError(err)
			return
		}

		resp.Score = types.Score{
			Id:        developerId,
			Score:     score.Score,
			UpdatedAt: time.Unix(score.DataUpdatedAt, 0).Format(time.RFC3339),
		}
	default:
		err = errno.InternalServiceError.WithMessage(fmt.Sprintf("Unexpected task state: %v", taskInfo.State.String()))
	}

	logx.Info("Successfully Get Score")
	return
}
