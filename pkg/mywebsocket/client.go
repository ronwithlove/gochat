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
	MyID     string `json:"myid"`
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
		_, p, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		var r RecivedString
		err = json.Unmarshal(p, &r)
		if err != nil {
			fmt.Printf(err.Error())
			return
		}

		if r.ChatType == "Private" {
			r.MyID=c.ID
			fmt.Println("private","c.ID:",c.ID,"myid:",r.MyID)
			c.Pool.PrivateTalk <- r

		} else {
			c.Pool.Broadcast <- r
			fmt.Printf("%s Broadcast Message: %+v\n", c.ID, r)
		}

	}
}
