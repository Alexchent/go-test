package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

// 设置websocket
// CheckOrigin防止跨站点的请求伪造
var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("hello"))

	//升级get请求为webSocket协议
	ws, err := upGrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer ws.Close() //返回前关闭

	for {
		//读取ws中的数据
		mt, message, err := ws.ReadMessage()
		if err != nil {
			break
		}
		log.Print(string(message))

		//写入ws数据
		err = ws.WriteMessage(mt, message)
		if err != nil {
			break
		}
	}
}

func main() {
	http.HandleFunc("/ws", wsHandler)

	fmt.Println("启动ws服务 :8080")
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		return
	}
}
