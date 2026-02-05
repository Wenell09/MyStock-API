package models

import "time"

type User struct {
	ID        string `gorm:"type:uuid;primaryKey"`
	Email     string `gorm:"type:varchar(255);uniqueIndex"`
	Name      string `gorm:"type:varchar(255)"`
	Role      string `gorm:"type:varchar(50);default:user"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
