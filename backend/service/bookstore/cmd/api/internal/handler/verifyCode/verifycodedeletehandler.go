package handler

import (
	"backend/common/errorx"
	"backend/service/bookstore/cmd/api/internal/logic/verifyCode"
	"backend/service/bookstore/cmd/api/internal/svc"
	"backend/service/bookstore/cmd/api/internal/types"
	"fmt"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

// 验证码 delete
func VerifyCodeDeleteHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.VerifyCodeDelReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, errorx.NewDefaultError(fmt.Sprintf("%v", err), ""))
			return
		}

		l := logic.NewVerifyCodeDeleteLogic(r.Context(), ctx)
		resp, err := l.VerifyCodeDelete(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
