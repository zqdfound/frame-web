package upload

import "mime/multipart"

// 上传工具工厂，目前只有阿里云oss
type OSS interface {
	UploadFile(file *multipart.FileHeader) (string, string, error)
	DeleteFile(key string) error
}

func NewOss() OSS {
	return &AliyunOSS{}
}
