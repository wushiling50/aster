package developer

import (
	"context"

	"github.com/wushiling50/aster/api/internal/svc"
	"github.com/wushiling50/aster/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLanguageUsageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetLanguageUsageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLanguageUsageLogic {
	return &GetLanguageUsageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetLanguageUsageLogic) GetLanguageUsage(req *types.GetLanguageUsageReq) (resp *types.GetLanguageUsageResp, err error) {
	// todo: add your logic here and delete this line

	return
}
