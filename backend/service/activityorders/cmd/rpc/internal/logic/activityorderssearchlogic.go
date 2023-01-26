package logic

import (
	"context"
	"errors"
	"github.com/jinzhu/copier"

	"backend/service/activityorders/cmd/rpc/internal/svc"
	"backend/service/activityorders/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ActivityOrdersSearchLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewActivityOrdersSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ActivityOrdersSearchLogic {
	return &ActivityOrdersSearchLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ActivityOrdersSearchLogic) ActivityOrdersSearch(in *pb.SearchActivityOrdersReq) (*pb.SearchActivityOrdersResp, error) {
	list, count, err := l.svcCtx.ActivityOrdersModel.List(l.ctx, in.Page, in.Limit, "")
	if err != nil {
		return &pb.SearchActivityOrdersResp{}, errors.New("查询失败：" + err.Error())
	}
	var resp []*pb.ActivityOrders
	if len(list) > 0 {
		for _, item := range list {
			var tmp pb.ActivityOrders
			copier.Copy(&tmp, item)
			resp = append(resp, &tmp)
		}
	}
	return &pb.SearchActivityOrdersResp{
		ActivityOrders: resp,
		Total:          count,
	}, nil
}
