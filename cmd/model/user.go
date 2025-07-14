package model

import (
	"time"

	"gorm.io/gorm"
)

type Roles []string

type User struct {
	ID           string `gorm:"primaryKey" json:"id"`     // UUID ou string
	Email        string `gorm:"uniqueIndex" json:"email"` // Email unique
	Phone        string `json:"phone"`                    // Numéro mobile
	FullName     string `json:"full_name"`                // Nom complet (optionnel)
	Password     string `json:"-"`                        // Hash du mot de passe
	PasswordSalt string `json:"-"`                        // Salt utilisé pour le hash

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Transactions []TransactionHistory `gorm:"foreignKey:UserID" json:"transactions"` // Association
}
