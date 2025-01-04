package function

import (
	"fmt"
	"github.com/gorilla/websocket"
	config "myapp/config/future"
)

func WsDetails(subscribeRequest []byte) []byte {

	// 建立 WebSocket 连接
	conn, _, err := websocket.DefaultDialer.Dial(config.TestWSUrl, nil)
	if err != nil {
		fmt.Println("无法连接到 WebSocket 服务器:", err)
	}
	defer conn.Close()

	// 发送订阅请求
	err = conn.WriteMessage(websocket.TextMessage, subscribeRequest)
	if err != nil {
		fmt.Println("发送订阅请求失败:", err)
	}

	// 读取 WebSocket 消息
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("读取消息时出错:", err)
		}
		return message
	}
}
