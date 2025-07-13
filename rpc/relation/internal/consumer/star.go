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

type StarConsumer struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStarConsumer(ctx context.Context, svcCtx *svc.ServiceContext) *StarConsumer {
	return &StarConsumer{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (c *StarConsumer) Consume(ctx context.Context, key string, value string) (err error) {
	logx.Info("Consume Message: ", value)

	var newStar *relation.Star

	if err = jsonx.UnmarshalFromString(value, &newStar); err != nil {
		err = errno.InternalJSONError.WithError(err)
		logx.Error(err)
		return
	}

	if newStar.DataId == constants.FetchStarredRepoCompletedDataId {
		err = c.updateStarUpdatedAt(newStar.DeveloperId)
		if err != nil {
			logx.Error(err)
			return
		}

		locksKey := c.svcCtx.Locks.GetNewLocksKey(constants.LockStarredRepo, newStar.DeveloperId)
		err = c.svcCtx.Locks.Unblock(c.ctx, locksKey)
		if err != nil {
			logx.Error(err)
			return
		}
	} else {
		err = c.addNewStar(newStar)
		if err != nil {
			logx.Error(err)
			return
		}
	}

	return
}

func (c *StarConsumer) updateStarUpdatedAt(developerId int64) error {
	starredRepoUpdatedAt, err := c.svcCtx.StarredRepoUpdatedAtModel.FindOneByDeveloperId(c.ctx, developerId)
	if err != nil {
		switch {
		case errors.Is(err, model_relation.ErrNotFound):
			logx.Info("No Found This Data!")

			dataId, err := c.svcCtx.StarredRepoUpdatedAtModel.CreateDataId()
			if err != nil {
				err = errno.InternalServiceError.WithError(err)
				return err
			}

			starredRepoUpdatedAt.DataId = dataId
			starredRepoUpdatedAt.DeveloperId = developerId

			if _, err = c.svcCtx.StarredRepoUpdatedAtModel.Insert(c.ctx, starredRepoUpdatedAt); err != nil {
				err = errno.InternalServiceError.WithError(err)
				return err
			}

			return nil
		default:
			err = errno.InternalServiceError.WithError(err)
			return err
		}
	}

	err = c.svcCtx.StarredRepoUpdatedAtModel.Update(c.ctx, starredRepoUpdatedAt)
	if err != nil {
		err = errno.InternalServiceError.WithError(err)
		return err
	}

	return nil
}

func (c *StarConsumer) addNewStar(star *relation.Star) error {
	l := logic.NewAddStarLogic(c.ctx, c.svcCtx)

	resp, err := l.AddStar(&relation.AddStarReq{
		DeveloperId: star.DeveloperId,
		RepoId:      star.RepoId,
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
