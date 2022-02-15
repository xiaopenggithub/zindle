package logic

import (
	"backend/common/errorx"
	"backend/service/bookstore/cmd/api/internal/svc"
	"backend/service/bookstore/cmd/api/internal/types"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
)

type VerifyCodeDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 验证码 delete
func NewVerifyCodeDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) VerifyCodeDeleteLogic {
	return VerifyCodeDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VerifyCodeDeleteLogic) VerifyCodeDelete(req types.VerifyCodeDelReq) (*types.VerifyCodeReply, error) {
	err := l.svcCtx.VerifyCodesModel.Delete(req.Id)
	if err != nil {
		return nil, errorx.NewCodeError(201, fmt.Sprintf("%v", err), "")
	}

	return nil, errorx.NewCodeError(200, fmt.Sprintf("编号[%d]删除成功", req.Id), "")
}
