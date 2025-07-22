package pack

import (
	"github.com/wushiling50/aster/api/internal/types"
)

type GetDeveloper struct {
	Developer types.Developer `json:"developer"`
}

type GetLanguageUsage struct {
	LanguageUsage types.LanguageUsage `json:"language_usage"`
}

type GetNation struct {
	Nation types.Nation `json:"nation"`
}

type GetScore struct {
	Score types.Score `json:"score"`
}

type GetSummary struct {
	Summary types.Summary `json:"summary"`
}

type PostTask struct {
	TaskId types.TaskId `json:"task_id"`
}

func BuildDeveloper(res *types.GetDeveloperResp) *GetDeveloper {
	return &GetDeveloper{
		Developer: *res.Developer,
	}
}

func BuildLanguageUsage(res *types.GetLanguageUsageResp) *GetLanguageUsage {
	return &GetLanguageUsage{
		LanguageUsage: *res.LanguageUsage,
	}
}

func BuildNation(res *types.GetNationResp) *GetNation {
	return &GetNation{
		Nation: *res.Nation,
	}
}

func BuildScore(res *types.GetScoreResp) *GetScore {
	return &GetScore{
		Score: *res.Score,
	}
}

func BuildSummary(res *types.GetSummaryResp) *GetSummary {
	return &GetSummary{
		Summary: *res.Summary,
	}
}

func BuildTaskId(id types.TaskId) *PostTask {
	return &PostTask{
		TaskId: id,
	}
}
