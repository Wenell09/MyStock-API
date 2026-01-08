package models

import "time"

type Supplier struct {
	ID        uint   `gorm:"primaryKey"`
	PublicID  string `gorm:"uniqueIndex;not null"`
	Name      string `gorm:"size:100;not null"`
	Items     []Item `gorm:"foreignKey:SupplierID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
