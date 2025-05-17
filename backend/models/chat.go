package models

import (
	"time"
)

type ChatType string

const (
	ChatTypeUser  ChatType = "user"
	ChatTypeGroup ChatType = "group"
)

type Chat struct {
	ID           string    `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Type         ChatType  `gorm:"type:varchar(10);not null"`
	Name         string    `gorm:"not null"`
	Avatar       string    `gorm:""`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
	Participants []User    `gorm:"many2many:chat_participants;constraint:OnDelete:CASCADE"`
	Messages     []Message `gorm:"foreignKey:RecipientID;constraint:OnDelete:CASCADE"`
}
