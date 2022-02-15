package logic

import (
	"backend/common/errorx"
	"backend/service/bookstore/cmd/api/internal/svc"
	"backend/service/bookstore/cmd/api/internal/types"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
)

type BookOrderFindOneLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 图书订单 findone
func NewBookOrderFindOneLogic(ctx context.Context, svcCtx *svc.ServiceContext) BookOrderFindOneLogic {
	return BookOrderFindOneLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BookOrderFindOneLogic) BookOrderFindOne(req types.BookOrderDelReq) (*types.BookOrderReply, error) {
	one, err := l.svcCtx.BookOrdersModel.FindOne(req.Id)
	if err != nil {
		return nil, errorx.NewCodeError(201, fmt.Sprintf("%v", err), "")
	}
	data := make(map[string]interface{})
	data["item"] = one

	return nil, errorx.NewCodeError(200, fmt.Sprintf("%s", "获取成功"), data)
}
