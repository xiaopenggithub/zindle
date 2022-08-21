package logic

import (
	"backend/common/errorx"
	"backend/service/activities/cmd/api/internal/svc"
	"backend/service/activities/cmd/api/internal/types"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
)

type ActivityOrdersFindOneLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 活动预订信息 findone
func NewActivityOrdersFindOneLogic(ctx context.Context, svcCtx *svc.ServiceContext) ActivityOrdersFindOneLogic {
	return ActivityOrdersFindOneLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ActivityOrdersFindOneLogic) ActivityOrdersFindOne(req types.ActivityOrdersDelReq) (*types.ActivityOrdersReply, error) {
	one, err := l.svcCtx.ActivityOrderssModel.FindOne(req.Id)
	if err != nil {
		return nil, errorx.NewCodeError(201, fmt.Sprintf("%v", err), "")
	}
	data := make(map[string]interface{})
	data["item"] = one

	return nil, errorx.NewCodeError(200, fmt.Sprintf("%s", "获取成功"), data)
}
