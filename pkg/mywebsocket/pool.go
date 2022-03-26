package mywebsocket

import "fmt"

type Pool struct {
	Register    chan *Client
	Unregister  chan *Client
	Clients     map[*Client]bool
	Broadcast   chan RecivedString
	PrivateTalk chan RecivedString
}

func NewPool() *Pool {
	return &Pool{
		Register:    make(chan *Client),
		Unregister:  make(chan *Client),
		Clients:     make(map[*Client]bool),
		Broadcast:   make(chan RecivedString),
		PrivateTalk: make(chan RecivedString),
	}
}

func (pool *Pool) Start() {
	for {
		select {
		case client := <-pool.Register:
			pool.Clients[client] = true
			fmt.Println("Server Size of Connection Pool: ", len(pool.Clients))
			for client, _ := range pool.Clients {
				fmt.Println(client.ID)
				err := client.Conn.WriteJSON(RecivedString{
					ChatType:     "Public",
					ClientID: client.ID,
					Message:     "New User Joined...",
				})
				if err != nil {
					fmt.Println(err)
				}
			}
			break
		case client := <-pool.Unregister:
			delete(pool.Clients, client)
			fmt.Println("Server Size of Connection Pool: ", len(pool.Clients))
			for client, _ := range pool.Clients {
				err := client.Conn.WriteJSON(RecivedString{
					ChatType:     "Public",
					ClientID: client.ID,
					Message:     "User Disconnected..."})
				if err != nil {
					fmt.Println(err)
				}
			}
			break
		case message := <-pool.Broadcast:
			fmt.Println("Server Broadcast: Sending message to all clients in Pool: ", message)
			for client, _ := range pool.Clients {
				if err := client.Conn.WriteJSON(message); err != nil {
					fmt.Println(err)
					return
				}
			}
		case privateMsg := <-pool.PrivateTalk:
			fmt.Println("Sending message to one client in Pool",privateMsg.MyID,":",privateMsg.ClientID)
			for client := range pool.Clients {
				if client.ID == privateMsg.ClientID {
					fmt.Println("match:", client.ID)
					if err := client.Conn.WriteJSON(privateMsg)
						err != nil {
						fmt.Println("private conn err:",err)
						return
					}
				}else if client.ID == privateMsg.MyID {
					fmt.Println("MyID:", client.ID)
					if err := client.Conn.WriteJSON(privateMsg)
						err != nil {
						fmt.Println("private conn err:",err)
						return
					}
				}
			}
		}
	}
}
