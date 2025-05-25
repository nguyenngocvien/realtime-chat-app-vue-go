package main

import (
	"log"
	"net/http"

	"github.com/nguyenngocvien/realtime-chat-app/config"
	"github.com/nguyenngocvien/realtime-chat-app/db"
	"github.com/nguyenngocvien/realtime-chat-app/handlers"
	"github.com/nguyenngocvien/realtime-chat-app/routes"
)

func main() {
	// Load config first
	config.LoadConfig()

	dbConn := db.InitDB()
	handlers.Init(dbConn)
	handlers.InitChatHandler(dbConn)

	router := routes.SetupRoutes()

	log.Println("Server started on :8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
