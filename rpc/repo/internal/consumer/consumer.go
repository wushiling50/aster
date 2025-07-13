package consumer

import (
	"context"

	"github.com/wushiling50/aster/gen/repo"
	"github.com/wushiling50/aster/pkg/constants"
	"github.com/wushiling50/aster/pkg/errno"
	"github.com/wushiling50/aster/pkg/utils"
	"github.com/wushiling50/aster/rpc/repo/internal/config"
	"github.com/wushiling50/aster/rpc/repo/internal/logic"
	"github.com/wushiling50/aster/rpc/repo/internal/svc"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
)

type RepoConsumer struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRepoConsumer(ctx context.Context, svcCtx *svc.ServiceContext) *RepoConsumer {
	return &RepoConsumer{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func Consumers(c config.Config, ctx context.Context, svc *svc.ServiceContext) []service.Service {
	return []service.Service{
		kq.MustNewQueue(c.KqRepoConsumerConf, NewRepoConsumer(ctx, svc)),
	}
}

func (c *RepoConsumer) Consume(ctx context.Context, key string, value string) (err error) {
	logx.Info("Consume Message: ", value)

	var (
		newRepo *repo.Repo
		exist   bool
	)

	if err = jsonx.UnmarshalFromString(value, &newRepo); err != nil {
		err = errno.InternalJSONError.WithError(err)
		logx.Error(err)
		return
	}

	if _, exist, err = c.getRepo(newRepo.Id); err != nil {
		logx.Error(err)
		return
	}

	if exist {
		err = c.updateOldRepo(newRepo)
		if err != nil {
			logx.Error(err)
			return
		}
	} else {
		err = c.addNewRepo(newRepo)
		if err != nil {
			logx.Error(err)
			return
		}
	}

	locksKey := c.svcCtx.Locks.GetNewLocksKey(constants.LockRepo, newRepo.Id)
	err = c.svcCtx.Locks.Unblock(c.ctx, locksKey)
	if err != nil {
		logx.Error(err)
		return
	}

	return
}

func (c *RepoConsumer) getRepo(repoId int64) (*repo.Repo, bool, error) {
	l := logic.NewGetRepoByIdLogic(c.ctx, c.svcCtx)

	resp, err := l.GetRepoById(&repo.GetRepoByIdReq{
		Id: repoId,
	})

	if err != nil {
		err = errno.InternalServiceError.WithError(err)
		return nil, false, err
	}

	if !utils.IsSuccess(resp.Base) {
		err = errno.BizError.WithMessage(resp.Base.Message)
		return nil, false, err
	}

	if resp.Repo == nil {
		logx.Info("No Found This Repo!")
		return nil, false, nil
	}

	return resp.Repo, true, nil
}

func (c *RepoConsumer) updateOldRepo(newRepo *repo.Repo) error {
	l := logic.NewUpdateRepoLogic(c.ctx, c.svcCtx)

	resp, err := l.UpdateRepo(&repo.UpdateRepoReq{
		Repo: newRepo,
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

func (c *RepoConsumer) addNewRepo(newRepo *repo.Repo) error {
	l := logic.NewAddRepoLogic(c.ctx, c.svcCtx)

	resp, err := l.AddRepo(&repo.AddRepoReq{
		Repo: newRepo,
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
