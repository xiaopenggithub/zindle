package svc

import (
	"backend/service/activities/cmd/rpc/pb"
	"context"
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

var SvcContext *ServiceContext

const (
	TypeCacheTestPro = "test:cache"
)

// PayloadCacheTestProTask 定时任务参数
type PayloadCacheTestProTask struct {
	Time time.Time
}

func NewTaskCacheTestPro(t time.Time) (*asynq.Task, error) {
	fmt.Println("----NewTaskCacheTestPro----")
	payload, err := json.Marshal(PayloadCacheTestProTask{Time: t})
	if err != nil {
		return nil, err
	}

	return asynq.NewTask(TypeCacheTestPro, payload), nil
}

func HandleTaskCacheTestPro(ctx context.Context, t *asynq.Task) error {
	logx.Info("========HandleTaskCacheTestPro====[" + time.Now().Format("2006-01-02 15:04:05") + "]")
	var p PayloadCacheTestProTask
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return err
	}
	res, _ := SvcContext.ActivityRPC.SearchActivities(ctx, &pb.SearchActivitiesReq{
		Limit: 100,
		Page:  1,
	})
	fmt.Printf("%v", res)
	logx.Info(" [*] Run 任务1 task time [%s]", p.Time.Format("2006-01-02 15:04:05"))
	return nil
}

// AsynqLoggingMiddleware 记录任务日志中间件
func AsynqLoggingMiddleware(h asynq.Handler) asynq.Handler {
	fmt.Println("----AsynqLoggingMiddleware 记录任务日志中间件----")
	return asynq.HandlerFunc(func(ctx context.Context, t *asynq.Task) error {
		start := time.Now()
		logx.Info("开始执行 ", t.Type())
		err := h.ProcessTask(ctx, t)
		if err != nil {
			return err
		}
		logx.Info("结束执行 %q: Elapsed Time = ", t.Type(), time.Since(start))
		return nil
	})
}
