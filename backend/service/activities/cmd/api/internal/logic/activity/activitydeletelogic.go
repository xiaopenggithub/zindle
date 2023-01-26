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

type ActivityDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewActivityDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ActivityDeleteLogic {
	return &ActivityDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ActivityDeleteLogic) ActivityDelete(req *types.ActivityDelReq) (resp *types.ActivityReply, err error) {
	_, err = l.svcCtx.ActivityRPC.DelActivities(l.ctx, &pb.DelActivitiesReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, errorx.NewCodeError(201, fmt.Sprintf("%v", err), "")
	}

	return nil, errorx.NewCodeError(200, fmt.Sprintf("编号[%d]删除成功", req.Id), "")
}
