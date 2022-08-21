package logic

import (
	"backend/common/errorx"
	"backend/service/activities/cmd/api/internal/svc"
	"backend/service/activities/cmd/api/internal/types"
	"context"
	"fmt"
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
	one, err := l.svcCtx.ActivitysModel.FindOne(req.Id)
	if err != nil {
		return nil, errorx.NewCodeError(201, fmt.Sprintf("%v", err), "")
	}
	data := make(map[string]interface{})
	data["item"] = one

	return nil, errorx.NewCodeError(200, fmt.Sprintf("%s", "获取成功"), data)
}
