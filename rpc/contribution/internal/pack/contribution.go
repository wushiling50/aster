package pack

import (
	"github.com/wushiling50/aster/gen/contribution"
	model_contribution "github.com/wushiling50/aster/pkg/model/contribution"
)

func BuildGenContribution(model *model_contribution.Contribution) *contribution.Contribution {
	return &contribution.Contribution{
		DataId:                model.DataId,
		DeveloperId:           model.DeveloperId,
		RepoId:                model.RepoId,
		Category:              model.Category,
		Content:               model.Content,
		ContributionCreatedAt: model.ContributionCreatedAt.Unix(),
		ContributionUpdatedAt: model.ContributionUpdatedAt.Unix(),
		ContributionId:        model.ContributionId,
	}
}
