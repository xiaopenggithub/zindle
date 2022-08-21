package logic

import (
	"backend/common/errorx"
	"backend/service/activities/cmd/api/internal/svc"
	"backend/service/activities/cmd/api/internal/types"
	"context"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type ActivityOrdersUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 活动预订信息 update
func NewActivityOrdersUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) ActivityOrdersUpdateLogic {
	return ActivityOrdersUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ActivityOrdersUpdateLogic) ActivityOrdersUpdate(req types.ActivityOrdersPostReq) (*types.ActivityOrdersReply, error) {
	oldData, err := l.svcCtx.ActivityOrderssModel.FindOne(req.Id)
	if err != nil {
		return nil, errorx.NewCodeError(201, fmt.Sprintf("%v", err), "")
	}

	copier.Copy(&oldData, &req)
	err = l.svcCtx.ActivityOrderssModel.Update(*oldData)

	if err != nil {
		return nil, errorx.NewCodeError(201, fmt.Sprintf("%v", err), "")
	}

	return nil, errorx.NewCodeError(200, fmt.Sprintf("%s", "修改成功"), "")
}
