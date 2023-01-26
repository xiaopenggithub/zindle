package logic

import (
	"context"
	"errors"

	"backend/service/activities/cmd/rpc/internal/svc"
	"backend/service/activities/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelActivitiesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelActivitiesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelActivitiesLogic {
	return &DelActivitiesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelActivitiesLogic) DelActivities(in *pb.DelActivitiesReq) (*pb.DelActivitiesResp, error) {
	err := l.svcCtx.ActivitiesModel.Delete(l.ctx, in.Id)
	if err != nil {
		return &pb.DelActivitiesResp{}, errors.New("删除失败：" + err.Error())
	}
	return &pb.DelActivitiesResp{}, nil
}
