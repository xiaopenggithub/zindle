package logic

import (
	"context"
	"errors"

	"backend/service/activityorders/cmd/rpc/internal/svc"
	"backend/service/activityorders/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ActivityOrdersDelLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewActivityOrdersDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ActivityOrdersDelLogic {
	return &ActivityOrdersDelLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ActivityOrdersDelLogic) ActivityOrdersDel(in *pb.DelActivityOrdersReq) (*pb.DelActivityOrdersResp, error) {
	err := l.svcCtx.ActivityOrdersModel.Delete(l.ctx, in.Id)
	if err != nil {
		return &pb.DelActivityOrdersResp{}, errors.New("删除失败：" + err.Error())
	}
	return &pb.DelActivityOrdersResp{}, nil
}
