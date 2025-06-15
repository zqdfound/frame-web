package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/hpcloud/tail"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func HandleWebSocket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("WebSocket upgrade error: %v", err)
		return
	}
	defer conn.Close()

	filename := c.Query("file")
	log.Printf("WebSocket 连接: %s", filename)

	if filename == "" {
		//return
		filename = "./log/log1.txt" // 默认日志文件
	}
	log.Printf("WebSocket 连接11111: %s", filename)

	t, err := tail.TailFile(filename, tail.Config{
		Follow:   true,
		ReOpen:   true,
		Location: &tail.SeekInfo{Offset: 0, Whence: 2}, // 从文件末尾开始
	})
	if err != nil {
		log.Printf("Tail file error: %v", err)
		return
	}

	for line := range t.Lines {
		err := conn.WriteMessage(websocket.TextMessage, []byte(line.Text))
		if err != nil {
			log.Printf("WebSocket write error: %v", err)
			return
		}
	}
}
