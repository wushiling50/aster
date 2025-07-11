package developer

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	githublangsgo "github.com/NDoolan360/github-langs-go"
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

type GetLanguageUsageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetLanguageUsageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLanguageUsageLogic {
	return &GetLanguageUsageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetLanguageUsageLogic) GetLanguageUsage(req *types.GetLanguageUsageReq) (resp *types.GetLanguageUsageResp, err error) {
	resp = new(types.GetLanguageUsageResp)

	reqId := req.TaskId

	developerId, err := github.GetIdByLogin(l.ctx, req.Login)
	if err != nil {
		logx.Errorf("applet.GetLanguageUsage: Failed To Get Id By Login :%v", err.Error())
		return
	}

	taskId := tasks.GetNewAPITaskKey(constants.APIGetLanguage, developerId, reqId)
	taskInfo, err := l.svcCtx.AsynqInspector.GetTaskInfo(constants.APITaskQueue, taskId)
	if err != nil {
		logx.Errorf("applet.GetLanguageUsage: Failed To Get Task Info %v", err.Error())
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
		var (
			languages = analysis.Languages{}
			usageMap  = make(map[string]float64)
			usageArr  []types.LanguageWithPercentage
		)

		err = json.Unmarshal(taskInfo.Result, &languages)
		if err != nil {
			logx.Errorf("applet.GetLanguageUsage: Failed To Unmarshal Task Result %v", err.Error())
			err = errno.InternalJSONError.WithError(err)
			return
		}

		err = json.Unmarshal([]byte(languages.Languages), &usageMap)
		if err != nil {
			logx.Errorf("applet.GetLanguageUsage: Failed To Unmarshal Languages %v", err.Error())
			err = errno.InternalJSONError.WithError(err)
			return
		}

		for name, percentage := range usageMap {
			color, err := l.getLanguageColor(name)
			if err != nil {
				err = errno.InternalServiceError.WithError(err)
				return nil, err
			}

			usageArr = append(usageArr, types.LanguageWithPercentage{
				Language: types.Language{
					Id:    strings.ReplaceAll(strings.ToLower(name), " ", "-"),
					Name:  name,
					Color: color,
				},
				Percentage: percentage,
			})
		}

		resp.LanguageUsage = &types.LanguageUsage{
			Id:        developerId,
			Languages: usageArr,
			UpdatedAt: time.Unix(languages.DataUpdatedAt, 0).Format(time.RFC3339),
		}
	default:
		err = errno.InternalServiceError.WithMessage(fmt.Sprintf("Unexpected Task State: %v", taskInfo.State.String()))
	}

	logx.Info("Successfully Get LanguageUsage")
	return
}

func (l *GetLanguageUsageLogic) getLanguageColor(name string) (color string, err error) {
	var language githublangsgo.Language

	if language, err = githublangsgo.GetLanguage(name); err != nil {
		err = errno.InternalLanguagesError.WithError(err)
		return
	}

	color = language.Color
	return
}
