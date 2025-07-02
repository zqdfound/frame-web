// @title Frame-Web API
// @version 1.0
// @description This is a sample server for Frame-Web.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
package main

import (
	"frame-web/core"
	"frame-web/global"
	"go.uber.org/zap"
)

func main() {
	// 初始化配置
	global.GLOBAL_VP = core.Viper()
	// 初始化zap日志
	global.LOG = core.ZapInit()
	zap.ReplaceGlobals(global.LOG)
	// 初始化Mysql数据库
	//global.DB = core.InitMysql()
	//global.LOG.Info("Mysql数据库连接成功")
	//if global.DB != nil {
	//	// 程序结束前关闭数据库链接
	//	db, _ := global.DB.DB()
	//	defer db.Close()
	//}
	core.RunWindowsServer()

}
