package logic

import (
	"backend/common/errorx"
	"backend/common/utils"
	"backend/service/bookstore/cmd/api/internal/svc"
	"backend/service/bookstore/cmd/api/internal/types"
	"context"
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
)

type BookOrderListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 图书订单 list
func NewBookOrderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) BookOrderListLogic {
	return BookOrderListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BookOrderListLogic) BookOrderList(req types.BookOrderListReq) (*types.BookOrderReply, error) {
	reqParam := utils.ListReq{}
	reqParam.Page = req.Page
	reqParam.PageSize = req.PageSize
	reqParam.Keyword = req.Keyword
	list, i, err := l.svcCtx.BookOrdersModel.List(reqParam)
	if err != nil {
		return nil, errorx.NewCodeError(201, fmt.Sprintf("%v", err), "")
	}
	data := make(map[string]interface{})
	data["page"] = req.Page
	data["pageSize"] = req.PageSize
	data["total"] = i
	//data["list"] = list
	// 获取所有书 todo: 后续放redis
	booksMap := make(map[int64]interface{})
	books, _, _ := l.svcCtx.BooksModel.List(utils.ListReq{
		Page:     1,
		PageSize: 200,
	})
	for _, v := range books {
		booksMap[v.Id] = v.Title
	}

	// 获取所有用户 todo: 后续放redis
	readersMap := make(map[int64]interface{})
	readers, _, _ := l.svcCtx.ReadersModel.List(utils.ListReq{
		Page:     1,
		PageSize: 200,
	})
	for _, v := range readers {
		readersMap[v.Id] = v.Username
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
		m["book_name"] = ""
		if bookName, ok := booksMap[v.BookId]; ok {
			m["book_name"] = bookName
		}
		//
		m["member_name"] = ""
		if readerName, ok := readersMap[v.MemberId]; ok {
			m["member_name"] = readerName
		}

		resultMap = append(resultMap, m)
	}
	data["list"] = resultMap

	return nil, errorx.NewCodeError(200, "ok", data)
}
