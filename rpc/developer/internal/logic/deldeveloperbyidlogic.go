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

type DelDeveloperByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelDeveloperByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelDeveloperByIdLogic {
	return &DelDeveloperByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelDeveloperByIdLogic) DelDeveloperById(in *developer.DelDeveloperByIdReq) (*developer.DelDeveloperByIdResp, error) {
	resp := new(developer.DelDeveloperByIdResp)

	err := l.delDeveloperById(in.Id)
	if err != nil {
		logx.Errorf("service.DelDeveloperById: Del Developer By Id Failed: %w", err)

		if errors.Is(err, model_developer.ErrNotFound) {
			err = errno.BizRepoNotFoundError.WithError(err)
		}

		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.Base = pack.BuildSuccessResp()

	return resp, nil
}

func (l *DelDeveloperByIdLogic) delDeveloperById(repoId int64) error {
	repo, err := l.svcCtx.DeveloperModel.FindOneById(l.ctx, repoId)
	if err != nil {
		return err
	}

	err = l.svcCtx.DeveloperModel.Delete(l.ctx, repo.DataId)
	if err != nil {
		return err
	}

	return nil
}
