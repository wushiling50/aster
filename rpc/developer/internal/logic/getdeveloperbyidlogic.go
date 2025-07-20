package logic

import (
	"context"
	"errors"

	"github.com/wushiling50/aster/gen/developer"
	model_developer "github.com/wushiling50/aster/pkg/model/developer"
	"github.com/wushiling50/aster/rpc/developer/internal/pack"
	"github.com/wushiling50/aster/rpc/developer/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDeveloperByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetDeveloperByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDeveloperByIdLogic {
	return &GetDeveloperByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetDeveloperByIdLogic) GetDeveloperById(in *developer.GetDeveloperByIdReq) (*developer.GetDeveloperByIdResp, error) {
	resp := new(developer.GetDeveloperByIdResp)

	oldDeveloper, err := l.getDeveloperById(in.Id)
	if err != nil {
		switch {
		case errors.Is(err, model_developer.ErrNotFound):
			resp.Base = pack.BuildSuccessResp()
			return resp, nil
		default:
			logx.Errorf("service.GetDeveloperById: Get Developer By Id Failed: %w", err)
			resp.Base = pack.BuildBaseResp(err)
			return resp, nil
		}
	}

	resp.Base = pack.BuildSuccessResp()
	resp.Developer = pack.BuildGenDeveloper(oldDeveloper)
	return resp, nil
}

func (l *GetDeveloperByIdLogic) getDeveloperById(developerId int64) (*model_developer.Developer, error) {
	developer, err := l.svcCtx.DeveloperModel.FindOneById(l.ctx, developerId)
	if err != nil {
		return nil, err
	}

	return developer, nil
}
