package models

type User struct {
	ID       string `gorm:"primaryKey"`
	Username string `gorm:"not null"`
}
