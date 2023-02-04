package main

import (
	"flag"
	"fmt"
	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/logx"
	"time"

	"backend/service/aq/internal/config"
	"backend/service/aq/internal/handler"
	"backend/service/aq/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/aq-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	svc.SvcContext = svc.NewServiceContext(c)
	handler.RegisterHandlers(server, svc.SvcContext)

	// 初始化 asynq worker
	mux := asynq.NewServeMux()
	mux.Use(svc.AsynqLoggingMiddleware)
	// 注册任务
	mux.HandleFunc(svc.TypeCacheTestPro, svc.HandleTaskCacheTestPro)
	go svc.SvcContext.AsynqSRV.Run(mux)

	// 初始化 asynq client
	t1, err := svc.NewTaskCacheTestPro(time.Now())
	if err != nil {
		logx.Error(err)
	}
	_, err = svc.SvcContext.Scheduler.Register("*/1 * * * *", t1)
	if err != nil {
		logx.Error(err)
	}

	// 初始化 asynq client2
	t2, err := svc.NewTaskCacheTestPro(time.Now())
	if err != nil {
		logx.Error(err)
	}
	_, err = svc.SvcContext.Scheduler.Register("*/2 * * * *", t2)
	if err != nil {
		logx.Error(err)
	}

	go svc.SvcContext.Scheduler.Run()

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
