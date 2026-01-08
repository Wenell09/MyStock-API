package models

import "time"

type Category struct {
	ID        uint   `gorm:"primaryKey"`
	PublicID  string `gorm:"uniqueIndex;not null" validate:"required"`
	Name      string `gorm:"size:100;not null" validate:"required"`
	Items     []Item `gorm:"foreignKey:CategoryID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
