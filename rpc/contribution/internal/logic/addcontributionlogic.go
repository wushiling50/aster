package logic

import (
	"context"
	"time"

	"github.com/wushiling50/aster/gen/contribution"
	model_contribution "github.com/wushiling50/aster/pkg/model/contribution"
	"github.com/wushiling50/aster/rpc/contribution/internal/pack"
	"github.com/wushiling50/aster/rpc/contribution/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddContributionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddContributionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddContributionLogic {
	return &AddContributionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddContributionLogic) AddContribution(in *contribution.AddContributionReq) (*contribution.AddContributionResp, error) {
	resp := new(contribution.AddContributionResp)

	err := l.addContribution(&model_contribution.Contribution{
		ContributionId: in.ContributionId,
		DeveloperId:    in.DeveloperId,
		RepoId:         in.RepoId,
		Category:       in.Category,
		Content:        in.Content,
		CreatedAt:      time.Unix(in.ContributionCreatedAt, 0),
		UpdatedAt:      time.Unix(in.ContributionUpdatedAt, 0),
	})

	if err != nil {
		logx.Errorf("service.AddContribution: Add Contribution Failed: %w", err)
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.Base = pack.BuildSuccessResp()

	return resp, nil
}

func (l *AddContributionLogic) addContribution(model *model_contribution.Contribution) error {
	dataId, err := l.svcCtx.ContributionModel.CreateDataId()
	if err != nil {
		return err
	}

	model.DataId = dataId
	_, err = l.svcCtx.ContributionModel.Insert(l.ctx, model)
	if err != nil {
		return err
	}

	return nil
}
