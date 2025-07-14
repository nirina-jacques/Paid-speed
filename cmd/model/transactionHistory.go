package model

import (
	"time"

	"gorm.io/gorm"
)

type TransactionHistory struct {
	ID string `gorm:"primaryKey" json:"id"`

	UserID string `gorm:"index" json:"user_id"`          // FK vers User
	User   User   `gorm:"foreignKey:UserID" json:"user"` // Relation GORM

	AmountEUR float64 `json:"amount_eur"`
	AmountAr  float64 `json:"amount_ar"`
	Rate      float64 `json:"rate"`

	MobileMoneyNumber   string `json:"mobile_money_number"`
	MobileMoneyProvider string `json:"mobile_money_provider"`

	PaypalTxnID  string `gorm:"index" json:"paypal_txn_id"`
	PaypalStatus string `json:"paypal_status"`

	PayoutStatus string `json:"payout_status"`
	Note         string `json:"note,omitempty"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
