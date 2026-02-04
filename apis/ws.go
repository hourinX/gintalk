package apis

import (
	"fmt"
	"gin-online-chat-backend/commons"
	"sync"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/websocket"
)

var onlineUsers = struct {
	sync.RWMutex
	conns map[string]*websocket.Conn
}{
	conns: make(map[string]*websocket.Conn),
}

func WsHandler(c *gin.Context) {
	userId := c.GetString("userId")
	if userId == "" {
		commons.Fail(c, commons.CodeUnauthorized, "用户ID不存在")
		return
	}
	websocket.Handler(func(ws *websocket.Conn) {
		defer func() {
			// 用户下线
			onlineUsers.Lock()
			delete(onlineUsers.conns, userId)
			onlineUsers.Unlock()
			ws.Close()
			fmt.Println("User disconnected:", userId)
		}()

		// 用户上线
		onlineUsers.Lock()
		onlineUsers.conns[userId] = ws
		onlineUsers.Unlock()
		fmt.Println("User connected:", userId)

		for {
			var msg string
			// 读取消息
			if err := websocket.Message.Receive(ws, &msg); err != nil {
				fmt.Println("Receive error:", err)
				break
			}
			fmt.Println("Received message from", userId, ":", msg)

			// 解析消息 JSON
			// 假设前端发送 {"to":"用户ID","message":"内容"}
			toUserId := "" // 这里你需要解析 msg JSON
			// TODO: 可以解析成 map[string]interface{} 或 struct

			onlineUsers.RLock()
			if toConn, ok := onlineUsers.conns[toUserId]; ok {
				if err := websocket.Message.Send(toConn, msg); err != nil {
					fmt.Println("Send error:", err)
				}
			}
			onlineUsers.RUnlock()
		}
	}).ServeHTTP(c.Writer, c.Request)
}
