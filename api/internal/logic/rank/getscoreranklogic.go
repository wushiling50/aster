package rank

import (
	"context"

	"github.com/wushiling50/aster/api/internal/svc"
	"github.com/wushiling50/aster/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetScoreRankLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetScoreRankLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetScoreRankLogic {
	return &GetScoreRankLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetScoreRankLogic) GetScoreRank(req *types.GetScoreRankReq) (resp *types.GetScoreRankResp, err error) {
	// todo: add your logic here and delete this line

	return
}
