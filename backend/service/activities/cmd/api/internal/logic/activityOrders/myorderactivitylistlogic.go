package logic

import (
	"backend/common/errorx"
	"backend/common/utils"
	"context"
	"encoding/json"
	"fmt"

	"backend/service/activities/cmd/api/internal/svc"
	"backend/service/activities/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MyOrderActivityListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMyOrderActivityListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MyOrderActivityListLogic {
	return &MyOrderActivityListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MyOrderActivityListLogic) MyOrderActivityList(req *types.ActivityOrdersListReq) (resp *types.ActivityOrdersReply, err error) {
	reqParam := utils.ListReq{}
	reqParam.Page = req.Page
	reqParam.PageSize = req.PageSize
	reqParam.Keyword = req.Keyword

	idNumber := json.Number(fmt.Sprintf("%v", l.ctx.Value("readerId")))
	id, err := idNumber.Int64()
	if err != nil {
		return nil, errorx.NewCodeError(401, "请重新登录再操作", "")
	}
	fmt.Println("---MyOrderActivityListLogic>id=---", id)

	list, i, err := l.svcCtx.ActivityOrderssModel.ListMy(reqParam, id)
	if err != nil {
		return nil, errorx.NewCodeError(201, fmt.Sprintf("%v", err), "")
	}
	data := make(map[string]interface{})
	data["page"] = req.Page
	data["pageSize"] = req.PageSize
	data["total"] = i
	data["list"] = list

	return nil, errorx.NewCodeError(200, "ok", data)
}
