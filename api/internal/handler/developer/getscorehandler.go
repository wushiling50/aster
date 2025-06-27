package developer

import (
	"net/http"

	"github.com/wushiling50/aster/api/internal/logic/developer"
	"github.com/wushiling50/aster/api/internal/pack"
	"github.com/wushiling50/aster/api/internal/svc"
	"github.com/wushiling50/aster/api/internal/types"
	"github.com/wushiling50/aster/pkg/errno"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetScoreHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetScoreReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJsonCtx(r.Context(), w, pack.RespError(errno.ParamError.WithError(err)))
			return
		}

		l := developer.NewGetScoreLogic(r.Context(), svcCtx)
		resp, err := l.GetScore(&req)
		if err != nil {
			httpx.OkJsonCtx(r.Context(), w, pack.RespError(err))
		} else {
			httpx.OkJsonCtx(r.Context(), w, pack.RespData(pack.BuildScore(resp)))
		}
	}
}
