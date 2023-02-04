// Code generated by goctl. DO NOT EDIT!
// Source: activitieOrder.proto

package activitieordersvc

import (
	"context"

	"backend/service/activityorders/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	ActivityOrders            = pb.ActivityOrders
	AddActivityOrdersReq      = pb.AddActivityOrdersReq
	AddActivityOrdersResp     = pb.AddActivityOrdersResp
	DelActivityOrdersBatchReq = pb.DelActivityOrdersBatchReq
	DelActivityOrdersReq      = pb.DelActivityOrdersReq
	DelActivityOrdersResp     = pb.DelActivityOrdersResp
	GetActivityOrdersByIdReq  = pb.GetActivityOrdersByIdReq
	GetActivityOrdersByIdResp = pb.GetActivityOrdersByIdResp
	SearchActivityOrdersReq   = pb.SearchActivityOrdersReq
	SearchActivityOrdersResp  = pb.SearchActivityOrdersResp
	UpdateActivityOrdersReq   = pb.UpdateActivityOrdersReq
	UpdateActivityOrdersResp  = pb.UpdateActivityOrdersResp

	ActivitieOrdersvc interface {
		// -----------------------活动订单表-----------------------
		ActivityOrdersAdd(ctx context.Context, in *AddActivityOrdersReq, opts ...grpc.CallOption) (*AddActivityOrdersResp, error)
		ActivityOrdersUpdate(ctx context.Context, in *UpdateActivityOrdersReq, opts ...grpc.CallOption) (*UpdateActivityOrdersResp, error)
		ActivityOrdersDel(ctx context.Context, in *DelActivityOrdersReq, opts ...grpc.CallOption) (*DelActivityOrdersResp, error)
		ActivityOrdersDelBatch(ctx context.Context, in *DelActivityOrdersBatchReq, opts ...grpc.CallOption) (*DelActivityOrdersResp, error)
		ActivityOrdersGetById(ctx context.Context, in *GetActivityOrdersByIdReq, opts ...grpc.CallOption) (*GetActivityOrdersByIdResp, error)
		ActivityOrdersSearch(ctx context.Context, in *SearchActivityOrdersReq, opts ...grpc.CallOption) (*SearchActivityOrdersResp, error)
	}

	defaultActivitieOrdersvc struct {
		cli zrpc.Client
	}
)

func NewActivitieOrdersvc(cli zrpc.Client) ActivitieOrdersvc {
	return &defaultActivitieOrdersvc{
		cli: cli,
	}
}

// -----------------------活动订单表-----------------------
func (m *defaultActivitieOrdersvc) ActivityOrdersAdd(ctx context.Context, in *AddActivityOrdersReq, opts ...grpc.CallOption) (*AddActivityOrdersResp, error) {
	client := pb.NewActivitieOrdersvcClient(m.cli.Conn())
	return client.ActivityOrdersAdd(ctx, in, opts...)
}

func (m *defaultActivitieOrdersvc) ActivityOrdersUpdate(ctx context.Context, in *UpdateActivityOrdersReq, opts ...grpc.CallOption) (*UpdateActivityOrdersResp, error) {
	client := pb.NewActivitieOrdersvcClient(m.cli.Conn())
	return client.ActivityOrdersUpdate(ctx, in, opts...)
}

func (m *defaultActivitieOrdersvc) ActivityOrdersDel(ctx context.Context, in *DelActivityOrdersReq, opts ...grpc.CallOption) (*DelActivityOrdersResp, error) {
	client := pb.NewActivitieOrdersvcClient(m.cli.Conn())
	return client.ActivityOrdersDel(ctx, in, opts...)
}

func (m *defaultActivitieOrdersvc) ActivityOrdersDelBatch(ctx context.Context, in *DelActivityOrdersBatchReq, opts ...grpc.CallOption) (*DelActivityOrdersResp, error) {
	client := pb.NewActivitieOrdersvcClient(m.cli.Conn())
	return client.ActivityOrdersDelBatch(ctx, in, opts...)
}

func (m *defaultActivitieOrdersvc) ActivityOrdersGetById(ctx context.Context, in *GetActivityOrdersByIdReq, opts ...grpc.CallOption) (*GetActivityOrdersByIdResp, error) {
	client := pb.NewActivitieOrdersvcClient(m.cli.Conn())
	return client.ActivityOrdersGetById(ctx, in, opts...)
}

func (m *defaultActivitieOrdersvc) ActivityOrdersSearch(ctx context.Context, in *SearchActivityOrdersReq, opts ...grpc.CallOption) (*SearchActivityOrdersResp, error) {
	client := pb.NewActivitieOrdersvcClient(m.cli.Conn())
	return client.ActivityOrdersSearch(ctx, in, opts...)
}