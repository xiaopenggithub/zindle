package logic

import (
	"backend/common/errorx"
	"backend/service/bookstore/cmd/api/internal/svc"
	"backend/service/bookstore/cmd/api/internal/types"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
)

type VerifyCodeDeleteBatchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 验证码 deletebatch
func NewVerifyCodeDeleteBatchLogic(ctx context.Context, svcCtx *svc.ServiceContext) VerifyCodeDeleteBatchLogic {
	return VerifyCodeDeleteBatchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VerifyCodeDeleteBatchLogic) VerifyCodeDeleteBatch(req types.VerifyCodeDelBatchReq) (*types.VerifyCodeReply, error) {
	err := l.svcCtx.VerifyCodesModel.DeleteBatch(req.Ids)
	if err != nil {
		return nil, errorx.NewCodeError(201, fmt.Sprintf("%v", err), "")
	}

	return nil, errorx.NewCodeError(200, fmt.Sprintf("编号[%s]删除成功", req.Ids), "")
}
