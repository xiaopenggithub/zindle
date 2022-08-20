package svc

import (
	"backend/service/activities/cmd/api/internal/config"
	"backend/service/activities/cmd/api/internal/middleware"
	"backend/service/activities/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config         config.Config
	CheckLogin     rest.Middleware
	ActivitysModel model.ActivitysModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:         c,
		CheckLogin:     middleware.NewCheckLoginMiddleware(c.Mysql.DataSource).Handle,
		ActivitysModel: model.NewActivitysModel(conn, c.CacheRedis),
	}
}
