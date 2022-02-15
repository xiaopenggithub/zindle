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
	"mime/multipart"
)

type BookAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 图书管理 create
func NewBookAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) BookAddLogic {
	return BookAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BookAddLogic) BookAdd(req types.BookPostReq, file *multipart.FileHeader) (*types.BookReply, error) {
	//检查是否重名
	checkItem, err2 := l.svcCtx.BooksModel.CheckDuplicate(req.Title)
	if checkItem.Id > 0 && err2 == nil {
		return nil, errorx.NewCodeError(201, "已经存在", "")
	}

	var item model.Books
	req.Cover = uploadfile(file)
	copier.Copy(&item, &req)

	_, err := l.svcCtx.BooksModel.Insert(item)
	if err != nil {
		return nil, errorx.NewCodeError(202, fmt.Sprintf("%v", err), "")
	}

	return nil, errorx.NewCodeError(200, "添加成功", "")
}
