package model

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	ID                  string `gorm:"primaryKey"`
	ClientEmail         string
	AmountEUR           float64
	AmountAr            float64
	Rate                float64
	CurrencyID          uint
	MobileMoneyNumber   string
	MobileMoneyProvider string
	PaypalTxnID         string
	PaypalStatus        string
	PayoutStatus        string
	Description         string
	Provider            string // paypal, mobilemoney
	TargetAccount       string // email PayPal ou numéro Mobile
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DeletedAt           gorm.DeletedAt `gorm:"index"` // soft delete option

	User     User     `gorm:"foreignKey:UserID"`
	Currency Currency `gorm:"foreignKey:CurrencyID"`
}
