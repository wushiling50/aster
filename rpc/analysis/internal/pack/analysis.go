package pack

import (
	"github.com/wushiling50/aster/gen/analysis"
	model_analysis "github.com/wushiling50/aster/pkg/model/analysis"
)

func BuildGenLanguages(model *model_analysis.Languages) *analysis.Languages {
	return &analysis.Languages{
		DataId:        model.DataId,
		DataCreatedAt: model.DataCreatedAt.Unix(),
		DataUpdatedAt: model.DataUpdatedAt.Unix(),
		Languages:     model.Language,
	}
}

func BuildGenNation(model *model_analysis.Nation) *analysis.Nation {
	return &analysis.Nation{
		DataId:        model.DataId,
		DataCreatedAt: model.DataCreatedAt.Unix(),
		DataUpdatedAt: model.DataUpdatedAt.Unix(),
		Nation:        model.Nation,
		Confidence:    model.Confidence,
	}
}

func BuildGenScore(model *model_analysis.Score) *analysis.Score {
	return &analysis.Score{
		DataId:        model.DataId,
		DataCreatedAt: model.DataCreatedAt.Unix(),
		DataUpdatedAt: model.DataUpdatedAt.Unix(),
		Score:         model.Score,
	}
}

func BuildGenSummary(model *model_analysis.Summary) *analysis.Summary {
	return &analysis.Summary{
		DataId:        model.DataId,
		DataCreatedAt: model.DataCreatedAt.Unix(),
		DataUpdatedAt: model.DataUpdatedAt.Unix(),
		Summary:       model.Summary,
	}
}
