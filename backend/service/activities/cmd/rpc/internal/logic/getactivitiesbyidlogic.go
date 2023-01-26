package logic

import (
	"context"
	"errors"
	"github.com/jinzhu/copier"

	"backend/service/activities/cmd/rpc/internal/svc"
	"backend/service/activities/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetActivitiesByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetActivitiesByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetActivitiesByIdLogic {
	return &GetActivitiesByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetActivitiesByIdLogic) GetActivitiesById(in *pb.GetActivitiesByIdReq) (*pb.GetActivitiesByIdResp, error) {
	one, err := l.svcCtx.ActivitiesModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return &pb.GetActivitiesByIdResp{}, errors.New("查询失败：" + err.Error())
	}
	var rpcActivities pb.Activities
	_ = copier.Copy(&rpcActivities, one)
	return &pb.GetActivitiesByIdResp{
		Activities: &rpcActivities,
	}, nil
}
