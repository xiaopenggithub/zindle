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

// 验证码 create
func VerifyCodeAddHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.VerifyCodePostReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, errorx.NewDefaultError(fmt.Sprintf("%v", err), ""))
			return
		}

		l := logic.NewVerifyCodeAddLogic(r.Context(), ctx)
		resp, err := l.VerifyCodeAdd(req)
		if err != nil {
			returnCode := 500
			switch e := err.(type) {
			case *errorx.CodeError:
				returnCode = e.DataInfo().Code
			default:
				returnCode = 500
			}
			httpx.Error(w, errorx.NewCodeError(returnCode, fmt.Sprintf("%v", err), ""))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
