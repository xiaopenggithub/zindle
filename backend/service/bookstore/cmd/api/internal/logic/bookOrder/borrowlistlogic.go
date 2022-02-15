package logic

import (
	"backend/common/errorx"
	"backend/common/utils"
	"backend/service/bookstore/cmd/api/internal/svc"
	"backend/service/bookstore/cmd/api/internal/types"
	"backend/service/bookstore/model"
	"context"
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

type BorrowListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBorrowListLogic(ctx context.Context, svcCtx *svc.ServiceContext) BorrowListLogic {
	return BorrowListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BorrowListLogic) BorrowList(req types.BookOrderListReq) (*types.BookOrderReply, error) {
	reqParam := utils.ListReq{}
	reqParam.Page = req.Page
	reqParam.PageSize = req.PageSize
	reqParam.Keyword = req.Keyword

	idNumber := json.Number(fmt.Sprintf("%v", l.ctx.Value("readerId")))
	id, err := idNumber.Int64()
	if err != nil {
		return nil, errorx.NewCodeError(401, "请重新登录再操作", "")
	}

	list, i, err := l.svcCtx.BookOrdersModel.ListBorrow(reqParam, id)
	if err != nil {
		return nil, errorx.NewCodeError(201, fmt.Sprintf("%v", err), "")
	}
	data := make(map[string]interface{})
	data["page"] = req.Page
	data["pageSize"] = req.PageSize
	data["total"] = i
	//data["list"] = list
	// 获取所有书 todo: 后续放redis
	booksMap := make(map[int64]*model.Books)
	books, _, _ := l.svcCtx.BooksModel.List(utils.ListReq{
		Page:     1,
		PageSize: 200,
	})
	for _, v := range books {
		booksMap[v.Id] = v
	}

	// todo:
	var resultMap []interface{}
	// 获取书名称
	// 获取用户名
	for _, v := range list {

		// 结构体转json
		sJson, err := json.Marshal(v)
		if err != nil {
			fmt.Println(err)
		}

		// json 转map
		m := make(map[string]interface{})
		json.Unmarshal(sJson, &m)

		// 补充字段
		m["title"] = ""
		m["description"] = ""
		m["cover"] = ""
		if m2, ok := booksMap[v.BookId]; ok {
			m["title"] = m2.Title
			m["description"] = m2.Description
			m["cover"] = "http://192.168.1.2:8080/uploads1/" + m2.Cover

		}
		dateFormat := "2006-01-02 15:04:05"
		// 应还：借出60天应还
		//t, err := time.Parse("2006-01-02T15:04:05.000+08:00", v.CreatedAt)
		t, err := time.Parse("2006-01-02T15:04:05+08:00", v.CreatedAt)
		add := t.Add(time.Hour * 24 * 60)
		//m["created_at"] = t.Format("2006-01-02 15:04:05")
		//m["shourld_return"] = add.Format("2006-01-02 15:04:05")
		m["created_at"] = t.Format("2006-01-02")
		m["shourld_return"] = add.Format("2006-01-02")

		t1, err := time.Parse("2006-01-02T15:04:05+08:00", v.ReturnDate)
		m["return_date"] = t1.Format(dateFormat)
		//fmt.Println(v.ReturnDate, m["return_date"], v.Id)

		//状态
		m["status"] = "已还"
		if m["return_date"] == dateFormat {
			m["status"] = "待还"
		}
		delete(m, "updated_at")
		delete(m, "deleted_at")
		resultMap = append(resultMap, m)
	}
	data["list"] = resultMap

	return nil, errorx.NewCodeError(200, "ok", data)
}
