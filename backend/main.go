package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nguyenngocvien/realtime-chat-app/config"
	"github.com/nguyenngocvien/realtime-chat-app/db"
	"github.com/nguyenngocvien/realtime-chat-app/handlers"
)

func main() {
	// Load config first
	config.LoadConfig()

	dbConn := db.InitDB()
	handlers.Init(dbConn)
	handlers.InitChatHandler(dbConn)

	router := mux.NewRouter()

	// CORS middleware
	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Printf("Received %s request to %s", r.Method, r.URL.Path)
			w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			w.Header().Set("Access-Control-Max-Age", "86400") // Cache preflight 24h
			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}
			next.ServeHTTP(w, r)
		})
	})

	// Routes without authorize
	router.HandleFunc("/api/login", handlers.Login).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/signup", handlers.Signup).Methods("POST", "OPTIONS")

	// Routes needs JWT
	protected := router.PathPrefix("/api").Subrouter()
	protected.Use(handlers.JWTMiddleware)
	protected.HandleFunc("/chats", handlers.CreateGroupChat).Methods("POST", "OPTIONS")
	protected.HandleFunc("/users/{userId}/chats", handlers.GetChatsByUserID).Methods("GET", "OPTIONS")
	protected.HandleFunc("/chats/{chatId}/messages", handlers.GetMessagesByChatID).Methods("GET", "OPTIONS")

	// Websocket
	router.HandleFunc("/ws", handlers.HandleWebSocket)

	log.Println("Server started on :8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
