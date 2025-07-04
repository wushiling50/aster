package logic

import (
	"context"

	"github.com/wushiling50/aster/gen/repo"
	model_repo "github.com/wushiling50/aster/pkg/model/repo"
	"github.com/wushiling50/aster/rpc/repo/internal/pack"
	"github.com/wushiling50/aster/rpc/repo/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddRepoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddRepoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddRepoLogic {
	return &AddRepoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddRepoLogic) AddRepo(in *repo.AddRepoReq) (*repo.AddRepoResp, error) {
	resp := new(repo.AddRepoResp)

	err := l.addRepo(pack.BuildModelRepo(in.Repo))
	if err != nil {
		logx.Errorf("service.AddRepo: Add Repo Failed: %w", err)
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.Base = pack.BuildSuccessResp()

	return resp, nil
}

func (l *AddRepoLogic) addRepo(model *model_repo.Repo) error {
	dataId, err := l.svcCtx.RepoModel.CreateDataId()
	if err != nil {
		return err
	}

	model.DataId = dataId
	_, err = l.svcCtx.RepoModel.Insert(l.ctx, model)
	if err != nil {
		return err
	}

	return nil
}
