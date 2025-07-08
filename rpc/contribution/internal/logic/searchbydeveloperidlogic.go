package logic

import (
	"context"

	"github.com/wushiling50/aster/gen/contribution"
	model_contribution "github.com/wushiling50/aster/pkg/model/contribution"
	"github.com/wushiling50/aster/rpc/contribution/internal/pack"
	"github.com/wushiling50/aster/rpc/contribution/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchByDeveloperIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchByDeveloperIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchByDeveloperIdLogic {
	return &SearchByDeveloperIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchByDeveloperIdLogic) SearchByDeveloperId(in *contribution.SearchByDeveloperIdReq) (*contribution.SearchByDeveloperIdResp, error) {
	resp := new(contribution.SearchByDeveloperIdResp)

	contributions, err := l.searchByDeveloperId(in.DeveloperId, in.Page, in.Limit)
	if err != nil {
		logx.Errorf("service.SearchByDeveloperId: Search By DeveloperId Failed: %w", err)
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	if len(contributions) == 0 {
		logx.Info("service.SearchByDeveloperId: No Found Contribution By DeveloperId")
	}

	resp.Contributions = contributions
	resp.Base = pack.BuildSuccessResp()

	return resp, nil
}

func (l *SearchByDeveloperIdLogic) searchByDeveloperId(developerId, page, limit int64) ([]*contribution.Contribution, error) {
	var modelContributions []*model_contribution.Contribution

	modelContributions, err := l.svcCtx.ContributionModel.SearchByDeveloperId(l.ctx, developerId, page, limit)
	if err != nil {
		return nil, err
	}

	contributions := make([]*contribution.Contribution, len(modelContributions))
	for i, modelContribution := range modelContributions {
		contributions[i] = pack.BuildGenContribution(modelContribution)
	}

	return contributions, nil
}
