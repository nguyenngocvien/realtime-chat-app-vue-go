package models

type User struct {
	ID       string `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	Username string `gorm:"type:text;unique" json:"username"`
	Password string `gorm:"type:text" json:"password"`
}
