package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/nguyenngocvien/realtime-chat-app/models"
	"gorm.io/gorm"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // Adjust for production
	},
}

var clients = make(map[*websocket.Conn]string) // Map connection to user ID
var broadcast = make(chan models.Message)
var dbConn *gorm.DB

func Init(db *gorm.DB) {
	dbConn = db
	go handleMessages()
}

type ClientMessage struct {
	RecipientID string `json:"recipientId"`
	Text        string `json:"text"`
}

func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}
	defer conn.Close()

	userID := r.URL.Query().Get("userId")
	if userID == "" {
		log.Println("Missing userId")
		return
	}
	clients[conn] = userID

	// Send existing messages
	var messages []models.Message
	dbConn.Where("sender_id = ? OR recipient_id = ?", userID, userID).Find(&messages)
	for _, msg := range messages {
		err := conn.WriteJSON(msg)
		if err != nil {
			log.Println("Write error:", err)
			conn.Close()
			delete(clients, conn)
			return
		}
	}

	for {
		var clientMsg ClientMessage
		err := conn.ReadJSON(&clientMsg)
		if err != nil {
			log.Println("Read error:", err)
			delete(clients, conn)
			break
		}

		msg := models.Message{
			SenderID:    userID,
			RecipientID: clientMsg.RecipientID,
			Text:        clientMsg.Text,
			Timestamp:   time.Now(),
		}
		dbConn.Create(&msg)
		broadcast <- msg
	}
}

func handleMessages() {
	for {
		msg := <-broadcast
		for client, userID := range clients {
			if userID == msg.SenderID || userID == msg.RecipientID {
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
