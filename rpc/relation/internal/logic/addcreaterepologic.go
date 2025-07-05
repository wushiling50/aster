package logic

import (
	"context"

	"github.com/wushiling50/aster/gen/relation"
	model_relation "github.com/wushiling50/aster/pkg/model/relation"
	"github.com/wushiling50/aster/rpc/relation/internal/pack"
	"github.com/wushiling50/aster/rpc/relation/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type AddCreateRepoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddCreateRepoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCreateRepoLogic {
	return &AddCreateRepoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------CreateRepo-----------------------
func (l *AddCreateRepoLogic) AddCreateRepo(in *relation.AddCreateRepoReq) (*relation.AddCreateRepoResp, error) {
	resp := new(relation.AddCreateRepoResp)

	err := l.addCreateRepo(&model_relation.CreateRepo{
		DeveloperId: in.DeveloperId,
		RepoId:      in.RepoId,
	})

	if err != nil {
		logx.Errorf("service.AddCreateRepo: Add Repo Failed: %w", err)
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.Base = pack.BuildSuccessResp()

	return resp, nil

}

func (l *AddCreateRepoLogic) addCreateRepo(model *model_relation.CreateRepo) error {
	dataId, err := l.svcCtx.CreateRepoModel.CreateDataId()
	if err != nil {
		return err
	}

	model.DataId = dataId
	_, err = l.svcCtx.CreateRepoModel.Insert(l.ctx, model)
	if err != nil {
		return err
	}

	return nil
}
