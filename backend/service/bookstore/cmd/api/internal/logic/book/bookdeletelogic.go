package logic

import (
	"backend/common/errorx"
	"backend/service/bookstore/cmd/api/internal/svc"
	"backend/service/bookstore/cmd/api/internal/types"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
)

type BookDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 图书管理 delete
func NewBookDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) BookDeleteLogic {
	return BookDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BookDeleteLogic) BookDelete(req types.BookDelReq) (*types.BookReply, error) {
	err := l.svcCtx.BooksModel.Delete(req.Id)
	if err != nil {
		return nil, errorx.NewCodeError(201, fmt.Sprintf("%v", err), "")
	}

	return nil, errorx.NewCodeError(200, fmt.Sprintf("编号[%d]删除成功", req.Id), "")
}
