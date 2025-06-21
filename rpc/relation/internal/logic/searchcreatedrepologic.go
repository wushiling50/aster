package logic

import (
	"context"

	"github.com/wushiling50/aster/gen/relation"
	"github.com/wushiling50/aster/rpc/relation/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchCreatedRepoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchCreatedRepoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchCreatedRepoLogic {
	return &SearchCreatedRepoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchCreatedRepoLogic) SearchCreatedRepo(in *relation.SearchCreatedRepoReq) (*relation.SearchCreatedRepoResp, error) {
	// todo: add your logic here and delete this line

	return &relation.SearchCreatedRepoResp{}, nil
}
