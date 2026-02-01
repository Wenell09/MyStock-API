package repository

import (
	"github.com/Wenell09/MyStock/internal/models"
	"gorm.io/gorm"
)

type ItemWarehouseRepository interface {
	FindByItemAndWarehouse(tx *gorm.DB, itemID uint, warehouseID uint) (models.ItemWarehouse, error)
	Create(tx *gorm.DB, data models.ItemWarehouse) error
	Update(tx *gorm.DB, data models.ItemWarehouse) error
}
