package handlers

import (
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/websocket"
	"github.com/nguyenngocvien/realtime-chat-app/config"
	"github.com/nguyenngocvien/realtime-chat-app/models"
	"gorm.io/gorm"
)

var clients = make(map[*websocket.Conn]string) // userId -> conn
var broadcast = make(chan models.Message)
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Cập nhật sau khi deploy
	},
}

func Init(db *gorm.DB) {
	go handleMessages(db)
}

func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	// Kiểm tra JWT từ query param
	tokenString := r.URL.Query().Get("token")
	if tokenString == "" {
		http.Error(w, "Token required", http.StatusUnauthorized)
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, http.ErrAbortHandler
		}
		return []byte(config.AppConfig.JWTSecretKey), nil
	})

	if err != nil || !token.Valid {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		http.Error(w, "Invalid token claims", http.StatusUnauthorized)
		return
	}

	userID, ok := claims["userId"].(string)
	if !ok {
		http.Error(w, "Invalid userId in token", http.StatusUnauthorized)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket upgrade error:", err)
		return
	}

	clients[conn] = userID

	for {
		var msg models.Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Println("WebSocket read error:", err)
			delete(clients, conn)
			break
		}

		msg.SenderID = userID
		broadcast <- msg
	}
}

func handleMessages(db *gorm.DB) {
	for {
		msg := <-broadcast
		for conn, userID := range clients {
			if msg.ChatID == userID || msg.SenderID == userID {
				err := conn.WriteJSON(msg)
				if err != nil {
					log.Println("WebSocket write error:", err)
					conn.Close()
					delete(clients, conn)
				}
			}
		}

		if err := db.Create(&msg).Error; err != nil {
			log.Println("Failed to save message:", err)
		}
	}
}
