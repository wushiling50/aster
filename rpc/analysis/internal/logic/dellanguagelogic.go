package logic

import (
	"context"

	"github.com/wushiling50/aster/gen/analysis"
	"github.com/wushiling50/aster/rpc/analysis/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelLanguageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelLanguageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelLanguageLogic {
	return &DelLanguageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelLanguageLogic) DelLanguage(in *analysis_analysis.DelAnalysisReq) (*analysis_analysis.DelAnalysisResp, error) {
	// todo: add your logic here and delete this line

	return &analysis_analysis.DelAnalysisResp{}, nil
}
