package logic

import (
	"context"
	"errors"

	"github.com/wushiling50/aster/gen/analysis"
	"github.com/wushiling50/aster/pkg/errno"
	model_analysis "github.com/wushiling50/aster/pkg/model/analysis"
	"github.com/wushiling50/aster/rpc/analysis/internal/pack"
	"github.com/wushiling50/aster/rpc/analysis/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetNationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetNationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetNationLogic {
	return &GetNationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetNationLogic) GetNation(in *analysis.GetAnalysisReq) (*analysis.GetNationResp, error) {
	resp := new(analysis.GetNationResp)

	nation, err := l.getNation(in.DeveloperId)
	if err != nil {
		if errors.Is(err, model_analysis.ErrNotFound) {
			err = errno.BizNotFoundError.WithError(err)
		}

		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.Base = pack.BuildSuccessResp()
	resp.Nation = nation
	return resp, nil
}

func (l *GetNationLogic) getNation(developerId int64) (*analysis.Nation, error) {
	var nation *model_analysis.Nation

	nation, err := l.svcCtx.NationModel.FindOneByDeveloperId(l.ctx, developerId)
	if err != nil {
		return nil, err
	}

	return pack.BuildGenNation(nation), err
}
