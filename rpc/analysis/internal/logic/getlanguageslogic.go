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

type GetLanguagesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLanguagesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLanguagesLogic {
	return &GetLanguagesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetLanguagesLogic) GetLanguages(in *analysis.GetAnalysisReq) (*analysis.GetLanguagesResp, error) {
	resp := new(analysis.GetLanguagesResp)

	languages, err := l.getLanguages(in.DeveloperId)
	if err != nil {
		if errors.Is(err, model_analysis.ErrNotFound) {
			err = errno.BizNotFoundError.WithError(err)
		}

		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.Base = pack.BuildSuccessResp()
	resp.Languages = languages
	return resp, nil
}

func (l *GetLanguagesLogic) getLanguages(developerId int64) (*analysis.Languages, error) {
	var languages *model_analysis.Languages

	languages, err := l.svcCtx.LanguagesModel.FindOneByDeveloperId(l.ctx, developerId)
	if err != nil {
		return nil, err
	}

	return pack.BuildGenLanguages(languages), err
}
