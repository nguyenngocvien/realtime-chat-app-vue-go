package handlers

import (
	"log"
	"net/http"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // Adjust for production
	},
}

var clients = make(map[*websocket.Conn]string) // Map client to user ID
var broadcast = make(chan Message)

type Message struct {
	Sender    string `json:"sender"`
	Recipient string `json:"recipient"`
	Text      string `json:"text"`
	Timestamp string `json:"timestamp"`
}

func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}
	defer conn.Close()

	userID := r.URL.Query().Get("userID") // Get userID from query
	clients[conn] = userID

	for {
		var msg Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Println("Read error:", err)
			delete(clients, conn)
			break
		}
		msg.Sender = userID
		broadcast <- msg
	}
}

func HandleMessages() {
	for {
		msg := <-broadcast
		for client, clientID := range clients {
			if clientID == msg.Recipient || clientID == msg.Sender {
				err := client.WriteJSON(msg)
				if err != nil {
					log.Println("Write error:", err)
					client.Close()
					delete(clients, client)
				}
			}
		}
	}
}