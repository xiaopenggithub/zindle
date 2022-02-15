package logic

import (
	"backend/common/errorx"
	logic "backend/service/bookstore/cmd/api/internal/logic/reader"
	"backend/service/bookstore/cmd/api/internal/svc"
	"backend/service/bookstore/cmd/api/internal/types"
	"context"
	"encoding/json"
	"fmt"
	"github.com/robfig/cron"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

type ReturnBookLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReturnBookLogic(ctx context.Context, svcCtx *svc.ServiceContext) ReturnBookLogic {
	return ReturnBookLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReturnBookLogic) ReturnBook(req types.BookOrderDelReq) (*types.BookOrderReply, error) {
	//获取用户订单
	//事务处理
	idNumber := json.Number(fmt.Sprintf("%v", l.ctx.Value("readerId")))
	id, err := idNumber.Int64()
	if err != nil {
		return nil, errorx.NewCodeError(401, "请重新登录再操作", "")
	}
	one, err := l.svcCtx.BookOrdersModel.FindOneOrder(req.Id, id)
	if err != nil {
		return nil, errorx.NewCodeError(201, fmt.Sprintf("%v", err), "")
	}
	err = l.svcCtx.BookOrdersModel.ReturnBook(*one)
	if err != nil {
		return nil, errorx.NewCodeError(500, fmt.Sprintf("%v", err), "")
	}
	return nil, errorx.NewCodeError(200, fmt.Sprintf("%s", "还书成功"), "")
}

func CheckExpire() {
	/*
		to := []string{"abc@qq.com"}
		subject := "邮箱验证"
		code := 1234
		body := fmt.Sprintf("您的验证码是%v,有效期5分钟", code)
		err := send(to, subject, body, l.svcCtx.Config)
		if err != nil {
			return
		}
	*/
	fmt.Println(">>>>-ReturnBook-Check10Second<<<<[", time.Now().Format("2006-01-02 15:04:05"), "]")
}
func CronInit(ctx *svc.ServiceContext) {
	//https://www.cnblogs.com/liuzhongchao/p/9521897.html
	/*
	   每隔5秒执行一次：#/5 * * * * ?
	   每隔1分钟执行一次：0 #/1 * * * ?
	   每天23点执行一次：0 0 23 * * ?
	   每天凌晨1点执行一次：0 0 1 * * ?
	   每月1号凌晨1点执行一次：0 0 1 1 * ?
	   在26分、29分、33分执行一次：0 26,29,33 * * * ?
	   每天的0点、13点、18点、21点都执行一次：0 0 0,13,18,21 * * ?
	*/

	go func() {
		crontab := cron.New()

		// 每天9点执行
		crontab.AddFunc("0 0 9 * * ?", func() {
			// todo 阶梯时间段提醒，(快到期有限次提醒/已经超过60天的)
			// 1. 获取需要提醒的订单（借60天还书，到期前7天开始提醒）
			// 2. 获取用户邮箱
			// 3. 发送邮件
			nearTime := time.Now().Add(-53 * 24 * time.Hour).Format("2006-01-02")
			bookOrders, err := ctx.BookOrdersModel.NearReturnBook(nearTime)
			if err != nil {
				// no
				fmt.Println(">>>>no data<<<<")
			} else {
				userSet := make(map[int64]bool)
				for _, v := range bookOrders {
					userSet[v.MemberId] = true
				}
				// 获取用户邮箱
				ids := ""
				for k, _ := range userSet {
					ids += fmt.Sprintf("%v,", k)
				}
				ids = ids[:len(ids)-1]
				readers, err := ctx.ReadersModel.FindByIds(ids)
				if err != nil {
					fmt.Println(">>>>no expired users<<<<")
				}
				for _, v := range readers {
					subject := "还书提醒"
					body := "你有图书待归还，如已还，请忽略此邮件"
					go logic.Send([]string{v.Email}, subject, body, ctx.Config)
				}
			}

		}) // 每秒 定时执行 myfunc 函数
		crontab.Start()
	}()
}
