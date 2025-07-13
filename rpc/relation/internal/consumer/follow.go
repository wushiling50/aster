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

type FollowConsumer struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFollowConsumer(ctx context.Context, svcCtx *svc.ServiceContext) *FollowConsumer {
	return &FollowConsumer{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (c *FollowConsumer) Consume(ctx context.Context, key string, value string) (err error) {
	logx.Info("Consume Message: ", value)

	var newFollow *relation.Follow

	if err = jsonx.UnmarshalFromString(value, &newFollow); err != nil {
		err = errno.InternalJSONError.WithError(err)
		logx.Error(err)
		return
	}

	switch newFollow.DataId {
	case constants.FetchFollowingCompletedDataId:
		err = c.updateFollowingUpdatedAt(newFollow.FollowerId)
		if err != nil {
			return
		}

		locksKey := c.svcCtx.Locks.GetNewLocksKey(constants.LockFollowing, newFollow.FollowerId)
		err = c.svcCtx.Locks.Unblock(c.ctx, locksKey)
		if err != nil {
			logx.Error(err)
			return
		}
	case constants.FetchFollowerCompletedDataId:
		err = c.updateFollowerUpdatedAt(newFollow.FollowingId)
		if err != nil {
			return
		}

		locksKey := c.svcCtx.Locks.GetNewLocksKey(constants.LockFollower, newFollow.FollowingId)
		err = c.svcCtx.Locks.Unblock(c.ctx, locksKey)
		if err != nil {
			logx.Error(err)
			return
		}
	default:
		err = c.addNewFollow(newFollow)
		if err != nil {
			logx.Error(err)
			return
		}
	}

	return
}

func (c *FollowConsumer) updateFollowerUpdatedAt(developerId int64) error {
	followerUpdatedAt, err := c.svcCtx.FollowerUpdatedAtModel.FindOneByDeveloperId(c.ctx, developerId)
	if err != nil {
		switch {
		case errors.Is(err, model_relation.ErrNotFound):
			logx.Info("No Found This Data!")

			dataId, err := c.svcCtx.FollowerUpdatedAtModel.CreateDataId()
			if err != nil {
				err = errno.InternalServiceError.WithError(err)
				return err
			}

			followerUpdatedAt.DataId = dataId
			followerUpdatedAt.DeveloperId = developerId

			if _, err = c.svcCtx.FollowerUpdatedAtModel.Insert(c.ctx, followerUpdatedAt); err != nil {
				err = errno.InternalServiceError.WithError(err)
				return err
			}
			return nil
		default:
			err = errno.InternalServiceError.WithError(err)
			return err
		}
	}

	err = c.svcCtx.FollowerUpdatedAtModel.Update(c.ctx, followerUpdatedAt)
	if err != nil {
		err = errno.InternalServiceError.WithError(err)
		return err
	}

	return nil
}

func (c *FollowConsumer) updateFollowingUpdatedAt(developerId int64) error {
	followingUpdatedAt, err := c.svcCtx.FollowingUpdatedAtModel.FindOneByDeveloperId(c.ctx, developerId)
	if err != nil {
		switch {
		case errors.Is(err, model_relation.ErrNotFound):
			logx.Info("No Found This Data!")

			dataId, err := c.svcCtx.FollowerUpdatedAtModel.CreateDataId()
			if err != nil {
				err = errno.InternalServiceError.WithError(err)
				return err
			}

			followingUpdatedAt.DataId = dataId
			followingUpdatedAt.DeveloperId = developerId

			if _, err = c.svcCtx.FollowingUpdatedAtModel.Insert(c.ctx, followingUpdatedAt); err != nil {
				err = errno.InternalServiceError.WithError(err)
				return err
			}

			return nil
		default:
			err = errno.InternalServiceError.WithError(err)
			return err
		}
	}

	err = c.svcCtx.FollowingUpdatedAtModel.Update(c.ctx, followingUpdatedAt)
	if err != nil {
		err = errno.InternalServiceError.WithError(err)
		return err
	}

	return nil
}

func (c *FollowConsumer) addNewFollow(newFollow *relation.Follow) error {
	l := logic.NewAddFollowLogic(c.ctx, c.svcCtx)

	resp, err := l.AddFollow(&relation.AddFollowReq{
		FollowerId:  newFollow.FollowerId,
		FollowingId: newFollow.FollowingId,
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
