package logic

import (
	"backend/common/errorx"
	"backend/service/bookstore/cmd/api/internal/svc"
	"backend/service/bookstore/cmd/api/internal/types"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
)

type VerifyCodeFindOneLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 验证码 findone
func NewVerifyCodeFindOneLogic(ctx context.Context, svcCtx *svc.ServiceContext) VerifyCodeFindOneLogic {
	return VerifyCodeFindOneLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VerifyCodeFindOneLogic) VerifyCodeFindOne(req types.VerifyCodeDelReq) (*types.VerifyCodeReply, error) {
	one, err := l.svcCtx.VerifyCodesModel.FindOne(req.Id)
	if err != nil {
		return nil, errorx.NewCodeError(201, fmt.Sprintf("%v", err), "")
	}
	data := make(map[string]interface{})
	data["item"] = one

	return nil, errorx.NewCodeError(200, fmt.Sprintf("%s", "获取成功"), data)
}
