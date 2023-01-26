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

type UpdateActivitiesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateActivitiesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateActivitiesLogic {
	return &UpdateActivitiesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateActivitiesLogic) UpdateActivities(in *pb.UpdateActivitiesReq) (*pb.UpdateActivitiesResp, error) {
	var data model.Activities
	_ = copier.Copy(&data, in)
	//old, err2 := l.svcCtx.ActivitiesModel.FindOne(l.ctx, in.Id)
	//fmt.Println("---adress---", in.Address)
	//old.Title = in.Title
	//old.Description = in.Description

	//if err2 != nil {
	//	return &pb.UpdateActivitiesResp{}, errors.New("修改错误：" + err2.Error())
	//}
	err := l.svcCtx.ActivitiesModel.Update(l.ctx, &data)
	if err != nil {
		return &pb.UpdateActivitiesResp{}, errors.New("修改失败：" + err.Error())
	}
	return &pb.UpdateActivitiesResp{}, nil
}
