package svc

import (
	"backend/service/bookstore/cmd/api/internal/config"
	"backend/service/bookstore/cmd/api/internal/middleware"
	"backend/service/bookstore/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config           config.Config
	CheckLogin       rest.Middleware
	BooksModel       model.BooksModel
	BookOrdersModel  model.BookOrdersModel
	ReadersModel     model.ReadersModel
	VerifyCodesModel model.VerifyCodesModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:           c,
		CheckLogin:       middleware.NewCheckLoginMiddleware(c.Mysql.DataSource).Handle,
		BooksModel:       model.NewBooksModel(conn, c.CacheRedis),
		BookOrdersModel:  model.NewBookOrdersModel(conn, c.CacheRedis),
		ReadersModel:     model.NewReadersModel(conn, c.CacheRedis),
		VerifyCodesModel: model.NewVerifyCodesModel(conn, c.CacheRedis),
	}
}
