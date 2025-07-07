package logic

import (
	"context"

	"github.com/wushiling50/aster/gen/relation"
	model_relation "github.com/wushiling50/aster/pkg/model/relation"
	"github.com/wushiling50/aster/rpc/relation/internal/pack"
	"github.com/wushiling50/aster/rpc/relation/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddStarLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddStarLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddStarLogic {
	return &AddStarLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------Star-----------------------
func (l *AddStarLogic) AddStar(in *relation.AddStarReq) (*relation.AddStarResp, error) {
	resp := new(relation.AddStarResp)

	err := l.addStar(&model_relation.Star{
		DeveloperId: in.DeveloperId,
		RepoId:      in.RepoId,
	})

	if err != nil {
		logx.Errorf("service.AddStar: Add Star Failed: %w", err)
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.Base = pack.BuildSuccessResp()

	return resp, nil

}

func (l *AddStarLogic) addStar(model *model_relation.Star) error {
	dataId, err := l.svcCtx.StarModel.CreateDataId()
	if err != nil {
		return err
	}

	model.DataId = dataId
	_, err = l.svcCtx.StarModel.Insert(l.ctx, model)
	if err != nil {
		return err
	}

	return nil
}
