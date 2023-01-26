package activityOrders

import (
	"backend/common/errorx"
	"backend/service/activityorders/cmd/rpc/pb"
	"context"
	"fmt"

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
	_, err = l.svcCtx.ActivityOrderRPC.ActivityOrdersDelBatch(l.ctx, &pb.DelActivityOrdersBatchReq{
		Ids: req.Ids,
	})

	if err != nil {
		return nil, errorx.NewCodeError(201, fmt.Sprintf("%v", err), "")
	}

	return nil, errorx.NewCodeError(200, fmt.Sprintf("编号[%s]删除成功", req.Ids), "")
}
