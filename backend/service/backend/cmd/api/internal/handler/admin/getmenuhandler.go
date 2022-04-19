package admin

import (
	"backend/common/errorx"
	logic "backend/service/backend/cmd/api/internal/logic/admin"
	"fmt"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"

	"backend/service/backend/cmd/api/internal/svc"
)

func GetMenuHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewGetMenuLogic(r.Context(), ctx)
		resp, err := l.GetMenu()
		if err != nil {
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
