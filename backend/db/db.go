package db

import (
	"log"

	"github.com/nguyenngocvien/realtime-chat-app/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dsn := "host=localhost user=postgres password=123456 dbname=chatdb port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	// AutoMigrate create or update table base on model
	err = db.AutoMigrate(&models.Message{}, &models.User{})
	if err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}

	log.Println("Database migration completed")
	return db
}
