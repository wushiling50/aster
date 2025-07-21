package pack

import (
	"strconv"
	"time"

	"github.com/wushiling50/aster/api/internal/types"
	"github.com/wushiling50/aster/pkg/errno"
	analysis "github.com/wushiling50/aster/rpc/analysis/analysisclient"
	developer "github.com/wushiling50/aster/rpc/developer/developerclient"
)

type Base struct {
	Code string `json:"code"`
	Msg  string `json:"message"`
}

type RespWithData struct {
	Code string `json:"code"`
	Msg  string `json:"message"`
	Data any    `json:"data"`
}

func RespSuccess() *Base {
	Errno := errno.Success
	return &Base{
		Code: strconv.FormatInt(Errno.ErrorCode, 10),
		Msg:  Errno.ErrorMsg,
	}
}

func RespData(data any) *RespWithData {
	return &RespWithData{
		Code: strconv.FormatInt(errno.SuccessCode, 10),
		Msg:  "Success",
		Data: data,
	}
}

func RespError(err error) *Base {
	Errno := errno.ConvertErr(err)
	return &Base{
		Code: strconv.FormatInt(Errno.ErrorCode, 10),
		Msg:  Errno.ErrorMsg,
	}
}

func BuildTypeDeveloper(res *developer.Developer) *types.Developer {
	return &types.Developer{
		Id:        res.Id,
		Name:      res.Name,
		Login:     res.Login,
		AvatarUrl: res.AvatarUrl,
		Company:   res.Company,
		Location:  res.Location,
		Bio:       res.Bio,
		Blog:      res.Blog,
		Email:     res.Email,
		Following: res.Following,
		Followers: res.Followers,
		Gists:     res.Gists,
		Stars:     res.Stars,
		Repos:     res.Repos,
	}
}

func BuildTypeScore(res *analysis.Score, id int64, score float64) *types.Score {
	return &types.Score{
		Id:        id,
		Score:     score,
		UpdatedAt: time.Unix(res.UpdatedAt, 0).Format(time.RFC3339),
	}
}
