package models

import "time"

type Message struct {
	ID        uint      `gorm:"primaryKey"`
	Sender    string    `gorm:"not null"`
	Recipient string    `gorm:"not null"`
	Text      string    `gorm:"not null"`
	Timestamp time.Time `gorm:"not null"`
}
