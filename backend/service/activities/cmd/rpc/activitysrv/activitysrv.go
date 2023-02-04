// Code generated by goctl. DO NOT EDIT!
// Source: activitie.proto

package activitysrv

import (
	"context"

	"backend/service/activities/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Activities            = pb.Activities
	AddActivitiesReq      = pb.AddActivitiesReq
	AddActivitiesResp     = pb.AddActivitiesResp
	DelActivitiesBatchReq = pb.DelActivitiesBatchReq
	DelActivitiesReq      = pb.DelActivitiesReq
	DelActivitiesResp     = pb.DelActivitiesResp
	GetActivitiesByIdReq  = pb.GetActivitiesByIdReq
	GetActivitiesByIdResp = pb.GetActivitiesByIdResp
	SearchActivitiesReq   = pb.SearchActivitiesReq
	SearchActivitiesResp  = pb.SearchActivitiesResp
	UpdateActivitiesReq   = pb.UpdateActivitiesReq
	UpdateActivitiesResp  = pb.UpdateActivitiesResp

	Activitysrv interface {
		// -----------------------活动信息表-----------------------
		AddActivities(ctx context.Context, in *AddActivitiesReq, opts ...grpc.CallOption) (*AddActivitiesResp, error)
		UpdateActivities(ctx context.Context, in *UpdateActivitiesReq, opts ...grpc.CallOption) (*UpdateActivitiesResp, error)
		DelActivities(ctx context.Context, in *DelActivitiesReq, opts ...grpc.CallOption) (*DelActivitiesResp, error)
		DelActivitiesBatch(ctx context.Context, in *DelActivitiesBatchReq, opts ...grpc.CallOption) (*DelActivitiesResp, error)
		GetActivitiesById(ctx context.Context, in *GetActivitiesByIdReq, opts ...grpc.CallOption) (*GetActivitiesByIdResp, error)
		SearchActivities(ctx context.Context, in *SearchActivitiesReq, opts ...grpc.CallOption) (*SearchActivitiesResp, error)
	}

	defaultActivitysrv struct {
		cli zrpc.Client
	}
)

func NewActivitysrv(cli zrpc.Client) Activitysrv {
	return &defaultActivitysrv{
		cli: cli,
	}
}

// -----------------------活动信息表-----------------------
func (m *defaultActivitysrv) AddActivities(ctx context.Context, in *AddActivitiesReq, opts ...grpc.CallOption) (*AddActivitiesResp, error) {
	client := pb.NewActivitysrvClient(m.cli.Conn())
	return client.AddActivities(ctx, in, opts...)
}

func (m *defaultActivitysrv) UpdateActivities(ctx context.Context, in *UpdateActivitiesReq, opts ...grpc.CallOption) (*UpdateActivitiesResp, error) {
	client := pb.NewActivitysrvClient(m.cli.Conn())
	return client.UpdateActivities(ctx, in, opts...)
}

func (m *defaultActivitysrv) DelActivities(ctx context.Context, in *DelActivitiesReq, opts ...grpc.CallOption) (*DelActivitiesResp, error) {
	client := pb.NewActivitysrvClient(m.cli.Conn())
	return client.DelActivities(ctx, in, opts...)
}

func (m *defaultActivitysrv) DelActivitiesBatch(ctx context.Context, in *DelActivitiesBatchReq, opts ...grpc.CallOption) (*DelActivitiesResp, error) {
	client := pb.NewActivitysrvClient(m.cli.Conn())
	return client.DelActivitiesBatch(ctx, in, opts...)
}

func (m *defaultActivitysrv) GetActivitiesById(ctx context.Context, in *GetActivitiesByIdReq, opts ...grpc.CallOption) (*GetActivitiesByIdResp, error) {
	client := pb.NewActivitysrvClient(m.cli.Conn())
	return client.GetActivitiesById(ctx, in, opts...)
}

func (m *defaultActivitysrv) SearchActivities(ctx context.Context, in *SearchActivitiesReq, opts ...grpc.CallOption) (*SearchActivitiesResp, error) {
	client := pb.NewActivitysrvClient(m.cli.Conn())
	return client.SearchActivities(ctx, in, opts...)
}