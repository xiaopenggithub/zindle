package logic

import (
	"backend/common/errorx"
	"backend/service/backend/cmd/rpc/systemuserget/systemusergeter"
	"context"
	"encoding/json"
	"fmt"

	"backend/service/backend/cmd/api/internal/svc"
	"backend/service/backend/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetUserInfoLogic {
	return GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo() (*types.AdminReply, error) {
	idNumber := json.Number(fmt.Sprintf("%v", l.ctx.Value("userId")))
	id, err := idNumber.Int64()
	if err != nil {
		return nil, errorx.NewCodeError(401, "请重新登录再操作", "")
	}
	systemuserResp, err := l.svcCtx.Systemusergeter.GetSystemuser(l.ctx, &systemusergeter.Request{
		Id: id,
	})
	if err != nil {
		fmt.Println("rpc----error", err)
	}
	//通过rpc获取nickname
	user := make(map[string]interface{})
	user["nickName"] = systemuserResp.NickName
	//user["headerImg"] = "/uploads/" + systemuserResp.Avatar
	authority := make(map[string]interface{})
	// menurouter-name
	authority["defaultRouter"] = "dashboard"
	user["authority"] = authority

	//user["sideMode"] = "dark"
	user["sideMode"] = "light"
	user["activeColor"] = "#1890ff"
	user["baseColor"] = "#fff"
	return nil, errorx.NewCodeError(200, "成功", user)
}
