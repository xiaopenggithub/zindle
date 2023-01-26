package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ActivityOrdersModel = (*customActivityOrdersModel)(nil)

type (
	// ActivityOrdersModel is an interface to be customized, add more methods here,
	// and implement the added methods in customActivityOrdersModel.
	ActivityOrdersModel interface {
		activityOrdersModel
	}

	customActivityOrdersModel struct {
		*defaultActivityOrdersModel
	}
)

// NewActivityOrdersModel returns a model for the database table.
func NewActivityOrdersModel(conn sqlx.SqlConn, c cache.CacheConf) ActivityOrdersModel {
	return &customActivityOrdersModel{
		defaultActivityOrdersModel: newActivityOrdersModel(conn, c),
	}
}
