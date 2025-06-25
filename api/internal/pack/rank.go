package pack

import (
	"time"

	"github.com/wushiling50/aster/api/internal/types"

	analysis "github.com/wushiling50/aster/rpc/analysis/analysisclient"
	developer "github.com/wushiling50/aster/rpc/developer/developerclient"
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

func BuildTypeDeveloper(res *developer.GetDeveloperByIdResp) *types.Developer {
	return &types.Developer{
		Id:        res.Developer.Id,
		Name:      res.Developer.Name,
		Login:     res.Developer.Login,
		AvatarUrl: res.Developer.AvatarUrl,
		Company:   res.Developer.Company,
		Location:  res.Developer.Location,
		Bio:       res.Developer.Bio,
		Blog:      res.Developer.Blog,
		Email:     res.Developer.Email,
		CreatedAt: time.Unix(res.Developer.CreatedAt, 0).Format(time.RFC3339),
		UpdatedAt: time.Unix(res.Developer.UpdatedAt, 0).Format(time.RFC3339),
		Following: res.Developer.Following,
		Followers: res.Developer.Followers,
		Gists:     res.Developer.Gists,
		Stars:     res.Developer.Stars,
		Repos:     res.Developer.Repos,
	}
}

func BuildTypeScore(res *analysis.GetScoreResp, id int64, score float64) *types.Score {
	return &types.Score{
		Id:        id,
		Score:     score,
		UpdatedAt: time.Unix(res.Score.DataUpdatedAt, 0).Format(time.RFC3339),
	}
}
