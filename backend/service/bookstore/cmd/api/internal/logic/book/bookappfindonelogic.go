package logic

import (
	"backend/common/errorx"
	"backend/service/bookstore/cmd/api/internal/svc"
	"backend/service/bookstore/cmd/api/internal/types"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
)

type BookAppFindOneLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBookAppFindOneLogic(ctx context.Context, svcCtx *svc.ServiceContext) BookAppFindOneLogic {
	return BookAppFindOneLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BookAppFindOneLogic) BookAppFindOne(req types.BookDelReq) (*types.BookReply, error) {
	one, err := l.svcCtx.BooksModel.FindOne(req.Id)
	if err != nil {
		return nil, errorx.NewCodeError(201, fmt.Sprintf("%v", err), "")
	}
	data := make(map[string]interface{})
	one.Cover = "http://192.168.1.2:8080/uploads1/" + one.Cover
	data["item"] = one

	return nil, errorx.NewCodeError(200, fmt.Sprintf("%s", "获取成功"), data)
}
