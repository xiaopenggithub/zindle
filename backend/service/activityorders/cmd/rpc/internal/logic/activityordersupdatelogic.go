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

type ActivityOrdersUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewActivityOrdersUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ActivityOrdersUpdateLogic {
	return &ActivityOrdersUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ActivityOrdersUpdateLogic) ActivityOrdersUpdate(in *pb.UpdateActivityOrdersReq) (*pb.UpdateActivityOrdersResp, error) {
	var data model.ActivityOrders
	_ = copier.Copy(&data, in)
	err := l.svcCtx.ActivityOrdersModel.Update(l.ctx, &data)
	if err != nil {
		return &pb.UpdateActivityOrdersResp{}, errors.New("修改失败：" + err.Error())
	}
	return &pb.UpdateActivityOrdersResp{}, nil
}
