package logic

import (
	"context"
	"errors"
	"github.com/jinzhu/copier"

	"backend/service/activityorders/cmd/rpc/internal/svc"
	"backend/service/activityorders/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ActivityOrdersGetByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewActivityOrdersGetByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ActivityOrdersGetByIdLogic {
	return &ActivityOrdersGetByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ActivityOrdersGetByIdLogic) ActivityOrdersGetById(in *pb.GetActivityOrdersByIdReq) (*pb.GetActivityOrdersByIdResp, error) {
	one, err := l.svcCtx.ActivityOrdersModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return &pb.GetActivityOrdersByIdResp{}, errors.New("查询失败：" + err.Error())
	}
	var rpcActivities pb.ActivityOrders
	_ = copier.Copy(&rpcActivities, one)
	return &pb.GetActivityOrdersByIdResp{
		ActivityOrders: &rpcActivities,
	}, nil
}
