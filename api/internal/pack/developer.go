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

func BuildNation(res *types.GetNationResp) *GetNation {
	return &GetNation{
		Nation: res.Nation,
	}
}
