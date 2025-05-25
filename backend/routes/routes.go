package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nguyenngocvien/realtime-chat-app/handlers"
)

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	// CORS middleware
	router.Use(corsMiddleware)

	// Public routes
	router.HandleFunc("/api/login", handlers.Login).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/signup", handlers.Signup).Methods("POST", "OPTIONS")

	// Protected routes
	protected := router.PathPrefix("/api").Subrouter()
	protected.Use(handlers.JWTMiddleware)

	protected.HandleFunc("/users/{userId}/chats", handlers.GetChatsByUserID).Methods("GET", "OPTIONS")
	protected.HandleFunc("/users/search", handlers.SearchUsers).Methods("GET", "OPTIONS")

	protected.HandleFunc("/chats", handlers.CreateGroupChat).Methods("POST", "OPTIONS")
	protected.HandleFunc("/chats/{chatId}/messages", handlers.GetMessagesByChatID).Methods("GET", "OPTIONS")

	protected.HandleFunc("/messages/search", handlers.SearchMessages).Methods("GET", "OPTIONS")

	// WebSocket
	router.HandleFunc("/ws", handlers.HandleWebSocket)

	return router
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Max-Age", "86400")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}
