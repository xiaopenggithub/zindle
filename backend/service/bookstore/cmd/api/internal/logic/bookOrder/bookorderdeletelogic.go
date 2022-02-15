package logic

import (
	"backend/common/errorx"
	"backend/service/bookstore/cmd/api/internal/svc"
	"backend/service/bookstore/cmd/api/internal/types"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
)

type BookOrderDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 图书订单 delete
func NewBookOrderDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) BookOrderDeleteLogic {
	return BookOrderDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BookOrderDeleteLogic) BookOrderDelete(req types.BookOrderDelReq) (*types.BookOrderReply, error) {
	err := l.svcCtx.BookOrdersModel.Delete(req.Id)
	if err != nil {
		return nil, errorx.NewCodeError(201, fmt.Sprintf("%v", err), "")
	}

	return nil, errorx.NewCodeError(200, fmt.Sprintf("编号[%d]删除成功", req.Id), "")
}
