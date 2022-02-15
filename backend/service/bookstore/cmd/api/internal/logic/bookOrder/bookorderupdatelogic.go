package logic

import (
	"backend/common/errorx"
	"backend/service/bookstore/cmd/api/internal/svc"
	"backend/service/bookstore/cmd/api/internal/types"
	"context"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type BookOrderUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 图书订单 update
func NewBookOrderUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) BookOrderUpdateLogic {
	return BookOrderUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BookOrderUpdateLogic) BookOrderUpdate(req types.BookOrderPostReq) (*types.BookOrderReply, error) {
	oldData, err := l.svcCtx.BookOrdersModel.FindOne(req.Id)
	if err != nil {
		return nil, errorx.NewCodeError(201, fmt.Sprintf("%v", err), "")
	}

	copier.Copy(&oldData, &req)
	err = l.svcCtx.BookOrdersModel.Update(*oldData)

	if err != nil {
		return nil, errorx.NewCodeError(201, fmt.Sprintf("%v", err), "")
	}

	return nil, errorx.NewCodeError(200, fmt.Sprintf("%s", "修改成功"), "")
}
