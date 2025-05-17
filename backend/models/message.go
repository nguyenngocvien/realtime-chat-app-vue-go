package models

import "time"

type Message struct {
	ID          uint      `gorm:"primaryKey"`
	SenderID    string    `gorm:"not null"`
	RecipientID string    `gorm:"not null"`
	Text        string    `gorm:"not null"`
	Timestamp   time.Time `gorm:"not null"`
	Read        bool      `gorm:"default:false"`
}
