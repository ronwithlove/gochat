package mywebsocket

import "fmt"

type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	Broadcast  chan Message
	PrivateTalk
}

type PrivateTalk struct {
	PrivateMsg chan Message
	Client
}

func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan Message),
	}
}

func (pool *Pool) Start() {
	for {
		select {
		case client := <-pool.Register:
			pool.Clients[client] = true
			fmt.Println("Size of Connection Pool: ", len(pool.Clients))
			for client, _ := range pool.Clients {
				fmt.Println(client.Conn.RemoteAddr().String())
				client.Conn.WriteJSON(Message{Type: 1, Body: "New user: " + client.Conn.RemoteAddr().String() + " Joined..."})
			}
			break
		case client := <-pool.Unregister:
			delete(pool.Clients, client)
			fmt.Println("Size of Connection Pool: ", len(pool.Clients))
			for client, _ := range pool.Clients {
				client.Conn.WriteJSON(Message{Type: 1, Body: "User: " + client.Conn.RemoteAddr().String() + " Disconnected..."})
			}
			break
		case message := <-pool.Broadcast:
			fmt.Println("Sending message to all clients in Pool: ",message)
			for client, _ := range pool.Clients {
				if err := client.Conn.WriteJSON(message); err != nil {
					fmt.Println(err)
					return
				}
			}
		case privateMsg := <-pool.PrivateTalk.PrivateMsg:
			fmt.Println("Sending message to all clients in Pool")
			if _, ok := pool.Clients[&pool.PrivateTalk.Client]; ok {
				if err := pool.PrivateTalk.Client.Conn.WriteJSON(privateMsg); err != nil {
					fmt.Println(err)
					return
				}
			}
		}
	}
}
