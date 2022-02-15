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
)

type ReaderAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 读者 create
func NewReaderAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) ReaderAddLogic {
	return ReaderAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReaderAddLogic) ReaderAdd(req types.ReaderPostReq) (*types.ReaderReply, error) {
	//检查是否重名
	checkItem, err2 := l.svcCtx.ReadersModel.CheckDuplicate(req.Username, req.Email)
	if checkItem.Id > 0 && err2 == nil {
		return nil, errorx.NewCodeError(201, "已经存在", "")
	}

	var item model.Readers
	copier.Copy(&item, &req)

	_, err := l.svcCtx.ReadersModel.Insert(item)
	if err != nil {
		return nil, errorx.NewCodeError(202, fmt.Sprintf("%v", err), "")
	}

	return nil, errorx.NewCodeError(200, "添加成功", "")
}
