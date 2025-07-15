package logic

import (
	"context"
	"math"

	"github.com/wushiling50/aster/gen/contribution"
	"github.com/wushiling50/aster/rpc/contribution/internal/pack"
	"github.com/wushiling50/aster/rpc/contribution/internal/svc"
	"golang.org/x/sync/errgroup"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelAllContributionInCategoryByDeveloperIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelAllContributionInCategoryByDeveloperIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelAllContributionInCategoryByDeveloperIdLogic {
	return &DelAllContributionInCategoryByDeveloperIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelAllContributionInCategoryByDeveloperIdLogic) DelAllContributionInCategoryByDeveloperId(in *contribution.DelAllContributionInCategoryByDeveloperIdReq) (*contribution.DelAllContributionInCategoryByDeveloperIdResp, error) {
	resp := new(contribution.DelAllContributionInCategoryByDeveloperIdResp)

	err := l.delAllContributionInCategoryByDeveloperIdLogic(in.Category, in.DeveloperId, 1, math.MaxInt64)
	if err != nil {
		logx.Errorf("service.DelAllContributionInCategoryByDeveloperId: Del All Contribution In Category By DeveloperId Failed: %w", err)
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.Base = pack.BuildSuccessResp()

	return resp, nil
}

func (l *DelAllContributionInCategoryByDeveloperIdLogic) delAllContributionInCategoryByDeveloperIdLogic(category string, developerId, page, limit int64) error {
	contributions, err := l.svcCtx.ContributionModel.SearchByDeveloperId(l.ctx, developerId, page, limit)
	if err != nil {
		return err
	}

	eg := errgroup.Group{}
	for _, contribution := range contributions {
		if contribution.Category != category {
			continue
		}

		dataId := contribution.DataId
		eg.Go(func() error {
			return l.svcCtx.ContributionModel.Delete(l.ctx, dataId)
		})
	}

	if err := eg.Wait(); err != nil {
		return err
	}

	return nil
}
