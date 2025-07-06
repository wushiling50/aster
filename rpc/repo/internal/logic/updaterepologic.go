package logic

import (
	"context"
	"errors"

	"github.com/wushiling50/aster/gen/repo"
	"github.com/wushiling50/aster/pkg/github"
	model_repo "github.com/wushiling50/aster/pkg/model/repo"
	"github.com/wushiling50/aster/rpc/repo/internal/pack"
	"github.com/wushiling50/aster/rpc/repo/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateRepoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateRepoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRepoLogic {
	return &UpdateRepoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateRepoLogic) UpdateRepo(in *repo.UpdateRepoReq) (*repo.UpdateRepoResp, error) {
	resp := new(repo.UpdateRepoResp)

	need, err := l.checkIfNeedUpdate(in.Repo.Id)
	if err != nil {
		logx.Errorf("service.UpdateRepo: Update Repo Failed: %w", err)
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	if need {
		err = l.updateRepo(pack.BuildModelRepo(in.Repo))
		if err != nil {
			logx.Errorf("service.UpdateRepo: Update Repo Failed: %w", err)
			resp.Base = pack.BuildBaseResp(err)
			return resp, nil
		}
	}

	resp.Base = pack.BuildSuccessResp()

	return resp, nil
}

func (l *UpdateRepoLogic) checkIfNeedUpdate(id int64) (bool, error) {
	repo, err := l.svcCtx.RepoModel.FindOneById(l.ctx, id)
	if err != nil {
		switch {
		case errors.Is(err, model_repo.ErrNotFound):
			return true, nil
		default:
			return false, err
		}
	}

	if github.CheckIfDataExpired(repo.DataUpdatedAt) {
		return true, nil
	} else {
		return false, nil
	}
}

func (l *UpdateRepoLogic) updateRepo(model *model_repo.Repo) error {
	repo, err := l.svcCtx.RepoModel.FindOneById(l.ctx, model.Id)
	if err != nil {
		return err
	}

	model.DataId = repo.DataId
	err = l.svcCtx.RepoModel.Update(l.ctx, model)
	if err != nil {
		return err
	}

	return nil
}
