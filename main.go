package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gochat/pkg/mywebsocket"
	"net/http"
)

func serveWs(pool *mywebsocket.Pool, w http.ResponseWriter, r *http.Request) {
	fmt.Println("WebSocket Endpoint Hit")
	conn, err := mywebsocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}

	client:=mywebsocket.NewClient(conn,pool)

	pool.Register <- client
	client.Read()
}

func setupRoutes() {
	pool := mywebsocket.NewPool()
	go pool.Start()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(pool, w, r)
	})
}

func main() {
	fmt.Println("Distributed Chat App v0.02")
	//setupRoutes()
	//http.ListenAndServe(":8080", nil)

	r := Setup()
	r.Run()

}

func Setup() *gin.Engine {
	r := gin.Default()

	r.GET("/ws", IndexHandler)

	return r
}

func IndexHandler(c *gin.Context) {
	pool := mywebsocket.NewPool()
	go pool.Start()

	fmt.Println("WebSocket Endpoint Hit")
	conn, err := mywebsocket.Upgrade(c.Writer, c.Request)
	if err != nil {
		fmt.Fprintf(c.Writer, "%+v\n", err)
	}

	client:=mywebsocket.NewClient(conn,pool)

	pool.Register <- client
	client.Read()
}
