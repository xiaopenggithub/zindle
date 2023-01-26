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

type ActivityListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewActivityListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ActivityListLogic {
	return &ActivityListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ActivityListLogic) ActivityList(req *types.ActivityListReq) (resp *types.ActivityReply, err error) {
	res, err := l.svcCtx.ActivityRPC.SearchActivities(l.ctx, &pb.SearchActivitiesReq{
		Title: req.Keyword,
		Page:  req.Page,
		Limit: req.PageSize,
	})
	if err != nil {
		return nil, errorx.NewCodeError(202, fmt.Sprintf("%v", err), "")
	}
	data := make(map[string]interface{})
	data["total"] = res.Total
	data["list"] = res.Activities

	return nil, errorx.NewCodeError(200, "ok", data)
}
