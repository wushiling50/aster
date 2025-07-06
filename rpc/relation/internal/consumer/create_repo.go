package consumer

import (
	"context"
	"errors"

	"github.com/wushiling50/aster/gen/relation"
	"github.com/wushiling50/aster/pkg/constants"
	"github.com/wushiling50/aster/pkg/errno"
	model_relation "github.com/wushiling50/aster/pkg/model/relation"
	"github.com/wushiling50/aster/pkg/utils"
	"github.com/wushiling50/aster/rpc/relation/internal/logic"
	"github.com/wushiling50/aster/rpc/relation/internal/svc"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateRepoConsumer struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateRepoConsumer(ctx context.Context, svcCtx *svc.ServiceContext) *CreateRepoConsumer {
	return &CreateRepoConsumer{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (c *CreateRepoConsumer) Consume(ctx context.Context, key string, value string) (err error) {
	logx.Info("Consume Message: ", value)

	var newCreateRepo *relation.CreateRepo

	if err = jsonx.UnmarshalFromString(value, &newCreateRepo); err != nil {
		err = errno.InternalJSONError.WithError(err)
		logx.Error(err)
		return
	}

	if newCreateRepo.DataId == constants.FetchCreatedRepoCompletedDataId {
		err = c.updateCreateRepoUpdatedAt(newCreateRepo.DeveloperId)
		if err != nil {
			logx.Error(err)
			return
		}
	} else {
		err = c.updateCreateRepo(newCreateRepo)
		if err != nil {
			logx.Error(err)
			return
		}
	}

	return
}

func (c *CreateRepoConsumer) updateCreateRepoUpdatedAt(developerId int64) error {
	createRepoUpdatedAt, err := c.svcCtx.CreatedRepoUpdatedAtModel.FindOneByDeveloperId(c.ctx, developerId)
	if err != nil {
		switch {
		case errors.Is(err, model_relation.ErrNotFound):
			logx.Info("No Found This Data!")

			dataId, err := c.svcCtx.CreatedRepoUpdatedAtModel.CreateDataId()
			if err != nil {
				err = errno.InternalServiceError.WithError(err)
				return err
			}

			createRepoUpdatedAt.DataId = dataId
			createRepoUpdatedAt.DeveloperId = developerId

			if _, err = c.svcCtx.CreatedRepoUpdatedAtModel.Insert(c.ctx, createRepoUpdatedAt); err != nil {
				err = errno.InternalServiceError.WithError(err)
				return err
			}
		default:
			err = errno.InternalServiceError.WithError(err)
			return err
		}
	}

	err = c.svcCtx.CreatedRepoUpdatedAtModel.Update(c.ctx, createRepoUpdatedAt)
	if err != nil {
		err = errno.InternalServiceError.WithError(err)
		return err
	}

	return nil
}

func (c *CreateRepoConsumer) updateCreateRepo(createRepo *relation.CreateRepo) error {
	l := logic.NewUpdateCreateRepoLogic(c.ctx, c.svcCtx)

	exist, err := l.CheckIfCreateRepoExist(createRepo.RepoId)
	if err != nil {
		err = errno.InternalServiceError.WithError(err)
		return err
	}

	if !exist {
		return c.addNewCreateRepo(createRepo)
	}

	resp, err := l.UpdateCreateRepo(&relation.UpdateCreateRepoReq{
		CreateRepo: createRepo,
	})

	if err != nil {
		err = errno.InternalServiceError.WithError(err)
		return err
	}

	if !utils.IsSuccess(resp.Base) {
		err = errno.BizError.WithMessage(resp.Base.Message)
		return err
	}

	return nil
}

func (c *CreateRepoConsumer) addNewCreateRepo(createRepo *relation.CreateRepo) error {
	l := logic.NewAddCreateRepoLogic(c.ctx, c.svcCtx)

	resp, err := l.AddCreateRepo(&relation.AddCreateRepoReq{
		DeveloperId: createRepo.DeveloperId,
		RepoId:      createRepo.RepoId,
	})

	if err != nil {
		err = errno.InternalServiceError.WithError(err)
		return err
	}

	if !utils.IsSuccess(resp.Base) {
		err = errno.BizError.WithMessage(resp.Base.Message)
		return err
	}

	return nil
}
