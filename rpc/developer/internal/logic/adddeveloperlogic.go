package logic

import (
	"context"

	"github.com/wushiling50/aster/gen/developer"
	model_developer "github.com/wushiling50/aster/pkg/model/developer"
	"github.com/wushiling50/aster/rpc/developer/internal/pack"
	"github.com/wushiling50/aster/rpc/developer/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddDeveloperLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddDeveloperLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddDeveloperLogic {
	return &AddDeveloperLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddDeveloperLogic) AddDeveloper(in *developer.AddDeveloperReq) (*developer.AddDeveloperResp, error) {
	resp := new(developer.AddDeveloperResp)

	err := l.addDeveloper(pack.BuildModelDeveloper(in.Developer))
	if err != nil {
		logx.Errorf("service.AddDeveloper: Add Developer Failed: %w", err)
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.Base = pack.BuildSuccessResp()

	return resp, nil
}

func (l *AddDeveloperLogic) addDeveloper(model *model_developer.Developer) error {
	dataId, err := l.svcCtx.DeveloperModel.CreateDataId()
	if err != nil {
		return err
	}

	model.DataId = dataId
	_, err = l.svcCtx.DeveloperModel.Insert(l.ctx, model)
	if err != nil {
		return err
	}

	return nil
}
