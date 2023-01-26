package svc

import (
	"backend/service/activities/cmd/api/internal/config"
	"backend/service/activities/cmd/api/internal/middleware"
	"backend/service/activities/cmd/rpc/activitysrv"
	"backend/service/activityorders/cmd/rpc/activitieordersvc"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config           config.Config
	CheckLogin       rest.Middleware
	ActivityRPC      activitysrv.Activitysrv
	ActivityOrderRPC activitieordersvc.ActivitieOrdersvc
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:           c,
		CheckLogin:       middleware.NewCheckLoginMiddleware(c.Mysql.DataSource).Handle,
		ActivityRPC:      activitysrv.NewActivitysrv(zrpc.MustNewClient(c.ActivityRpcConf)),
		ActivityOrderRPC: activitieordersvc.NewActivitieOrdersvc(zrpc.MustNewClient(c.ActivityOrderRpcConf)),
	}
}
