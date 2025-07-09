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

type GetSummaryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSummaryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSummaryLogic {
	return &GetSummaryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetSummaryLogic) GetSummary(in *analysis.GetAnalysisReq) (*analysis.GetSummaryResp, error) {
	resp := new(analysis.GetSummaryResp)

	summary, err := l.getSummary(in.DeveloperId)
	if err != nil {
		if errors.Is(err, model_analysis.ErrNotFound) {
			err = errno.BizNotFoundError.WithError(err)
		}

		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.Base = pack.BuildSuccessResp()
	resp.Summary = summary
	return resp, nil
}

func (l *GetSummaryLogic) getSummary(developerId int64) (*analysis.Summary, error) {
	var summary *model_analysis.Summary

	summary, err := l.svcCtx.SummaryModel.FindOneByDeveloperId(l.ctx, developerId)
	if err != nil {
		return nil, err
	}

	return pack.BuildGenSummary(summary), err
}
