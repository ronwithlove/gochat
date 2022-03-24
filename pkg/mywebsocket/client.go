package mywebsocket

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
)

type Client struct {
	ID   string
	Conn *websocket.Conn
	Pool *Pool
}

type Message struct {
	Type     int    `json:"type"`
	ClientID string `json:"clientid"`
	Body     string `json:"body"`
}

func NewClient(conn *websocket.Conn, pool *Pool) *Client {
	return &Client{
		Conn: conn,
		Pool: pool,
	}
}

func (c *Client) Read() {
	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	for {
		messageType, p, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		message := Message{Type: messageType, Body: string(p)}
		c.Pool.Broadcast <- message
		//c.Pool.PrivateTalk.PrivateMsg <-message
		fmt.Printf("Message Received: %+v\n", message)
	}
}
