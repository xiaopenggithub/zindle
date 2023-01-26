package activityOrders

import (
	"context"

	"backend/service/activities/cmd/api/internal/svc"
	"backend/service/activities/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ActivityOrdersDeleteBatchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewActivityOrdersDeleteBatchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ActivityOrdersDeleteBatchLogic {
	return &ActivityOrdersDeleteBatchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ActivityOrdersDeleteBatchLogic) ActivityOrdersDeleteBatch(req *types.ActivityOrdersDelBatchReq) (resp *types.ActivityOrdersReply, err error) {
	// todo: add your logic here and delete this line

	return
}
