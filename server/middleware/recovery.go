package middleware

import (
	zlog "frame-web/zap"

	"github.com/gin-gonic/gin"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// 使用zap.go中封装的方法记录错误
				zlog.Error("panic recovered", "error", err)
				// c.AbortWithStatusJSON(500, gin.H{
				// 	"code":    500,
				// 	"message": "Internal Server Error",
				// })
			}
		}()

		c.Next()
	}
}
