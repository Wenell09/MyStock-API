package repository

import (
	"github.com/Wenell09/MyStock/internal/models"
	"gorm.io/gorm"
)

type ItemWarehouseRepositoryImpl struct {
	DB *gorm.DB
}

func NewItemWarehouseRepository(db *gorm.DB) ItemWarehouseRepository {
	return &ItemWarehouseRepositoryImpl{
		DB: db,
	}
}

// Create implements [ItemWarehouseRepository].
func (i *ItemWarehouseRepositoryImpl) Create(tx *gorm.DB, data models.ItemWarehouse) error {
	return tx.Create(&data).Error
}

// FindByItemAndWarehouse implements [ItemWarehouseRepository].
func (i *ItemWarehouseRepositoryImpl) FindByItemAndWarehouse(tx *gorm.DB, itemID uint, warehouseID uint) (models.ItemWarehouse, error) {
	var data models.ItemWarehouse
	err := tx.Where("item_id = ? AND warehouse_id = ?", itemID, warehouseID).First(&data).Error
	if err != nil {
		return models.ItemWarehouse{}, err
	}
	return data, nil
}

// Update implements [ItemWarehouseRepository].
func (i *ItemWarehouseRepositoryImpl) Update(tx *gorm.DB, data models.ItemWarehouse) error {
	return tx.Model(&models.ItemWarehouse{}).Where("item_id = ? AND warehouse_id = ?", data.ItemID, data.WarehouseID).Update("stock", data.Stock).Error
}
