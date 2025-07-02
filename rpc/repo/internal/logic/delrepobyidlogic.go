package logic

import (
	"context"
	"errors"

	"github.com/wushiling50/aster/gen/repo"
	"github.com/wushiling50/aster/pkg/errno"
	model_repo "github.com/wushiling50/aster/pkg/model/repo"
	"github.com/wushiling50/aster/rpc/repo/internal/pack"
	"github.com/wushiling50/aster/rpc/repo/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelRepoByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelRepoByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelRepoByIdLogic {
	return &DelRepoByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelRepoByIdLogic) DelRepoById(in *repo.DelRepoByIdReq) (*repo.DelRepoByIdResp, error) {
	resp := new(repo.DelRepoByIdResp)

	err := l.delRepoById(in.Id)
	if err != nil {
		logx.Errorf("service.DelRepoById: Del Repo By Id Failed: %w", err)

		if errors.Is(err, model_repo.ErrNotFound) {
			err = errno.BizRepoNotFoundError.WithError(err)
		}

		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.Base = pack.BuildSuccessResp()

	return resp, nil
}

func (l *DelRepoByIdLogic) delRepoById(repoId int64) error {
	repo, err := l.svcCtx.RepoModel.FindOneById(l.ctx, repoId)
	if err != nil {
		return err
	}

	err = l.svcCtx.RepoModel.Delete(l.ctx, repo.DataId)
	if err != nil {
		return err
	}

	return nil
}
