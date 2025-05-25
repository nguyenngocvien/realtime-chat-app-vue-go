package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/nguyenngocvien/realtime-chat-app/models"
)

func SearchUsers(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	if query == "" {
		http.Error(w, "Query parameter 'q' is required", http.StatusBadRequest)
		return
	}

	var users []models.User
	if err := db.Where("username ILIKE ?", "%"+query+"%").Find(&users).Error; err != nil {
		http.Error(w, "Failed to search users", http.StatusInternalServerError)
		return
	}

	response := []struct {
		ID       string `json:"id"`
		Username string `json:"username"`
	}{}
	for _, user := range users {
		response = append(response, struct {
			ID       string `json:"id"`
			Username string `json:"username"`
		}{ID: user.ID, Username: user.Username})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"users": response,
	})
}
