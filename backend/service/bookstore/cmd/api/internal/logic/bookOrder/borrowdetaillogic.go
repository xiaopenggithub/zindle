package logic

import (
	"backend/common/errorx"
	"backend/service/bookstore/cmd/api/internal/svc"
	"backend/service/bookstore/cmd/api/internal/types"
	"context"
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

type BorrowDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBorrowDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) BorrowDetailLogic {
	return BorrowDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BorrowDetailLogic) BorrowDetail(req types.BookOrderDelReq) (*types.BookOrderReply, error) {
	idNumber := json.Number(fmt.Sprintf("%v", l.ctx.Value("readerId")))
	id, err := idNumber.Int64()
	if err != nil {
		return nil, errorx.NewCodeError(401, "请重新登录再操作", "")
	}
	one, err := l.svcCtx.BookOrdersModel.FindOneOrder(req.Id, id)
	if err != nil {
		return nil, errorx.NewCodeError(201, fmt.Sprintf("%v", err), "")
	}
	data := make(map[string]interface{})
	//one.Cover = "http://192.168.1.2:8080/uploads1/" + one.Cover

	// 结构体转json
	sJson, err := json.Marshal(one)
	if err != nil {
		fmt.Println(err)
	}
	// json 转map
	item := make(map[string]interface{})
	json.Unmarshal(sJson, &item)

	t, err := time.Parse("2006-01-02T15:04:05+08:00", one.CreatedAt)
	add := t.Add(time.Hour * 24 * 60)
	item["created_at"] = t.Format("2006-01-02")
	item["shourld_return"] = add.Format("2006-01-02")

	dateFormat := "2006-01-02 15:04:05"
	t1, err := time.Parse("2006-01-02T15:04:05+08:00", one.ReturnDate)
	item["return_date"] = t1.Format(dateFormat)

	//状态
	item["status"] = "已还"
	fmt.Println(one.ReturnDate, item["return_date"], one.Id)
	if item["return_date"] == dateFormat {
		item["status"] = "待还"
	}

	// 添加其它字段
	book, _ := l.svcCtx.BooksModel.FindOne(one.BookId)
	item["title"] = book.Title
	item["cover"] = "http://192.168.1.2:8080/uploads1/" + book.Cover
	item["description"] = book.Description

	data["item"] = item

	return nil, errorx.NewCodeError(200, fmt.Sprintf("%s", "获取成功"), data)
}
