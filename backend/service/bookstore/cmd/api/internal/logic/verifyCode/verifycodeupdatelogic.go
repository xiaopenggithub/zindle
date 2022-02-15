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

type VerifyCodeUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 验证码 update
func NewVerifyCodeUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) VerifyCodeUpdateLogic {
	return VerifyCodeUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VerifyCodeUpdateLogic) VerifyCodeUpdate(req types.VerifyCodePostReq) (*types.VerifyCodeReply, error) {
	oldData, err := l.svcCtx.VerifyCodesModel.FindOne(req.Id)
	if err != nil {
		return nil, errorx.NewCodeError(201, fmt.Sprintf("%v", err), "")
	}

	copier.Copy(&oldData, &req)
	err = l.svcCtx.VerifyCodesModel.Update(*oldData)

	if err != nil {
		return nil, errorx.NewCodeError(201, fmt.Sprintf("%v", err), "")
	}

	return nil, errorx.NewCodeError(200, fmt.Sprintf("%s", "修改成功"), "")
}
