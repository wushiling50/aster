package logic

import (
	"context"

	"github.com/wushiling50/aster/gen/relation"
	model_relation "github.com/wushiling50/aster/pkg/model/relation"
	"github.com/wushiling50/aster/rpc/relation/internal/pack"
	"github.com/wushiling50/aster/rpc/relation/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddForkLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddForkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddForkLogic {
	return &AddForkLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------Fork-----------------------
func (l *AddForkLogic) AddFork(in *relation.AddForkReq) (*relation.AddForkResp, error) {
	resp := new(relation.AddForkResp)

	err := l.addFork(&model_relation.Fork{
		OriginalRepoId: in.OriginalRepoId,
		ForkRepoId:     in.ForkRepoId,
	})

	if err != nil {
		logx.Errorf("service.AddFork: Add Fork Failed: %w", err)
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.Base = pack.BuildSuccessResp()

	return resp, nil

}

func (l *AddForkLogic) addFork(model *model_relation.Fork) error {
	dataId, err := l.svcCtx.ForkModel.CreateDataId()
	if err != nil {
		return err
	}

	model.DataId = dataId
	_, err = l.svcCtx.ForkModel.Insert(l.ctx, model)
	if err != nil {
		return err
	}

	return nil
}
