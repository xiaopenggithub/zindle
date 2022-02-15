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

// 图书管理 list
func BookListHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.BookListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, errorx.NewDefaultError(fmt.Sprintf("%v", err), ""))
			return
		}

		l := logic.NewBookListLogic(r.Context(), ctx)
		resp, err := l.BookList(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
