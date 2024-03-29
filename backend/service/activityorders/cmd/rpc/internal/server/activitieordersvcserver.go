// Code generated by goctl. DO NOT EDIT!
// Source: activitieOrder.proto

package server

import (
	"context"

	"backend/service/activityorders/cmd/rpc/internal/logic"
	"backend/service/activityorders/cmd/rpc/internal/svc"
	"backend/service/activityorders/cmd/rpc/pb"
)

type ActivitieOrdersvcServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedActivitieOrdersvcServer
}

func NewActivitieOrdersvcServer(svcCtx *svc.ServiceContext) *ActivitieOrdersvcServer {
	return &ActivitieOrdersvcServer{
		svcCtx: svcCtx,
	}
}

// -----------------------活动订单表-----------------------
func (s *ActivitieOrdersvcServer) ActivityOrdersAdd(ctx context.Context, in *pb.AddActivityOrdersReq) (*pb.AddActivityOrdersResp, error) {
	l := logic.NewActivityOrdersAddLogic(ctx, s.svcCtx)
	return l.ActivityOrdersAdd(in)
}

func (s *ActivitieOrdersvcServer) ActivityOrdersUpdate(ctx context.Context, in *pb.UpdateActivityOrdersReq) (*pb.UpdateActivityOrdersResp, error) {
	l := logic.NewActivityOrdersUpdateLogic(ctx, s.svcCtx)
	return l.ActivityOrdersUpdate(in)
}

func (s *ActivitieOrdersvcServer) ActivityOrdersDel(ctx context.Context, in *pb.DelActivityOrdersReq) (*pb.DelActivityOrdersResp, error) {
	l := logic.NewActivityOrdersDelLogic(ctx, s.svcCtx)
	return l.ActivityOrdersDel(in)
}

func (s *ActivitieOrdersvcServer) ActivityOrdersDelBatch(ctx context.Context, in *pb.DelActivityOrdersBatchReq) (*pb.DelActivityOrdersResp, error) {
	l := logic.NewActivityOrdersDelBatchLogic(ctx, s.svcCtx)
	return l.ActivityOrdersDelBatch(in)
}

func (s *ActivitieOrdersvcServer) ActivityOrdersGetById(ctx context.Context, in *pb.GetActivityOrdersByIdReq) (*pb.GetActivityOrdersByIdResp, error) {
	l := logic.NewActivityOrdersGetByIdLogic(ctx, s.svcCtx)
	return l.ActivityOrdersGetById(in)
}

func (s *ActivitieOrdersvcServer) ActivityOrdersSearch(ctx context.Context, in *pb.SearchActivityOrdersReq) (*pb.SearchActivityOrdersResp, error) {
	l := logic.NewActivityOrdersSearchLogic(ctx, s.svcCtx)
	return l.ActivityOrdersSearch(in)
}
