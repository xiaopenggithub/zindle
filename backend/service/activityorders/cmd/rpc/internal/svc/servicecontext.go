package svc

import (
	"backend/service/activityorders/cmd/rpc/internal/config"
	"backend/service/activityorders/cmd/rpc/internal/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config              config.Config
	ActivityOrdersModel model.ActivityOrdersModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DataSource)
	return &ServiceContext{
		Config:              c,
		ActivityOrdersModel: model.NewActivityOrdersModel(conn, c.CacheRedis),
	}
}
