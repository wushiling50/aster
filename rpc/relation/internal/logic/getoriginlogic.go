package logic

import (
	"context"
	"errors"

	"github.com/wushiling50/aster/gen/relation"
	model_relation "github.com/wushiling50/aster/pkg/model/relation"
	"github.com/wushiling50/aster/rpc/relation/internal/pack"
	"github.com/wushiling50/aster/rpc/relation/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOriginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOriginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOriginLogic {
	return &GetOriginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOriginLogic) GetOrigin(in *relation.GetOriginReq) (*relation.GetOriginResp, error) {
	resp := new(relation.GetOriginResp)

	originalRepoId, err := l.getOrigin(in.ForkRepoId)
	if err != nil {
		switch {
		case errors.Is(err, model_relation.ErrNotFound):
			resp.Base = pack.BuildSuccessResp()
			return resp, nil
		default:
			logx.Errorf("service.GetOrigin: Get Origin Failed: %w", err)
			resp.Base = pack.BuildBaseResp(err)
			return resp, nil
		}
	}

	resp.Base = pack.BuildSuccessResp()
	resp.OriginalRepoId = originalRepoId
	return resp, nil
}

func (l *GetOriginLogic) getOrigin(forkRepoId int64) (int64, error) {
	fork, err := l.svcCtx.ForkModel.FindOneByForkRepoId(l.ctx, forkRepoId)
	if err != nil {
		return 0, err
	}

	return fork.OriginalRepoId, nil
}
