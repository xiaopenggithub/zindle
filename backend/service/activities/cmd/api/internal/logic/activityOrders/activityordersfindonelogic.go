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

type ActivityOrdersFindOneLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewActivityOrdersFindOneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ActivityOrdersFindOneLogic {
	return &ActivityOrdersFindOneLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ActivityOrdersFindOneLogic) ActivityOrdersFindOne(req *types.ActivityOrdersDelReq) (resp *types.ActivityOrdersReply, err error) {
	res, err := l.svcCtx.ActivityOrderRPC.ActivityOrdersGetById(l.ctx, &pb.GetActivityOrdersByIdReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, errorx.NewCodeError(202, fmt.Sprintf("%v", err), "")
	}

	return nil, errorx.NewCodeError(200, fmt.Sprintf("%v", "ok"), res.ActivityOrders)
}
