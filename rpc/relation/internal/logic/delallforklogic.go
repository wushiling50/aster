package logic

import (
	"context"
	"math"

	"github.com/wushiling50/aster/gen/relation"
	"github.com/wushiling50/aster/rpc/relation/internal/pack"
	"github.com/wushiling50/aster/rpc/relation/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelAllForkLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelAllForkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelAllForkLogic {
	return &DelAllForkLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelAllForkLogic) DelAllFork(in *relation.DelAllForkReq) (*relation.DelAllForkResp, error) {
	resp := new(relation.DelAllForkResp)

	err := l.delAllFork(in.OriginalRepoId, 1, math.MaxInt64)
	if err != nil {
		logx.Errorf("service.DelAllFork: Del All Fork Failed: %w", err)
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.Base = pack.BuildSuccessResp()

	return resp, nil
}

func (l *DelAllForkLogic) delAllFork(originalRepoId, page, limit int64) error {
	forks, err := l.svcCtx.ForkModel.SearchFork(l.ctx, originalRepoId, page, limit)
	if err != nil {
		return err
	}

	for _, fork := range forks {
		err := l.svcCtx.ForkModel.Delete(l.ctx, fork.DataId)
		if err != nil {
			return err
		}
	}

	return nil
}
