package main

import (
	"backend/common/errorx"
	logic "backend/service/bookstore/cmd/api/internal/logic/bookOrder"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"

	"backend/service/bookstore/cmd/api/internal/config"
	"backend/service/bookstore/cmd/api/internal/handler"
	"backend/service/bookstore/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/books-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	// 开启定时任务,发送还书提醒邮件
	logic.CronInit(ctx)

	handler.RegisterHandlers(server, ctx)

	// 自定义错误
	httpx.SetErrorHandler(func(err error) (int, interface{}) {

		switch e := err.(type) {
		case *errorx.CodeError:
			return http.StatusOK, e.DataInfo()
		default:
			return http.StatusInternalServerError, nil
		}
	})

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
