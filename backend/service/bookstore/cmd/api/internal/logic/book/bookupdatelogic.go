package logic

import (
	"backend/common/errorx"
	"backend/common/utils"
	"backend/service/bookstore/cmd/api/internal/svc"
	"backend/service/bookstore/cmd/api/internal/types"
	"context"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"io"
	"mime/multipart"
	"os"
	"path"
	"strings"
	"time"
)

type BookUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 图书管理 update
func NewBookUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) BookUpdateLogic {
	return BookUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BookUpdateLogic) BookUpdate(req types.BookPostReq, file *multipart.FileHeader) (*types.BookReply, error) {
	oldData, err := l.svcCtx.BooksModel.FindOne(req.Id)
	if err != nil {
		return nil, errorx.NewCodeError(201, fmt.Sprintf("%v", err), "")
	}
	oldCover := oldData.Cover

	fmt.Printf("%v\n", oldData)
	fmt.Println("-------oldCover-------", oldCover)

	req.Cover = uploadfile(file)

	copier.Copy(&oldData, &req)
	err = l.svcCtx.BooksModel.Update(*oldData)

	if err != nil {
		return nil, errorx.NewCodeError(201, fmt.Sprintf("%v", err), "")
	}
	// 删除oldcover
	// delete oldCover begin
	//oldPath := "uploads" + "/" + oldCover
	oldPath := oldCover

	if _, err := os.Stat(oldPath); os.IsNotExist(err) {
		// nofile
		fmt.Println("没有找到文件", oldCover)
	} else {
		fmt.Println("准备删除文件", oldCover)
		if err := os.Remove(oldPath); err != nil {
			fmt.Println("删除文件error")
		}
	}
	// delete oldCover end

	return nil, errorx.NewCodeError(200, fmt.Sprintf("%s", "修改成功"), "")
}
func uploadfile(file *multipart.FileHeader) string {
	filePath := ""
	if file != nil {
		// 读取文件后缀
		ext := path.Ext(file.Filename)
		ext = strings.ToLower(ext)
		fmt.Println("------file-ext..", ext)
		// 读取文件名并加密
		fileName := strings.TrimSuffix(file.Filename, ext)
		fileName = utils.MD5V(fileName)
		// 拼接新文件名
		lastName := fileName + "_" + time.Now().Format("20060102150405") + ext
		fmt.Println("------filename..", fileName)
		fmt.Println("------newname..", lastName)
		// todo:读取全局变量的定义路径
		savePath := "uploads"
		// 尝试创建此路径
		err := os.MkdirAll(savePath, os.ModePerm)
		if err != nil {
			return ""
		}
		// 拼接路径和文件名
		filePath = savePath + "/" + lastName

		// 上传逻辑 begin
		// 打开文件 defer 关闭
		src, err := file.Open()
		if err != nil {
			return ""
		}
		defer src.Close()
		// 创建文件 defer 关闭
		out, err := os.Create(filePath)
		if err != nil {
			return ""
		}
		defer out.Close()
		// 传输（拷贝）文件
		_, err = io.Copy(out, src)
		if err != nil {
			return ""
		}
		// 上传逻辑 end

	}
	return filePath
}
