package logic

import (
	"context"

	"github.com/wushiling50/aster/gen/analysis"
	"github.com/wushiling50/aster/rpc/analysis/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLanguageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateLanguageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLanguageLogic {
	return &UpdateLanguageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateLanguageLogic) UpdateLanguage(in *analysis.UpdateAnalysisReq) (*analysis.UpdateAnalysisResp, error) {
	// todo: add your logic here and delete this line

	return &analysis.UpdateAnalysisResp{}, nil
}
