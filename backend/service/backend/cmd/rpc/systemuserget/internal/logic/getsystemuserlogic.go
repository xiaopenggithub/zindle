package logic

import (
	"context"
	"database/sql"
	"fmt"

	"backend/service/backend/cmd/rpc/systemuserget/internal/svc"
	"backend/service/backend/cmd/rpc/systemuserget/systemuserget"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSystemuserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSystemuserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSystemuserLogic {
	return &GetSystemuserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetSystemuserLogic) GetSystemuser(in *systemuserget.Request) (*systemuserget.Response, error) {
	fmt.Println("GetSystemuser-Id:", in.Id)
	one, err := l.svcCtx.SystemUserModel.FindOne(in.Id)
	if err != nil {
		return nil, err
	}

	nullUsername := sql.NullString{Valid: false}
	if in.Id > 0 {
		nullUsername = sql.NullString{String: fmt.Sprintf("%d", in.Id), Valid: true}
	}
	fmt.Println(nullUsername)

	return &systemuserget.Response{
		NickName: one.NickName,
	}, nil
}
