package logic

import (
	"backend/common/errorx"
	"backend/service/bookstore/cmd/api/internal/svc"
	"backend/service/bookstore/cmd/api/internal/types"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
)

type BookFindOneLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 图书管理 findone
func NewBookFindOneLogic(ctx context.Context, svcCtx *svc.ServiceContext) BookFindOneLogic {
	return BookFindOneLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BookFindOneLogic) BookFindOne(req types.BookDelReq) (*types.BookReply, error) {
	one, err := l.svcCtx.BooksModel.FindOne(req.Id)
	if err != nil {
		return nil, errorx.NewCodeError(201, fmt.Sprintf("%v", err), "")
	}
	data := make(map[string]interface{})
	data["item"] = one

	return nil, errorx.NewCodeError(200, fmt.Sprintf("%s", "获取成功"), data)
}
