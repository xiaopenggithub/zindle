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

type MyOrderActivityListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMyOrderActivityListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MyOrderActivityListLogic {
	return &MyOrderActivityListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MyOrderActivityListLogic) MyOrderActivityList(req *types.ActivityOrdersListReq) (resp *types.ActivityOrdersReply, err error) {
	res, err := l.svcCtx.ActivityOrderRPC.ActivityOrdersSearch(l.ctx, &pb.SearchActivityOrdersReq{
		Page:  req.Page,
		Limit: req.PageSize,
	})
	if err != nil {
		return nil, errorx.NewCodeError(202, fmt.Sprintf("%v", err), "")
	}
	data := make(map[string]interface{})
	data["total"] = res.Total
	data["list"] = res.ActivityOrders

	return nil, errorx.NewCodeError(200, "ok", data)
}
