package logic

import (
	"backend/common/errorx"
	"backend/service/bookstore/cmd/api/internal/svc"
	"backend/service/bookstore/cmd/api/internal/types"
	"context"
	"encoding/json"
	"fmt"
	"github.com/tal-tech/go-zero/core/logx"
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
