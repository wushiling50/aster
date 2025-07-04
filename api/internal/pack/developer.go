package pack

import (
	"github.com/wushiling50/aster/api/internal/types"
)

type GetDeveloper struct {
	Developer types.Developer `json:"developer"`
}

type GetLanguageUsage struct {
	LanguageUsage types.LanguageUsage `json:"language_usage"`
	TaskState     types.TaskState     `json:"task_state"`
}

type GetNation struct {
	Nation    types.Nation    `json:"nation"`
	TaskState types.TaskState `json:"task_state"`
}

type GetScore struct {
	Score     types.Score     `json:"score"`
	TaskState types.TaskState `json:"task_state"`
}

type GetSummary struct {
	Summary   types.Summary   `json:"summary"`
	TaskState types.TaskState `json:"task_state"`
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
		TaskState:     res.TaskState,
	}
}

func BuildNation(res *types.GetNationResp) *GetNation {
	return &GetNation{
		Nation:    *res.Nation,
		TaskState: res.TaskState,
	}
}

func BuildScore(res *types.GetScoreResp) *GetScore {
	return &GetScore{
		Score:     *res.Score,
		TaskState: res.TaskState,
	}
}

func BuildSummary(res *types.GetSummaryResp) *GetSummary {
	return &GetSummary{
		Summary:   *res.Summary,
		TaskState: res.TaskState,
	}
}

func BuildTaskId(id types.TaskId) *PostTask {
	return &PostTask{
		TaskId: id,
	}
}
