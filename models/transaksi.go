package models

type Transaction struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	UserID      uint   `json:"user_id"`
	Description string `json:"description"`
	Amount      uint   `json:"amount"`
}

type UserTransaction struct {
	ID           uint `gorm:"primaryKey" json:"id"`
	UserID       uint `json:"user_id"`
	TransactionID uint `json:"transaction_id"`
}
