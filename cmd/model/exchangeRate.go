package model

import "time"

type ExchangeRate struct {
	ID             uint `gorm:"primaryKey"`
	FromCurrencyID uint
	ToCurrencyID   uint
	Rate           float64
	LastUpdated    time.Time

	FromCurrency Currency `gorm:"foreignKey:FromCurrencyID"`
	ToCurrency   Currency `gorm:"foreignKey:ToCurrencyID"`
}
