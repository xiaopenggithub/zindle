package logic

import (
	"backend/common/errorx"
	"backend/service/activities/cmd/api/internal/svc"
	"backend/service/activities/cmd/api/internal/types"
	"backend/service/activities/model"
	"context"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type ActivityOrdersAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 活动预订信息 create
func NewActivityOrdersAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) ActivityOrdersAddLogic {
	return ActivityOrdersAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ActivityOrdersAddLogic) ActivityOrdersAdd(req types.ActivityOrdersPostReq) (*types.ActivityOrdersReply, error) {
	//检查是否重名
	//checkItem, err2 := l.svcCtx.ActivityOrderssModel.CheckDuplicate(req.Name)
	//if checkItem.Id > 0 && err2 == nil {
	//	return nil, errorx.NewCodeError(201, "已经存在", "")
	//}

	var item model.ActivityOrderss
	copier.Copy(&item, &req)

	_, err := l.svcCtx.ActivityOrderssModel.Insert(item)
	if err != nil {
		return nil, errorx.NewCodeError(202, fmt.Sprintf("%v", err), "")
	}

	return nil, errorx.NewCodeError(200, "添加成功", "")
}
