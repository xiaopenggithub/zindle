package logic

import (
	"backend/common/errorx"
	"backend/service/bookstore/cmd/api/internal/svc"
	"backend/service/bookstore/cmd/api/internal/types"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
)

type CountsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCountsLogic(ctx context.Context, svcCtx *svc.ServiceContext) CountsLogic {
	return CountsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CountsLogic) Counts(req types.BookDelReq) (*types.BookReply, error) {
	totalBook, totalBorrow, totalShouldReturn, err := l.svcCtx.BooksModel.Counts(req.Id)
	if err != nil {
		return nil, errorx.NewCodeError(201, fmt.Sprintf("%v", err), "")
	}
	data := make(map[string]interface{})
	counts := make(map[string]int64)
	counts["totalBook"] = totalBook
	counts["totalBorrow"] = totalBorrow
	counts["totalShouldReturn"] = totalShouldReturn
	data["counts"] = counts

	return nil, errorx.NewCodeError(200, fmt.Sprintf("%s", "获取成功"), data)
}
