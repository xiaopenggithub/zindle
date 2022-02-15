package logic

import (
	"backend/common/errorx"
	"backend/service/bookstore/cmd/api/internal/svc"
	"backend/service/bookstore/cmd/api/internal/types"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
)

type ReaderFindOneLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 读者 findone
func NewReaderFindOneLogic(ctx context.Context, svcCtx *svc.ServiceContext) ReaderFindOneLogic {
	return ReaderFindOneLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReaderFindOneLogic) ReaderFindOne(req types.ReaderDelReq) (*types.ReaderReply, error) {
	one, err := l.svcCtx.ReadersModel.FindOne(req.Id)
	if err != nil {
		return nil, errorx.NewCodeError(201, fmt.Sprintf("%v", err), "")
	}
	data := make(map[string]interface{})
	data["item"] = one

	return nil, errorx.NewCodeError(200, fmt.Sprintf("%s", "获取成功"), data)
}
