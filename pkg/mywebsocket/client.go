package mywebsocket

import (
	"encoding/json"
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
		ID:   conn.RemoteAddr().String(),
		Pool: pool,
	}
}

type RecivedString struct {
	ChatType string `json:"chatType"`
	ClientID string `json:"clientId"`
	Message  string `json:"message"`
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
		message := Message{
			Type:     messageType,
			ClientID: c.ID,
			Body:     string(p),
		}

		var r RecivedString
		err = json.Unmarshal(p, &r)
		if err != nil {
			fmt.Printf(err.Error())
			return
		}

		if r.ChatType == "Private" {
			fmt.Println("private")
			c.Pool.PrivateTalk <- r

		} else {
			c.Pool.Broadcast <- message
			fmt.Printf("%s Broadcast Message: %+v\n", c.ID, message)
		}

	}
}
