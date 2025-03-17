package models

type Order struct {
	ID     uint    `json:"id" gorm:"primaryKey"`
	Total  float64 `json:"total" gorm:"not null"`
	UserID uint    `json:"user_id" gorm:"not null"`
}
