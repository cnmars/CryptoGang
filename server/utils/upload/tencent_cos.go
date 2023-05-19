package upload

import (
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/url"
	"time"

	"golddigger/global"

	"github.com/tencentyun/cos-go-sdk-v5"
	"go.uber.org/zap"
)

type TencentCOS struct{}

// UploadFile upload file to COS
func (*TencentCOS) UploadFile(file *multipart.FileHeader) (string, string, error) {
	client := NewClient()
	f, openError := file.Open()
	if openError != nil {
		global.GD_LOG.Error("function file.Open() Filed", zap.Any("err", openError.Error()))
		return "", "", errors.New("function file.Open() Filed, err:" + openError.Error())
	}
	defer f.Close() // 创建文件 defer 关闭
	fileKey := fmt.Sprintf("%d%s", time.Now().Unix(), file.Filename)

	_, err := client.Object.Put(context.Background(), global.GD_CONFIG.TencentCOS.PathPrefix+"/"+fileKey, f, nil)
	if err != nil {
		panic(err)
	}
	return global.GD_CONFIG.TencentCOS.BaseURL + "/" + global.GD_CONFIG.TencentCOS.PathPrefix + "/" + fileKey, fileKey, nil
}

// DeleteFile delete file form COS
func (*TencentCOS) DeleteFile(key string) error {
	client := NewClient()
	name := global.GD_CONFIG.TencentCOS.PathPrefix + "/" + key
	_, err := client.Object.Delete(context.Background(), name)
	if err != nil {
		global.GD_LOG.Error("function bucketManager.Delete() Filed", zap.Any("err", err.Error()))
		return errors.New("function bucketManager.Delete() Filed, err:" + err.Error())
	}
	return nil
}

// NewClient init COS client
func NewClient() *cos.Client {
	urlStr, _ := url.Parse("https://" + global.GD_CONFIG.TencentCOS.Bucket + ".cos." + global.GD_CONFIG.TencentCOS.Region + ".myqcloud.com")
	baseURL := &cos.BaseURL{BucketURL: urlStr}
	client := cos.NewClient(baseURL, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  global.GD_CONFIG.TencentCOS.SecretID,
			SecretKey: global.GD_CONFIG.TencentCOS.SecretKey,
		},
	})
	return client
}
