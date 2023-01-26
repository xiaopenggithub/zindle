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

type ActivityFindOneLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewActivityFindOneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ActivityFindOneLogic {
	return &ActivityFindOneLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ActivityFindOneLogic) ActivityFindOne(req *types.ActivityDelReq) (resp *types.ActivityReply, err error) {
	res, err := l.svcCtx.ActivityRPC.GetActivitiesById(l.ctx, &pb.GetActivitiesByIdReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, errorx.NewCodeError(202, fmt.Sprintf("%v", err), "")
	}

	return nil, errorx.NewCodeError(200, fmt.Sprintf("%v", "ok"), res.Activities)
}
