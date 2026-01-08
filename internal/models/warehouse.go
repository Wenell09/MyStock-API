package models

import "time"

type Warehouse struct {
	ID             uint               `gorm:"primaryKey"`
	PublicID       string             `gorm:"uniqueIndex;not null"`
	Name           string             `gorm:"size:100;not null"`
	ItemWarehouses []ItemWarehouse    `gorm:"foreignKey:WarehouseID"`
	Transactions   []StockTransaction `gorm:"foreignKey:WarehouseID"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
