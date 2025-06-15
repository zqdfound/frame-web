package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"go-log-v/server/db"
	mid "go-log-v/server/middleware"
	zlog "go-log-v/server/zap"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化zap日志
	zlog.InitLogger()
	// 初始化数据库
	db.InitDB()
	r := gin.Default()
	// 使用跨域中间件
	r.Use(mid.Cors())
	// 静态文件服务
	r.Static("/static", "../frontend/dist")
	// 根路径重定向到前端
	//r.GET("/", func(c *gin.Context) {
	//	c.File("./frontend/dist/index.html")
	//})

	// API路由
	r.GET("/api/logs", listLogs)

	// WebSocket路由
	// r.GET("/ws", HandleWebSocket)

	r.POST("/api/sn", HandleSnForm)

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
