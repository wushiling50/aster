package logic

import (
	"context"
	"errors"

	"github.com/wushiling50/aster/gen/repo"
	model_repo "github.com/wushiling50/aster/pkg/model/repo"
	"github.com/wushiling50/aster/rpc/repo/internal/pack"
	"github.com/wushiling50/aster/rpc/repo/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetRepoByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRepoByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRepoByIdLogic {
	return &GetRepoByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetRepoByIdLogic) GetRepoById(in *repo.GetRepoByIdReq) (*repo.GetRepoByIdResp, error) {
	resp := new(repo.GetRepoByIdResp)

	oldRepo, err := l.getRepoById(in.Id)
	if err != nil {
		switch {
		case errors.Is(err, model_repo.ErrNotFound):
			resp.Base = pack.BuildSuccessResp()
			return resp, nil
		default:
			logx.Errorf("service.GetRepoById: Get Repo By Id Failed: %w", err)
			resp.Base = pack.BuildBaseResp(err)
			return resp, nil
		}
	}

	resp.Base = pack.BuildSuccessResp()
	resp.Repo = pack.BuildGenRepo(oldRepo)
	return resp, nil
}

func (l *GetRepoByIdLogic) getRepoById(repoId int64) (*model_repo.Repo, error) {
	repo, err := l.svcCtx.RepoModel.FindOneById(l.ctx, repoId)
	if err != nil {
		return nil, err
	}

	return repo, nil
}
