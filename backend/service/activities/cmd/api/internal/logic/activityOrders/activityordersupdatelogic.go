package activityOrders

import (
	"backend/common/errorx"
	"backend/service/activityorders/cmd/rpc/pb"
	"context"
	"fmt"
	"github.com/jinzhu/copier"

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
	var newdata pb.UpdateActivityOrdersReq
	_ = copier.Copy(&newdata, req)
	_, err = l.svcCtx.ActivityOrderRPC.ActivityOrdersUpdate(l.ctx, &newdata)
	if err != nil {
		return nil, errorx.NewCodeError(201, fmt.Sprintf("%v", err), "")
	}

	return nil, errorx.NewCodeError(200, fmt.Sprintf("%s", "修改成功"), "")
}
