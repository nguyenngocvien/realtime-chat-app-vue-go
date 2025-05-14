package main

import (
	"log"
	"net/http"

	"github.com/nguyenngocvien/realtime-chat-app/handlers"
)

func main() {
	go handlers.HandleMessages()
	http.HandleFunc("/ws", handlers.HandleWebSocket)
	log.Println("Server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
