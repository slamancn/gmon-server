package ws

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

func Index(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

func InitWebSocket(c *gin.Context) {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			log.Println("升级协议", r.Header["User-Agent"])
			return true
		},
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()
	userId := c.Query("id")
	log.Println("用户id:", userId)
	// 如果用户id为空 则无法连接成功
	if userId == "" {
		log.Println("无用户id")
		return
	}

	// 查询redis中用户id的链接状态,如果已连接则关掉之前的链接使用新的链接
	
	for {
		mt, message, err := conn.ReadMessage()

		log.Println("获取客户端发送的消息:" + string(message))
		fmt.Println(mt)
		if err != nil {
			log.Println(err)
			break
		}

		var msg = "春风得意马蹄疾,一日看尽长安花"
		err2 := conn.WriteMessage(websocket.TextMessage, []byte(msg))
		if err2 != nil {
			log.Println("write:", err2)
			break
		}
	}

}
