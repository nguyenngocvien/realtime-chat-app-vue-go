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

	// Routes without authorize
	router.HandleFunc("/api/login", handlers.Login).Methods("POST")

	// Routes needs JWT
	protected := router.PathPrefix("/api").Subrouter()
	protected.Use(handlers.JWTMiddleware)
	protected.HandleFunc("/users/{userId}/chats", handlers.GetChatsByUserID).Methods("GET")

	// Websocket
	router.HandleFunc("/ws", handlers.HandleWebSocket)

	// router.Use(func(next http.Handler) http.Handler {
	// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	// 		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	// 		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	// 		if r.Method == "OPTIONS" {
	// 			w.WriteHeader(http.StatusOK)
	// 			return
	// 		}
	// 		next.ServeHTTP(w, r)
	// 	})
	// })

	log.Println("Server started on :8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
