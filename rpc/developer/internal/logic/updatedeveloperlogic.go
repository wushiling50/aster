package logic

import (
	"context"

	"github.com/wushiling50/aster/gen/developer"
	model_developer "github.com/wushiling50/aster/pkg/model/developer"
	"github.com/wushiling50/aster/rpc/developer/internal/pack"
	"github.com/wushiling50/aster/rpc/developer/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateDeveloperLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateDeveloperLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateDeveloperLogic {
	return &UpdateDeveloperLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateDeveloperLogic) UpdateDeveloper(in *developer.UpdateDeveloperReq) (*developer.UpdateDeveloperResp, error) {
	resp := new(developer.UpdateDeveloperResp)

	err := l.updateDeveloper(pack.BuildModelDeveloper(in.Developer))
	if err != nil {
		logx.Errorf("service.UpdateDeveloper: Update Developer Failed: %w", err)
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.Base = pack.BuildSuccessResp()

	return resp, nil
}

func (l *UpdateDeveloperLogic) updateDeveloper(model *model_developer.Developer) error {
	developer, err := l.svcCtx.DeveloperModel.FindOneById(l.ctx, model.Id)
	if err != nil {
		return err
	}

	model.DataId = developer.DataId
	err = l.svcCtx.DeveloperModel.Update(l.ctx, model)
	if err != nil {
		return err
	}

	return nil
}
