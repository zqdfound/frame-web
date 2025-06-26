package initialize

import (
	"frame-web/api"
	"frame-web/global"
	"frame-web/middleware"
	"frame-web/model/response"
	"github.com/gin-gonic/gin"
)

// 初始化总路由
func Routers() *gin.Engine {
	Router := gin.New()
	//Router.Use(gin.Recovery())
	Router.Use(middleware.GinRecovery(false)) // 使用自定义的恢复中间件
	//if gin.Mode() == gin.DebugMode {
	//	Router.Use(gin.Logger())
	//}
	Router.Use(middleware.Cors()) // 直接放行全部跨域请求
	// 公共路由组 - 不需要鉴权
	SetAuthRoutes(Router)   // 需要鉴权的路由
	SetNoAuthRoutes(Router) // 不需要鉴权的路由
	Router.GET("/health", func(c *gin.Context) {
		response.Ok(c)
	})
	Router.GET("/panic", func(c *gin.Context) {
		panic("测试panic")
	})
	return Router
}

// 需要鉴权的路由
func SetAuthRoutes(router *gin.Engine) {
	apiGroup := router.Group(global.CONFIG.System.RouterPrefix + "/api")
	apiGroup.Use(middleware.JWTAuth2())
	SetupUserRoutes(apiGroup)
	SetupFileRoutes(apiGroup)
}

// 不需要鉴权的路由
func SetNoAuthRoutes(router *gin.Engine) {
	apiGroup := router.Group(global.CONFIG.System.RouterPrefix)
	SetupUserRoutes(apiGroup)
}

// //////////////////////////////////////////////////////////////////////
// 用户相关
func SetupUserRoutes(apiGroup *gin.RouterGroup) {
	userGroup := apiGroup.Group("/users")
	userApi := api.UserApi{}
	userGroup.GET("/get", userApi.GetUser)          // 获取用户
	userGroup.GET("/page", userApi.GetUsersPage)    // 获取用户分页
	userGroup.POST("/create", userApi.CreatUsers)   // 创建用户
	userGroup.DELETE("/remove", userApi.DeleteUser) // 删除用户
	userGroup.POST("/update", userApi.UpdateUser)   // 更新用户信息
	userGroup.POST("/diy", userApi.GetDiy)          // 更新用户信息
}

// 文件操作
func SetupFileRoutes(group *gin.RouterGroup) {
	fileGroup := group.Group("/files")
	fileGroup.POST("/upload", api.UploadFile)   // 上传文件
	fileGroup.DELETE("/delete", api.DeleteFile) // 删除文件
}
