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

type CheckBorrowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCheckBorrowLogic(ctx context.Context, svcCtx *svc.ServiceContext) CheckBorrowLogic {
	return CheckBorrowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckBorrowLogic) CheckBorrow(req types.BookOrderDelReq) (*types.BookOrderReply, error) {
	idNumber := json.Number(fmt.Sprintf("%v", l.ctx.Value("readerId")))
	id, err := idNumber.Int64()
	if err != nil {
		return nil, errorx.NewCodeError(401, "请重新登录再操作", "")
	}
	_, err = l.svcCtx.BookOrdersModel.CheckBorrow(req.Id, id)
	if err != nil {
		return nil, errorx.NewCodeError(200, fmt.Sprintf("%v", "获取成功"), true)
	}
	return nil, errorx.NewCodeError(200, fmt.Sprintf("%s", "获取成功"), false)
}
