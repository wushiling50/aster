package rank

import (
	"net/http"

	"github.com/wushiling50/aster/api/internal/logic/rank"
	"github.com/wushiling50/aster/api/internal/pack"
	"github.com/wushiling50/aster/api/internal/svc"
	"github.com/wushiling50/aster/api/internal/types"
	"github.com/wushiling50/aster/pkg/errno"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetScoreRankHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetScoreRankReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJsonCtx(r.Context(), w, pack.RespError(errno.ParamError.WithError(err)))
			return
		}

		if req.Limit < 10 {
			req.Limit = 10
		} else if req.Limit > 100 {
			req.Limit = 100
		}

		l := rank.NewGetScoreRankLogic(r.Context(), svcCtx)
		resp, err := l.GetScoreRank(&req)
		if err != nil {
			httpx.OkJsonCtx(r.Context(), w, pack.RespError(err))
		} else {
			httpx.OkJsonCtx(r.Context(), w, pack.RespData(pack.BuildRank(resp)))
		}
	}
}
