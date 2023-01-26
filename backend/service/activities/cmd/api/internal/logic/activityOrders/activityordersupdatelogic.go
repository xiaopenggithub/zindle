package activityOrders

import (
	"context"

	"backend/service/activities/cmd/api/internal/svc"
	"backend/service/activities/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ActivityOrdersUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewActivityOrdersUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ActivityOrdersUpdateLogic {
	return &ActivityOrdersUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ActivityOrdersUpdateLogic) ActivityOrdersUpdate(req *types.ActivityOrdersPostReq) (resp *types.ActivityOrdersReply, err error) {
	// todo: add your logic here and delete this line

	return
}
