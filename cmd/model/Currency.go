package model

import "time"

type Currency struct {
	ID        uint   `gorm:"primaryKey"`
	Code      string `gorm:"uniqueIndex"` // ex: USD, EUR, MGA
	Name      string
	Symbol    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
