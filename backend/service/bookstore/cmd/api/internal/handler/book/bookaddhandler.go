package handler

import (
	"backend/common/errorx"
	logic "backend/service/bookstore/cmd/api/internal/logic/book"
	"backend/service/bookstore/cmd/api/internal/svc"
	"backend/service/bookstore/cmd/api/internal/types"
	"fmt"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

// 图书管理 create
func BookAddHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.BookPostReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, errorx.NewDefaultError(fmt.Sprintf("%v", err), ""))
			return
		}

		l := logic.NewBookAddLogic(r.Context(), ctx)
		_, header, _ := r.FormFile("file")
		resp, err := l.BookAdd(req, header)
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
