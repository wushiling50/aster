package consumer

import (
	"context"
	"errors"

	"github.com/wushiling50/aster/gen/contribution"
	"github.com/wushiling50/aster/pkg/constants"
	"github.com/wushiling50/aster/pkg/errno"
	model_contribution "github.com/wushiling50/aster/pkg/model/contribution"
	"github.com/wushiling50/aster/pkg/utils"
	"github.com/wushiling50/aster/rpc/contribution/internal/config"
	"github.com/wushiling50/aster/rpc/contribution/internal/logic"
	"github.com/wushiling50/aster/rpc/contribution/internal/svc"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
)

type ContributionConsumer struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewContributionConsumer(ctx context.Context, svcCtx *svc.ServiceContext) *ContributionConsumer {
	return &ContributionConsumer{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func Consumers(c config.Config, ctx context.Context, svc *svc.ServiceContext) []service.Service {
	return []service.Service{
		kq.MustNewQueue(c.KqContributionConsumerConf, NewContributionConsumer(ctx, svc)),
	}
}

func (c *ContributionConsumer) Consume(ctx context.Context, key string, value string) (err error) {
	logx.Info("Consume Message: ", value)

	var newContribution *contribution.Contribution

	if err = jsonx.UnmarshalFromString(value, &newContribution); err != nil {
		err = errno.InternalJSONError.WithError(err)
		logx.Error(err)
		return
	}

	switch newContribution.DataId {
	case constants.FetchIssuePROfUserCompletedDataId:
		err = c.updateIssuePROfUserUpdatedAt(newContribution.DeveloperId)
		if err != nil {
			return
		}

		locksKey := c.svcCtx.Locks.GetNewLocksKey(constants.LockIssuePROfUser, newContribution.DeveloperId)
		err = c.svcCtx.Locks.Unblock(c.ctx, locksKey)
		if err != nil {
			logx.Error(err)
			return
		}
	case constants.FetchCommentOfUserCompletedDataId:
		err = c.updateCommentOfUserUpdatedAt(newContribution.DeveloperId)
		if err != nil {
			return
		}

		locksKey := c.svcCtx.Locks.GetNewLocksKey(constants.LockCommentOfUser, newContribution.DeveloperId)
		err = c.svcCtx.Locks.Unblock(c.ctx, locksKey)
		if err != nil {
			logx.Error(err)
			return
		}
	case constants.FetchReviewOfUserCompletedDataId:
		err = c.updateReviewOfUserUpdatedAt(newContribution.DeveloperId)
		if err != nil {
			return
		}

		locksKey := c.svcCtx.Locks.GetNewLocksKey(constants.LockReviewOfUser, newContribution.DeveloperId)
		err = c.svcCtx.Locks.Unblock(c.ctx, locksKey)
		if err != nil {
			logx.Error(err)
			return
		}
	default:
		err = c.addNewContribution(newContribution)
		if err != nil {
			logx.Error(err)
			return
		}
	}

	return
}

func (c *ContributionConsumer) updateIssuePROfUserUpdatedAt(developerId int64) error {
	issuePrOfUserUpdatedAt, err := c.svcCtx.IssuePrOfUserUpdatedAtModel.FindOneByDeveloperId(c.ctx, developerId)
	if err != nil {
		switch {
		case errors.Is(err, model_contribution.ErrNotFound):
			logx.Info("No Found This Data!")

			dataId, err := c.svcCtx.IssuePrOfUserUpdatedAtModel.CreateDataId()
			if err != nil {
				err = errno.InternalServiceError.WithError(err)
				return err
			}

			issuePrOfUserUpdatedAt := &model_contribution.IssuePrOfUserUpdatedAt{
				DataId:      dataId,
				DeveloperId: developerId,
			}

			if _, err = c.svcCtx.IssuePrOfUserUpdatedAtModel.Insert(c.ctx, issuePrOfUserUpdatedAt); err != nil {
				err = errno.InternalServiceError.WithError(err)
				return err
			}

			return nil
		default:
			err = errno.InternalServiceError.WithError(err)
			return err
		}
	}

	err = c.svcCtx.IssuePrOfUserUpdatedAtModel.Update(c.ctx, issuePrOfUserUpdatedAt)
	if err != nil {
		err = errno.InternalServiceError.WithError(err)
		return err
	}

	return nil
}

func (c *ContributionConsumer) updateCommentOfUserUpdatedAt(developerId int64) error {
	commentOfUserUpdatedAt, err := c.svcCtx.CommentOfUserUpdatedAtModel.FindOneByDeveloperId(c.ctx, developerId)
	if err != nil {
		switch {
		case errors.Is(err, model_contribution.ErrNotFound):
			logx.Info("No Found This Data!")

			dataId, err := c.svcCtx.CommentOfUserUpdatedAtModel.CreateDataId()
			if err != nil {
				err = errno.InternalServiceError.WithError(err)
				return err
			}

			commentOfUserUpdatedAt := &model_contribution.CommentOfUserUpdatedAt{
				DataId:      dataId,
				DeveloperId: developerId,
			}

			if _, err = c.svcCtx.CommentOfUserUpdatedAtModel.Insert(c.ctx, commentOfUserUpdatedAt); err != nil {
				err = errno.InternalServiceError.WithError(err)
				return err
			}

			return nil
		default:
			err = errno.InternalServiceError.WithError(err)
			return err
		}
	}

	err = c.svcCtx.CommentOfUserUpdatedAtModel.Update(c.ctx, commentOfUserUpdatedAt)
	if err != nil {
		err = errno.InternalServiceError.WithError(err)
		return err
	}

	return nil
}

func (c *ContributionConsumer) updateReviewOfUserUpdatedAt(developerId int64) error {
	reviewOfUserUpdatedAt, err := c.svcCtx.ReviewOfUserUpdatedAtModel.FindOneByDeveloperId(c.ctx, developerId)
	if err != nil {
		switch {
		case errors.Is(err, model_contribution.ErrNotFound):
			logx.Info("No Found This Data!")

			dataId, err := c.svcCtx.ReviewOfUserUpdatedAtModel.CreateDataId()
			if err != nil {
				err = errno.InternalServiceError.WithError(err)
				return err
			}

			reviewOfUserUpdatedAt := &model_contribution.ReviewOfUserUpdatedAt{
				DataId:      dataId,
				DeveloperId: developerId,
			}

			if _, err = c.svcCtx.ReviewOfUserUpdatedAtModel.Insert(c.ctx, reviewOfUserUpdatedAt); err != nil {
				err = errno.InternalServiceError.WithError(err)
				return err
			}

			return nil
		default:
			err = errno.InternalServiceError.WithError(err)
			return err
		}
	}

	err = c.svcCtx.ReviewOfUserUpdatedAtModel.Update(c.ctx, reviewOfUserUpdatedAt)
	if err != nil {
		err = errno.InternalServiceError.WithError(err)
		return err
	}

	return nil
}

func (c *ContributionConsumer) addNewContribution(newContribution *contribution.Contribution) error {
	l := logic.NewAddContributionLogic(c.ctx, c.svcCtx)

	resp, err := l.AddContribution(&contribution.AddContributionReq{
		DeveloperId:           newContribution.DeveloperId,
		RepoId:                newContribution.RepoId,
		Category:              newContribution.Category,
		Content:               newContribution.Content,
		ContributionCreatedAt: newContribution.ContributionCreatedAt,
		ContributionUpdatedAt: newContribution.ContributionUpdatedAt,
		ContributionId:        newContribution.ContributionId,
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
