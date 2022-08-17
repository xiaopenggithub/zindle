package logic

import (
	"backend/common/errorx"
	"backend/service/bookstore/cmd/api/internal/svc"
	"backend/service/bookstore/cmd/api/internal/types"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"os"
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
	oldData, err := l.svcCtx.BooksModel.FindOne(req.Id)
	oldCover := oldData.Cover
	err = l.svcCtx.BooksModel.Delete(req.Id)
	if err != nil {
		return nil, errorx.NewCodeError(201, fmt.Sprintf("%v", err), "")
	}
	// 删除文件
	if _, err := os.Stat(oldCover); os.IsNotExist(err) {
		// nofile
		fmt.Println("没有找到文件", oldCover)
	} else {
		fmt.Println("准备删除文件", oldCover)
		if err := os.Remove(oldCover); err != nil {
			fmt.Println("删除文件error")
		}
	}
	return nil, errorx.NewCodeError(200, fmt.Sprintf("编号[%d]删除成功", req.Id), "")
}
