package logic

import (
	"backend/common/errorx"
	"backend/service/activities/cmd/api/internal/svc"
	"backend/service/activities/cmd/api/internal/types"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
)

type ActivityOrdersDeleteBatchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 活动预订信息 deletebatch
func NewActivityOrdersDeleteBatchLogic(ctx context.Context, svcCtx *svc.ServiceContext) ActivityOrdersDeleteBatchLogic {
	return ActivityOrdersDeleteBatchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ActivityOrdersDeleteBatchLogic) ActivityOrdersDeleteBatch(req types.ActivityOrdersDelBatchReq) (*types.ActivityOrdersReply, error) {
	err := l.svcCtx.ActivityOrderssModel.DeleteBatch(req.Ids)
	if err != nil {
		return nil, errorx.NewCodeError(201, fmt.Sprintf("%v", err), "")
	}

	return nil, errorx.NewCodeError(200, fmt.Sprintf("编号[%s]删除成功", req.Ids), "")
}
