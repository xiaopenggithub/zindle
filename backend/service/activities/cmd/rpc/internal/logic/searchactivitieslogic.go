package logic

import (
	"context"
	"errors"
	"github.com/jinzhu/copier"

	"backend/service/activities/cmd/rpc/internal/svc"
	"backend/service/activities/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchActivitiesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchActivitiesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchActivitiesLogic {
	return &SearchActivitiesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchActivitiesLogic) SearchActivities(in *pb.SearchActivitiesReq) (*pb.SearchActivitiesResp, error) {
	list, count, err := l.svcCtx.ActivitiesModel.List(l.ctx, in.Page, in.Limit, in.Title)
	if err != nil {
		return &pb.SearchActivitiesResp{}, errors.New("查询失败：" + err.Error())
	}
	var resp []*pb.Activities
	if len(list) > 0 {
		for _, item := range list {
			var tmp pb.Activities
			copier.Copy(&tmp, item)
			resp = append(resp, &tmp)
		}
	}
	return &pb.SearchActivitiesResp{
		Activities: resp,
		Total:      count,
	}, nil
}
