package admin

import (
	"backend/common/errorx"
	"fmt"
	"net/http"

	"backend/service/backend/cmd/api/internal/logic/admin"
	"backend/service/backend/cmd/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetUserInfoHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("======GetUserInfoHandler===>")

		l := logic.NewGetUserInfoLogic(r.Context(), ctx)

		resp, err := l.GetUserInfo()
		if err != nil {
			//httpx.Error(w, err)
			returnCode := 500
			var data interface{}
			data = ""
			switch e := err.(type) {
			case *errorx.CodeError:
				returnCode = e.DataInfo().Code
				data = e.DataInfo().Data
			default:
				returnCode = 500
			}
			httpx.Error(w, errorx.NewCodeError(returnCode, fmt.Sprintf("%v", err), data))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
