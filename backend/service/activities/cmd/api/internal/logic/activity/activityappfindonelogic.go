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

type ActivityAppFindOneLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewActivityAppFindOneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ActivityAppFindOneLogic {
	return &ActivityAppFindOneLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ActivityAppFindOneLogic) ActivityAppFindOne(req *types.ActivityDelReq) (resp *types.ActivityReply, err error) {
	res, err := l.svcCtx.ActivityRPC.GetActivitiesById(l.ctx, &pb.GetActivitiesByIdReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, errorx.NewCodeError(202, fmt.Sprintf("%v", err), "")
	}

	return nil, errorx.NewCodeError(200, fmt.Sprintf("%v", "ok"), res.Activities)
}
