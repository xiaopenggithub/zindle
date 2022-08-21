package handler

import (
	"backend/common/errorx"
	"backend/service/activities/cmd/api/internal/logic/activityOrders"
	"backend/service/activities/cmd/api/internal/svc"
	"backend/service/activities/cmd/api/internal/types"
	"fmt"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

// 活动预订信息 create
func ActivityOrdersAddHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ActivityOrdersPostReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, errorx.NewDefaultError(fmt.Sprintf("%v", err), ""))
			return
		}

		l := logic.NewActivityOrdersAddLogic(r.Context(), ctx)
		resp, err := l.ActivityOrdersAdd(req)
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
