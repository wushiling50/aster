package pack

import (
	"github.com/wushiling50/aster/api/internal/types"
)

type GetScoreRank struct {
	Rank  []*types.DeveloperWithScore `json:"rank"`
	Total int64                       `json:"total"`
}

func BuildRank(res *types.GetScoreRankResp) *GetScoreRank {
	return &GetScoreRank{
		Rank:  res.Rank,
		Total: res.Total,
	}
}
