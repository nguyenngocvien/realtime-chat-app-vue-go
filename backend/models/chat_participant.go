package models

type ChatParticipant struct {
	ChatID string `gorm:"type:uuid;primaryKey"`
	UserID string `gorm:"primaryKey"`
}
