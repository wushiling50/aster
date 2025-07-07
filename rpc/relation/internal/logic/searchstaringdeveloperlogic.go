package logic

import (
	"context"

	"github.com/wushiling50/aster/gen/relation"
	"github.com/wushiling50/aster/rpc/relation/internal/pack"
	"github.com/wushiling50/aster/rpc/relation/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchStaringDeveloperLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchStaringDeveloperLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchStaringDeveloperLogic {
	return &SearchStaringDeveloperLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchStaringDeveloperLogic) SearchStaringDeveloper(in *relation.SearchStaringDeveloperReq) (*relation.SearchStaringDeveloperResp, error) {
	resp := new(relation.SearchStaringDeveloperResp)

	staringDeveloperIds, err := l.searchStaringDeveloper(in.RepoId, in.Page, in.Limit)
	if err != nil {
		logx.Errorf("service.SearchStaringDeveloper: Search Staring Developer Failed: %w", err)
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	if len(staringDeveloperIds) == 0 {
		logx.Info("service.SearchStaringDeveloper: No Found DeveloperId")
	}

	resp.Base = pack.BuildSuccessResp()
	resp.DeveloperIds = staringDeveloperIds

	return resp, nil
}

func (l *SearchStaringDeveloperLogic) searchStaringDeveloper(repoId, page, limit int64) ([]int64, error) {
	stars, err := l.svcCtx.StarModel.SearchStaringDeveloper(l.ctx, repoId, page, limit)
	if err != nil {
		return nil, err
	}

	staringDeveloperIds := make([]int64, len(stars))
	for i, star := range stars {
		staringDeveloperIds[i] = star.DeveloperId
	}

	return staringDeveloperIds, nil
}
