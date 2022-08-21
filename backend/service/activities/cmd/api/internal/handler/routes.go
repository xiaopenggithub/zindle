// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	activity "backend/service/activities/cmd/api/internal/handler/activity"
	activityOrders "backend/service/activities/cmd/api/internal/handler/activityOrders"
	"backend/service/activities/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.CheckLogin},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/activity/list",
					Handler: activity.ActivityListHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/activity/delete",
					Handler: activity.ActivityDeleteHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/activity/deleteBatch",
					Handler: activity.ActivityDeleteBatchHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/activity/find",
					Handler: activity.ActivityFindOneHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/activity/add",
					Handler: activity.ActivityAddHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/activity/update",
					Handler: activity.ActivityUpdateHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/activity/appList",
				Handler: activity.ActivityAppListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/activity/appFind",
				Handler: activity.ActivityAppFindOneHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.CheckLogin},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/activityOrders/list",
					Handler: activityOrders.ActivityOrdersListHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/activityOrders/delete",
					Handler: activityOrders.ActivityOrdersDeleteHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/activityOrders/deleteBatch",
					Handler: activityOrders.ActivityOrdersDeleteBatchHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/activityOrders/find",
					Handler: activityOrders.ActivityOrdersFindOneHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/activityOrders/add",
					Handler: activityOrders.ActivityOrdersAddHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/activityOrders/update",
					Handler: activityOrders.ActivityOrdersUpdateHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)
}