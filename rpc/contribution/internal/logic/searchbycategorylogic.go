package logic

import (
	"context"

	"github.com/wushiling50/aster/gen/contribution"
	model_contribution "github.com/wushiling50/aster/pkg/model/contribution"
	"github.com/wushiling50/aster/rpc/contribution/internal/pack"
	"github.com/wushiling50/aster/rpc/contribution/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type SearchByCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchByCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchByCategoryLogic {
	return &SearchByCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchByCategoryLogic) SearchByCategory(in *contribution.SearchByCategoryReq) (*contribution.SearchByCategoryResp, error) {
	resp := new(contribution.SearchByCategoryResp)

	contributions, err := l.searchByCategory(in.Category, in.Page, in.Limit)
	if err != nil {
		logx.Errorf("service.SearchByCategory: Search By Category Failed: %w", err)
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	if len(contributions) == 0 {
		logx.Info("service.SearchByCategory: No Found Contribution By Category")
	}

	resp.Contributions = contributions
	resp.Base = pack.BuildSuccessResp()

	return resp, nil
}

func (l *SearchByCategoryLogic) searchByCategory(category string, page, limit int64) ([]*contribution.Contribution, error) {
	var modelContributions []*model_contribution.Contribution

	modelContributions, err := l.svcCtx.ContributionModel.SearchByCategory(l.ctx, category, page, limit)
	if err != nil {
		return nil, err
	}

	contributions := make([]*contribution.Contribution, len(modelContributions))
	for i, modelContribution := range modelContributions {
		contributions[i] = pack.BuildGenContribution(modelContribution)
	}

	return contributions, nil
}
