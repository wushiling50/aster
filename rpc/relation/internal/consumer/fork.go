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

type ForkConsumer struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewForkConsumer(ctx context.Context, svcCtx *svc.ServiceContext) *ForkConsumer {
	return &ForkConsumer{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (c *ForkConsumer) Consume(ctx context.Context, key string, value string) (err error) {
	logx.Info("Consume Message: ", value)

	var newFork *relation.Fork

	if err = jsonx.UnmarshalFromString(value, &newFork); err != nil {
		err = errno.InternalJSONError.WithError(err)
		logx.Error(err)
		return
	}

	if newFork.DataId == constants.FetchForkCompletedDataId {
		err = c.updateForkUpdatedAt(newFork.OriginalRepoId)
		if err != nil {
			logx.Error(err)
			return
		}
	} else {
		err = c.addNewFork(newFork)
		if err != nil {
			logx.Error(err)
			return
		}
	}

	return
}

func (c *ForkConsumer) updateForkUpdatedAt(repoId int64) error {
	forkUpdatedAt, err := c.svcCtx.ForkUpdatedAtModel.FindOneByRepoId(c.ctx, repoId)
	if err != nil {
		switch {
		case errors.Is(err, model_relation.ErrNotFound):
			logx.Info("No Found This Data!")

			dataId, err := c.svcCtx.ForkUpdatedAtModel.CreateDataId()
			if err != nil {
				err = errno.InternalServiceError.WithError(err)
				return err
			}

			forkUpdatedAt.DataId = dataId
			forkUpdatedAt.RepoId = repoId

			if _, err = c.svcCtx.ForkUpdatedAtModel.Insert(c.ctx, forkUpdatedAt); err != nil {
				err = errno.InternalServiceError.WithError(err)
				return err
			}

			return nil
		default:
			err = errno.InternalServiceError.WithError(err)
			return err
		}
	}

	err = c.svcCtx.ForkUpdatedAtModel.Update(c.ctx, forkUpdatedAt)
	if err != nil {
		err = errno.InternalServiceError.WithError(err)
		return err
	}

	return nil
}

func (c *ForkConsumer) addNewFork(fork *relation.Fork) error {
	l := logic.NewAddForkLogic(c.ctx, c.svcCtx)

	resp, err := l.AddFork(&relation.AddForkReq{
		OriginalRepoId: fork.OriginalRepoId,
		ForkRepoId:     fork.ForkRepoId,
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
