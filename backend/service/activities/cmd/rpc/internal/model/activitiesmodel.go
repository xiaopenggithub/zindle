package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ActivitiesModel = (*customActivitiesModel)(nil)

type (
	// ActivitiesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customActivitiesModel.
	ActivitiesModel interface {
		activitiesModel
	}

	customActivitiesModel struct {
		*defaultActivitiesModel
	}
)

// NewActivitiesModel returns a model for the database table.
func NewActivitiesModel(conn sqlx.SqlConn, c cache.CacheConf) ActivitiesModel {
	return &customActivitiesModel{
		defaultActivitiesModel: newActivitiesModel(conn, c),
	}
}
