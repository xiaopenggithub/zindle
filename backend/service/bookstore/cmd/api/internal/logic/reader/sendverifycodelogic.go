package logic

import (
	"backend/common/errorx"
	"backend/service/bookstore/cmd/api/internal/config"
	"backend/service/bookstore/cmd/api/internal/svc"
	"backend/service/bookstore/cmd/api/internal/types"
	"backend/service/bookstore/model"
	"bytes"
	"context"
	"crypto/rand"
	"crypto/tls"
	"fmt"
	"github.com/jordan-wright/email"
	"github.com/zeromicro/go-zero/core/logx"
	"math/big"
	"net/smtp"
)

type SendVerifyCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendVerifyCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) SendVerifyCodeLogic {
	return SendVerifyCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendVerifyCodeLogic) SendVerifyCode(req types.VerifyReq) (*types.ReaderReply, error) {
	fmt.Println("-------SendVerifyCode-----")
	fmt.Println(req.Account)
	to := []string{req.Account}
	subject := "邮箱验证"
	code := createRandomNumber(4)
	body := fmt.Sprintf("您的验证码是%v,有效期5分钟", code)
	err := Send(to, subject, body, l.svcCtx.Config)
	if err != nil {
		return nil, errorx.NewCodeError(201, fmt.Sprintf("%v", err), "")
	}
	// 发送成功,验证码入库
	var verifyCode model.VerifyCodes

	verifyCode.Account = req.Account
	verifyCode.Type = req.Type
	verifyCode.Code = code

	_, err1 := l.svcCtx.VerifyCodesModel.Insert(verifyCode)

	if err1 != nil {
		return nil, errorx.NewCodeError(202, fmt.Sprintf("%v", err1), "")
	}

	return nil, errorx.NewCodeError(200, "发送成功", "")
}
func Send(to []string, subject string, body string, c config.Config) error {
	from := c.Email.From
	nickname := c.Email.Nickname
	secret := c.Email.Secret
	host := c.Email.Host
	port := c.Email.Port
	isSSL := c.Email.IsSSL

	auth := smtp.PlainAuth("", from, secret, host)
	e := email.NewEmail()
	if nickname != "" {
		e.From = fmt.Sprintf("%s <%s>", nickname, from)
	} else {
		e.From = from
	}
	e.To = to
	e.Subject = subject
	e.HTML = []byte(body)
	var err error
	hostAddr := fmt.Sprintf("%s:%d", host, port)
	if isSSL {
		err = e.SendWithTLS(hostAddr, auth, &tls.Config{ServerName: host})
	} else {
		err = e.Send(hostAddr, auth)
	}
	return err
}

//https://www.cnblogs.com/liujie-php/p/10736266.html
func createRandomNumber(len int) string {
	var numbers = []byte{1, 2, 3, 4, 5, 7, 8, 9}
	var container string
	length := bytes.NewReader(numbers).Len()

	for i := 1; i <= len; i++ {
		random, err := rand.Int(rand.Reader, big.NewInt(int64(length)))
		if err != nil {

		}
		container += fmt.Sprintf("%d", numbers[random.Int64()])
	}
	return container
}

func createRandomString(len int) string {
	var container string
	var str = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	b := bytes.NewBufferString(str)
	length := b.Len()
	bigInt := big.NewInt(int64(length))
	for i := 0; i < len; i++ {
		randomInt, _ := rand.Int(rand.Reader, bigInt)
		container += string(str[randomInt.Int64()])
	}
	return container
}
