package models

import "time"

type Item struct {
	ID             uint   `gorm:"primaryKey"`
	PublicID       string `gorm:"uniqueIndex;not null"`
	Name           string `gorm:"size:100;not null"`
	CategoryID     uint
	Category       Category `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	SupplierID     uint
	Supplier       Supplier           `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	ItemWarehouses []ItemWarehouse    `gorm:"foreignKey:ItemID"`
	Transactions   []StockTransaction `gorm:"foreignKey:ItemID"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
