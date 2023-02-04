package handler

import (
	"net/http"

	"backend/service/aq/internal/logic"
	"backend/service/aq/internal/svc"
	"backend/service/aq/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DoTaskHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Req
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewDoTaskLogic(r.Context(), svcCtx)
		resp, err := l.DoTask(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
