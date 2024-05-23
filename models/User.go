package models

// User adalah model untuk entitas pengguna
type User struct {
	ID           uint   `gorm:"primaryKey" json:"id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	PasswordHash string `json:"-"`
}

