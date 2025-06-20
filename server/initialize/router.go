package initialize

import (
	"frame-web/global"
	"frame-web/middleware"
	"frame-web/model/response"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 初始化总路由
func Routers() *gin.Engine {
	Router := gin.New()
	//Router.Use(gin.Recovery())
	Router.Use(middleware.GinRecovery(false)) // 使用自定义的恢复中间件
	if gin.Mode() == gin.DebugMode {
		Router.Use(gin.Logger())
	}
	// 公共路由组 - 不需要鉴权
	PublicGroup := Router.Group(global.CONFIG.System.RouterPrefix)
	{
		// 在这里添加不需要鉴权的路由
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
	}

	// 私有路由组 - 需要鉴权
	PrivateGroup := Router.Group(global.CONFIG.System.RouterPrefix)
	//PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())

	PrivateGroup.Use(middleware.JWTAuth2())
	{
		// 在这里添加需要鉴权的路由
		// 例如：PrivateGroup.GET("/user/info", handler.GetUserInfo)
		PrivateGroup.GET("/test/panic", func(c *gin.Context) {
			panic("测试panic处理")
		})
		PrivateGroup.GET("/test/redis", func(c *gin.Context) {
			err := global.REDIS.Set(c, "aaaa", "bbb", 60*time.Second).Err()
			if err != nil {
				response.FailWithMessage("Redis操作失败: "+err.Error(), c)
			}
			val, err := global.REDIS.Get(c, "aaaa").Result()
			response.OkWithDetailed(val, "Redis操作成功"+val, c)
		})
	}

	return Router
}
