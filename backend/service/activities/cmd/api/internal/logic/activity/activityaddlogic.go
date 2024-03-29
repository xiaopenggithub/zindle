package activity

import (
	"backend/common/errorx"
	"backend/service/activities/cmd/rpc/pb"
	"context"
	"crypto/md5"
	"fmt"
	"github.com/jinzhu/copier"

	"backend/service/activities/cmd/api/internal/svc"
	"backend/service/activities/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ActivityAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewActivityAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ActivityAddLogic {
	return &ActivityAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ActivityAddLogic) ActivityAdd(req *types.ActivityPostReq) (resp *types.ActivityReply, err error) {
	var addActivitiesReq pb.AddActivitiesReq
	_ = copier.Copy(&addActivitiesReq, req)

	result, err := l.svcCtx.ActivityRPC.AddActivities(l.ctx, &addActivitiesReq)
	if err != nil {
		return nil, errorx.NewCodeError(202, fmt.Sprintf("%v", err), "")
	}
	data := make(map[string]interface{})
	data["id"] = result.Id

	return nil, errorx.NewCodeError(200, fmt.Sprintf("%v", "添加成功"), data)
}

var Md5PasswordSalt = "20240216"

func hasPassword(password string) string {
	hasedBytes := md5.Sum([]byte(password + Md5PasswordSalt))
	return fmt.Sprintf("%x", hasedBytes)
}
