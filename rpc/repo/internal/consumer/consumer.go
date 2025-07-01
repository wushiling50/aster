package consumer

import (
	"context"

	"github.com/wushiling50/aster/gen/repo"
	"github.com/wushiling50/aster/pkg/errno"
	model_repo "github.com/wushiling50/aster/pkg/model/repo"
	"github.com/wushiling50/aster/rpc/repo/internal/config"
	"github.com/wushiling50/aster/rpc/repo/internal/logic"
	"github.com/wushiling50/aster/rpc/repo/internal/pack"
	"github.com/wushiling50/aster/rpc/repo/internal/svc"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
)

type RepoConsumer struct {
	ctx context.Context
	svc *svc.ServiceContext
}

func NewRepoConsumer(ctx context.Context, svc *svc.ServiceContext) *RepoConsumer {
	return &RepoConsumer{
		ctx: ctx,
		svc: svc,
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
		newRepo *model_repo.Repo
		oldRepo *model_repo.Repo
		exist   bool
	)

	if err = jsonx.UnmarshalFromString(value, &newRepo); err != nil {
		err = errno.InternalJSONError.WithError(err)
		logx.Error(err)
		return
	}

	if oldRepo, exist, err = c.getRepo(newRepo.Id); err != nil {
		logx.Error(err)
		return
	}

	if exist {
		err = c.updateOldRepo(oldRepo, newRepo)
		if err != nil {
			logx.Error(err)
			return
		}
	}

	err = c.insertNewRepo(newRepo)
	if err != nil {
		logx.Error(err)
		return
	}

	// if err = unblockRepoUpdateLock(c, newRepo.Id); err != nil {
	// 	return
	// }

	return
}

func (c *RepoConsumer) getRepo(repoId int64) (*model_repo.Repo, bool, error) {
	l := logic.NewGetRepoByIdLogic(c.ctx, c.svc)

	resp, err := l.GetRepoById(&repo.GetRepoByIdReq{
		Id: repoId,
	})

	if err != nil {
		logx.Error(err)
		return nil, false, err
	}

	if resp.Repo == nil {
		logx.Info("No Found This Repo!")
		return nil, false, nil
	}

	modelRepo := pack.BuildModelRepo(resp.Repo)

	return modelRepo, true, nil
}

func (c *RepoConsumer) updateOldRepo(oldRepo *model_repo.Repo, newRepo *model_repo.Repo) error {
	return nil
}

func (c *RepoConsumer) insertNewRepo(newRepo *model_repo.Repo) error {
	return nil
}
