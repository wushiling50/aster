package logic

import (
	"context"

	"github.com/wushiling50/aster/gen/relation"
	model_relation "github.com/wushiling50/aster/pkg/model/relation"
	"github.com/wushiling50/aster/rpc/relation/internal/pack"
	"github.com/wushiling50/aster/rpc/relation/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type AddFollowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddFollowLogic {
	return &AddFollowLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------Follow-----------------------
func (l *AddFollowLogic) AddFollow(in *relation.AddFollowReq) (*relation.AddFollowResp, error) {
	resp := new(relation.AddFollowResp)

	err := l.addFollow(&model_relation.Follow{
		FollowerId:  in.FollowerId,
		FollowingId: in.FollowingId,
	})

	if err != nil {
		logx.Errorf("service.AddFollow: Add Follow Failed: %w", err)
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.Base = pack.BuildSuccessResp()

	return resp, nil

}

func (l *AddFollowLogic) addFollow(model *model_relation.Follow) error {
	dataId, err := l.svcCtx.FollowModel.CreateDataId()
	if err != nil {
		return err
	}

	model.DataId = dataId
	_, err = l.svcCtx.FollowModel.Insert(l.ctx, model)
	if err != nil {
		return err
	}

	return nil
}
