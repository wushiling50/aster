package logic

import (
	"context"
	"errors"

	"github.com/hibiken/asynq"
	"github.com/wushiling50/aster/gen/relation"
	"github.com/wushiling50/aster/pkg/constants"
	"github.com/wushiling50/aster/pkg/github"
	model_relation "github.com/wushiling50/aster/pkg/model/relation"
	"github.com/wushiling50/aster/pkg/tasks"
	"github.com/wushiling50/aster/rpc/relation/internal/pack"
	"github.com/wushiling50/aster/rpc/relation/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCreateRepoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCreateRepoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCreateRepoLogic {
	return &UpdateCreateRepoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateCreateRepoLogic) UpdateCreateRepo(in *relation.UpdateCreateRepoReq) (*relation.UpdateCreateRepoResp, error) {
	resp := new(relation.UpdateCreateRepoResp)

	needUpdate, err := l.checkIfNeedUpdateCreateRepo(in.DeveloperId)
	if err != nil {
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	if needUpdate {
		err = l.pushCreateRepoTask(in.DeveloperId)
		if err != nil {
			resp.Base = pack.BuildBaseResp(err)
			return resp, err
		}
	}

	resp.Base = pack.BuildSuccessResp()

	return resp, nil
}

func (l *UpdateCreateRepoLogic) checkIfNeedUpdateCreateRepo(developerId int64) (bool, error) {
	createRepoUpdatedAt, err := l.svcCtx.CreatedRepoUpdatedAtModel.FindOneByDeveloperId(l.ctx, developerId)
	if err != nil {
		switch {
		case errors.Is(err, model_relation.ErrNotFound):
			return true, nil
		default:
			return false, err
		}
	}

	if github.CheckIfDataExpired(createRepoUpdatedAt.DataUpdatedAt) {
		return true, nil
	} else {
		return false, nil
	}

}

func (l *UpdateCreateRepoLogic) pushCreateRepoTask(id int64) (err error) {
	var (
		task   *asynq.Task
		taskId string
	)

	if task, taskId, err = tasks.NewFetcherTask(constants.FetchCreatedRepo, id, "", 0); err != nil {
		return
	}

	_, err = l.svcCtx.AsynqClient.Enqueue(
		task,
		asynq.TaskID(taskId),
		asynq.Queue(constants.FetcherTaskQueue),
		asynq.MaxRetry(constants.FetchMaxRetry))
	if err != nil {
		return
	}

	return
}
