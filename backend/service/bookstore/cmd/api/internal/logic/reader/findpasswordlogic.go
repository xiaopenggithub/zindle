package logic

import (
	"backend/common/errorx"
	"backend/common/utils"
	"backend/service/bookstore/cmd/api/internal/svc"
	"backend/service/bookstore/cmd/api/internal/types"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

type FindPasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) FindPasswordLogic {
	return FindPasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindPasswordLogic) FindPassword(req types.FindPasswordReq) (*types.ReaderReply, error) {
	// 验证码检查
	code := req.VerificationCode
	if len(code) < 4 {
		return nil, errorx.NewCodeError(202, "请输入正确的验证码", "")
	}
	// 5分钟有效
	availableTime := time.Now().Add(-time.Minute * 5).Format("2006-01-02 15:04:05")
	checkCode, err3 := l.svcCtx.VerifyCodesModel.CheckCode(req.Email, req.VerificationCode, availableTime)
	if err3 != nil {
		// sql: no rows in result set
		return nil, errorx.NewCodeError(203, fmt.Sprintf("%v", "验证码:错误/失效"), "")
	}
	// 验证正确
	checkCode.Status = 1
	l.svcCtx.VerifyCodesModel.Update(checkCode)
	// 获取用户信息
	reader, err3 := l.svcCtx.ReadersModel.FindOneByUserEmail(req.Email)
	if err3 != nil {
		return nil, errorx.NewCodeError(204, "信息有误,请核查", req)
	}
	reader.Password = utils.MD5V(req.Password)
	err4 := l.svcCtx.ReadersModel.Update(reader)
	if err4 != nil {
		return nil, errorx.NewCodeError(205, "操作失败,请重试", req)
	}
	return nil, errorx.NewCodeError(200, "操作成功", req)
}
