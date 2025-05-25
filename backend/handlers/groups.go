package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/nguyenngocvien/realtime-chat-app/models"
)

func CreateGroupChat(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name         string   `json:"name"`
		Type         string   `json:"type"`
		Participants []string `json:"participants"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	if input.Name == "" || input.Type != "group" || len(input.Participants) == 0 {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	chat := models.Chat{
		ID:   uuid.New().String(),
		Name: input.Name,
		Type: models.ChatType(input.Type),
	}
	if err := db.Create(&chat).Error; err != nil {
		http.Error(w, "Failed to create chat", http.StatusInternalServerError)
		return
	}

	for _, userID := range input.Participants {
		participant := models.ChatParticipant{
			ChatID: chat.ID,
			UserID: userID,
		}
		if err := db.Create(&participant).Error; err != nil {
			http.Error(w, "Failed to add participant", http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"chat": chat,
	})
}
