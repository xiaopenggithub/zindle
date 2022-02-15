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

type VerifyCodeAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 验证码 create
func NewVerifyCodeAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) VerifyCodeAddLogic {
	return VerifyCodeAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VerifyCodeAddLogic) VerifyCodeAdd(req types.VerifyCodePostReq) (*types.VerifyCodeReply, error) {
	var item model.VerifyCodes
	copier.Copy(&item, &req)

	_, err := l.svcCtx.VerifyCodesModel.Insert(item)
	if err != nil {
		return nil, errorx.NewCodeError(202, fmt.Sprintf("%v", err), "")
	}

	return nil, errorx.NewCodeError(200, "添加成功", "")
}
