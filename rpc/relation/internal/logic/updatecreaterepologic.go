package logic

import (
	"context"
	"errors"

	"github.com/wushiling50/aster/gen/relation"
	"github.com/wushiling50/aster/pkg/github"
	model_relation "github.com/wushiling50/aster/pkg/model/relation"
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

	need, err := l.checkIfNeedUpdate(in.CreateRepo.DeveloperId)
	if err != nil {
		logx.Errorf("service.UpdateCreateRepo: Update Create Repo Failed: %w", err)
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	if need {
		err = l.updateRepo(pack.BuildModelCreateRepo(in.CreateRepo))
		if err != nil {
			logx.Errorf("service.UpdateCreateRepo: Update Create Repo Failed: %w", err)
			resp.Base = pack.BuildBaseResp(err)
			return resp, nil
		}
	}

	resp.Base = pack.BuildSuccessResp()

	return resp, nil
}

func (l *UpdateCreateRepoLogic) checkIfNeedUpdate(developerId int64) (bool, error) {
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

func (l *UpdateCreateRepoLogic) updateRepo(model *model_relation.CreateRepo) error {
	createRepo, err := l.svcCtx.CreateRepoModel.FindOneByRepoId(l.ctx, model.RepoId)
	if err != nil {
		return err
	}

	model.DataId = createRepo.DataId
	err = l.svcCtx.CreateRepoModel.Update(l.ctx, model)
	if err != nil {
		return err
	}

	return nil
}
