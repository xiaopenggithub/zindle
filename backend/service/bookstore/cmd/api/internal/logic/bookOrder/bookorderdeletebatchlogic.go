package logic

import (
	"backend/common/errorx"
	"backend/service/bookstore/cmd/api/internal/svc"
	"backend/service/bookstore/cmd/api/internal/types"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
)

type BookOrderDeleteBatchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 图书订单 deletebatch
func NewBookOrderDeleteBatchLogic(ctx context.Context, svcCtx *svc.ServiceContext) BookOrderDeleteBatchLogic {
	return BookOrderDeleteBatchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BookOrderDeleteBatchLogic) BookOrderDeleteBatch(req types.BookOrderDelBatchReq) (*types.BookOrderReply, error) {
	err := l.svcCtx.BookOrdersModel.DeleteBatch(req.Ids)
	if err != nil {
		return nil, errorx.NewCodeError(201, fmt.Sprintf("%v", err), "")
	}

	return nil, errorx.NewCodeError(200, fmt.Sprintf("编号[%s]删除成功", req.Ids), "")
}
