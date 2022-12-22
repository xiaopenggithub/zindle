package logic

import (
	"backend/common/errorx"
	"context"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/copier"

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
	// 检查自己订单
	idNumber := json.Number(fmt.Sprintf("%v", l.ctx.Value("readerId")))
	id, err := idNumber.Int64()
	if err != nil {
		return nil, errorx.NewCodeError(401, "请重新登录再操作", "")
	}
	oldData, err := l.svcCtx.ActivityOrderssModel.FindOneOrder(req.Id, id)
	if err != nil {
		return nil, errorx.NewCodeError(201, fmt.Sprintf("%v", err), "")
	}
	fmt.Printf("--OrdersUpdateLogic--oldData-%v", oldData)
	fmt.Printf("--OrdersUpdateLogic--req-%v", req)
	oldData.MemberId = id

	copier.Copy(&oldData, &req)
	err = l.svcCtx.ActivityOrderssModel.Update(*oldData)

	if err != nil {
		return nil, errorx.NewCodeError(201, fmt.Sprintf("%v", err), "")
	}

	return nil, errorx.NewCodeError(200, fmt.Sprintf("%s", "修改成功"), "")
}
