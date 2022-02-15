package logic

import (
	"backend/common/errorx"
	"backend/common/utils"
	"backend/service/bookstore/cmd/api/internal/svc"
	"backend/service/bookstore/cmd/api/internal/types"
	"backend/service/bookstore/model"
	"context"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) RegisterLogic {
	return RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req types.RegisterReq) (*types.ReaderReply, error) {
	//检查是否重名
	checkItem, err2 := l.svcCtx.ReadersModel.CheckDuplicate(req.Username, req.Email)
	if checkItem.Id > 0 && err2 == nil {
		return nil, errorx.NewCodeError(201, "用户名或邮箱已经存在", "")
	}

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

	var item model.Readers
	copier.Copy(&item, &req)
	item.LoginDate = "2006-01-02 15:04:05"
	item.Password = utils.MD5V(item.Password)

	_, err := l.svcCtx.ReadersModel.Insert(item)
	if err != nil {
		return nil, errorx.NewCodeError(202, fmt.Sprintf("%v", err), "")
	}

	return nil, errorx.NewCodeError(200, "添加成功", "")
}
