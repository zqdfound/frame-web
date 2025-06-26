package thirdService

import (
	"frame-web/utils/upload"
	"mime/multipart"
)

func UploadFile(header *multipart.FileHeader) (url string, relativePath string, err error) {
	oss := upload.NewOss()
	filePath, key, uploadErr := oss.UploadFile(header)
	if uploadErr != nil {
		return "", "", uploadErr
	}
	return filePath, key, nil
}

// DeleteFile 删除文件
// 传入文件的相对路径
func DeleteFile(relativePath string) error {
	oss := upload.NewOss()
	return oss.DeleteFile(relativePath)
}
