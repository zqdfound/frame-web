package api

import (
	"frame-web/global"
	"frame-web/model/response"
	"frame-web/svc/thirdService"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func UploadFile(c *gin.Context) {
	header, err := c.FormFile("file")
	if err != nil {
		response.FailWithMessage("获取文件失败:"+err.Error(), c)
		return
	}
	url, relativePath, uploadErr := thirdService.UploadFile(header)
	if uploadErr != nil {
		response.FailWithMessage("上传文件失败:"+uploadErr.Error(), c)
		return
	}
	global.LOG.Info("文件上传成功", zap.String("url", url), zap.String("relativePath", relativePath))
	response.OkWithData(map[string]string{
		"url":          url,
		"relativePath": relativePath,
	}, c)
}

func DeleteFile(c *gin.Context) {
	relativePath := c.Query("relativePath")
	if relativePath == "" {
		response.FailWithMessage("relativePath不能为空", c)
		return
	}
	err := thirdService.DeleteFile(relativePath)
	if err != nil {
		response.FailWithMessage("删除文件失败:"+err.Error(), c)
		return
	}
	global.LOG.Info("文件删除成功", zap.String("relativePath", relativePath))
	response.OkWithMessage("删除成功", c)
}
