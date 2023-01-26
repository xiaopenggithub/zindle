package logic

import (
	"context"
	"errors"

	"backend/service/activityorders/cmd/rpc/internal/svc"
	"backend/service/activityorders/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ActivityOrdersDelBatchLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewActivityOrdersDelBatchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ActivityOrdersDelBatchLogic {
	return &ActivityOrdersDelBatchLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ActivityOrdersDelBatchLogic) ActivityOrdersDelBatch(in *pb.DelActivityOrdersBatchReq) (*pb.DelActivityOrdersResp, error) {
	err := l.svcCtx.ActivityOrdersModel.DeleteBatch(l.ctx, in.Ids)
	if err != nil {
		return &pb.DelActivityOrdersResp{}, errors.New("删除失败：" + err.Error())
	}
	return &pb.DelActivityOrdersResp{}, nil
}
