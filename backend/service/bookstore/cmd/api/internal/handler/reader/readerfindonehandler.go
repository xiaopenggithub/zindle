package handler

import (
	"backend/common/errorx"
	"backend/service/bookstore/cmd/api/internal/logic/reader"
	"backend/service/bookstore/cmd/api/internal/svc"
	"backend/service/bookstore/cmd/api/internal/types"
	"fmt"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

// 读者 findOne
func ReaderFindOneHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ReaderDelReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, errorx.NewDefaultError(fmt.Sprintf("%v     %v", err, req), ""))
			return
		}

		l := logic.NewReaderFindOneLogic(r.Context(), ctx)
		resp, err := l.ReaderFindOne(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
