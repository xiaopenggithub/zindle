package svc

import (
	"backend/service/activities/cmd/rpc/activitysrv"
	"backend/service/aq/internal/config"
	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/zrpc"
	"time"
)

type ServiceContext struct {
	Config      config.Config
	AsynqSRV    *asynq.Server
	Scheduler   *asynq.Scheduler
	ActivityRPC activitysrv.Activitysrv
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 初始化定时任务 server 用来启动 worker 处理任务
	asynqSRV := asynq.NewServer(
		&asynq.RedisClientOpt{
			Addr:     c.Redis.Host,
			Password: c.Redis.Pass,
			DB:       9,
		},
		asynq.Config{Concurrency: 20})

	// 初始化定时任务 client scheduler 用来注册定时任务送入队列
	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic(err)
	}
	scheduler := asynq.NewScheduler(
		asynq.RedisClientOpt{
			Addr:     c.Redis.Host,
			Password: c.Redis.Pass,
			DB:       9,
		}, &asynq.SchedulerOpts{
			Location: location,
		})
	return &ServiceContext{
		Config:      c,
		ActivityRPC: activitysrv.NewActivitysrv(zrpc.MustNewClient(c.ActivityRpcConf)),
		AsynqSRV:    asynqSRV,
		Scheduler:   scheduler,
	}
}
