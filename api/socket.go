package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

//webSocket 请求 ping 返回 pong
func ping(c *gin.Context) {
	// 升级 get 请求为 webSocket 协议
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer ws.Close()
	for {
		log.Print("get message...")
		// // 读取 ws 中的数据
		mt, message, err := ws.ReadMessage()
		if err != nil {
			break
		}
		if string(message) == "ping" {
			message = []byte("pong")
		}
		log.Printf("收到消息：type: %s , mes: %s", mt, message)

		// processMsg := make(chan []byte, 5)
		go func(wsConn *websocket.Conn, msg []byte) {
			//mt, message, _ := ws.ReadMessage()
			log.Printf("收到消息：mes: %s", msg)
			e := wsConn.WriteMessage(1, msg)
			if e != nil {
				log.Printf("发消息: err  %s", e)
			}
			// for {
			// 	select {
			// 	case msg := <-processMsg:
			// 		err = ws.WriteMessage(mt, msg)
			// 		if err != nil {
			// 			log.Printf("err %s", err)
			// 		}
			// 	default:
			// 		//ws.WriteMessage(mt, []byte("nothing..."))
			// 	}
			// }
		}(ws, message)
		log.Print("get over...")
	}
}
