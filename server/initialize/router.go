package initialize

import (
	"frame-web/global"
	"frame-web/middleware"
	"frame-web/model/request"
	"frame-web/model/response"
	"frame-web/svc/models"
	userService "frame-web/svc/service"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 初始化总路由
func Routers() *gin.Engine {
	Router := gin.New()
	//Router.Use(gin.Recovery())
	Router.Use(middleware.GinRecovery(false)) // 使用自定义的恢复中间件
	if gin.Mode() == gin.DebugMode {
		Router.Use(gin.Logger())
	}
	// 跨域，如需跨域可以打开下面的注释
	Router.Use(middleware.Cors()) // 直接放行全部跨域请求
	global.LOG.Info("use middleware cors")
	// 公共路由组 - 不需要鉴权
	PublicGroup := Router.Group(global.CONFIG.System.RouterPrefix)
	{
		// 在这里添加不需要鉴权的路由
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
		PublicGroup.POST("/users", func(c *gin.Context) {
			var user models.User
			if err := c.ShouldBindJSON(&user); err != nil {
				response.Fail(c)
				return
			}
			if err := userService.CreateUser(&user); err != nil {
				global.LOG.Error("新增用户失败!", zap.Error(err))
				response.FailWithMessage("新增用户失败:"+err.Error(), c)
				return
			}
			global.LOG.Info("新增用户成功!", zap.Any("user", user))
			response.Ok(c)
		})
		PublicGroup.GET("/users", func(c *gin.Context) {
			pageNum, _ := strconv.Atoi(c.Query("pageNum"))
			pageSize, _ := strconv.Atoi(c.Query("pageSize"))

			list, total, err := userService.GetAllUsersPage(&userService.UserPageReq{
				PageInfo: request.PageInfo{
					Page:     pageNum,
					PageSize: pageSize,
				},
				User: &models.User{
					Username: c.Query("username"),
				},
			})
			if err != nil {
				global.LOG.Error("获取失败!", zap.Error(err))
				response.FailWithMessage("获取失败:"+err.Error(), c)
				return
			}
			response.OkWithDetailed(response.PageResult{
				List:     list,
				Total:    total,
				Page:     1,
				PageSize: 10,
			}, "获取成功", c)
		})

		PublicGroup.POST("/device", func(c *gin.Context) {
			type DInfo struct {
				Sn  string `json:"sn"`
				Pwd string `json:"pwd"`
			}
			// DInfo 是类型，不能直接取地址，需要创建该类型的实例
			var dInfo DInfo
			if err := c.ShouldBindJSON(&dInfo); err != nil {
				response.FailWithMessage(err.Error(), c)
				return
			}
			sn := dInfo.Sn
			pwd := dInfo.Pwd
			if "woailiming" != pwd {
				response.FailWithMessage("wrong password", c)
				return
			}
			device, err := userService.GetDeviceInfo(sn)
			if err != nil {
				response.FailWithMessage(err.Error(), c)
				return
			}
			response.OkWithData(device, c)
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
