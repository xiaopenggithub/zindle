package logic

import (
	"backend/common/errorx"
	"backend/common/utils"
	model2 "backend/service/backend/model"
	"backend/service/bookstore/cmd/api/internal/svc"
	"backend/service/bookstore/cmd/api/internal/types"
	"backend/service/bookstore/model"
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/zeromicro/go-zero/core/logx"
	"strings"
	"time"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) LoginLogic {
	return LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req types.LoginReq, ip string) (*types.ReaderReply, error) {
	if len(strings.TrimSpace(req.Username)) == 0 || len(strings.TrimSpace(req.Password)) == 0 {
		return nil, errorx.NewDefaultError("参数错误", "")
	}
	fmt.Println("username----", req.Username)
	fmt.Println("Password----", req.Password)
	reader, err := l.svcCtx.ReadersModel.FindOneByUserName(req.Username)
	switch err {
	case nil:
	case model2.ErrNotFound:
		return nil, errorx.NewDefaultError("用户名密码不正确", "")
		//return nil, errorx.NewDefaultError("用户名不存在", "")
	default:
		return nil, err
	}
	fmt.Println("old Password----", utils.MD5V(req.Password))
	fmt.Println("current Password----", reader.Password)
	if reader.Password != utils.MD5V(req.Password) {
		//fmt.Println("req.password", utils.MD5V(req.Password))
		//fmt.Println("req.Username", req.Username)
		//fmt.Printf("userInfo.Password[%s]\n", userInfo.Password)
		//fmt.Printf("userInfo.Name[%s]\n", userInfo.Name)
		//fmt.Printf("userInfo:%v\n", userInfo)
		return nil, errorx.NewDefaultError("用户名或密码不正确", "")
	}

	// ---start---
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.Auth.AccessExpire
	fmt.Println(accessExpire)
	jwtToken, err := l.getReaderJwtToken(l.svcCtx.Config.Auth.AccessSecret, now, l.svcCtx.Config.Auth.AccessExpire, reader)
	if err != nil {
		return nil, errorx.NewCodeError(201, fmt.Sprintf("[授权错误]%v", err), "")
	}
	// ---end---
	tokenInfo := make(map[string]interface{})
	tokenInfo["token"] = jwtToken
	tokenInfo["username"] = reader.Username
	//tokenInfo["menus"] = l.getAdminsMenu(userInfo.RoleId)

	//更新登录信息
	reader.LoginDate = time.Now().Format("2006-01-02 15:04:05")
	reader.LoginIp = ip
	err = l.svcCtx.ReadersModel.Update(reader)
	if err != nil {
		return nil, errorx.NewCodeError(202, fmt.Sprintf("[数据处理错误]%v", err), "")
	}
	return nil, errorx.NewCodeError(200, "登录成功", tokenInfo)
}
func (l *LoginLogic) getReaderJwtToken(secretKey string, iat, seconds int64, reader model.Readers) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["readerId"] = reader.Id
	claims["readerUsername"] = reader.Username
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
