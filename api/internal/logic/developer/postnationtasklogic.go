package developer

import (
	"context"

	"github.com/wushiling50/aster/api/internal/svc"
	"github.com/wushiling50/aster/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PostNationTaskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPostNationTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostNationTaskLogic {
	return &PostNationTaskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PostNationTaskLogic) PostNationTask(req *types.PostTaskReq) (resp *types.PostTaskResp, err error) {
	// todo: add your logic here and delete this line

	return
}
