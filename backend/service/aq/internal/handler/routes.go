// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"backend/service/aq/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/chat/doTask",
				Handler: DoTaskHandler(serverCtx),
			},
		},
	)
}
