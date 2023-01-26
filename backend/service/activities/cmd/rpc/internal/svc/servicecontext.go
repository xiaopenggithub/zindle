package svc

import (
	"backend/service/activities/cmd/rpc/internal/config"
	"backend/service/activities/cmd/rpc/internal/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config          config.Config
	ActivitiesModel model.ActivitiesModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DataSource)
	return &ServiceContext{
		Config:          c,
		ActivitiesModel: model.NewActivitiesModel(conn, c.CacheRedis),
	}
}
