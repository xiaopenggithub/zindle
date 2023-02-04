package activityOrders

import (
	"backend/common/errorx"
	"backend/service/activities/cmd/api/internal/svc"
	"backend/service/activities/cmd/api/internal/types"
	pb2 "backend/service/activities/cmd/rpc/pb"
	"backend/service/activityorders/cmd/rpc/pb"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/mr"

	"github.com/zeromicro/go-zero/core/logx"
)

type ActivityOrdersListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewActivityOrdersListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ActivityOrdersListLogic {
	return &ActivityOrdersListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ActivityOrdersListLogic) ActivityOrdersList(req *types.ActivityOrdersListReq) (resp *types.ActivityOrdersReply, err error) {
	res, err := l.svcCtx.ActivityOrderRPC.ActivityOrdersSearch(l.ctx, &pb.SearchActivityOrdersReq{
		Page:  req.Page,
		Limit: req.PageSize,
	})
	if err != nil {
		return nil, errorx.NewCodeError(202, fmt.Sprintf("%v", err), "")
	}
	data := make(map[string]interface{})
	data["total"] = res.Total
	data["list"] = res.ActivityOrders
	//resData := res.ActivityOrders

	if res.Total > 0 {
		ps, err := mr.MapReduce(func(source chan<- interface{}) {
			for _, item := range res.ActivityOrders {
				//fmt.Println(item.ActivityId, item.MemberId)
				//source <- [3]int64{item.ActivityId, item.MemberId,item.Id}
				tmp := make(map[string]interface{})
				tmp["id"] = item.Id
				tmp["member_id"] = item.MemberId // todo:根据用户id调用rpc获取用户信息
				tmp["activity_id"] = item.ActivityId
				tmp["status"] = item.Status
				tmp["seat_number"] = item.SeatNumber
				source <- tmp
			}
		}, func(item interface{}, writer mr.Writer, cancel func(error)) {
			//ids := item.([3]int64)
			obj := item.(map[string]interface{})
			activity, err2 := l.svcCtx.ActivityRPC.GetActivitiesById(l.ctx, &pb2.GetActivitiesByIdReq{
				Id: obj["activity_id"].(int64),
			})
			if err2 != nil {
				cancel(err2)
				return
			}
			// todo:根据用户id调用rpc获取用户信息
			obj["activity"] = activity.Activities
			delete(obj, "activity_id")

			writer.Write(obj)
		}, func(pipe <-chan interface{}, writer mr.Writer, cancel func(error)) {
			var r []map[string]interface{}
			for p := range pipe {
				r = append(r, p.(map[string]interface{}))
			}
			writer.Write(r)
		})
		if err != nil {
			return nil, errorx.NewCodeError(204, fmt.Sprintf("%v", err), "")
		}
		data["list"] = ps
	}

	return nil, errorx.NewCodeError(200, "ok", data)
}
