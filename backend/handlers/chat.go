package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/nguyenngocvien/realtime-chat-app/models"
	"gorm.io/gorm"
)

type ChatResponse struct {
	ID           string           `json:"id"`
	Type         models.ChatType  `json:"type"`
	Name         string           `json:"name"`
	Avatar       string           `json:"avatar,omitempty"`
	LastMessage  *MessageResponse `json:"lastMessage,omitempty"`
	UnreadCount  int              `json:"unreadCount"`
	Participants []UserResponse   `json:"participants,omitempty"`
}

type MessageResponse struct {
	ID          uint   `json:"id"`
	SenderID    string `json:"senderId"`
	RecipientID string `json:"recipientId"`
	Text        string `json:"text"`
	Timestamp   string `json:"timestamp"`
}

type UserResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

var db *gorm.DB

func InitChatHandler(database *gorm.DB) {
	db = database
}

func GetChatsByUserID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["userId"]

	var chats []models.Chat
	err := db.
		Preload("Participants").
		Preload("Messages", func(db *gorm.DB) *gorm.DB {
			return db.Order("timestamp DESC").Limit(1)
		}).
		Joins("JOIN chat_participants ON chat_participants.chat_id = chats.id").
		Where("chat_participants.user_id = ?", userID).
		Find(&chats).Error
	if err != nil {
		log.Println("Error fetching chats:", err)
		http.Error(w, "Failed to fetch chats", http.StatusInternalServerError)
		return
	}

	response := []ChatResponse{}
	for _, chat := range chats {
		var unreadCount int64
		db.Model(&models.Message{}).
			Where("recipient_id = ? AND read = ? AND sender_id != ?", chat.ID, false, userID).
			Count(&unreadCount)

		var lastMessage *MessageResponse
		if len(chat.Messages) > 0 {
			msg := chat.Messages[0]
			lastMessage = &MessageResponse{
				ID:          msg.ID,
				SenderID:    msg.SenderID,
				RecipientID: msg.RecipientID,
				Text:        msg.Text,
				Timestamp:   msg.Timestamp.Format(time.RFC3339),
			}
		}

		participants := []UserResponse{}
		for _, p := range chat.Participants {
			participants = append(participants, UserResponse{
				ID:       p.ID,
				Username: p.Username,
			})
		}

		response = append(response, ChatResponse{
			ID:           chat.ID,
			Type:         chat.Type,
			Name:         chat.Name,
			Avatar:       chat.Avatar,
			LastMessage:  lastMessage,
			UnreadCount:  int(unreadCount),
			Participants: participants,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"chats": response,
	})
}
