package models

import "time"

type StockTransactionType string

const (
	StockIn  StockTransactionType = "IN"
	StockOut StockTransactionType = "OUT"
)

type StockTransaction struct {
	ID          uint   `gorm:"primaryKey"`
	PublicID    string `gorm:"uniqueIndex;not null"`
	ItemID      uint
	Item        Item `gorm:"constraint:OnDelete:CASCADE;"`
	WarehouseID uint
	Warehouse   Warehouse            `gorm:"constraint:OnDelete:CASCADE;"`
	Type        StockTransactionType `gorm:"type:stock_transaction_type;not null"`
	Quantity    int                  `gorm:"not null;check:quantity > 0"`
	CreatedAt   time.Time
}
