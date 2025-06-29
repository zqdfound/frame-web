package core

import (
	"fmt"
	"frame-web/global"
	"frame-web/initialize"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"time"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {

	//init Redis
	//utils.Redis()
	Router := initialize.Routers()
	address := fmt.Sprintf(":%d", global.CONFIG.System.Addr)
	//address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)
	s := initServer(address, Router)

	//global.LOG.Info("server run success on ", zap.String("address",global.CONFIG.System.Addr ))

	fmt.Printf(`
	--------------------------------------启动成功-------------------------------------
`, address)
	global.LOG.Error(s.ListenAndServe().Error())
}

func initServer(address string, router *gin.Engine) server {
	s := endless.NewServer(address, router)
	s.ReadHeaderTimeout = 10 * time.Minute
	s.WriteTimeout = 10 * time.Minute
	s.MaxHeaderBytes = 1 << 20
	return s
}
