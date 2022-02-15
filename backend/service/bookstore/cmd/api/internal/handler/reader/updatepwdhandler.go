package handler

import (
	"net/http"

	"backend/service/bookstore/cmd/api/internal/logic/reader"
	"backend/service/bookstore/cmd/api/internal/svc"
	"backend/service/bookstore/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func UpdatePwdHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdatePwdReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUpdatePwdLogic(r.Context(), ctx)
		resp, err := l.UpdatePwd(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
