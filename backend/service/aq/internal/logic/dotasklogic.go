package logic

import (
	"context"

	"backend/service/aq/internal/svc"
	"backend/service/aq/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DoTaskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDoTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DoTaskLogic {
	return &DoTaskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DoTaskLogic) DoTask(req *types.Req) (resp *types.Reply, err error) {
	// todo: add your logic here and delete this line

	return
}
