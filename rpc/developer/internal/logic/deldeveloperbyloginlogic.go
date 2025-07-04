package logic

import (
	"context"
	"errors"

	"github.com/wushiling50/aster/gen/developer"
	"github.com/wushiling50/aster/pkg/errno"
	model_developer "github.com/wushiling50/aster/pkg/model/developer"
	"github.com/wushiling50/aster/rpc/developer/internal/pack"
	"github.com/wushiling50/aster/rpc/developer/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelDeveloperByLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelDeveloperByLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelDeveloperByLoginLogic {
	return &DelDeveloperByLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelDeveloperByLoginLogic) DelDeveloperByLogin(in *developer.DelDeveloperByLoginReq) (*developer.DelDeveloperByLoginResp, error) {
	resp := new(developer.DelDeveloperByLoginResp)

	err := l.delDeveloperByLogin(in.Login)
	if err != nil {
		logx.Errorf("service.DelDeveloperByLogin: Del Developer By Login Failed: %w", err)

		if errors.Is(err, model_developer.ErrNotFound) {
			err = errno.BizRepoNotFoundError.WithError(err)
		}

		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.Base = pack.BuildSuccessResp()

	return resp, nil
}

func (l *DelDeveloperByLoginLogic) delDeveloperByLogin(login string) error {
	repo, err := l.svcCtx.DeveloperModel.FindOneByLogin(l.ctx, login)
	if err != nil {
		return err
	}

	err = l.svcCtx.DeveloperModel.Delete(l.ctx, repo.DataId)
	if err != nil {
		return err
	}

	return nil
}
