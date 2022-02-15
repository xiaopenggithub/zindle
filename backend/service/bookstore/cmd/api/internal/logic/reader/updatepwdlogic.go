package logic

import (
	"backend/common/errorx"
	"backend/common/utils"
	"backend/service/bookstore/cmd/api/internal/svc"
	"backend/service/bookstore/cmd/api/internal/types"
	"context"
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"strings"
)

type UpdatePwdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdatePwdLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdatePwdLogic {
	return UpdatePwdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdatePwdLogic) UpdatePwd(req types.UpdatePwdReq) (*types.ReaderReply, error) {
	if len(strings.TrimSpace(req.Password)) == 0 || len(strings.TrimSpace(req.OldPassword)) == 0 {
		return nil, errorx.NewDefaultError("参数错误", "")
	}

	idNumber := json.Number(fmt.Sprintf("%v", l.ctx.Value("readerId")))
	id, err := idNumber.Int64()
	if err != nil {
		return nil, errorx.NewCodeError(401, "请重新登录再操作", "")
	}
	// 获取用户信息
	oldData, err := l.svcCtx.ReadersModel.FindOne(id)
	if err != nil {
		return nil, errorx.NewCodeError(201, "数据错误", "")
	}
	//验证原密码
	if oldData.Password != utils.MD5V(req.OldPassword) {
		return nil, errorx.NewCodeError(202, "原密码不正确", "")
	}
	// 执行修改
	oldData.Password = utils.MD5V(req.Password)
	fmt.Printf("%v\n", oldData)
	err = l.svcCtx.ReadersModel.Update(*oldData)
	if err != nil {
		return nil, errorx.NewCodeError(203, fmt.Sprintf("[%v]修改失败，请重试", err), "")
	}

	return nil, errorx.NewCodeError(200, "修改成功", "")
}
