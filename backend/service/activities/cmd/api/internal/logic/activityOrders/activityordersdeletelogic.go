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

type ActivityOrdersDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewActivityOrdersDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ActivityOrdersDeleteLogic {
	return &ActivityOrdersDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ActivityOrdersDeleteLogic) ActivityOrdersDelete(req *types.ActivityOrdersDelReq) (resp *types.ActivityOrdersReply, err error) {
	_, err = l.svcCtx.ActivityOrderRPC.ActivityOrdersDel(l.ctx, &pb.DelActivityOrdersReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, errorx.NewCodeError(201, fmt.Sprintf("%v", err), "")
	}

	return nil, errorx.NewCodeError(200, fmt.Sprintf("编号[%d]删除成功", req.Id), "")
}
