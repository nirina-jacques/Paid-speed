package model

import (
	"time"

	"gorm.io/gorm"
)

type Transfer struct {
	ID                  string `gorm:"primaryKey"`
	ClientEmail         string
	AmountEUR           float64
	AmountAr            float64
	Rate                float64
	MobileMoneyNumber   string
	MobileMoneyProvider string
	PaypalTxnID         string
	PaypalStatus        string
	PayoutStatus        string
	Note                string
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DeletedAt           gorm.DeletedAt `gorm:"index"` // soft delete option
}
