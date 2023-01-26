package activity

import (
	"backend/common/errorx"
	"backend/service/activities/cmd/rpc/pb"
	"context"
	"fmt"
	"github.com/jinzhu/copier"

	"backend/service/activities/cmd/api/internal/svc"
	"backend/service/activities/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ActivityUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewActivityUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ActivityUpdateLogic {
	return &ActivityUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ActivityUpdateLogic) ActivityUpdate(req *types.ActivityPostReq) (resp *types.ActivityReply, err error) {
	//id, err := l.svcCtx.ActivityRPC.GetActivitiesById(l.ctx, &pb.GetActivitiesByIdReq{
	//	Id: req.Id,
	//})
	var newdata pb.UpdateActivitiesReq
	_ = copier.Copy(&newdata, req)
	_, err = l.svcCtx.ActivityRPC.UpdateActivities(l.ctx, &newdata)
	if err != nil {
		return nil, errorx.NewCodeError(201, fmt.Sprintf("%v", err), "")
	}

	return nil, errorx.NewCodeError(200, fmt.Sprintf("%s", "修改成功"), "")
}
