package rank

import (
	"net/http"

	"github.com/wushiling50/aster/api/internal/logic/rank"
	"github.com/wushiling50/aster/api/internal/svc"
	"github.com/wushiling50/aster/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetScoreRankHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetScoreRankReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := rank.NewGetScoreRankLogic(r.Context(), svcCtx)
		resp, err := l.GetScoreRank(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
