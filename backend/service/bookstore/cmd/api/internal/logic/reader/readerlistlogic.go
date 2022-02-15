package logic

import (
	"backend/common/errorx"
	"backend/common/utils"
	"backend/service/bookstore/cmd/api/internal/svc"
	"backend/service/bookstore/cmd/api/internal/types"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
)

type ReaderListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 读者 list
func NewReaderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) ReaderListLogic {
	return ReaderListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReaderListLogic) ReaderList(req types.ReaderListReq) (*types.ReaderReply, error) {
	reqParam := utils.ListReq{}
	reqParam.Page = req.Page
	reqParam.PageSize = req.PageSize
	reqParam.Keyword = req.Keyword
	list, i, err := l.svcCtx.ReadersModel.List(reqParam)
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
