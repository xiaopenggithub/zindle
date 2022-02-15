package logic

import (
	"backend/common/errorx"
	"backend/service/bookstore/cmd/api/internal/svc"
	"backend/service/bookstore/cmd/api/internal/types"
	"backend/service/bookstore/model"
	"context"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type BookOrderAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 图书订单 create
func NewBookOrderAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) BookOrderAddLogic {
	return BookOrderAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BookOrderAddLogic) BookOrderAdd(req types.BookOrderPostReq) (*types.BookOrderReply, error) {
	//检查是否重名
	//checkItem, err2 := l.svcCtx.BookOrdersModel.CheckDuplicate(req.)
	//if checkItem.Id > 0 && err2 == nil {
	//	return nil, errorx.NewCodeError(201, "已经存在", "")
	//}

	var item model.BookOrders
	copier.Copy(&item, &req)

	_, err := l.svcCtx.BookOrdersModel.Insert(item)
	if err != nil {
		return nil, errorx.NewCodeError(202, fmt.Sprintf("%v", err), "")
	}

	return nil, errorx.NewCodeError(200, "添加成功", "")
}
