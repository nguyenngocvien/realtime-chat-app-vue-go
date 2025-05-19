package db

import (
	"fmt"
	"log"

	"github.com/nguyenngocvien/realtime-chat-app/config"
	"github.com/nguyenngocvien/realtime-chat-app/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	c := config.AppConfig

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		c.DBHost, c.DBUser, c.DBPassword, c.DBName, c.DBPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	// Enable uuid-ossp extension for UUID
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")

	// AutoMigrate create or update table base on model
	err = db.AutoMigrate(&models.Chat{}, &models.User{}, &models.Message{}, &models.ChatParticipant{})
	if err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}

	log.Println("Database migration completed")
	return db
}
