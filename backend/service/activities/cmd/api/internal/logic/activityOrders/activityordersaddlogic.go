package activityOrders

import (
	"context"

	"backend/service/activities/cmd/api/internal/svc"
	"backend/service/activities/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ActivityOrdersAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewActivityOrdersAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ActivityOrdersAddLogic {
	return &ActivityOrdersAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ActivityOrdersAddLogic) ActivityOrdersAdd(req *types.ActivityOrdersPostReq) (resp *types.ActivityOrdersReply, err error) {
	// todo: add your logic here and delete this line

	return
}
