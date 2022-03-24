package main

import (
	"fmt"
	"github.com/gochat/pkg/mywebsocket"
	"net/http"
)

// 定义 WebSocket 服务处理函数
func serveWs(w http.ResponseWriter, r *http.Request) {
	ws, err := mywebsocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+V\n", err)
	}
	go mywebsocket.Writer(ws)
	mywebsocket.Reader(ws)
}

func setupRoutes() {
	http.HandleFunc("/ws", serveWs)
}

func main() {
	fmt.Println("Distributed Chat App v0.01")
	setupRoutes()
	http.ListenAndServe(":8080", nil)
}