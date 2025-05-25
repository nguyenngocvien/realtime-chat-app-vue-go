package models

import "time"

type Message struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	ChatID    string    `gorm:"not null" json:"chatId"`
	ChatName  string    `gorm:"not null" json:"chatName"`
	SenderID  string    `gorm:"not null" json:"senderId"`
	Text      string    `gorm:"not null" json:"text"`
	Timestamp time.Time `gorm:"not null" json:"timestamp"`
	Read      bool      `gorm:"default:false" json:"read"`
}
