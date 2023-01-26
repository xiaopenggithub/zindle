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
	var addActivityOrdersReq pb.AddActivityOrdersReq
	_ = copier.Copy(&addActivityOrdersReq, req)
	result, err := l.svcCtx.ActivityOrderRPC.ActivityOrdersAdd(l.ctx, &addActivityOrdersReq)
	if err != nil {
		return nil, errorx.NewCodeError(202, fmt.Sprintf("%v", err), "")
	}
	data := make(map[string]interface{})
	data["id"] = result.Id

	return nil, errorx.NewCodeError(200, fmt.Sprintf("%v", "添加成功"), data)
}
