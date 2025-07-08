package logic

import (
	"context"

	"github.com/wushiling50/aster/gen/contribution"
	model_contribution "github.com/wushiling50/aster/pkg/model/contribution"
	"github.com/wushiling50/aster/rpc/contribution/internal/pack"
	"github.com/wushiling50/aster/rpc/contribution/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchByRepoIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchByRepoIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchByRepoIdLogic {
	return &SearchByRepoIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchByRepoIdLogic) SearchByRepoId(in *contribution.SearchByRepoIdReq) (*contribution.SearchByRepoIdResp, error) {
	resp := new(contribution.SearchByRepoIdResp)

	contributions, err := l.searchByRepoId(in.RepoId, in.Page, in.Limit)
	if err != nil {
		logx.Errorf("service.SearchByRepoId: Search By RepoId Failed: %w", err)
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	if len(contributions) == 0 {
		logx.Info("service.SearchByRepoId: No Found Contribution By RepoId")
	}

	resp.Contributions = contributions
	resp.Base = pack.BuildSuccessResp()

	return resp, nil
}

func (l *SearchByRepoIdLogic) searchByRepoId(repoId, page, limit int64) ([]*contribution.Contribution, error) {
	var modelContributions []*model_contribution.Contribution

	modelContributions, err := l.svcCtx.ContributionModel.SearchByRepoId(l.ctx, repoId, page, limit)
	if err != nil {
		return nil, err
	}

	contributions := make([]*contribution.Contribution, len(modelContributions))
	for i, modelContribution := range modelContributions {
		contributions[i] = pack.BuildGenContribution(modelContribution)
	}

	return contributions, nil
}
