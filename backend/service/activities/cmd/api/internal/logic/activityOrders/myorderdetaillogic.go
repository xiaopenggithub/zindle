package activityOrders

import (
	"context"

	"backend/service/activities/cmd/api/internal/svc"
	"backend/service/activities/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MyOrderDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMyOrderDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MyOrderDetailLogic {
	return &MyOrderDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MyOrderDetailLogic) MyOrderDetail(req *types.ActivityOrdersDelReq) (resp *types.ActivityOrdersReply, err error) {
	// todo: add your logic here and delete this line

	return
}
