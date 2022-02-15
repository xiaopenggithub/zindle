package logic

import (
	"backend/common/errorx"
	"backend/service/bookstore/cmd/api/internal/svc"
	"backend/service/bookstore/cmd/api/internal/types"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
)

type ReaderDeleteBatchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 读者 deletebatch
func NewReaderDeleteBatchLogic(ctx context.Context, svcCtx *svc.ServiceContext) ReaderDeleteBatchLogic {
	return ReaderDeleteBatchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReaderDeleteBatchLogic) ReaderDeleteBatch(req types.ReaderDelBatchReq) (*types.ReaderReply, error) {
	err := l.svcCtx.ReadersModel.DeleteBatch(req.Ids)
	if err != nil {
		return nil, errorx.NewCodeError(201, fmt.Sprintf("%v", err), "")
	}

	return nil, errorx.NewCodeError(200, fmt.Sprintf("编号[%s]删除成功", req.Ids), "")
}
