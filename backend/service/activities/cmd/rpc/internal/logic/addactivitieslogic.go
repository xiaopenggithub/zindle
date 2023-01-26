package logic

import (
	"backend/service/activities/cmd/rpc/internal/model"
	"backend/service/activities/cmd/rpc/internal/svc"
	"backend/service/activities/cmd/rpc/pb"
	"context"
	"errors"
	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddActivitiesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddActivitiesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddActivitiesLogic {
	return &AddActivitiesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------活动信息表-----------------------
func (l *AddActivitiesLogic) AddActivities(in *pb.AddActivitiesReq) (*pb.AddActivitiesResp, error) {
	var data model.Activities
	_ = copier.Copy(&data, in)
	result, err := l.svcCtx.ActivitiesModel.Insert(l.ctx, &data)
	if err != nil {
		return &pb.AddActivitiesResp{}, errors.New("添加失败:" + err.Error())
	}
	id, _ := result.LastInsertId()
	return &pb.AddActivitiesResp{
		Id: id,
	}, nil
}
