package logic

import (
	"backend/common/errorx"
	"backend/service/bookstore/cmd/api/internal/svc"
	"backend/service/bookstore/cmd/api/internal/types"
	"context"
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
)

type BorrowBookLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBorrowBookLogic(ctx context.Context, svcCtx *svc.ServiceContext) BorrowBookLogic {
	return BorrowBookLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BorrowBookLogic) BorrowBook(req types.BookOrderDelReq) (*types.BookOrderReply, error) {
	//1.生成用户订单
	//2.减少库存
	//事务处理
	idNumber := json.Number(fmt.Sprintf("%v", l.ctx.Value("readerId")))
	id, err := idNumber.Int64()
	if err != nil {
		return nil, errorx.NewCodeError(401, "请重新登录再操作", "")
	}
	//再次检查是否本人已借
	_, err = l.svcCtx.BookOrdersModel.CheckBorrow(req.Id, id)
	if err != nil {
		// todo 处理并发借书情况
		err = l.svcCtx.BookOrdersModel.BorrowBook(req.Id, id)
		if err != nil {
			return nil, errorx.NewCodeError(201, fmt.Sprintf("%v", err), "")
		}
		return nil, errorx.NewCodeError(200, fmt.Sprintf("%v", "借书成功"), "")
	}
	return nil, errorx.NewCodeError(201, fmt.Sprintf("%s", "些书已借待归还"), "")
}
