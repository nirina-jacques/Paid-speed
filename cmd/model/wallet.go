package model

type Wallet struct {
	ID         uint    `gorm:"primaryKey"`
	UserID     string  `gorm:"index"`
	CurrencyID uint    `gorm:"index"`
	Balance    float64 `gorm:"default:0"`

	User     User     `gorm:"foreignKey:UserID"`
	Currency Currency `gorm:"foreignKey:CurrencyID"`
}
