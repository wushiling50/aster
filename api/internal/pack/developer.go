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

func BuildDeveloper(res *types.GetDeveloperResp) *GetDeveloper {
	return &GetDeveloper{
		Developer: res.Developer,
	}
}

func BuildLanguageUsage(res *types.GetLanguageUsageResp) *GetLanguageUsage {
	return &GetLanguageUsage{
		LanguageUsage: res.LanguageUsage,
	}
}
