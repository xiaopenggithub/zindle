package logic

import (
	"backend/common/errorx"
	"backend/common/utils"
	"backend/service/backend/cmd/api/internal/svc"
	"backend/service/backend/cmd/api/internal/types"
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path"
	"strings"
	"time"
)

type ChangeAvatarLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 连接到 MinIO 服务器
var endpoint = "127.0.0.1:9000" // MinIO 服务器地址
var accessKey = "minioadmin"    // MinIO 访问密钥
var secretKey = "minioadmin"    // MinIO 访问密钥密钥
var useSSL = false              // 是否使用 SSL 连接

var bucketName = "avatar"    // 存储桶名称
var location = "sh-20240220" // 存储桶的地理位置

func NewChangeAvatarLogic(ctx context.Context, svcCtx *svc.ServiceContext) ChangeAvatarLogic {
	return ChangeAvatarLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChangeAvatarLogic) ChangeAvatar(req types.AdminChangeAvatarReq, file *multipart.FileHeader) (*types.AdminReply, error) {
	dst := ""
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
			return nil, errorx.NewCodeError(202, fmt.Sprintf("%v", err), "")
		}
		// 拼接路径和文件名
		dst = savePath + "/" + lastName

		// 上传逻辑 begin
		// 打开文件 defer 关闭
		src, err := file.Open()
		if err != nil {
			return nil, errorx.NewCodeError(202, fmt.Sprintf("%v", err), "")
		}
		defer src.Close()
		// 创建文件 defer 关闭
		out, err := os.Create(dst)
		if err != nil {
			return nil, errorx.NewCodeError(202, fmt.Sprintf("%v", err), "")
		}
		defer out.Close()
		// 传输（拷贝）文件
		_, err = io.Copy(out, src)
		if err != nil {
			return nil, errorx.NewCodeError(202, fmt.Sprintf("%v", err), "")
		}
		// 上传逻辑 end

		// begin transaction
		// update avatar begin
		userIdNumber := json.Number(fmt.Sprintf("%v", l.ctx.Value("userId")))
		userId, err := userIdNumber.Int64()
		if err != nil {
			return nil, errorx.NewCodeError(401, "请重新登录再操作", "")
		}
		// 获取用户信息
		oldUser, err := l.svcCtx.SystemUserModel.FindOne(userId)
		if err != nil {
			return nil, errorx.NewCodeError(201, "数据错误", "")
		}
		oldAvatarPath := savePath + "/" + oldUser.Avatar
		// 执行修改,只存储文件名
		oldUser.Avatar = lastName
		fmt.Printf("%v\n", oldUser)
		err = l.svcCtx.SystemUserModel.Update(*oldUser)
		if err != nil {
			return nil, errorx.NewCodeError(203, fmt.Sprintf("[%v]修改失败，请重试", err), "")
		}
		// update avatar end

		// delete old avatar begin
		if _, err := os.Stat(oldAvatarPath); os.IsNotExist(err) {
			// nofile
		} else {
			if err := os.Remove(oldAvatarPath); err != nil {
				fmt.Println("删除文件error")
			}
		}
		// delete old avatar end

	}
	return nil, errorx.NewCodeError(200, fmt.Sprintf("%v", "上传完成"), "/"+dst)

}

func MD5V(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
func MinIOUpload(objectName string, file multipart.File, header *multipart.FileHeader) {
	// 连接到 MinIO 服务器
	//endpoint := "127.0.0.1:9000" // MinIO 服务器地址
	//accessKey := "minioadmin"    // MinIO 访问密钥
	//secretKey := "minioadmin"    // MinIO 访问密钥密钥
	//useSSL := false              // 是否使用 SSL 连接

	// 初始化 MinIO 客户端对象
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}

	// 创建一个新的存储桶
	//bucketName := "mybucket" // 存储桶名称
	//location := "us-east-1"  // 存储桶的地理位置
	err = minioClient.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{Region: location})
	if err != nil {
		// 如果存储桶已经存在，则会报错，可以忽略该错误
		exists, errBucketExists := minioClient.BucketExists(context.Background(), bucketName)
		if errBucketExists == nil && exists {
			log.Printf("Bucket '%s' already exists\n", bucketName)
		} else {
			log.Fatalln(err)
		}
	} else {
		log.Printf("Successfully created bucket '%s'\n", bucketName)
	}

	// 将文件上传到 MinIO
	//n, err := minioClient.PutObject(context.Background(), bucketName, objectName, bytes.NewReader(fileData), int64(len(fileData)), minio.PutObjectOptions{})
	n, err := minioClient.PutObject(context.Background(), bucketName, objectName, file, header.Size, minio.PutObjectOptions{})
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Successfully uploaded %s of size %d\n", objectName, n)
}

func MinIORemove(objectName string) {
	// 连接到 MinIO 服务器
	//endpoint := "127.0.0.1:9000" // MinIO 服务器地址
	//accessKey := "minioadmin"    // MinIO 访问密钥
	//secretKey := "minioadmin"    // MinIO 访问密钥密钥
	//useSSL := false              // 是否使用 SSL 连接

	// 初始化 MinIO 客户端对象
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}

	// 创建一个新的存储桶
	//bucketName := "mybucket" // 存储桶名称
	//location := "us-east-1"  // 存储桶的地理位置
	err = minioClient.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{Region: location})
	if err != nil {
		// 如果存储桶已经存在，则会报错，可以忽略该错误
		exists, errBucketExists := minioClient.BucketExists(context.Background(), bucketName)
		if errBucketExists == nil && exists {
			log.Printf("Bucket '%s' already exists\n", bucketName)
		} else {
			log.Fatalln(err)
		}
	} else {
		log.Printf("Successfully created bucket '%s'\n", bucketName)
	}
	//// 删除文件
	err = minioClient.RemoveObject(context.Background(), bucketName, objectName, minio.RemoveObjectOptions{})
	if err != nil {
		log.Fatalln(err)
	}
}
