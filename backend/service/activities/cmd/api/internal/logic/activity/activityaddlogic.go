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

type ActivityAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 付款信息 create
func NewActivityAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) ActivityAddLogic {
	return ActivityAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ActivityAddLogic) ActivityAdd(req types.ActivityPostReq) (*types.ActivityReply, error) {
	//检查是否重名
	checkItem, err2 := l.svcCtx.ActivitysModel.CheckDuplicate(req.Title)
	if checkItem.Id > 0 && err2 == nil {
		return nil, errorx.NewCodeError(201, "已经存在", "")
	}

	var item model.Activitys
	copier.Copy(&item, &req)

	_, err := l.svcCtx.ActivitysModel.Insert(item)
	if err != nil {
		return nil, errorx.NewCodeError(202, fmt.Sprintf("%v", err), "")
	}

	return nil, errorx.NewCodeError(200, "添加成功", "")
}
