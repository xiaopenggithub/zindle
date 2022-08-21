package logic

import (
	"backend/common/errorx"
	"backend/service/activities/cmd/api/internal/svc"
	"backend/service/activities/cmd/api/internal/types"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
)

type ActivityOrdersDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 活动预订信息 delete
func NewActivityOrdersDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) ActivityOrdersDeleteLogic {
	return ActivityOrdersDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ActivityOrdersDeleteLogic) ActivityOrdersDelete(req types.ActivityOrdersDelReq) (*types.ActivityOrdersReply, error) {
	err := l.svcCtx.ActivityOrderssModel.Delete(req.Id)
	if err != nil {
		return nil, errorx.NewCodeError(201, fmt.Sprintf("%v", err), "")
	}

	return nil, errorx.NewCodeError(200, fmt.Sprintf("编号[%d]删除成功", req.Id), "")
}
