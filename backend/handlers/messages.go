package handlers

import (
	"encoding/json"
	"net/http"
)

func SearchMessages(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userId").(string)
	query := r.URL.Query().Get("q")
	if query == "" {
		http.Error(w, "Query parameter 'q' is required", http.StatusBadRequest)
		return
	}

	var messages []struct {
		ID        string `json:"id"`
		ChatID    string `json:"chatId"`
		ChatName  string `json:"chatName"`
		Content   string `json:"content"`
		Timestamp string `json:"timestamp"`
	}
	if err := db.Table("messages").
		Select("messages.id, messages.chat_id, chats.name as chat_name, messages.text as content, messages.timestamp").
		Joins("JOIN chats ON messages.chat_id = chats.id").
		Joins("JOIN chat_participants ON chats.id = chat_participants.chat_id").
		Where("chat_participants.user_id = ? AND messages.text ILIKE ?", userID, "%"+query+"%").
		Order("messages.timestamp DESC").
		Limit(20).
		Find(&messages).Error; err != nil {
		http.Error(w, "Failed to search messages", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"messages": messages,
	})
}
