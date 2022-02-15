package logic

import (
	"backend/common/errorx"
	"backend/service/bookstore/cmd/api/internal/svc"
	"backend/service/bookstore/cmd/api/internal/types"
	"context"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type ReaderUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 读者 update
func NewReaderUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) ReaderUpdateLogic {
	return ReaderUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReaderUpdateLogic) ReaderUpdate(req types.ReaderPostReq) (*types.ReaderReply, error) {
	oldData, err := l.svcCtx.ReadersModel.FindOne(req.Id)
	if err != nil {
		return nil, errorx.NewCodeError(201, fmt.Sprintf("%v", err), "")
	}

	copier.Copy(&oldData, &req)
	err = l.svcCtx.ReadersModel.Update(*oldData)

	if err != nil {
		return nil, errorx.NewCodeError(201, fmt.Sprintf("%v", err), "")
	}

	return nil, errorx.NewCodeError(200, fmt.Sprintf("%s", "修改成功"), "")
}
