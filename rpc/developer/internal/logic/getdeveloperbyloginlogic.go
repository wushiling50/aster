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

type GetDeveloperByLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetDeveloperByLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDeveloperByLoginLogic {
	return &GetDeveloperByLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetDeveloperByLoginLogic) GetDeveloperByLogin(in *developer.GetDeveloperByLoginReq) (*developer.GetDeveloperByLoginResp, error) {
	resp := new(developer.GetDeveloperByLoginResp)

	oldDeveloper, err := l.getDeveloperByLogin(in.Login)
	if err != nil {
		switch {
		case errors.Is(err, model_developer.ErrNotFound):
			resp.Base = pack.BuildSuccessResp()
			return resp, nil
		default:
			logx.Errorf("service.GetDeveloperByLogin: Get Developer By Login Failed: %w", err)
			resp.Base = pack.BuildBaseResp(err)
			return resp, nil
		}
	}

	resp.Base = pack.BuildSuccessResp()
	resp.Developer = pack.BuildGenDeveloper(oldDeveloper)
	return resp, nil
}

func (l *GetDeveloperByLoginLogic) getDeveloperByLogin(login string) (*model_developer.Developer, error) {
	developer, err := l.svcCtx.DeveloperModel.FindOneByLogin(l.ctx, login)
	if err != nil {
		return nil, err
	}

	return developer, nil
}
