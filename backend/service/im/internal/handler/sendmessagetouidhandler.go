package handler

import (
	"backend/common/errorx"
	"fmt"
	"net/http"

	"backend/service/im/internal/logic"
	"backend/service/im/internal/svc"
	"backend/service/im/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func SendMessageToUidHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MessageReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, errorx.NewDefaultError(fmt.Sprintf("%v", err), ""))
			return
		}

		l := logic.NewSendMessageToUidLogic(r.Context(), svcCtx)
		resp, err := l.SendMessageToUid(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
