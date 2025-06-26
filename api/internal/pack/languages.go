package pack

import "github.com/wushiling50/aster/api/internal/types"

type GetLanguages struct {
	LanguageList []types.Language `json:"language_list"`
}

func BuildLanguages(res *types.GetLanguagesResp) *GetLanguages {
	return &GetLanguages{
		LanguageList: res.LanguageList,
	}
}
