package activityOrders

import (
	"context"

	"backend/service/activities/cmd/api/internal/svc"
	"backend/service/activities/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrdersUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrdersUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrdersUpdateLogic {
	return &OrdersUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrdersUpdateLogic) OrdersUpdate(req *types.ActivityOrdersPostReq) (resp *types.ActivityOrdersReply, err error) {
	// todo: add your logic here and delete this line

	return
}
