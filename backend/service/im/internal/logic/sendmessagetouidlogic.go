package logic

import (
	"backend/service/im/internal"
	"backend/service/im/internal/svc"
	"backend/service/im/internal/types"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendMessageToUidLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendMessageToUidLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendMessageToUidLogic {
	return &SendMessageToUidLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendMessageToUidLogic) SendMessageToUid(req *types.MessageReq) (resp *types.MessageReply, err error) {
	internal.GlobalHub.SendMessageToUid([]byte(req.Message), req.Cid)
	return
}
