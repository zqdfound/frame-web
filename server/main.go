package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"frame-web/config"
	"frame-web/db"
	mid "frame-web/middleware"
	"frame-web/model/response"
	"frame-web/utils"
	zlog "frame-web/zap"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var (
	ctx       = context.Background()
	jwtConfig = mid.JWTConfig{
		SigningKey:  "woailiming",
		WhiteList:   []string{"/api/public", "/test/set/jwt"},
		ContextKey:  "user",
		TokenLookup: "header:Authorization",
	}
)

func main() {
	// 初始化配置
	config.InitViper()
	//dbDsn := viper.GetString("database.dsn")

	// 初始化zap日志
	zlog.InitLogger()
	// 初始化数据库
	db.InitDB()
	//初始化redis
	utils.NewRedisHelper()

	r := gin.Default()
	r.Use(mid.JWTAuth(jwtConfig))
	r.Use(mid.Recovery())
	// 使用跨域中间件
	r.Use(mid.Cors())
	// 静态文件服务
	r.Static("/static", "../frontend/dist")

	// API路由
	r.GET("/api/logs", listLogs)

	// WebSocket路由
	// r.GET("/ws", HandleWebSocket)

	r.POST("/api/sn", HandleSnForm)

	// r.GET("/test/panic", func(c *gin.Context) {
	// 	panic("测试panic处理")
	// })
	r.GET("/test/redis", func(c *gin.Context) {
		name, err := utils.GetRedisHelper().Set(c, "aaa", "1111", 10*time.Minute).Result()
		if err != nil {
			zlog.Error("Failed to set redis",
				"err", err,
			)
		}
		zlog.Info("Success to set redis",
			"name", name,
		)
	})
	r.GET("/test/set/jwt", func(c *gin.Context) {
		claims := jwt.MapClaims{
			"user_id":  123,
			"username": "testuser",
			"exp":      time.Now().Add(time.Hour * 24).Unix(), // 24小时后过期
		}
		jwtStr, err := mid.GenerateToken(jwtConfig.SigningKey, claims)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "未授权"})
			return
		}
		response.OkWithData(jwtStr, c)
	})

	r.Run(":8080")

	zlog.Info("Server started",
		"port", "8080",
	)

}

func listLogs(c *gin.Context) {
	fileNames := []string{
		"./log/log1.txt",
	}
	// 实现获取日志文件列表的逻辑
	c.JSON(http.StatusOK, gin.H{"files": fileNames})
}

// KGF9JNHJHQ
func HandleSnForm(c *gin.Context) {
	var snForm struct {
		Sn  string `json:"sn"`
		Pwd string `json:"pwd"`
	}

	if err := c.ShouldBindJSON(&snForm); err != nil {
		zlog.Error("Failed to bind JSON",
			"err", err,
		)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	zlog.Info("Received SN form",
		"sn", snForm.Sn,
		"pwd", snForm.Pwd, // 出于安全考虑，不记录实际密码
	)
	if snForm.Pwd != "nevermore" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "pwd error"})
		return
	}
	//1 查询设备信息 dep配置
	deviceInfo, err := reqDeviceInfo(snForm.Sn)
	if err != nil {
		zlog.Error("Failed to req device info",
			"err", err,
		)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 实现获取日志文件列表的逻辑
	c.JSON(http.StatusOK, gin.H{"device": deviceInfo})
	//2 移除激活锁，移除mdm
	//3 去除痕迹

}

// KGF9JNHJHQ
func reqDeviceInfo(sn string) (map[string]interface{}, error) {
	// 准备Basic Auth
	username := "aaa"
	password := "bbbbb"
	auth := username + ":" + password
	basicAuth := "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))

	// 创建请求
	req, err := http.NewRequest("POST", "https://xxxxxxxxxxxxxxx.com/mdm/facade/deviceInfo/"+sn, nil)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %w", err)
	}

	// 设置请求头
	req.Header.Add("Authorization", basicAuth)
	fmt.Printf("req.Header: %v\n", req.Header)
	req.Header.Add("Content-Type", "application/json")

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}
	defer resp.Body.Close()
	// 检查响应状态码
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("请求失败，状态码: %d", resp.StatusCode)
	}
	// 解析响应
	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, fmt.Errorf("解析响应失败: %v\n", err)

	}
	return result, nil
}
