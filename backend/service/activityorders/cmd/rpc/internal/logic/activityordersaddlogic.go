package logic

import (
	"backend/service/activityorders/cmd/rpc/internal/model"
	"context"
	"errors"
	"github.com/jinzhu/copier"

	"backend/service/activityorders/cmd/rpc/internal/svc"
	"backend/service/activityorders/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ActivityOrdersAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewActivityOrdersAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ActivityOrdersAddLogic {
	return &ActivityOrdersAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------活动订单表-----------------------
func (l *ActivityOrdersAddLogic) ActivityOrdersAdd(in *pb.AddActivityOrdersReq) (*pb.AddActivityOrdersResp, error) {
	var data model.ActivityOrders
	_ = copier.Copy(&data, in)
	result, err := l.svcCtx.ActivityOrdersModel.Insert(l.ctx, &data)
	if err != nil {
		return &pb.AddActivityOrdersResp{}, errors.New("添加失败:" + err.Error())
	}
	id, err := result.LastInsertId()
	return &pb.AddActivityOrdersResp{
		Id: id,
	}, nil
}
