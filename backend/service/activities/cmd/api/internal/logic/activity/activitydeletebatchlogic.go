package activity

import (
	"backend/common/errorx"
	"backend/service/activities/cmd/rpc/pb"
	"context"
	"fmt"

	"backend/service/activities/cmd/api/internal/svc"
	"backend/service/activities/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ActivityDeleteBatchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewActivityDeleteBatchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ActivityDeleteBatchLogic {
	return &ActivityDeleteBatchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ActivityDeleteBatchLogic) ActivityDeleteBatch(req *types.ActivityDelBatchReq) (resp *types.ActivityReply, err error) {
	_, err = l.svcCtx.ActivityRPC.DelActivitiesBatch(l.ctx, &pb.DelActivitiesBatchReq{
		Ids: req.Ids,
	})
	
	if err != nil {
		return nil, errorx.NewCodeError(201, fmt.Sprintf("%v", err), "")
	}

	return nil, errorx.NewCodeError(200, fmt.Sprintf("编号[%s]删除成功", req.Ids), "")
}
