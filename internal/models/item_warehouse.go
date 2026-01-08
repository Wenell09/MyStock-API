package models

import "time"

type ItemWarehouse struct {
	ID          uint   `gorm:"primaryKey"`
	PublicID    string `gorm:"uniqueIndex;not null" `
	ItemID      uint
	Item        Item `gorm:"constraint:OnDelete:CASCADE;"`
	WarehouseID uint
	Warehouse   Warehouse `gorm:"constraint:OnDelete:CASCADE;"`
	Stock       int       `gorm:"not null;default:0"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
