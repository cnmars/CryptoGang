package upload

import (
	"mime/multipart"

	"golddigger/global"
)

// OSS 对象存储接口
type OSS interface {
	UploadFile(file *multipart.FileHeader) (string, string, error)
	DeleteFile(key string) error
}

// NewOss OSS的实例化方法
func NewOss() OSS {
	switch global.GD_CONFIG.System.OssType {
	case "local":
		return &Local{}
	case "qiniu":
		return &Qiniu{}
	case "tencent-cos":
		return &TencentCOS{}
	case "aliyun-oss":
		return &AliyunOSS{}
	case "huawei-obs":
		return HuaWeiObs
	case "aws-s3":
		return &AwsS3{}
	default:
		return &Local{}
	}
}
