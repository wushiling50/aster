package pack

import (
	"github.com/wushiling50/aster/api/internal/types"
)

type GetScoreRank struct {
	Rank       []*types.DeveloperWithScore `json:"rank"`
	QueryTotal int64                       `json:"query_total"`
	DataTotal  int64                       `json:"data_total"`
}

func BuildRank(res *types.GetScoreRankResp) *GetScoreRank {
	return &GetScoreRank{
		Rank:       res.Rank,
		QueryTotal: res.QueryTotal,
		DataTotal:  res.DataTotal,
	}
}
