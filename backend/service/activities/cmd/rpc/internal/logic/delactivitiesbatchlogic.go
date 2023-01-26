package logic

import (
	"context"
	"errors"

	"backend/service/activities/cmd/rpc/internal/svc"
	"backend/service/activities/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelActivitiesBatchLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelActivitiesBatchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelActivitiesBatchLogic {
	return &DelActivitiesBatchLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelActivitiesBatchLogic) DelActivitiesBatch(in *pb.DelActivitiesBatchReq) (*pb.DelActivitiesResp, error) {
	err := l.svcCtx.ActivitiesModel.DeleteBatch(l.ctx, in.Ids)
	if err != nil {
		return &pb.DelActivitiesResp{}, errors.New("删除失败：" + err.Error())
	}
	return &pb.DelActivitiesResp{}, nil
}
