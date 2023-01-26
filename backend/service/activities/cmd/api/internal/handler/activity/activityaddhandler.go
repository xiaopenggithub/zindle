package activity

import (
	"backend/common/errorx"
	"fmt"
	"net/http"

	"backend/service/activities/cmd/api/internal/logic/activity"
	"backend/service/activities/cmd/api/internal/svc"
	"backend/service/activities/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ActivityAddHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ActivityPostReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, errorx.NewDefaultError(fmt.Sprintf("%v", err), ""))
			return
		}

		l := activity.NewActivityAddLogic(r.Context(), svcCtx)
		resp, err := l.ActivityAdd(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
